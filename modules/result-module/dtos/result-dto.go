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

type GetResultsRequest struct {
	StartDate     string   `json:"startDate" binding:"required"`
	EndDate       string   `json:"endDate" binding:"required"`
	Score         float64  `json:"score"`
	ScoreMax      float64  `json:"scoreMax"`
	SubjectIds    []string `json:"subjectIds" binding:"required"`
	StudentId     string   `json:"studentId" binding:"required"`
	TeacherId     string   `json:"teacherId" binding:"required"`
	ClassRoomId   string   `json:"classRoomId" binding:"required"`
	AssessmentId  string   `json:"assessmentId"`
	DesignationId string   `json:"designationId"`
}
