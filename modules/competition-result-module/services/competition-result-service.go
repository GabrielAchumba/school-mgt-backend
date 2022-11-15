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
	"github.com/GabrielAchumba/school-mgt-backend/modules/competition-result-module/dtos"
	"github.com/GabrielAchumba/school-mgt-backend/modules/competition-result-module/models"
	"github.com/GabrielAchumba/school-mgt-backend/modules/competition-result-module/utils"
	gradeDTOPackage "github.com/GabrielAchumba/school-mgt-backend/modules/grade-module/dtos"
	gradeServicePackage "github.com/GabrielAchumba/school-mgt-backend/modules/grade-module/services"
	levelServicePackage "github.com/GabrielAchumba/school-mgt-backend/modules/level-module/services"
	sessionServicePackage "github.com/GabrielAchumba/school-mgt-backend/modules/session-module/services"
	staffServicePackage "github.com/GabrielAchumba/school-mgt-backend/modules/staff-module/services"
	studentServicePackage "github.com/GabrielAchumba/school-mgt-backend/modules/student-module/services"
	subjectServicePackage "github.com/GabrielAchumba/school-mgt-backend/modules/subject-module/services"
	userServicePackage "github.com/GabrielAchumba/school-mgt-backend/modules/user-module/services"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CompetitionResultService interface {
	CreateCompetitionResult(userId string, requestModel dtos.CreateCompetitionResultRequest) (interface{}, error)
	CreateCompetitionResults(userId string, _models []dtos.CreateCompetitionResultRequest) (interface{}, error)
	DeleteCompetitionResult(id string, schoolId string) (int64, error)
	GetCompetitionResult(id string, schoolId string) (dtos.CompetitionResultResponse, error)
	GetCompetitionResults(schoolId string) ([]dtos.CompetitionResultResponse, error)
	PutCompetitionResult(id string, item dtos.UpdateCompetitionResultRequest) (interface{}, error)
	ComputeSummaryCompetitionResults(req dtos.GetCompetitionResultsRequest) (interface{}, error)
	ComputeSummaryCompetitionResults2(req dtos.GetCompetitionResultsRequest) (interface{}, error)
	ComputeStudentsSummaryCompetitionResults(req dtos.GetCompetitionResultsRequest) (interface{}, error)
	SummaryStudentsPositions(req dtos.GetCompetitionResultsRequest) (interface{}, error)
	SummaryStudentsPositions2(req dtos.GetCompetitionResultsRequest) (interface{}, error)
	ComputeStudentsCompetitionResultsByDateRange(req dtos.GetCompetitionResultsRequest) (interface{}, error)
}

type serviceImpl struct {
	ctx                    context.Context
	collection             *mongo.Collection
	userService            userServicePackage.UserService
	studentService         studentServicePackage.StudentService
	subjectService         subjectServicePackage.SubjectService
	classRoomService       classRoomServicePackage.ClassRoomService
	assessmentService      assessmentServicePackage.AssessmentService
	staffService           staffServicePackage.StaffService
	sessionService         sessionServicePackage.SessionService
	gradeService           gradeServicePackage.GradeService
	levelService           levelServicePackage.LevelService
	CompetitionResultUtils utils.CompetitionResultUtils
}

func New(mongoClient *mongo.Client, config config.Settings, ctx context.Context,
	userService userServicePackage.UserService,
	studentService studentServicePackage.StudentService,
	subjectService subjectServicePackage.SubjectService,
	classRoomService classRoomServicePackage.ClassRoomService,
	assessmentService assessmentServicePackage.AssessmentService,
	staffService staffServicePackage.StaffService,
	sessionService sessionServicePackage.SessionService,
	gradeService gradeServicePackage.GradeService,
	levelService levelServicePackage.LevelService) CompetitionResultService {
	collection := mongoClient.Database(config.Database.DatabaseName).Collection(config.TableNames.CompetitionResult)

	return &serviceImpl{
		collection:             collection,
		ctx:                    ctx,
		userService:            userService,
		studentService:         studentService,
		subjectService:         subjectService,
		classRoomService:       classRoomService,
		assessmentService:      assessmentService,
		staffService:           staffService,
		sessionService:         sessionService,
		gradeService:           gradeService,
		levelService:           levelService,
		CompetitionResultUtils: utils.CompetitionResultUtilsImpl{},
	}
}

