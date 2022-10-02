package services

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/GabrielAchumba/school-mgt-backend/common/config"
	"github.com/GabrielAchumba/school-mgt-backend/common/conversion"
	errors "github.com/GabrielAchumba/school-mgt-backend/common/errors"
	"github.com/GabrielAchumba/school-mgt-backend/modules/student-module/dtos"
	"github.com/GabrielAchumba/school-mgt-backend/modules/student-module/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type StudentService interface {
	CreateStudent(userId string, requestModel dtos.CreateStudentRequest) (interface{}, error)
	DeleteStudent(id string) (int64, error)
	GetStudent(id string) (dtos.StudentResponse, error)
	GetStudents() ([]dtos.StudentResponse, error)
	PutStudent(id string, item dtos.UpdateStudentRequest) (interface{}, error)
}

type serviceImpl struct {
	ctx        context.Context
	collection *mongo.Collection
}

func New(mongoClient *mongo.Client, config config.Settings, ctx context.Context) StudentService {
	collection := mongoClient.Database(config.Database.DatabaseName).Collection(config.TableNames.Student)

	return &serviceImpl{
		collection: collection,
		ctx:        ctx,
	}
}

func (impl serviceImpl) DeleteStudent(id string) (int64, error) {

	log.Print("Call to delete type of student by id started.")
	objId := conversion.GetMongoId(id)
	filter := bson.D{bson.E{Key: "_id", Value: objId}}

	result, err := impl.collection.DeleteOne(impl.ctx, filter)

	if err != nil {
		return result.DeletedCount, errors.Error("Error in deleting type of student.")
	}

	if result.DeletedCount < 1 {
		return result.DeletedCount, errors.Error("Type of student with specified ID not found!")
	}

	log.Print("Call to delete type of student by id completed.")
	return result.DeletedCount, nil
}

func (impl serviceImpl) GetStudent(id string) (dtos.StudentResponse, error) {

	log.Print("Get Type of Student called")
	objId := conversion.GetMongoId(id)
	var Student dtos.StudentResponse

	filter := bson.D{bson.E{Key: "_id", Value: objId}}

	err := impl.collection.FindOne(impl.ctx, filter).Decode(&Student)
	if err != nil {
		return Student, errors.Error("could not find type of student by id")
	}

	log.Print("Get type of student completed")
	return Student, err

}

func (impl serviceImpl) GetStudents() ([]dtos.StudentResponse, error) {

	log.Print("Call to get all types of student started.")

	var Students []dtos.StudentResponse
	filter := bson.D{}
	cur, err := impl.collection.Find(impl.ctx, filter)

	if err != nil {
		Students = make([]dtos.StudentResponse, 0)
		return Students, errors.Error("Types of student not found!")
	}

	err = cur.All(impl.ctx, &Students)
	if err != nil {
		return Students, err
	}

	cur.Close(impl.ctx)
	if len(Students) == 0 {
		Students = make([]dtos.StudentResponse, 0)
	}

	log.Print("Call to get all types of student completed.")
	return Students, err
}

func (impl serviceImpl) CreateStudent(userId string, model dtos.CreateStudentRequest) (interface{}, error) {

	log.Print("Call to create student started.")

	var modelObj models.Student
	conversion.Convert(model, &modelObj)

	modelObj.CreatedBy = userId
	modelObj.CreatedAt = time.Now()

	if modelObj.FirstName == "" {
		return nil, errors.Error("FirstName of Student cannot be empty.")
	}
	if modelObj.LastName == "" {
		return nil, errors.Error("LastName of Student cannot be empty.")
	}

	filter := bson.D{bson.E{Key: "firstname", Value: modelObj.FirstName},
		bson.E{Key: "lastname", Value: modelObj.LastName},
		bson.E{Key: "birthday", Value: modelObj.BirthDay},
		bson.E{Key: "birthmonth", Value: modelObj.BirthMonth},
		bson.E{Key: "birthyear", Value: modelObj.BirthYear}}
	count, err := impl.collection.CountDocuments(impl.ctx, filter)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.Error(fmt.Sprintf("Student '%v'already exist.",
			strings.Join([]string{model.FirstName, model.LastName}, " ")))
	}
	_, er := impl.collection.InsertOne(impl.ctx, modelObj)
	if er != nil {
		return nil, errors.Error("Error in creating student.")
	}
	log.Print("Call to create student completed.")
	return modelObj, er
}

func (impl serviceImpl) PutStudent(id string, User dtos.UpdateStudentRequest) (interface{}, error) {

	log.Print("PutStudent started")

	objId := conversion.GetMongoId(id)
	var updatedStudent dtos.UpdateStudentRequest
	conversion.Convert(User, &updatedStudent)
	filter := bson.D{bson.E{Key: "_id", Value: objId}}
	var modelObj models.Student

	update := bson.D{bson.E{Key: "firstname", Value: updatedStudent.FirstName},
		bson.E{Key: "lastname", Value: updatedStudent.LastName},
		bson.E{Key: "birthday", Value: updatedStudent.BirthDay},
		bson.E{Key: "birthmonth", Value: updatedStudent.BirthMonth},
		bson.E{Key: "birthyear", Value: updatedStudent.BirthYear}}
	_, err := impl.collection.UpdateOne(impl.ctx, filter, bson.D{bson.E{Key: "$set", Value: update}})

	if err != nil {
		return modelObj, errors.Error("Could not upadte student")
	}

	log.Print("PutStudent completed")
	return modelObj, nil
}
