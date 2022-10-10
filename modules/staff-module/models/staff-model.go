package models

import "time"

type Staff struct {
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy string    `json:"createdBy"`
	Type      string    `json:"type" binding:"required"`
	SchoolId  string    `json:"schoolId" binding:"required"`
}
