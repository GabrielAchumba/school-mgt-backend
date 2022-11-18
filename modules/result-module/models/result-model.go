package models

import "time"

type Result struct {
	CreatedAt     time.Time `json:"createdAt"`
	CreatedBy     string    `json:"createdBy"`
	CreatedYear   int       `json:"createdYear" binding:"required"`
	CreatedMonth  int       `json:"createdMonth" binding:"required"`
	CreatedDay    int       `json:"createdDay" binding:"required"`    // 9
	Score         float64   `json:"score" binding:"required"`         // 7
	ScoreMax      float64   `json:"scoreMax" binding:"required"`      // 8
	SubjectId     string    `json:"subjectId" binding:"required"`     //2
	StudentId     string    `json:"studentId" binding:"required"`     // 3
	TeacherId     string    `json:"teacherId" binding:"required"`     // 5
	ClassRoomId   string    `json:"classRoomId" binding:"required"`   //1
	AssessmentId  string    `json:"assessmentId" binding:"required"`  // 6
	DesignationId string    `json:"designationId" binding:"required"` // 4
	LevelId       string    `json:"levelId"`                          //10
	SchoolId      string    `json:"schoolId" binding:"required"`
}
