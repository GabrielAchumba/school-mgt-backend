package dtos

import "time"

type CreateAssessmentRequest struct {
	Type       string  `json:"type" binding:"required"`
	Percentage float64 `json:"percentage" binding:"required"`
	SubjectId  string  `json:"subjectId"`
	Name       string  `json:"name"`
	SchoolId   string  `json:"schoolId" binding:"required"`
}

type UpdateAssessmentRequest struct {
	Type       string  `json:"type" binding:"required"`
	Percentage float64 `json:"percentage" binding:"required"`
	SubjectId  string  `json:"subjectId"`
	Name       string  `json:"name"`
	SchoolId   string  `json:"schoolId" binding:"required"`
}

type AssessmentResponse struct {
	Id         string    `json:"id"  bson:"_id"`
	CreatedAt  time.Time `json:"createdAt"`
	CreatedBy  string    `json:"createdBy"`
	Type       string    `json:"type" binding:"required"`
	Percentage float64   `json:"percentage" binding:"required"`
	Name       string    `json:"name"`
	SubjectId  string    `json:"subjectId"`
	SchoolId   string    `json:"schoolId" binding:"required"`
}