func (impl serviceImpl) DeleteCompetitionResult(id string, schoolId string) (int64, error) {

	log.Print("Call to delete type of CompetitionResult by id started.")
	objId := conversion.GetMongoId(id)
	filter := bson.D{bson.E{Key: "_id", Value: objId},
		bson.E{Key: "schoolid", Value: schoolId}}

	CompetitionResult, err := impl.collection.DeleteOne(impl.ctx, filter)

	if err != nil {
		return CompetitionResult.DeletedCount, errors.Error("Error in deleting type of CompetitionResult.")
	}

	if CompetitionResult.DeletedCount < 1 {
		return CompetitionResult.DeletedCount, errors.Error("Type of CompetitionResult with specified ID not found!")
	}

	log.Print("Call to delete type of CompetitionResult by id completed.")
	return CompetitionResult.DeletedCount, nil
}

func (impl serviceImpl) GetCompetitionResult(id string, schoolId string) (dtos.CompetitionResultResponse, error) {

	log.Print("Get Type of CompetitionResult called")
	objId := conversion.GetMongoId(id)
	var CompetitionResult dtos.CompetitionResultResponse

	filter := bson.D{bson.E{Key: "_id", Value: objId},
		bson.E{Key: "schoolid", Value: schoolId}}

	err := impl.collection.FindOne(impl.ctx, filter).Decode(&CompetitionResult)
	if err != nil {
		return CompetitionResult, errors.Error("could not find type of CompetitionResult by id")
	}

	student, _ := impl.studentService.GetStudent(CompetitionResult.StudentId, CompetitionResult.SchoolId)
	subject, _ := impl.subjectService.GetSubject(CompetitionResult.SubjectId, CompetitionResult.SchoolId)
	teacher, _ := impl.userService.GetUser(CompetitionResult.TeacherId, CompetitionResult.SchoolId)
	_classRoom, _ := impl.classRoomService.GetClassRoom(CompetitionResult.ClassRoomId, CompetitionResult.SchoolId)
	assessment, _ := impl.assessmentService.GetAssessment(CompetitionResult.AssessmentId, CompetitionResult.SchoolId)
	designation, _ := impl.staffService.GetStaff(CompetitionResult.DesignationId, CompetitionResult.SchoolId)

	CompetitionResult.StudentFullName = student.FirstName + " " + student.LastName
	CompetitionResult.SubjectFullName = subject.Type
	CompetitionResult.TeacherFullName = teacher.FirstName + " " + teacher.LastName
	CompetitionResult.ClassRoomFullName = _classRoom.Type
	CompetitionResult.AssessmentFullName = assessment.Type
	CompetitionResult.DesignationFullName = designation.Type

	log.Print("Get type of CompetitionResult completed")
	return CompetitionResult, err

}

func (impl serviceImpl) GetCompetitionResults(schoolId string) ([]dtos.CompetitionResultResponse, error) {

	log.Print("Call to get all types of CompetitionResult started.")

	var CompetitionResults []dtos.CompetitionResultResponse
	filter := bson.D{bson.E{Key: "schoolid", Value: schoolId}}
	cur, err := impl.collection.Find(impl.ctx, filter)

	if err != nil {
		CompetitionResults = make([]dtos.CompetitionResultResponse, 0)
		return CompetitionResults, errors.Error("Types of CompetitionResult not found!")
	}

	err = cur.All(impl.ctx, &CompetitionResults)
	if err != nil {
		return CompetitionResults, err
	}

	cur.Close(impl.ctx)
	length := len(CompetitionResults)
	if length == 0 {
		CompetitionResults = make([]dtos.CompetitionResultResponse, 0)
	}

	for i := 0; i < length; i++ {
		designation, _ := impl.staffService.GetStaff(CompetitionResults[i].DesignationId, CompetitionResults[i].SchoolId)
		student, _ := impl.studentService.GetStudent(CompetitionResults[i].StudentId, CompetitionResults[i].SchoolId)
		subject, _ := impl.subjectService.GetSubject(CompetitionResults[i].SubjectId, CompetitionResults[i].SchoolId)
		teacher, _ := impl.userService.GetUser(CompetitionResults[i].TeacherId, CompetitionResults[i].SchoolId)
		classRoom, _ := impl.classRoomService.GetClassRoom(CompetitionResults[i].ClassRoomId, CompetitionResults[i].SchoolId)
		assessment, _ := impl.assessmentService.GetAssessment(CompetitionResults[i].AssessmentId, CompetitionResults[i].SchoolId)

		CompetitionResults[i].StudentFullName = student.FirstName + " " + student.LastName
		CompetitionResults[i].SubjectFullName = subject.Type
		CompetitionResults[i].TeacherFullName = teacher.FirstName + " " + teacher.LastName
		CompetitionResults[i].ClassRoomFullName = classRoom.Type
		CompetitionResults[i].AssessmentFullName = assessment.Type
		CompetitionResults[i].DesignationFullName = designation.Type

	}

	log.Print("Call to get all types of CompetitionResult completed.")
	return CompetitionResults, err
}

