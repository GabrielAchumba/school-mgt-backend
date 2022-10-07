package dtos

import "time"

type CreateAssessmentRequest struct {
	Type       string  `json:"type" binding:"required"`
	Percentage float64 `json:"percentage" binding:"required"`
}

type UpdateAssessmentRequest struct {
	Type       string  `json:"type" binding:"required"`
	Percentage float64 `json:"percentage" binding:"required"`
}

type AssessmentResponse struct {
	Id         string    `json:"id"  bson:"_id"`
	CreatedAt  time.Time `json:"createdAt"`
	CreatedBy  string    `json:"createdBy"`
	Type       string    `json:"type" binding:"required"`
	Percentage float64   `json:"percentage" binding:"required"`
}
