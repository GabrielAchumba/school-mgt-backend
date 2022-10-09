package dtos

import "time"

type CreateResultRequest struct {
	Score         float64 `json:"score" binding:"required"`
	ScoreMax      float64 `json:"scoreMax" binding:"required"`
	SubjectId     string  `json:"subjectId" binding:"required"`
	StudentId     string  `json:"studentId" binding:"required"`
	TeacherId     string  `json:"teacherId" binding:"required"`
	ClassRoomId   string  `json:"classRoomId" binding:"required"`
	AssessmentId  string  `json:"assessmentId" binding:"required"`
	DesignationId string  `json:"designationId" binding:"required"`
	CreatedAt     string  `json:"createdAt" binding:"required"`
}

type UpdateResultRequest struct {
	Score         float64 `json:"score" binding:"required"`
	ScoreMax      float64 `json:"scoreMax" binding:"required"`
	SubjectId     string  `json:"subjectId" binding:"required"`
	StudentId     string  `json:"studentId" binding:"required"`
	TeacherId     string  `json:"teacherId" binding:"required"`
	ClassRoomId   string  `json:"classRoomId" binding:"required"`
	AssessmentId  string  `json:"assessmentId" binding:"required"`
	DesignationId string  `json:"designationId" binding:"required"`
	UpdatedAt     string  `json:"updatedAt" binding:"required"`
}

type ResultResponse struct {
	Id                  string    `json:"id"  bson:"_id"`
	CreatedAt           time.Time `json:"createdAt"`
	CreatedBy           string    `json:"createdBy"`
	Score               float64   `json:"score" binding:"required"`
	ScoreMax            float64   `json:"scoreMax" binding:"required"`
	SubjectId           string    `json:"subjectId" binding:"required"`
	StudentId           string    `json:"studentId" binding:"required"`
	TeacherId           string    `json:"teacherId" binding:"required"`
	ClassRoomId         string    `json:"classRoomId" binding:"required"`
	AssessmentId        string    `json:"assessmentId" binding:"required"`
	DesignationId       string    `json:"designationId" binding:"required"`
	SubjectFullName     string    `json:"subjectFullName" binding:"required"`
	StudentFullName     string    `json:"studentFullName" binding:"required"`
	TeacherFullName     string    `json:"teacherFullName" binding:"required"`
	ClassRoomFullName   string    `json:"classRoomFullName" binding:"required"`
	AssessmentFullName  string    `json:"assessmentFullName" binding:"required"`
	DesignationFullName string    `json:"designationFullName" binding:"required"`
}

type RangeOfScore struct {
	From             float64 `json:"from"`
	To               float64 `json:"to"`
	NumberOfStudents int     `json:"numberOfStudents"`
}

type GetResultsRequest struct {
	StartDate     string         `json:"startDate"`
	EndDate       string         `json:"endDate"`
	Score         float64        `json:"score"`
	ScoreMax      float64        `json:"scoreMax"`
	SubjectIds    []string       `json:"subjectIds"`
	StudentId     string         `json:"studentId"`
	StudentIds    []string       `json:"studentIds"`
	TeacherId     string         `json:"teacherId"`
	TeacherIds    []string       `json:"teacherIds"`
	ClassRoomId   string         `json:"classRoomId"`
	AssessmentId  string         `json:"assessmentId"`
	DesignationId string         `json:"designationId"`
	RangeOfScores []RangeOfScore `json:"rangeOfScores"`
}

type SubJectResult struct {
	Assessments  map[string]float64 `json:"assessments"`
	SubjectScore float64            `json:"subjectScore"`
}

type StudentResults struct {
	StudentId string                   `json:"studentId"`
	Subjects  map[string]SubJectResult `json:"subjects"`
}