func (impl serviceImpl) ComputeSummaryCompetitionResults(req dtos.GetCompetitionResultsRequest) (interface{}, error) {

	log.Print("Call ComputeSummaryCompetitionResults started")

	assessments, _ := impl.assessmentService.GetAssessments(req.SchoolId)
	subjects, _ := impl.subjectService.GetSubjects(req.SchoolId)
	grades, _ := impl.gradeService.GetGrades(req.SchoolId)

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

	var CompetitionResults []dtos.CompetitionResultResponse
	filter := bson.D{bson.E{Key: "createdat", Value: bson.D{bson.E{Key: "$gte", Value: startDate}}},
		bson.E{Key: "createdat", Value: bson.D{bson.E{Key: "$lte", Value: endDate}}},
		bson.E{Key: "subjectid", Value: bson.D{bson.E{Key: "$in", Value: req.SubjectIds}}},
		bson.E{Key: "teacherid", Value: req.TeacherId},
		bson.E{Key: "studentid", Value: req.StudentId},
		bson.E{Key: "classroomid", Value: req.ClassRoomId},
		bson.E{Key: "schoolid", Value: req.SchoolId}}

	cur, err := impl.collection.Find(impl.ctx, filter)

	if err != nil {
		CompetitionResults = make([]dtos.CompetitionResultResponse, 0)
		return CompetitionResults, errors.Error("CompetitionResults not found!")
	}

	err = cur.All(impl.ctx, &CompetitionResults)
	if err != nil {
		return CompetitionResults, err
	}

	cur.Close(impl.ctx)
	length := len(CompetitionResults)
	if length == 0 {
		CompetitionResults = make([]dtos.CompetitionResultResponse, 0)
	}

	subjectsScores := make(map[string]dtos.SubJectCompetitionResult, 0)
	for _, subject := range subjects {
		var subjectScore float64 = 0
		isSubject := false
		subjectAssessments := make(map[string]dtos.AssesmentGroup, 0)
		for _, assessment := range assessments {

			if subject.Id == assessment.SubjectId {

				var assessmentScore float64 = 0
				var assessmentCounter float64 = 0
				for _, CompetitionResult := range CompetitionResults {
					if CompetitionResult.AssessmentId == assessment.Id &&
						CompetitionResult.SubjectId == subject.Id {
						if CompetitionResult.ScoreMax > 0 {
							assessmentScore = assessmentScore + (CompetitionResult.Score / CompetitionResult.ScoreMax)
							assessmentCounter++
						}
					}
				}

				if assessmentCounter > 0 {
					var totalAssementScore float64 = (assessmentScore / assessmentCounter) * assessment.Percentage
					subjectAssessments[assessment.Type] = dtos.AssesmentGroup{
						AssessmentScore: totalAssementScore,
						ScoreMax:        assessment.Percentage,
					}
					subjectScore = subjectScore + totalAssementScore
					isSubject = true
				}

			}
		}

		if isSubject {
			grade, point := gradeDTOPackage.GetGradeAndPoint(grades, subjectScore)
			subjectsScores[subject.Type] = dtos.SubJectCompetitionResult{
				SubjectScore: subjectScore,
				Assessments:  subjectAssessments,
				Grade:        grade,
				Point:        point,
			}
		}
	}

	log.Print("Call ComputeSummaryCompetitionResults completed")
	return subjectsScores, err

}

