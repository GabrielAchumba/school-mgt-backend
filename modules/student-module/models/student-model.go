package models

import "time"

type Student struct {
	CreatedAt               time.Time `json:"createdAt"`
	CreatedBy               string    `json:"createdBy"`
	FirstName               string    `json:"firstName" binding:"required"`
	LastName                string    `json:"lastName" binding:"required"`
	BirthYear               int       `json:"birthYear" binding:"required"`
	BirthMonth              int       `json:"birthMonth" binding:"required"`
	BirthDay                int       `json:"birthDay" binding:"required"`
	UserType                string    `json:"userType" binding:"required"`
	Token                   int       `json:"token" binding:"required"`
	UserName                string    `json:"userName,omitempty" binding:"required"`
	Password                string    `json:"password,omitempty" binding:"required"`
	SubscriptionType        int       `json:"subscriptionType" binding:"required"`
	CreatedSubscriptionDate time.Time `json:"createdSubscriptionDate"`
	SchoolId                string    `json:"schoolId" binding:"required"`
}
