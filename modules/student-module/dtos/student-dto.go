package dtos

import "time"

type CreateStudentRequest struct {
	Type string `json:"type" binding:"required"`
}

type UpdateStudentRequest struct {
	Type string `json:"type" binding:"required"`
}

type StudentResponse struct {
	Id        string    `json:"id"  bson:"_id"`
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy string    `json:"createdBy"`
	Type      string    `json:"type" binding:"required"`
}
