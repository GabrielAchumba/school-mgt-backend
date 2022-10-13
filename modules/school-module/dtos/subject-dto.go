package dtos

import "time"

type CreateSchoolRequest struct {
	SchoolName string `json:"schoolName" binding:"required"`
	Address    string `json:"address" binding:"required"`
	ReferedBy  string `json:"referedBy" binding:"required"`
}

type UpdateSchoolRequest struct {
	SchoolName string `json:"schoolName" binding:"required"`
	Address    string `json:"address" binding:"required"`
}

type SchoolResponse struct {
	Id         string    `json:"id"  bson:"_id"`
	CreatedAt  time.Time `json:"createdAt"`
	SchoolName string    `json:"schoolName" binding:"required"`
	Address    string    `json:"address" binding:"required"`
	ReferedBy  string    `json:"referedBy" binding:"required"`
}
