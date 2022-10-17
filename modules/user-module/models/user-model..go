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
	CountryCode          string    `json:"countryCode"`
	FirstName            string    `json:"firstName" binding:"required"`
	LastName             string    `json:"lastName" binding:"required"`
	PhoneNumber          string    `json:"phoneNumber"`
	Email                string    `json:"email"`
	UserType             string    `json:"userType,omitempty" binding:"required"`
	DesignationId        string    `json:"designationId"`
	UserName             string    `json:"userName,omitempty" binding:"required"`
	Password             string    `json:"password,omitempty" binding:"required"`
	PasswordResetToken   string    `json:"passwordResetToken"`
	PasswordResetAt      time.Time `json:"passwordResetAt"`
	SchoolId             string    `json:"schoolId" binding:"required"`
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
