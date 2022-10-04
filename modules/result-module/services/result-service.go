package services

import (
	"context"
	"log"
	"time"

	"github.com/GabrielAchumba/school-mgt-backend/common/config"
	"github.com/GabrielAchumba/school-mgt-backend/common/conversion"
	errors "github.com/GabrielAchumba/school-mgt-backend/common/errors"
	classRoomServicePackage "github.com/GabrielAchumba/school-mgt-backend/modules/classroom-module/services"
	"github.com/GabrielAchumba/school-mgt-backend/modules/result-module/dtos"
	"github.com/GabrielAchumba/school-mgt-backend/modules/result-module/models"
	studentServicePackage "github.com/GabrielAchumba/school-mgt-backend/modules/student-module/services"
	subjectServicePackage "github.com/GabrielAchumba/school-mgt-backend/modules/subject-module/services"
	userServicePackage "github.com/GabrielAchumba/school-mgt-backend/modules/user-module/services"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ResultService interface {
	CreateResult(userId string, requestModel dtos.CreateResultRequest) (interface{}, error)
	DeleteResult(id string) (int64, error)
	GetResult(id string) (dtos.ResultResponse, error)
	GetResults() ([]dtos.ResultResponse, error)
	PutResult(id string, item dtos.UpdateResultRequest) (interface{}, error)
}

type serviceImpl struct {
	ctx              context.Context
	collection       *mongo.Collection
	userService      userServicePackage.UserService
	studentService   studentServicePackage.StudentService
	subjectService   subjectServicePackage.SubjectService
	classRoomService classRoomServicePackage.ClassRoomService
}

func New(mongoClient *mongo.Client, config config.Settings, ctx context.Context,
	userService userServicePackage.UserService,
	studentService studentServicePackage.StudentService,
	subjectService subjectServicePackage.SubjectService,
	classRoomService classRoomServicePackage.ClassRoomService) ResultService {
	collection := mongoClient.Database(config.Database.DatabaseName).Collection(config.TableNames.Result)

	return &serviceImpl{
		collection:       collection,
		ctx:              ctx,
		userService:      userService,
		studentService:   studentService,
		subjectService:   subjectService,
		classRoomService: classRoomService,
	}
}

func (impl serviceImpl) DeleteResult(id string) (int64, error) {

	log.Print("Call to delete type of Result by id started.")
	objId := conversion.GetMongoId(id)
	filter := bson.D{bson.E{Key: "_id", Value: objId}}

	result, err := impl.collection.DeleteOne(impl.ctx, filter)

	if err != nil {
		return result.DeletedCount, errors.Error("Error in deleting type of Result.")
	}

	if result.DeletedCount < 1 {
		return result.DeletedCount, errors.Error("Type of Result with specified ID not found!")
	}

	log.Print("Call to delete type of Result by id completed.")
	return result.DeletedCount, nil
}

func (impl serviceImpl) GetResult(id string) (dtos.ResultResponse, error) {

	log.Print("Get Type of Result called")
	objId := conversion.GetMongoId(id)
	var Result dtos.ResultResponse

	filter := bson.D{bson.E{Key: "_id", Value: objId}}

	err := impl.collection.FindOne(impl.ctx, filter).Decode(&Result)
	if err != nil {
		return Result, errors.Error("could not find type of Result by id")
	}

	student, _ := impl.studentService.GetStudent(Result.StudentId)
	subject, _ := impl.subjectService.GetSubject(Result.SubjectId)
	teacher, _ := impl.userService.GetUser(Result.TeacherId)
	_classRoom, _ := impl.classRoomService.GetClassRoom(Result.ClassRoomId)

	Result.StudentFullName = student.FirstName + " " + student.LastName
	Result.SubjectFullName = subject.Type
	Result.TeacherFullName = teacher.FirstName + " " + teacher.LastName
	Result.ClassRoomFullName = _classRoom.Type

	log.Print("Get type of Result completed")
	return Result, err

}

func (impl serviceImpl) GetResults() ([]dtos.ResultResponse, error) {

	log.Print("Call to get all types of Result started.")

	var Results []dtos.ResultResponse
	filter := bson.D{}
	cur, err := impl.collection.Find(impl.ctx, filter)

	if err != nil {
		Results = make([]dtos.ResultResponse, 0)
		return Results, errors.Error("Types of Result not found!")
	}

	err = cur.All(impl.ctx, &Results)
	if err != nil {
		return Results, err
	}

	cur.Close(impl.ctx)
	length := len(Results)
	if length == 0 {
		Results = make([]dtos.ResultResponse, 0)
	}

	for i := 0; i < length; i++ {
		student, _ := impl.studentService.GetStudent(Results[i].StudentId)
		subject, _ := impl.subjectService.GetSubject(Results[i].SubjectId)
		teacher, _ := impl.userService.GetUser(Results[i].TeacherId)
		classRoom, _ := impl.classRoomService.GetClassRoom(Results[i].ClassRoomId)

		Results[i].StudentFullName = student.FirstName + " " + student.LastName
		Results[i].SubjectFullName = subject.Type
		Results[i].TeacherFullName = teacher.FirstName + " " + teacher.LastName
		Results[i].ClassRoomFullName = classRoom.Type

	}

	log.Print("Call to get all types of Result completed.")
	return Results, err
}

func (impl serviceImpl) ViewSelectdResults(req dtos.GetResultsRequest) ([]dtos.ResultResponse, error) {

	log.Print("Call to get results by range of date started")

	/* { tags: ["red", "blank"] } */

	var Results []dtos.ResultResponse
	filter := bson.D{bson.E{Key: "createdAt", Value: bson.D{bson.E{Key: "$gte", Value: req.StartDate}}},
		bson.E{Key: "nleveloneroomonechildren", Value: bson.D{bson.E{Key: "$lte", Value: req.EndDate}}},
		bson.E{Key: "subjectsid", Value: req.SubjectIds},
		bson.E{Key: "teacherid", Value: req.TeacherId},
		bson.E{Key: "studentid", Value: req.StudentId},
		bson.E{Key: "classroomid", Value: req.ClassRoomId}}
	cur, err := impl.collection.Find(impl.ctx, filter)

	if err != nil {
		Results = make([]dtos.ResultResponse, 0)
		return Results, errors.Error("Results not found!")
	}

	err = cur.All(impl.ctx, &Results)
	if err != nil {
		return Results, err
	}

	cur.Close(impl.ctx)
	length := len(Results)
	if length == 0 {
		Results = make([]dtos.ResultResponse, 0)
	}

	for i := 0; i < length; i++ {
		student, _ := impl.studentService.GetStudent(Results[i].StudentId)
		subject, _ := impl.subjectService.GetSubject(Results[i].SubjectId)
		teacher, _ := impl.userService.GetUser(Results[i].TeacherId)
		classRoom, _ := impl.classRoomService.GetClassRoom(Results[i].ClassRoomId)

		Results[i].StudentFullName = student.FirstName + " " + student.LastName
		Results[i].SubjectFullName = subject.Type
		Results[i].TeacherFullName = teacher.FirstName + " " + teacher.LastName
		Results[i].ClassRoomFullName = classRoom.Type

	}

	log.Print("Call to get results by range of date completed")
	return Results, err
}

func (impl serviceImpl) CreateResult(userId string, model dtos.CreateResultRequest) (interface{}, error) {

	log.Print("Call to create Result started.")

	var modelObj models.Result
	conversion.Convert(model, &modelObj)

	modelObj.CreatedBy = userId
	modelObj.CreatedAt = time.Now()
	modelObj.CreatedDay = modelObj.CreatedAt.Day()
	modelObj.CreatedMonth = int(modelObj.CreatedAt.Month())
	modelObj.CreatedYear = modelObj.CreatedAt.Year()

	filter := bson.D{bson.E{Key: "createdat", Value: modelObj.CreatedAt},
		bson.E{Key: "studentid", Value: modelObj.StudentId},
		bson.E{Key: "subjectid", Value: modelObj.SubjectId},
		bson.E{Key: "teacherid", Value: modelObj.TeacherId},
		bson.E{Key: "classroomid", Value: modelObj.ClassRoomId}}
	count, err := impl.collection.CountDocuments(impl.ctx, filter)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.Error("Result already exist.")
	}
	_, er := impl.collection.InsertOne(impl.ctx, modelObj)
	if er != nil {
		return nil, errors.Error("Error in creating Result.")
	}
	log.Print("Call to create Result completed.")
	return modelObj, er
}

