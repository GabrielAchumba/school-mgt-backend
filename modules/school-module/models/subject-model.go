package models

import "time"

type School struct {
	CreatedAt  time.Time `json:"createdAt"`
	CreatedBy  string    `json:"createdBy"`
	SchoolName string    `json:"schoolName" binding:"required"`
	Address    string    `json:"address" binding:"required"`
}
