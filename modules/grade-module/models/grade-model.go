package models

import "time"

type Grade struct {
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy string    `json:"createdBy"`
	Type      string    `json:"type" binding:"required"`
	Point     float64   `json:"point"`
	From      float64   `json:"from"`
	To        float64   `json:"to"`
	SchoolId  string    `json:"schoolId" binding:"required"`
}
