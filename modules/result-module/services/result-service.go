package services

import (
	"context"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/GabrielAchumba/school-mgt-backend/common/config"
	"github.com/GabrielAchumba/school-mgt-backend/common/conversion"
	errors "github.com/GabrielAchumba/school-mgt-backend/common/errors"
	assessmentServicePackage "github.com/GabrielAchumba/school-mgt-backend/modules/assessment-module/services"
	classRoomServicePackage "github.com/GabrielAchumba/school-mgt-backend/modules/classroom-module/services"
	"github.com/GabrielAchumba/school-mgt-backend/modules/result-module/dtos"
	"github.com/GabrielAchumba/school-mgt-backend/modules/result-module/models"
	"github.com/GabrielAchumba/school-mgt-backend/modules/result-module/utils"
	staffServicePackage "github.com/GabrielAchumba/school-mgt-backend/modules/staff-module/services"
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
	ComputeSummaryResults(req dtos.GetResultsRequest) (interface{}, error)
	ComputeStudentsSummaryResults(req dtos.GetResultsRequest) (interface{}, error)
	SummaryStudentsPositions(req dtos.GetResultsRequest) (interface{}, error)
}

type serviceImpl struct {
	ctx               context.Context
	collection        *mongo.Collection
	userService       userServicePackage.UserService
	studentService    studentServicePackage.StudentService
	subjectService    subjectServicePackage.SubjectService
	classRoomService  classRoomServicePackage.ClassRoomService
	assessmentService assessmentServicePackage.AssessmentService
	staffService      staffServicePackage.StaffService
	resultUtils       utils.ResultUtils
}

func New(mongoClient *mongo.Client, config config.Settings, ctx context.Context,
	userService userServicePackage.UserService,
	studentService studentServicePackage.StudentService,
	subjectService subjectServicePackage.SubjectService,
	classRoomService classRoomServicePackage.ClassRoomService,
	assessmentService assessmentServicePackage.AssessmentService,
	staffService staffServicePackage.StaffService) ResultService {
	collection := mongoClient.Database(config.Database.DatabaseName).Collection(config.TableNames.Result)

	return &serviceImpl{
		collection:        collection,
		ctx:               ctx,
		userService:       userService,
		studentService:    studentService,
		subjectService:    subjectService,
		classRoomService:  classRoomService,
		assessmentService: assessmentService,
		staffService:      staffService,
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
	assessment, _ := impl.assessmentService.GetAssessment(Result.AssessmentId)
	designation, _ := impl.staffService.GetStaff(Result.DesignationId)

	Result.StudentFullName = student.FirstName + " " + student.LastName
	Result.SubjectFullName = subject.Type
	Result.TeacherFullName = teacher.FirstName + " " + teacher.LastName
	Result.ClassRoomFullName = _classRoom.Type
	Result.AssessmentFullName = assessment.Type
	Result.DesignationFullName = designation.Type

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
		designation, _ := impl.staffService.GetStaff(Results[i].DesignationId)
		student, _ := impl.studentService.GetStudent(Results[i].StudentId)
		subject, _ := impl.subjectService.GetSubject(Results[i].SubjectId)
		teacher, _ := impl.userService.GetUser(Results[i].TeacherId)
		classRoom, _ := impl.classRoomService.GetClassRoom(Results[i].ClassRoomId)
		assessment, _ := impl.assessmentService.GetAssessment(Results[i].AssessmentId)

		Results[i].StudentFullName = student.FirstName + " " + student.LastName
		Results[i].SubjectFullName = subject.Type
		Results[i].TeacherFullName = teacher.FirstName + " " + teacher.LastName
		Results[i].ClassRoomFullName = classRoom.Type
		Results[i].AssessmentFullName = assessment.Type
		Results[i].DesignationFullName = designation.Type

	}

	log.Print("Call to get all types of Result completed.")
	return Results, err
}

