package dtos

import (
	"time"
)

type LoginUserRequest struct {
	UserName string `json:"userName,omitempty" `
	Password string `json:"password" binding:"required"`
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

type CreateUserRequest struct {
	CreatedAt        time.Time `json:"createdAt"`
	CreatedBy        string    `json:"createdBy"`
	CountryCode      string    `json:"countryCode"`
	FullName         string    `json:"fullName" binding:"required"`
	PhoneNumber      string    `json:"phoneNumber"`
	UserType         string    `json:"userType,omitempty" binding:"required"`
	UserName         string    `json:"userName,omitempty" binding:"required"`
	Password         string    `json:"password,omitempty" binding:"required"`
	FileUrl          string    `json:"fileUrl"`
	FileName         string    `json:"fileName"`
	OriginalFileName string    `json:"originalFileName"`
}

type UserInternalOperation struct {
	ID               string    `json:"id"  bson:"_id"`
	CreatedAt        time.Time `json:"createdAt"`
	CreatedBy        string    `json:"createdBy"`
	CountryCode      string    `json:"countryCode"`
	FullName         string    `json:"fullName" binding:"required"`
	PhoneNumber      string    `json:"phoneNumber"`
	UserType         string    `json:"userType,omitempty" binding:"required"`
	UserName         string    `json:"userName,omitempty" binding:"required"`
	Password         string    `json:"password,omitempty" binding:"required"`
	FileUrl          string    `json:"fileUrl"`
	FileName         string    `json:"fileName"`
	OriginalFileName string    `json:"originalFileName"`
}
