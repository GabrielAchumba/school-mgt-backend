package models

import "time"

type Assessment struct {
	CreatedAt  time.Time `json:"createdAt"`
	CreatedBy  string    `json:"createdBy"`
	Type       string    `json:"type" binding:"required"`
	Percentage float64   `json:"percentage" binding:"required"`
}
