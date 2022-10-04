package dtos

import (
	"time"
)

type LoginUserRequest struct {
	UserName string `json:"userName,omitempty" `
	Password string `json:"password" binding:"required"`
}

type LoginUserResponse struct {
	Token     string       `json:"token"`
	ExpiresAt time.Time    `json:"expiresAt"`
	User      UserResponse `json:"user"`
}
type UserInternalOperation struct {
	ID            string    `json:"id"  bson:"_id"`
	PhoneNumber   string    `json:"phoneNumber" `
	FirstName     string    `json:"firstName" `
	LastName      string    `json:"lastName" `
	UserType      string    `json:"userType" `
	Designation   string    `json:"designation" `
	DesignationId string    `json:"designationId"`
	UserName      string    `json:"userName,omitempty"`
	Password      string    `json:"password"`
	CreatedAt     time.Time `json:"createdAt"`
	Base64String  string    `json:"base64String"`
}

type CreateUserRequest struct {
	Base64String         string `json:"base64String,omitempty"`
	IsPhotographUploaded int64  `json:"isPhotographUploaded,omitempty"`
	FirstName            string `json:"firstName"  binding:"required"`
	LastName             string `json:"lastName"  binding:"required"`
	PhoneNumber          string `json:"phoneNumber" binding:"required"`
	CountryCode          string `json:"countryCode" binding:"required"`
	Email                string `json:"email"`
	UserType             string `json:"userType" binding:"required"`
	DesignationId        string `json:"designationId" binding:"required"`
	UserName             string `json:"userName"  binding:"required"`
	Password             string `json:"password"  binding:"required"`
}

type UpdateUserRequest struct {
	Base64String         string `json:"base64String"`
	IsPhotographUploaded int    `json:"isPhotographUploaded"`
	FirstName            string `json:"firstName"`
	LastName             string `json:"lastName"`
	PhoneNumber          string `json:"phoneNumber"`
	CountryCode          string `json:"countryCode"`
	Email                string `json:"email"`
	UserType             string `json:"userType"`
	DesignationId        string `json:"designationId"`
	UserName             string `json:"userName"`
	Password             string `json:"password"`
}

type UserResponse struct {
	Id                   string    `json:"id"  bson:"_id"`
	Base64String         string    `json:"base64String"`
	IsPhotographUploaded int       `json:"isPhotographUploaded"`
	CreatedAt            time.Time `json:"createdAt"`
	CreatedBy            string    `json:"createdBy"`
	CountryCode          string    `json:"countryCode"`
	FirstName            string    `json:"firstName"`
	LastName             string    `json:"lastName"`
	PhoneNumber          string    `json:"phoneNumber"`
	Email                string    `json:"email"`
	UserType             string    `json:"userType,omitempty"`
	Designation          string    `json:"designation,omitempty"`
	DesignationId        string    `json:"designationId,omitempty"`
	UserName             string    `json:"userName,omitempty"`
	Password             string    `json:"password,omitempty"`
}

type ForgotPasswordInput struct {
	Email      string `json:"email"`
	ResetToken string `json:"resetToken"`
	ExpiryTime int    `json:"expiryTime"`
	Message    string `json:"message"`
	UserName   string `json:"userName"`
}

type ResetPasswordInput struct {
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
	ResetToken      string `json:"resetToken"`
	UserName        string `json:"userName"`
}
