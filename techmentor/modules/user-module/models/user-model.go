package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
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
