package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Base64String         string    `json:"base64String"`
	IsPhotographUploaded int       `json:"isPhotographUploaded"`
	CreatedAt            time.Time `json:"createdAt"`
	CreatedBy            string    `json:"createdBy"`
	ConfirmedBy          string    `json:"confirmedBy"`
	BlockedBy            string    `json:"blockedBy"`
	Confirmed            bool      `json:"confirmed"`
	CountryCode          string    `json:"countryCode"`
	FirstName            string    `json:"firstName" binding:"required"`
	LastName             string    `json:"lastName" binding:"required"`
	PhoneNumber          string    `json:"phoneNumber"`
	Email                string    `json:"email"`
	UserType             string    `json:"userType,omitempty" binding:"required"`
	DesignationId        string    `json:"designationId"`
	UserName             string    `json:"userName,omitempty" binding:"required"`
	Password             string    `json:"password,omitempty" binding:"required"`
	ClassRoomId          string    `json:"classRoomId"`
	LevelId              string    `json:"levelId"`
	SessionId            string    `json:"sessionId"`
	Token                int       `json:"token"`
	PasswordResetToken   string    `json:"passwordResetToken"`
	PasswordResetAt      time.Time `json:"passwordResetAt"`
	SchoolId             string    `json:"schoolId" binding:"required"`
	FileUrl              string    `json:"fileUrl"`
	FileName             string    `json:"fileName"`
	OriginalFileName     string    `json:"originalFileName"`
	ReferralID           string    `json:"referralID"`
	ReferalSchoolId      string    `json:"referalSchoolId"`
}

func (user *User) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func CheckPassword(savedPassword, providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(savedPassword), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
