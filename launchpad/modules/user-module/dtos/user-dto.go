package dtos

import (
	"time"
)

type LoginUserRequest struct {
	UserName string `json:"userName,omitempty" `
	Password string `json:"password"`
}

type LoginUserResponse struct {
	Token     string       `json:"token"`
	ExpiresAt time.Time    `json:"expiresAt"`
	User      UserResponse `json:"user"`
}
type UserInternalOperation struct {
	ID           string    `json:"id"  bson:"_id"`
	PhoneNumber  string    `json:"phoneNumber" `
	Title        string    `json:"title" `
	FirstName    string    `json:"firstName" `
	MiddleName   string    `json:"middleName"`
	LastName     string    `json:"lastName" `
	UserType     string    `json:"userType" `
	Designation  string    `json:"designation" `
	Description  string    `json:"description"`
	Region       string    `json:"region"`
	UserName     string    `json:"userName,omitempty"`
	Password     string    `json:"password"`
	CreatedAt    time.Time `json:"createdAt"`
	Base64String string    `json:"base64String"`
}

type CreateUserRequest struct {
	Base64String         string `json:"base64String,omitempty"`
	IsPhotographUploaded int64  `json:"isPhotographUploaded,omitempty"`
	CreatedDay           int    `json:"createdDay"`
	CreatedMonth         int    `json:"createdMonth"`
	CreatedYear          int    `json:"createdYear"`

	Title       string `json:"title"`
	FirstName   string `json:"firstName"`
	MiddleName  string `json:"middleName"`
	LastName    string `json:"lastName"`
	PhoneNumber string `json:"phoneNumber"`

	UserType    string `json:"userType"`
	Designation string `json:"designation"`
	Description string `json:"description"`
	Region      string `json:"region"`

	UserName       string `json:"userName"`
	Password       string `json:"password"`
	ContributorId  string `json:"contributorId"`
	ParentUserName string `json:"parentUserName"`
	CountryCode    string `json:"countryCode"`
}

type UpdateUserRequest struct {
	Base64String         string `json:"base64String"`
	IsPhotographUploaded int    `json:"isPhotographUploaded"`
	CreatedDay           int    `json:"createdDay"`
	CreatedMonth         int    `json:"createdMonth"`
	CreatedYear          int    `json:"createdYear"`

	Title       string `json:"title"`
	FirstName   string `json:"firstName"`
	MiddleName  string `json:"middleName"`
	LastName    string `json:"lastName"`
	PhoneNumber string `json:"phoneNumber"`

	UserType    string `json:"userType"`
	Designation string `json:"designation"`
	Description string `json:"description"`
	Region      string `json:"region"`

	UserName    string `json:"userName"`
	Password    string `json:"password"`
	CountryCode string `json:"countryCode"`
}

type UserResponse struct {
	Id            string `json:"id"  bson:"_id"`
	Title         string `json:"title"`
	FirstName     string `json:"firstName"`
	MiddleName    string `json:"middleName"`
	LastName      string `json:"lastName"`
	FullName      string `json:"fullName"`
	Gender        string `json:"gender"`
	ContributorId string `json:"contributorId"`
	PhoneNumber   string `json:"phoneNumber"`

	UserType    string `json:"userType"`
	Designation string `json:"designation"`
	Description string `json:"description"`
	Region      string `json:"region"`

	UserName  string    `json:"userName"`
	CreatedAt time.Time `json:"createdAt"`

	Base64String         string `json:"base64String"`
	IsPhotographUploaded int    `json:"isPhotographUploaded"`
	CreatedBy            string `json:"createdBy"`
	CreatedDay           int    `json:"createdDay"`
	CreatedMonth         int    `json:"createdMonth"`
	CreatedYear          int    `json:"createdYear"`
	Address              string `json:"address"`
	ResidentialCity      string `json:"residentialCity"`
	ResidentialState     string `json:"residentialState"`
	Email                string `json:"email"`
	BloodGroup           string `json:"bloodGroup"`
	Genotype             string `json:"genotype"`
	MaritalStatus        string `json:"maritalStatus"`
	LGAOfOrigin          string `json:"lGAOfOrigin"`
	StateOfOrigin        string `json:"stateOfOrigin"`
	Country              string `json:"country"`
	NOKNames             string `json:"nOKNames"`
	NOKAddress           string `json:"nOKAddress"`
	NOKPhoneNumber       string `json:"nOKPhoneNumber"`
	NOKRelationship      string `json:"nOKRelationship"`
	BankName             string `json:"bankName"`
	AccountName          string `json:"accountName"`
	AccountNumber        string `json:"accountNumber"`
	BVN                  string `json:"bVN"`

	Password    string `json:"password,omitempty"`
	CountryCode string `json:"countryCode"`
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
	UserName        string `json:"userName"`
}
