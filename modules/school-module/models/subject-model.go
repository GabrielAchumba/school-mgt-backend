package models

import "time"

type School struct {
	CreatedAt  time.Time `json:"createdAt"`
	SchoolName string    `json:"schoolName" binding:"required"`
	Address    string    `json:"address" binding:"required"`
	ReferedBy  string    `json:"referedBy" binding:"required"`
}