func (impl serviceImpl) ComputeSummaryCompetitionResults2(req dtos.GetCompetitionResultsRequest) (interface{}, error) {

	log.Print("Call ComputeSummaryCompetitionResults2 started")

	assessments, _ := impl.assessmentService.GetAssessments(req.SchoolId)
	subjects, _ := impl.subjectService.GetSubjects(req.SchoolId)
	grades, _ := impl.gradeService.GetGrades(req.SchoolId)

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

	var CompetitionResults []dtos.CompetitionResultResponse
	filter := bson.D{bson.E{Key: "createdat", Value: bson.D{bson.E{Key: "$gte", Value: startDate}}},
		bson.E{Key: "createdat", Value: bson.D{bson.E{Key: "$lte", Value: endDate}}},
		bson.E{Key: "subjectid", Value: bson.D{bson.E{Key: "$in", Value: req.SubjectIds}}},
		bson.E{Key: "levelid", Value: req.LevelId},
		bson.E{Key: "studentid", Value: req.StudentId},
		bson.E{Key: "classroomid", Value: req.ClassRoomId},
		bson.E{Key: "schoolid", Value: req.SchoolId}}

	cur, err := impl.collection.Find(impl.ctx, filter)

	if err != nil {
		CompetitionResults = make([]dtos.CompetitionResultResponse, 0)
		return CompetitionResults, errors.Error("CompetitionResults not found!")
	}

	err = cur.All(impl.ctx, &CompetitionResults)
	if err != nil {
		return CompetitionResults, err
	}

	cur.Close(impl.ctx)
	length := len(CompetitionResults)
	if length == 0 {
		CompetitionResults = make([]dtos.CompetitionResultResponse, 0)
	}

	subjectsScores := make(map[string]dtos.SubJectCompetitionResult, 0)
	for _, subject := range subjects {
		var subjectScore float64 = 0
		isSubject := false
		subjectAssessments := make(map[string]dtos.AssesmentGroup, 0)
		for _, assessment := range assessments {

			if subject.Id == assessment.SubjectId {

				var assessmentScore float64 = 0
				var assessmentCounter float64 = 0
				for _, CompetitionResult := range CompetitionResults {
					if CompetitionResult.AssessmentId == assessment.Id &&
						CompetitionResult.SubjectId == subject.Id {
						if CompetitionResult.ScoreMax > 0 {
							assessmentScore = assessmentScore + (CompetitionResult.Score / CompetitionResult.ScoreMax)
							assessmentCounter++
						}
					}
				}

				if assessmentCounter > 0 {
					var totalAssementScore float64 = (assessmentScore / assessmentCounter) * assessment.Percentage
					subjectAssessments[assessment.Type] = dtos.AssesmentGroup{
						AssessmentScore: totalAssementScore,
						ScoreMax:        assessment.Percentage,
					}
					subjectScore = subjectScore + totalAssementScore
					isSubject = true
				}

			}
		}

		if isSubject {
			grade, point := gradeDTOPackage.GetGradeAndPoint(grades, subjectScore)
			subjectsScores[subject.Type] = dtos.SubJectCompetitionResult{
				SubjectScore: subjectScore,
				Assessments:  subjectAssessments,
				Grade:        grade,
				Point:        point,
			}
		}
	}

	log.Print("Call ComputeSummaryCompetitionResults2 completed")
	return subjectsScores, err

}

