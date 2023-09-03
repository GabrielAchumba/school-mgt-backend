package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	CreatedAt        time.Time `json:"createdAt"`
	CreatedBy        string    `json:"createdBy"`
	FirstName        string    `json:"firstName" binding:"required"`
	LastName         string    `json:"lastName" binding:"required"`
	PhoneNumber      string    `json:"phoneNumber"`
	Email            string    `json:"email"`
	MaritalStatus    string    `json:"maritalStatus"`
	Gender           string    `json:"gender"`
	UserType         string    `json:"userType,omitempty" binding:"required"`
	DesignationId    string    `json:"designationId"`
	UserName         string    `json:"userName,omitempty" binding:"required"`
	Password         string    `json:"password,omitempty" binding:"required"`
	FileUrl          string    `json:"fileUrl"`
	FileName         string    `json:"fileName"`
	OriginalFileName string    `json:"originalFileName"`
	StateOfResidence string    `json:"stateOfResidence"`
	City             string    `json:"city"`
	Address          string    `json:"address"`
	NOKFullName      string    `json:"nOKFullName"`
	NOKRelationship  string    `json:"nOKRelationship"`
	NOKEmail         string    `json:"nOKEmail"`
	NOKPhone         string    `json:"nOKPhone"`
	NOKAddress       string    `json:"nOKAddress"`
	AccountType      string    `json:"accountType"`
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
