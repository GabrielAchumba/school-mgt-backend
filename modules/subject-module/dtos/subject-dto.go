package dtos

import "time"

type CreateSubjectRequest struct {
	Type     string `json:"type" binding:"required"`
	SchoolId string `json:"schoolId" binding:"required"`
}

type UpdateSubjectRequest struct {
	Type     string `json:"type" binding:"required"`
	SchoolId string `json:"schoolId" binding:"required"`
}

type SubjectResponse struct {
	Id        string    `json:"id"  bson:"_id"`
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy string    `json:"createdBy"`
	Type      string    `json:"type" binding:"required"`
	SchoolId  string    `json:"schoolId" binding:"required"`
}
