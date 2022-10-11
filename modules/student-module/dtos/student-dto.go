package dtos

import "time"

type CreateStudentRequest struct {
	FirstName        string `json:"firstName" binding:"required"`
	LastName         string `json:"lastName" binding:"required"`
	BirthYear        int    `json:"birthYear" binding:"required"`
	BirthMonth       int    `json:"birthMonth" binding:"required"`
	BirthDay         int    `json:"birthDay" binding:"required"`
	SubscriptionType int    `json:"subscriptionType" binding:"required"`
	SchoolId         string `json:"schoolId"`
}

type UpdateStudentRequest struct {
	FirstName        string   `json:"firstName"`
	LastName         string   `json:"lastName"`
	BirthYear        int      `json:"birthYear"`
	BirthMonth       int      `json:"birthMonth"`
	BirthDay         int      `json:"birthDay"`
	Token            int      `json:"token"`
	SubscriptionType int      `json:"subscriptionType"`
	StudentIds       []string `json:"studentIds"`
	SchoolId         string   `json:"schoolId"`
}

type StudentResponse struct {
	Id                        string    `json:"id"  bson:"_id"`
	CreatedAt                 time.Time `json:"createdAt"`
	CreatedBy                 string    `json:"createdBy"`
	FirstName                 string    `json:"firstName" binding:"required"`
	LastName                  string    `json:"lastName" binding:"required"`
	BirthYear                 int       `json:"birthYear" binding:"required"`
	BirthMonth                int       `json:"birthMonth" binding:"required"`
	BirthDay                  int       `json:"birthDay" binding:"required"`
	UserType                  string    `json:"userType" binding:"required"`
	Token                     int       `json:"token" binding:"required"`
	SubscriptionType          int       `json:"subscriptionType" binding:"required"`
	RemainingSubscriptionDays int       `json:"remainingSubscriptionDays"`
	SchoolId                  string    `json:"schoolId"`
}