func (impl serviceImpl) SummaryStudentsPositions(req dtos.GetCompetitionResultsRequest) (interface{}, error) {

	log.Print("Call SummaryStudentsPositions started")

	assessments, _ := impl.assessmentService.GetAssessments(req.SchoolId)
	subjects, _ := impl.subjectService.GetSubjects(req.SchoolId)
	grades, _ := impl.gradeService.GetGrades(req.SchoolId)

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

	var CompetitionResults []dtos.CompetitionResultResponse
	filter := bson.D{bson.E{Key: "createdat", Value: bson.D{bson.E{Key: "$gte", Value: startDate}}},
		bson.E{Key: "createdat", Value: bson.D{bson.E{Key: "$lte", Value: endDate}}},
		bson.E{Key: "subjectid", Value: bson.D{bson.E{Key: "$in", Value: req.SubjectIds}}},
		bson.E{Key: "teacherid", Value: bson.D{bson.E{Key: "$in", Value: req.TeacherIds}}},
		bson.E{Key: "studentid", Value: bson.D{bson.E{Key: "$in", Value: req.StudentIds}}},
		bson.E{Key: "classroomid", Value: req.ClassRoomId},
		bson.E{Key: "schoolid", Value: req.SchoolId}}

	cur, err := impl.collection.Find(impl.ctx, filter)

	if err != nil {
		CompetitionResults = make([]dtos.CompetitionResultResponse, 0)
		return CompetitionResults, errors.Error("CompetitionResults not found!")
	}

	err = cur.All(impl.ctx, &CompetitionResults)
	if err != nil {
		return CompetitionResults, err
	}

	cur.Close(impl.ctx)
	length := len(CompetitionResults)
	if length == 0 {
		CompetitionResults = make([]dtos.CompetitionResultResponse, 0)
	}

	grouppedStudents := impl.CompetitionResultUtils.GroupByStudents(CompetitionResults)
	studentIds := make([]string, 0)
	for studentId := range grouppedStudents {
		studentIds = append(studentIds, studentId)
	}
	selectedStudents, _ := impl.studentService.GetSelecedStudents(studentIds)
	students := make([]dtos.StudentCompetitionResults, 0)

	i := -1
	for studentId, grouppedStudent := range grouppedStudents {
		i++
		var overallScore float64 = 0
		var overallScoreMax float64 = 0
		subjectsScores := make(map[string]dtos.SubJectCompetitionResult, 0)
		for _, subject := range subjects {
			var subjectScore float64 = 0
			isSubject := false
			subjectAssessments := make(map[string]dtos.AssesmentGroup, 0)
			for _, assessment := range assessments {

				if subject.Id == assessment.SubjectId {

					var assessmentScore float64 = 0
					var assessmentCounter float64 = 0
					for _, CompetitionResult := range grouppedStudent {
						if CompetitionResult.AssessmentId == assessment.Id &&
							CompetitionResult.SubjectId == subject.Id {
							if CompetitionResult.ScoreMax > 0 {
								assessmentScore = assessmentScore + (CompetitionResult.Score / CompetitionResult.ScoreMax)
								assessmentCounter++
							}
						}
					}

					if assessmentCounter > 0 {
						var totalAssementScore float64 = (assessmentScore / assessmentCounter) * assessment.Percentage
						subjectAssessments[assessment.Type] = dtos.AssesmentGroup{
							AssessmentScore: totalAssementScore,
							ScoreMax:        assessment.Percentage,
						}
						subjectScore = subjectScore + totalAssementScore
						isSubject = true
					}
				}

			}

			if isSubject {
				grade, point := gradeDTOPackage.GetGradeAndPoint(grades, subjectScore)
				overallScore = overallScore + subjectScore
				overallScoreMax = overallScoreMax + 100
				subjectsScores[subject.Type] = dtos.SubJectCompetitionResult{
					SubjectScore: subjectScore,
					Assessments:  subjectAssessments,
					Grade:        grade,
					Point:        point,
				}
			}
		}

		student := dtos.StudentCompetitionResults{
			StudentId:       studentId,
			FullName:        selectedStudents[i].FirstName + " " + selectedStudents[i].LastName,
			OverallScore:    overallScore,
			OverallScoreMax: overallScoreMax,
			Subjects:        subjectsScores,
		}

		students = append(students, student)
	}

	log.Print("Call SummaryStudentsPositions completed")
	return utils.SortPositionCompetitionResults(students), err

}

