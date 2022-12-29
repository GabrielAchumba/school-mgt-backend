package dtos

import "time"

type CreateLevelRequest struct {
	Type     string `json:"type" binding:"required"`
	SchoolId string `json:"schoolId" binding:"required"`
}

type UpdateLevelRequest struct {
	Type     string `json:"type" binding:"required"`
	SchoolId string `json:"schoolId" binding:"required"`
}

type LevelResponse struct {
	Id        string    `json:"id"  bson:"_id"`
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy string    `json:"createdBy"`
	Type      string    `json:"type" binding:"required"`
	SchoolId  string    `json:"schoolId" binding:"required"`
}

type LevelIds struct {
	Ids      []string `json:"ids"`
	SchoolId string   `json:"schoolId"`
}
