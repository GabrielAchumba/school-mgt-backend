package dtos

import "time"

type CreateClassRoomRequest struct {
	Type string `json:"type" binding:"required"`
}

type UpdateClassRoomRequest struct {
	Type string `json:"type" binding:"required"`
}

type ClassRoomResponse struct {
	Id        string    `json:"id"  bson:"_id"`
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy string    `json:"createdBy"`
	Type      string    `json:"type" binding:"required"`
}
