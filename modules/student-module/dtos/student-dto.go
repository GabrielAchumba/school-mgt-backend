package dtos

import "time"

type CreateStudentRequest struct {
	FirstName  string `json:"firstName" binding:"required"`
	LastName   string `json:"lastName" binding:"required"`
	BirthYear  int    `json:"birthYear" binding:"required"`
	BirthMonth int    `json:"birthMonth" binding:"required"`
	BirthDay   int    `json:"birthDay" binding:"required"`
}

type UpdateStudentRequest struct {
	FirstName  string `json:"firstName" binding:"required"`
	LastName   string `json:"lastName" binding:"required"`
	BirthYear  int    `json:"birthYear" binding:"required"`
	BirthMonth int    `json:"birthMonth" binding:"required"`
	BirthDay   int    `json:"birthDay" binding:"required"`
}

type StudentResponse struct {
	Id         string    `json:"id"  bson:"_id"`
	CreatedAt  time.Time `json:"createdAt"`
	CreatedBy  string    `json:"createdBy"`
	FirstName  string    `json:"firstName" binding:"required"`
	LastName   string    `json:"lastName" binding:"required"`
	BirthYear  int       `json:"birthYear" binding:"required"`
	BirthMonth int       `json:"birthMonth" binding:"required"`
	BirthDay   int       `json:"birthDay" binding:"required"`
}