func (impl serviceImpl) ComputeSummaryResults(req dtos.GetResultsRequest) (interface{}, error) {

	log.Print("Call ComputeSummaryResults started")

	assessments, _ := impl.assessmentService.GetAssessments()
	subjects, _ := impl.subjectService.GetSubjects()

	splitStartDte := strings.Split(req.StartDate, "/")
	startDay, _ := strconv.Atoi(splitStartDte[2])
	startMonth, _ := strconv.Atoi(splitStartDte[1])
	startYear, _ := strconv.Atoi(splitStartDte[0])

	splitEndDate := strings.Split(req.EndDate, "/")
	endDay, _ := strconv.Atoi(splitEndDate[2])
	endMonth, _ := strconv.Atoi(splitEndDate[1])
	endYear, _ := strconv.Atoi(splitEndDate[0])

	startDate := time.Date(startYear, time.Month(startMonth), startDay, 1, 10, 30, 0, time.UTC)
	endDate := time.Date(endYear, time.Month(endMonth), endDay, 1, 10, 30, 0, time.UTC)

	var Results []dtos.ResultResponse
	filter := bson.D{bson.E{Key: "createdat", Value: bson.D{bson.E{Key: "$gte", Value: startDate}}},
		bson.E{Key: "createdat", Value: bson.D{bson.E{Key: "$lte", Value: endDate}}},
		bson.E{Key: "subjectid", Value: bson.D{bson.E{Key: "$in", Value: req.SubjectIds}}},
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

	subjectsScores := make(map[string]dtos.SubJectResult, 0)
	for _, subject := range subjects {
		var subjectScore float64 = 0
		isSubject := false
		subjectAssessments := make(map[string]float64, 0)
		for _, assessment := range assessments {
			var assessmentScore float64 = 0
			var assessmentCounter float64 = 0
			for _, Result := range Results {
				if Result.AssessmentId == assessment.Id &&
					Result.SubjectId == subject.Id {
					if Result.ScoreMax > 0 {
						assessmentScore = assessmentScore + (Result.Score / Result.ScoreMax)
						assessmentCounter++
					}
				}
			}

			if assessmentCounter > 0 {
				var totalAssementScore float64 = (assessmentScore / assessmentCounter) * assessment.Percentage
				subjectAssessments[assessment.Type] = totalAssementScore
				subjectScore = subjectScore + totalAssementScore
				isSubject = true
			}
		}

		if isSubject {
			subjectsScores[subject.Type] = dtos.SubJectResult{
				SubjectScore: subjectScore,
				Assessments:  subjectAssessments,
			}
		}
	}

	log.Print("Call ComputeSummaryResults completed")
	return subjectsScores, err

}

func (impl serviceImpl) SummaryStudentsPositions(req dtos.GetResultsRequest) (interface{}, error) {

	log.Print("Call SummaryStudentsPositions started")

	assessments, _ := impl.assessmentService.GetAssessments()
	subjects, _ := impl.subjectService.GetSubjects()

	splitStartDte := strings.Split(req.StartDate, "/")
	startDay, _ := strconv.Atoi(splitStartDte[2])
	startMonth, _ := strconv.Atoi(splitStartDte[1])
	startYear, _ := strconv.Atoi(splitStartDte[0])

	splitEndDate := strings.Split(req.EndDate, "/")
	endDay, _ := strconv.Atoi(splitEndDate[2])
	endMonth, _ := strconv.Atoi(splitEndDate[1])
	endYear, _ := strconv.Atoi(splitEndDate[0])

	startDate := time.Date(startYear, time.Month(startMonth), startDay, 1, 10, 30, 0, time.UTC)
	endDate := time.Date(endYear, time.Month(endMonth), endDay, 1, 10, 30, 0, time.UTC)

	var Results []dtos.ResultResponse
	filter := bson.D{bson.E{Key: "createdat", Value: bson.D{bson.E{Key: "$gte", Value: startDate}}},
		bson.E{Key: "createdat", Value: bson.D{bson.E{Key: "$lte", Value: endDate}}},
		bson.E{Key: "subjectid", Value: bson.D{bson.E{Key: "$in", Value: req.SubjectIds}}},
		bson.E{Key: "teacherid", Value: bson.D{bson.E{Key: "$in", Value: req.TeacherIds}}},
		bson.E{Key: "studentid", Value: bson.D{bson.E{Key: "$in", Value: req.StudentIds}}},
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

	grouppedStudents := impl.resultUtils.GroupByStudents(Results)
	students := make([]dtos.StudentResults, 0)

	for studentId, grouppedStudent := range grouppedStudents {
		subjectsScores := make(map[string]dtos.SubJectResult, 0)
		for _, subject := range subjects {
			var subjectScore float64 = 0
			isSubject := false
			subjectAssessments := make(map[string]float64, 0)
			for _, assessment := range assessments {
				var assessmentScore float64 = 0
				var assessmentCounter float64 = 0
				for _, Result := range grouppedStudent {
					if Result.AssessmentId == assessment.Id &&
						Result.SubjectId == subject.Id {
						if Result.ScoreMax > 0 {
							assessmentScore = assessmentScore + (Result.Score / Result.ScoreMax)
							assessmentCounter++
						}
					}
				}

				if assessmentCounter > 0 {
					var totalAssementScore float64 = (assessmentScore / assessmentCounter) * assessment.Percentage
					subjectAssessments[assessment.Type] = totalAssementScore
					subjectScore = subjectScore + totalAssementScore
					isSubject = true
				}
			}

			if isSubject {
				subjectsScores[subject.Type] = dtos.SubJectResult{
					SubjectScore: subjectScore,
					Assessments:  subjectAssessments,
				}
			}
		}

		student := dtos.StudentResults{
			StudentId: studentId,
			Subjects:  subjectsScores,
		}

		students = append(students, student)
	}

	log.Print("Call SummaryStudentsPositions completed")
	return students, err

}

func (impl serviceImpl) ComputeStudentsSummaryResults(req dtos.GetResultsRequest) (interface{}, error) {

	log.Print("Call ComputeStudentsSummaryResults started")

	assessments, _ := impl.assessmentService.GetAssessments()

	splitStartDte := strings.Split(req.StartDate, "/")
	startDay, _ := strconv.Atoi(splitStartDte[2])
	startMonth, _ := strconv.Atoi(splitStartDte[1])
	startYear, _ := strconv.Atoi(splitStartDte[0])

	splitEndDate := strings.Split(req.EndDate, "/")
	endDay, _ := strconv.Atoi(splitEndDate[2])
	endMonth, _ := strconv.Atoi(splitEndDate[1])
	endYear, _ := strconv.Atoi(splitEndDate[0])

	startDate := time.Date(startYear, time.Month(startMonth), startDay, 1, 10, 30, 0, time.UTC)
	endDate := time.Date(endYear, time.Month(endMonth), endDay, 1, 10, 30, 0, time.UTC)

	var Results []dtos.ResultResponse
	filter := bson.D{bson.E{Key: "createdat", Value: bson.D{bson.E{Key: "$gte", Value: startDate}}},
		bson.E{Key: "createdat", Value: bson.D{bson.E{Key: "$lte", Value: endDate}}},
		bson.E{Key: "teacherid", Value: bson.D{bson.E{Key: "$in", Value: req.TeacherIds}}},
		bson.E{Key: "subjectid", Value: req.StudentId},
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

	grouppedStudents := impl.resultUtils.GroupByStudents(Results)
	subjectScores := make([]float64, 0)

	for _, StudentResults := range grouppedStudents {
		var subjectScore float64 = 0
		isSubject := false
		for _, assessment := range assessments {
			var assessmentScore float64 = 0
			var assessmentCounter float64 = 0
			for _, Result := range StudentResults {
				if Result.AssessmentId == assessment.Id {
					if Result.ScoreMax > 0 {
						assessmentScore = assessmentScore + (Result.Score / Result.ScoreMax)
						assessmentCounter++
					}
				}
			}

			if assessmentCounter > 0 {
				var totalAssementScore float64 = (assessmentScore / assessmentCounter) * assessment.Percentage
				subjectScore = subjectScore + totalAssementScore
				isSubject = true
			}
		}

		if isSubject {
			subjectScores = append(subjectScores, subjectScore)
		}
	}

	RangeOfScores := make([]dtos.RangeOfScore, 0)
	copy(RangeOfScores, req.RangeOfScores)
	counter := -1
	counter2 := 0
	for _, RangeOfScore := range RangeOfScores {
		counter++
		for _, subjectScore := range subjectScores {
			if subjectScore >= RangeOfScore.From &&
				subjectScore <= RangeOfScore.To {
				counter2++
			}
		}

		RangeOfScores[counter].NumberOfStudents = counter2
	}

	log.Print("Call ComputeStudentsSummaryResults completed")
	return RangeOfScores, err

}

func (impl serviceImpl) CreateResult(userId string, model dtos.CreateResultRequest) (interface{}, error) {

	log.Print("Call to create Result started.")

	var modelObj models.Result
	conversion.Convert(model, &modelObj)

	modelObj.CreatedBy = userId

	splitCreatedAt := strings.Split(model.CreatedAt, "/")
	modelObj.CreatedDay, _ = strconv.Atoi(splitCreatedAt[2])
	modelObj.CreatedMonth, _ = strconv.Atoi(splitCreatedAt[1])
	modelObj.CreatedYear, _ = strconv.Atoi(splitCreatedAt[0])

	modelObj.CreatedAt = time.Date(modelObj.CreatedYear,
		time.Month(modelObj.CreatedMonth), modelObj.CreatedDay, 1, 10, 30, 0, time.UTC)

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

func (impl serviceImpl) PutResult(id string, model dtos.UpdateResultRequest) (interface{}, error) {

	log.Print("PutResult started")

	objId := conversion.GetMongoId(id)
	var updatedResult dtos.UpdateResultRequest
	conversion.Convert(model, &updatedResult)
	filter := bson.D{bson.E{Key: "_id", Value: objId}}
	var modelObj models.Result

	splitCreatedAt := strings.Split(model.UpdatedAt, "/")
	modelObj.CreatedDay, _ = strconv.Atoi(splitCreatedAt[2])
	modelObj.CreatedMonth, _ = strconv.Atoi(splitCreatedAt[1])
	modelObj.CreatedYear, _ = strconv.Atoi(splitCreatedAt[0])

	modelObj.CreatedAt = time.Date(modelObj.CreatedYear,
		time.Month(modelObj.CreatedMonth), modelObj.CreatedDay, 1, 10, 30, 0, time.UTC)

	update := bson.D{bson.E{Key: "createdat", Value: modelObj.CreatedAt},
		bson.E{Key: "studentid", Value: updatedResult.StudentId},
		bson.E{Key: "subjectid", Value: updatedResult.SubjectId},
		bson.E{Key: "teacherid", Value: updatedResult.TeacherId},
		bson.E{Key: "classroomid", Value: updatedResult.ClassRoomId},
		bson.E{Key: "assessmentid", Value: updatedResult.AssessmentId},
		bson.E{Key: "designationid", Value: updatedResult.DesignationId},
		bson.E{Key: "score", Value: updatedResult.Score},
		bson.E{Key: "scoremax", Value: updatedResult.ScoreMax},
		bson.E{Key: "createdyear", Value: modelObj.CreatedYear},
		bson.E{Key: "createdmonth", Value: modelObj.CreatedMonth},
		bson.E{Key: "createdday", Value: modelObj.CreatedDay}}
	_, err := impl.collection.UpdateOne(impl.ctx, filter, bson.D{bson.E{Key: "$set", Value: update}})

	if err != nil {
		return updatedResult, errors.Error("Could not upadte Result")
	}

	log.Print("PutResult completed")
	return updatedResult, nil
}
