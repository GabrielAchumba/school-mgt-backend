package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Base64String      string    `json:"base64String"`
	CreatedAt         time.Time `json:"createdAt"`
	CreatedBy         string    `json:"createdBy"`
	ConfirmedBy       string    `json:"confirmedBy"`
	BlockedBy         string    `json:"blockedBy"`
	Confirmed         bool      `json:"confirmed"`
	CountryCode       string    `json:"countryCode"`
	FirstName         string    `json:"firstName" binding:"required"`
	LastName          string    `json:"lastName" binding:"required"`
	PhoneNumber       string    `json:"phoneNumber"`
	UserType          string    `json:"userType" binding:"required"`
	DesignationId     string    `json:"designationId"`
	UserName          string    `json:"userName" binding:"required"`
	Password          string    `json:"password" binding:"required"`
	RealestateCompany string    `json:"realestateCompany"`
	FileUrl           string    `json:"fileUrl"`
	FileName          string    `json:"fileName"`
	OriginalFileName  string    `json:"originalFileName"`
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
