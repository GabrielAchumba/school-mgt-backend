package dtos

import "time"

type CreateStudentRequest struct {
	FirstName        string `json:"firstName" binding:"required"`
	LastName         string `json:"lastName" binding:"required"`
	BirthYear        int    `json:"birthYear" binding:"required"`
	BirthMonth       int    `json:"birthMonth" binding:"required"`
	BirthDay         int    `json:"birthDay" binding:"required"`
	UserName         string `json:"userName"`
	Password         string `json:"password"`
	ClassRoomId      string `json:"classRoomId"`
	LevelId          string `json:"levelId"`
	SessionId        string `json:"sessionId"`
	SubscriptionType int    `json:"subscriptionType" binding:"required"`
	SchoolId         string `json:"schoolId"`
}

type UpdateStudentRequest struct {
	FirstName        string   `json:"firstName"`
	LastName         string   `json:"lastName"`
	BirthYear        int      `json:"birthYear"`
	BirthMonth       int      `json:"birthMonth"`
	BirthDay         int      `json:"birthDay"`
	UserName         string   `json:"userName"`
	Password         string   `json:"password"`
	ClassRoomId      string   `json:"classRoomId"`
	LevelId          string   `json:"levelId"`
	SessionId        string   `json:"sessionId"`
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
	UserName                  string    `json:"userName"`
	Password                  string    `json:"password"`
	ClassRoomId               string    `json:"classRoomId"`
	LevelId                   string    `json:"levelId"`
	SessionId                 string    `json:"sessionId"`
	Token                     int       `json:"token" binding:"required"`
	SubscriptionType          int       `json:"subscriptionType" binding:"required"`
	RemainingSubscriptionDays int       `json:"remainingSubscriptionDays"`
	SchoolId                  string    `json:"schoolId"`
}

type LoginStudentResponse struct {
	Token     string          `json:"token"`
	ExpiresAt time.Time       `json:"expiresAt"`
	User      StudentResponse `json:"user"`
}