func (impl serviceImpl) PutResult(id string, User dtos.UpdateResultRequest) (interface{}, error) {

	log.Print("PutResult started")

	objId := conversion.GetMongoId(id)
	var updatedResult dtos.UpdateResultRequest
	conversion.Convert(User, &updatedResult)
	filter := bson.D{bson.E{Key: "_id", Value: objId}}
	var modelObj models.Result

	update := bson.D{bson.E{Key: "createdat", Value: modelObj.CreatedAt},
		bson.E{Key: "studentid", Value: modelObj.StudentId},
		bson.E{Key: "subjectid", Value: modelObj.SubjectId},
		bson.E{Key: "teacherid", Value: modelObj.TeacherId},
		bson.E{Key: "classroomid", Value: modelObj.ClassRoomId},
		bson.E{Key: "score", Value: modelObj.Score},
		bson.E{Key: "scoremax", Value: modelObj.ScoreMax},
		bson.E{Key: "createdyear", Value: modelObj.CreatedYear},
		bson.E{Key: "createdmonth", Value: modelObj.CreatedMonth},
		bson.E{Key: "createdday", Value: modelObj.CreatedDay}}
	_, err := impl.collection.UpdateOne(impl.ctx, filter, bson.D{bson.E{Key: "$set", Value: update}})

	if err != nil {
		return modelObj, errors.Error("Could not upadte Result")
	}

	log.Print("PutResult completed")
	return modelObj, nil
}
