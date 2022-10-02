package models

import "time"

type Student struct {
	CreatedAt  time.Time `json:"createdAt"`
	CreatedBy  string    `json:"createdBy"`
	FirstName  string    `json:"firstName" binding:"required"`
	LastName   string    `json:"lastName" binding:"required"`
	BirthYear  int       `json:"birthYear" binding:"required"`
	BirthMonth int       `json:"birthMonth" binding:"required"`
	BirthDay   int       `json:"birthDay" binding:"required"`
}
