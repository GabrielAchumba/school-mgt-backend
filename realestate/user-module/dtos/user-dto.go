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
	ID                string    `json:"id"  bson:"_id"`
	PhoneNumber       string    `json:"phoneNumber" `
	FirstName         string    `json:"firstName" `
	LastName          string    `json:"lastName" `
	UserType          string    `json:"userType" `
	Designation       string    `json:"designation" `
	DesignationId     string    `json:"designationId"`
	UserName          string    `json:"userName,omitempty"`
	Password          string    `json:"password"`
	CreatedAt         time.Time `json:"createdAt"`
	Base64String      string    `json:"base64String"`
	CountryCode       string    `json:"countryCode"`
	ConfirmedBy       string    `json:"confirmedBy"`
	BlockedBy         string    `json:"blockedBy"`
	Confirmed         bool      `json:"confirmed"`
	FileUrl           string    `json:"fileUrl"`
	FileName          string    `json:"fileName"`
	OriginalFileName  string    `json:"originalFileName"`
	RealestateCompany string    `json:"realestateCompany"`
}

type CreateUserRequest struct {
	CreatedBy         string `json:"createdBy"`
	Base64String      string `json:"base64String,omitempty"`
	FirstName         string `json:"firstName"`
	LastName          string `json:"lastName"`
	PhoneNumber       string `json:"phoneNumber"`
	CountryCode       string `json:"countryCode"`
	UserType          string `json:"userType"`
	DesignationId     string `json:"designationId"`
	UserName          string `json:"userName"`
	Password          string `json:"password"`
	ConfirmedBy       string `json:"confirmedBy"`
	BlockedBy         string `json:"blockedBy"`
	Confirmed         bool   `json:"confirmed"`
	FileUrl           string `json:"fileUrl"`
	FileName          string `json:"fileName"`
	OriginalFileName  string `json:"originalFileName"`
	RealestateCompany string `json:"realestateCompany"`
}

type UpdateUserRequest struct {
	Base64String         string `json:"base64String"`
	IsPhotographUploaded int    `json:"isPhotographUploaded"`
	FirstName            string `json:"firstName"`
	LastName             string `json:"lastName"`
	PhoneNumber          string `json:"phoneNumber"`
	CountryCode          string `json:"countryCode"`
	UserType             string `json:"userType"`
	DesignationId        string `json:"designationId"`
	UserName             string `json:"userName"`
	Password             string `json:"password"`
	ConfirmedBy          string `json:"confirmedBy"`
	BlockedBy            string `json:"blockedBy"`
	Confirmed            bool   `json:"confirmed"`
	FileUrl              string `json:"fileUrl"`
	FileName             string `json:"fileName"`
	OriginalFileName     string `json:"originalFileName"`
	RealestateCompany    string `json:"realestateCompany"`
}

type UserResponse struct {
	Id                string    `json:"id"  bson:"_id"`
	Base64String      string    `json:"base64String"`
	CreatedAt         time.Time `json:"createdAt"`
	CreatedBy         string    `json:"createdBy"`
	ConfirmedBy       string    `json:"confirmedBy"`
	BlockedBy         string    `json:"blockedBy"`
	Confirmed         bool      `json:"confirmed"`
	CountryCode       string    `json:"countryCode"`
	FirstName         string    `json:"firstName"`
	LastName          string    `json:"lastName"`
	PhoneNumber       string    `json:"phoneNumber"`
	UserType          string    `json:"userType,omitempty"`
	Designation       string    `json:"designation,omitempty"`
	DesignationId     string    `json:"designationId,omitempty"`
	UserName          string    `json:"userName,omitempty"`
	Password          string    `json:"password,omitempty"`
	FileUrl           string    `json:"fileUrl"`
	FileName          string    `json:"fileName"`
	OriginalFileName  string    `json:"originalFileName"`
	RealestateCompany string    `json:"realestateCompany"`
}

type UserResponsePaginated struct {
	TotalNumberOfUsers int            `json:"totalNumberOfUsers"`
	PaginatedUsers     []UserResponse `json:"paginatedUsers"`
	Limit              int            `json:"limit"`
}

type ForgotPasswordInput struct {
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
