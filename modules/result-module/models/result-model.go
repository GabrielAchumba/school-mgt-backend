package models

import "time"

type Result struct {
	CreatedAt     time.Time `json:"createdAt"`
	CreatedBy     string    `json:"createdBy"`
	CreatedYear   int       `json:"createdYear" binding:"required"`
	CreatedMonth  int       `json:"createdMonth" binding:"required"`
	CreatedDay    int       `json:"createdDay" binding:"required"`
	Score         float64   `json:"score" binding:"required"`
	ScoreMax      float64   `json:"scoreMax" binding:"required"`
	SubjectId     string    `json:"subjectId" binding:"required"`
	StudentId     string    `json:"studentId" binding:"required"`
	TeacherId     string    `json:"teacherId" binding:"required"`
	ClassRoomId   string    `json:"classRoomId" binding:"required"`
	AssessmentId  string    `json:"assessmentId" binding:"required"`
	DesignationId string    `json:"designationId" binding:"required"`
	SchoolId      string    `json:"schoolId" binding:"required"`
}
