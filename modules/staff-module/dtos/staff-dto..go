package dtos

import "time"

type CreateStaffRequest struct {
	Type     string `json:"type" binding:"required"`
	SchoolId string `json:"schoolId" binding:"required"`
}

type UpdateStaffRequest struct {
	Type     string `json:"type" binding:"required"`
	SchoolId string `json:"schoolId" binding:"required"`
}

type StaffResponse struct {
	Id        string    `json:"id"  bson:"_id"`
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy string    `json:"createdBy"`
	Type      string    `json:"type" binding:"required"`
	SchoolId  string    `json:"schoolId" binding:"required"`
}