func (impl serviceImpl) SummaryStudentsPositions2(req dtos.GetCompetitionResultsRequest) (interface{}, error) {

	log.Print("Call SummaryStudentsPositions started")

	assessments, _ := impl.assessmentService.GetAssessments(req.SchoolId)
	subjects, _ := impl.subjectService.GetSubjects(req.SchoolId)
	grades, _ := impl.gradeService.GetGrades(req.SchoolId)

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

	var CompetitionResults []dtos.CompetitionResultResponse
	filter := bson.D{bson.E{Key: "createdat", Value: bson.D{bson.E{Key: "$gte", Value: startDate}}},
		bson.E{Key: "createdat", Value: bson.D{bson.E{Key: "$lte", Value: endDate}}},
		bson.E{Key: "subjectid", Value: bson.D{bson.E{Key: "$in", Value: req.SubjectIds}}},
		bson.E{Key: "classroomid", Value: bson.D{bson.E{Key: "$in", Value: req.ClassRoomIds}}},
		bson.E{Key: "studentid", Value: bson.D{bson.E{Key: "$in", Value: req.StudentIds}}},
		bson.E{Key: "leveid", Value: req.LevelId},
		bson.E{Key: "schoolid", Value: req.SchoolId}}

	cur, err := impl.collection.Find(impl.ctx, filter)

	if err != nil {
		CompetitionResults = make([]dtos.CompetitionResultResponse, 0)
		return CompetitionResults, errors.Error("CompetitionResults not found!")
	}

	err = cur.All(impl.ctx, &CompetitionResults)
	if err != nil {
		return CompetitionResults, err
	}

	cur.Close(impl.ctx)
	length := len(CompetitionResults)
	if length == 0 {
		CompetitionResults = make([]dtos.CompetitionResultResponse, 0)
	}

	grouppedStudents := impl.CompetitionResultUtils.GroupByStudents(CompetitionResults)
	studentIds := make([]string, 0)
	for studentId := range grouppedStudents {
		studentIds = append(studentIds, studentId)
	}
	selectedStudents, _ := impl.studentService.GetSelecedStudents(studentIds)
	students := make([]dtos.StudentCompetitionResults, 0)

	i := -1
	for studentId, grouppedStudent := range grouppedStudents {
		i++
		var overallScore float64 = 0
		var overallScoreMax float64 = 0
		subjectsScores := make(map[string]dtos.SubJectCompetitionResult, 0)
		for _, subject := range subjects {
			var subjectScore float64 = 0
			isSubject := false
			subjectAssessments := make(map[string]dtos.AssesmentGroup, 0)
			for _, assessment := range assessments {

				if subject.Id == assessment.SubjectId {

					var assessmentScore float64 = 0
					var assessmentCounter float64 = 0
					for _, CompetitionResult := range grouppedStudent {
						if CompetitionResult.AssessmentId == assessment.Id &&
							CompetitionResult.SubjectId == subject.Id {
							if CompetitionResult.ScoreMax > 0 {
								assessmentScore = assessmentScore + (CompetitionResult.Score / CompetitionResult.ScoreMax)
								assessmentCounter++
							}
						}
					}

					if assessmentCounter > 0 {
						var totalAssementScore float64 = (assessmentScore / assessmentCounter) * assessment.Percentage
						subjectAssessments[assessment.Type] = dtos.AssesmentGroup{
							AssessmentScore: totalAssementScore,
							ScoreMax:        assessment.Percentage,
						}
						subjectScore = subjectScore + totalAssementScore
						isSubject = true
					}

				}
			}

			if isSubject {
				grade, point := gradeDTOPackage.GetGradeAndPoint(grades, subjectScore)
				overallScore = overallScore + subjectScore
				overallScoreMax = overallScoreMax + 100
				subjectsScores[subject.Type] = dtos.SubJectCompetitionResult{
					SubjectScore: subjectScore,
					Assessments:  subjectAssessments,
					Grade:        grade,
					Point:        point,
				}
			}
		}

		student := dtos.StudentCompetitionResults{
			StudentId:       studentId,
			FullName:        selectedStudents[i].FirstName + " " + selectedStudents[i].LastName,
			OverallScore:    overallScore,
			OverallScoreMax: overallScoreMax,
			Subjects:        subjectsScores,
		}

		students = append(students, student)
	}

	log.Print("Call SummaryStudentsPositions completed")
	return utils.SortPositionCompetitionResults(students), err

}

