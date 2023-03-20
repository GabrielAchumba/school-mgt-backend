package models

import (
	"time"
)

type TechRequest struct {
	CreatedAt      time.Time `json:"createdAt"`
	CreatedBy      string    `json:"createdBy"`
	UserId         string    `json:"userId"`
	Title          string    `json:"title"`
	TimeZone       string    `json:"timeZone"`
	CountryCode    string    `json:"countryCode"`
	ExpectedBudget string    `json:"expectedBudget"`
	RequestStatus  string    `json:"requestStatus"`
	RequestDetails string    `json:"requestDetails"`
	Deliverables   string    `json:"deliverables"`
	TypeOfProject  string    `json:"typeOfProject"`
}
