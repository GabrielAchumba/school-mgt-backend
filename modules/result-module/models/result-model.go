package models

import "time"

type Result struct {
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy string    `json:"createdBy"`
	Score     float64   `json:"score" binding:"required"`
	SubjectId string    `json:"subjectId" binding:"required"`
	StudentId string    `json:"studentId" binding:"required"`
	TeacherId string    `json:"teacherId" binding:"required"`
}