func (impl serviceImpl) ComputeStudentsSummaryCompetitionResults(req dtos.GetCompetitionResultsRequest) (interface{}, error) {

	log.Print("Call ComputeStudentsSummaryCompetitionResults started")

	assessments, _ := impl.assessmentService.GetAssessments(req.SchoolId)

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

	var CompetitionResults []dtos.CompetitionResultResponse
	filter := bson.D{bson.E{Key: "createdat", Value: bson.D{bson.E{Key: "$gte", Value: startDate}}},
		bson.E{Key: "createdat", Value: bson.D{bson.E{Key: "$lte", Value: endDate}}},
		bson.E{Key: "teacherid", Value: bson.D{bson.E{Key: "$in", Value: req.TeacherIds}}},
		bson.E{Key: "subjectid", Value: req.SubjectId},
		bson.E{Key: "classroomid", Value: req.ClassRoomId},
		bson.E{Key: "schoolid", Value: req.SchoolId}}

	cur, err := impl.collection.Find(impl.ctx, filter)

	if err != nil {
		CompetitionResults = make([]dtos.CompetitionResultResponse, 0)
		return CompetitionResults, errors.Error("CompetitionResults not found!")
	}

	err = cur.All(impl.ctx, &CompetitionResults)
	if err != nil {
		return CompetitionResults, err
	}

	cur.Close(impl.ctx)
	length := len(CompetitionResults)
	if length == 0 {
		CompetitionResults = make([]dtos.CompetitionResultResponse, 0)
	}

	grouppedStudents := impl.CompetitionResultUtils.GroupByStudents(CompetitionResults)
	subjectScores := make([]float64, 0)

	for _, StudentCompetitionResults := range grouppedStudents {
		var subjectScore float64 = 0
		isSubject := false
		for _, assessment := range assessments {
			var assessmentScore float64 = 0
			var assessmentCounter float64 = 0
			for _, CompetitionResult := range StudentCompetitionResults {
				if CompetitionResult.AssessmentId == assessment.Id {
					if CompetitionResult.ScoreMax > 0 {
						assessmentScore = assessmentScore + (CompetitionResult.Score / CompetitionResult.ScoreMax)
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

	log.Print("Call ComputeStudentsSummaryCompetitionResults completed")
	return RangeOfScores, err

}

func (impl serviceImpl) ComputeStudentsCompetitionResultsByDateRange(req dtos.GetCompetitionResultsRequest) (interface{}, error) {

	log.Print("Call ComputeStudentsCompetitionResultsByDateRange started")

	assessments, _ := impl.assessmentService.GetAssessments(req.SchoolId)

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

	var CompetitionResults []dtos.CompetitionResultResponse
	filter := bson.D{bson.E{Key: "createdat", Value: bson.D{bson.E{Key: "$gte", Value: startDate}}},
		bson.E{Key: "createdat", Value: bson.D{bson.E{Key: "$lte", Value: endDate}}},
		bson.E{Key: "teacherid", Value: bson.D{bson.E{Key: "$in", Value: req.TeacherIds}}},
		bson.E{Key: "subjectid", Value: req.StudentId},
		bson.E{Key: "classroomid", Value: req.ClassRoomId},
		bson.E{Key: "schoolid", Value: req.SchoolId}}

	cur, err := impl.collection.Find(impl.ctx, filter)

	if err != nil {
		CompetitionResults = make([]dtos.CompetitionResultResponse, 0)
		return CompetitionResults, errors.Error("CompetitionResults not found!")
	}

	err = cur.All(impl.ctx, &CompetitionResults)
	if err != nil {
		return CompetitionResults, err
	}

	cur.Close(impl.ctx)
	length := len(CompetitionResults)
	if length == 0 {
		CompetitionResults = make([]dtos.CompetitionResultResponse, 0)
	}

	grouppedMonthYears := make(map[string]dtos.MonthYear, 0)
	grouppedMonthYearsModified := make(map[string]dtos.MonthYear, 0)
	if req.IsMonthly {
		grouppedMonthYears = impl.CompetitionResultUtils.GroupByStudentsAndMonthYear(CompetitionResults, req.MonthYears)
	} else {
		grouppedMonthYears = impl.CompetitionResultUtils.GroupByStudentsAndYear(CompetitionResults, req.MonthYears)
	}

	for grouppedMonthYearKey, grouppedMonthYear := range grouppedMonthYears {

		subjectScores := make([]float64, 0)

		for _, StudentCompetitionResults := range grouppedMonthYear.Students {
			var subjectScore float64 = 0
			isSubject := false
			for _, assessment := range assessments {
				var assessmentScore float64 = 0
				var assessmentCounter float64 = 0
				for _, CompetitionResult := range StudentCompetitionResults {
					if CompetitionResult.AssessmentId == assessment.Id {
						if CompetitionResult.ScoreMax > 0 {
							assessmentScore = assessmentScore + (CompetitionResult.Score / CompetitionResult.ScoreMax)
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

		grouppedMonthYear.RangeOfScores = RangeOfScores
		grouppedMonthYearsModified[grouppedMonthYearKey] = grouppedMonthYear
	}

	log.Print("Call ComputeStudentsCompetitionResultsByDateRange completed")
	return grouppedMonthYearsModified, err

}

func (impl serviceImpl) CreateCompetitionResult(userId string, model dtos.CreateCompetitionResultRequest) (interface{}, error) {

	log.Print("Call to create CompetitionResult started.")

	var modelObj models.CompetitionResult
	conversion.Convert(model, &modelObj)

	modelObj.CreatedBy = userId
	modelObj.SchoolId = model.SchoolId

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
		bson.E{Key: "classroomid", Value: modelObj.ClassRoomId},
		bson.E{Key: "schoolid", Value: modelObj.SchoolId}}
	count, err := impl.collection.CountDocuments(impl.ctx, filter)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.Error("CompetitionResult already exist.")
	}
	_, er := impl.collection.InsertOne(impl.ctx, modelObj)
	if er != nil {
		return nil, errors.Error("Error in creating CompetitionResult.")
	}
	log.Print("Call to create CompetitionResult completed.")
	return modelObj, er
}

func (impl serviceImpl) CreateCompetitionResults(userId string, _models []dtos.CreateCompetitionResultRequest) (interface{}, error) {

	log.Print("Call to create CompetitionResults started.")

	modelObjs := make([]interface{}, 0)
	for _, model := range _models {
		var modelObj models.CompetitionResult
		modelObj.CreatedBy = userId
		modelObj.CreatedAt = time.Now()
		conversion.Convert(model, &modelObj)
		modelObjs = append(modelObjs, modelObj)
	}

	_, er := impl.collection.InsertMany(impl.ctx, modelObjs)
	if er != nil {
		return nil, errors.Error("Error in creating CompetitionResults.")
	}
	log.Print("Call to create CompetitionResults completed.")
	return modelObjs, er
}

func (impl serviceImpl) PutCompetitionResult(id string, model dtos.UpdateCompetitionResultRequest) (interface{}, error) {

	log.Print("PutCompetitionResult started")

	objId := conversion.GetMongoId(id)
	var updatedCompetitionResult dtos.UpdateCompetitionResultRequest
	conversion.Convert(model, &updatedCompetitionResult)
	updatedCompetitionResult.SchoolId = model.SchoolId
	filter := bson.D{bson.E{Key: "_id", Value: objId}}
	var modelObj models.CompetitionResult

	splitCreatedAt := strings.Split(model.UpdatedAt, "/")
	modelObj.CreatedDay, _ = strconv.Atoi(splitCreatedAt[2])
	modelObj.CreatedMonth, _ = strconv.Atoi(splitCreatedAt[1])
	modelObj.CreatedYear, _ = strconv.Atoi(splitCreatedAt[0])

	modelObj.CreatedAt = time.Date(modelObj.CreatedYear,
		time.Month(modelObj.CreatedMonth), modelObj.CreatedDay, 1, 10, 30, 0, time.UTC)

	update := bson.D{bson.E{Key: "createdat", Value: modelObj.CreatedAt},
		bson.E{Key: "studentid", Value: updatedCompetitionResult.StudentId},
		bson.E{Key: "subjectid", Value: updatedCompetitionResult.SubjectId},
		bson.E{Key: "teacherid", Value: updatedCompetitionResult.TeacherId},
		bson.E{Key: "classroomid", Value: updatedCompetitionResult.ClassRoomId},
		bson.E{Key: "assessmentid", Value: updatedCompetitionResult.AssessmentId},
		bson.E{Key: "designationid", Value: updatedCompetitionResult.DesignationId},
		bson.E{Key: "score", Value: updatedCompetitionResult.Score},
		bson.E{Key: "scoremax", Value: updatedCompetitionResult.ScoreMax},
		bson.E{Key: "createdyear", Value: modelObj.CreatedYear},
		bson.E{Key: "createdmonth", Value: modelObj.CreatedMonth},
		bson.E{Key: "createdday", Value: modelObj.CreatedDay},
		bson.E{Key: "schoolid", Value: updatedCompetitionResult.SchoolId}}
	_, err := impl.collection.UpdateOne(impl.ctx, filter, bson.D{bson.E{Key: "$set", Value: update}})

	if err != nil {
		return updatedCompetitionResult, errors.Error("Could not upadte CompetitionResult")
	}

	log.Print("PutCompetitionResult completed")
	return updatedCompetitionResult, nil
}
