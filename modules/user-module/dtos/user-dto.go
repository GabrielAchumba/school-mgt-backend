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
	ID               string    `json:"id"  bson:"_id"`
	PhoneNumber      string    `json:"phoneNumber" `
	FirstName        string    `json:"firstName" `
	LastName         string    `json:"lastName" `
	UserType         string    `json:"userType" `
	Designation      string    `json:"designation" `
	DesignationId    string    `json:"designationId"`
	UserName         string    `json:"userName,omitempty"`
	Password         string    `json:"password"`
	Token            int       `json:"token"`
	CreatedAt        time.Time `json:"createdAt"`
	Base64String     string    `json:"base64String"`
	SchoolId         string    `json:"schoolId" binding:"required"`
	CountryCode      string    `json:"countryCode"`
	ClassRoomId      string    `json:"classRoomId"`
	LevelId          string    `json:"levelId"`
	SessionId        string    `json:"sessionId"`
	ConfirmedBy      string    `json:"confirmedBy"`
	BlockedBy        string    `json:"blockedBy"`
	Confirmed        bool      `json:"confirmed"`
	FileUrl          string    `json:"fileUrl"`
	FileName         string    `json:"fileName"`
	OriginalFileName string    `json:"originalFileName"`
	ReferralID       string    `json:"referralID"`
	ReferalSchoolId  string    `json:"referalSchoolId"`
}

type CreateUserRequest struct {
	CreatedBy            string   `json:"createdBy"`
	Base64String         string   `json:"base64String,omitempty"`
	IsPhotographUploaded int64    `json:"isPhotographUploaded,omitempty"`
	FirstName            string   `json:"firstName"`
	LastName             string   `json:"lastName"`
	PhoneNumber          string   `json:"phoneNumber"`
	CountryCode          string   `json:"countryCode"`
	Email                string   `json:"email"`
	UserType             string   `json:"userType"`
	DesignationId        string   `json:"designationId"`
	UserName             string   `json:"userName"`
	Password             string   `json:"password"`
	Token                int      `json:"token"`
	SchoolId             string   `json:"schoolId"`
	ClassRoomId          string   `json:"classRoomId"`
	ClassRoomIds         []string `json:"classRoomIds"`
	LevelId              string   `json:"levelId"`
	SessionId            string   `json:"sessionId"`
	ConfirmedBy          string   `json:"confirmedBy"`
	BlockedBy            string   `json:"blockedBy"`
	Confirmed            bool     `json:"confirmed"`
	FileUrl              string   `json:"fileUrl"`
	FileName             string   `json:"fileName"`
	OriginalFileName     string   `json:"originalFileName"`
	ReferralID           string   `json:"referralID"`
	ReferalSchoolId      string   `json:"referalSchoolId"`
}

type UpdateUserRequest struct {
	Base64String         string   `json:"base64String"`
	IsPhotographUploaded int      `json:"isPhotographUploaded"`
	FirstName            string   `json:"firstName"`
	LastName             string   `json:"lastName"`
	PhoneNumber          string   `json:"phoneNumber"`
	CountryCode          string   `json:"countryCode"`
	Email                string   `json:"email"`
	UserType             string   `json:"userType"`
	DesignationId        string   `json:"designationId"`
	UserName             string   `json:"userName"`
	Password             string   `json:"password"`
	Token                int      `json:"token"`
	SchoolId             string   `json:"schoolId"`
	StudentIds           []string `json:"studentIds"`
	ClassRoomId          string   `json:"classRoomId"`
	LevelId              string   `json:"levelId"`
	SessionId            string   `json:"sessionId"`
	ConfirmedBy          string   `json:"confirmedBy"`
	BlockedBy            string   `json:"blockedBy"`
	Confirmed            bool     `json:"confirmed"`
	FileUrl              string   `json:"fileUrl"`
	FileName             string   `json:"fileName"`
	OriginalFileName     string   `json:"originalFileName"`
	ReferralID           string   `json:"referralID"`
	ReferalSchoolId      string   `json:"referalSchoolId"`
}

type UserResponse struct {
	Id                   string    `json:"id"  bson:"_id"`
	Base64String         string    `json:"base64String"`
	IsPhotographUploaded int       `json:"isPhotographUploaded"`
	CreatedAt            time.Time `json:"createdAt"`
	CreatedBy            string    `json:"createdBy"`
	ConfirmedBy          string    `json:"confirmedBy"`
	BlockedBy            string    `json:"blockedBy"`
	Confirmed            bool      `json:"confirmed"`
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
	SchoolId             string    `json:"schoolId" binding:"required"`
	Token                int       `json:"token"`
	ClassRoomId          string    `json:"classRoomId"`
	LevelId              string    `json:"levelId"`
	SessionId            string    `json:"sessionId"`
	FileUrl              string    `json:"fileUrl"`
	FileName             string    `json:"fileName"`
	OriginalFileName     string    `json:"originalFileName"`
	ReferralID           string    `json:"referralID"`
	ReferalSchoolId      string    `json:"referalSchoolId"`
}

type UserResponsePaginated struct {
	TotalNumberOfUsers int            `json:"totalNumberOfUsers"`
	PaginatedUsers     []UserResponse `json:"paginatedUsers"`
	Limit              int            `json:"limit"`
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
