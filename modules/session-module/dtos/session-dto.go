package dtos

import "time"

type CreateSessionRequest struct {
	Type     string `json:"type" binding:"required"`
	SchoolId string `json:"schoolId" binding:"required"`
}

type UpdateSessionRequest struct {
	Type     string `json:"type" binding:"required"`
	SchoolId string `json:"schoolId" binding:"required"`
}

type SessionResponse struct {
	Id        string    `json:"id"  bson:"_id"`
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy string    `json:"createdBy"`
	Type      string    `json:"type" binding:"required"`
	SchoolId  string    `json:"schoolId" binding:"required"`
}
