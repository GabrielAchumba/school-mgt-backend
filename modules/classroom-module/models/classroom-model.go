package models

import "time"

type ClassRoom struct {
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy string    `json:"createdBy"`
	Type      string    `json:"type" binding:"required"`
	LevelId   string    `json:"levelId"`
	SchoolId  string    `json:"schoolId" binding:"required"`
}
