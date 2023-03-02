package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Base64String         string    `json:"base64String"`
	IsPhotographUploaded int       `json:"isPhotographUploaded"`
	CreatedBy            string    `json:"createdBy"`
	CreatedDay           int       `json:"createdDay"`
	CreatedMonth         int       `json:"createdMonth"`
	CreatedYear          int       `json:"createdYear"`
	CreatedAt            time.Time `json:"createdAt"`
	Title                string    `json:"title"`
	FirstName            string    `json:"firstName" binding:"required"`
	MiddleName           string    `json:"middleName"`
	LastName             string    `json:"lastName" binding:"required"`
	Address              string    `json:"address"`
	ResidentialCity      string    `json:"residentialCity"`
	ResidentialState     string    `json:"residentialState"`
	Email                string    `json:"email"`
	PhoneNumber          string    `json:"phoneNumber"`
	BloodGroup           string    `json:"bloodGroup"`
	Genotype             string    `json:"genotype"`
	MaritalStatus        string    `json:"maritalStatus"`
	LGAOfOrigin          string    `json:"lGAOfOrigin"`
	StateOfOrigin        string    `json:"stateOfOrigin"`
	Country              string    `json:"country"`
	NOKNames             string    `json:"nOKNames"`
	NOKAddress           string    `json:"nOKAddress"`
	NOKPhoneNumber       string    `json:"nOKPhoneNumber"`
	NOKRelationship      string    `json:"nOKRelationship"`
	BankName             string    `json:"bankName"`
	AccountName          string    `json:"accountName"`
	AccountNumber        string    `json:"accountNumber"`
	BVN                  string    `json:"bVN"`

	UserType    string `json:"userType,omitempty" binding:"required"`
	Designation string `json:"designation"`
	Description string `json:"description"`
	Region      string `json:"region"`

	UserName           string    `json:"userName,omitempty" binding:"required"`
	Password           string    `json:"password,omitempty" binding:"required"`
	PasswordResetToken string    `json:"passwordResetToken"`
	PasswordResetAt    time.Time `json:"passwordResetAt"`
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
