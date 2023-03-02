package dtos

import (
	"time"

	basemodule "github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/base-module"
	categoryDTOPackage "github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/category-module/dtos"
)

type AccountDTO struct {
	CreatedDay   int `json:"createdDay"`
	CreatedMonth int `json:"createdMonth"`
	CreatedYear  int `json:"createdYear"`

	ParentId              string  `json:"parentId"`
	ContributorId         string  `json:"contributorId"`
	Contribution          float64 `json:"contribution"`
	Message               string  `json:"message"`
	Reference             string  `json:"reference"`
	Status                string  `json:"status"`
	Trans                 string  `json:"trans"`
	Transactions          string  `json:"transactions"`
	Trxref                string  `json:"trxref"`
	IsComfirmed           bool    `json:"isComfirmed"`
	TypeOfStream          string  `json:"typeOfStream"`
	FileName              string  `json:"fileName"`
	Base64String          string  `json:"base64String"`
	BankNamePaidFrom      string  `json:"bankNamePaidFrom"`
	AccountNamePaidFrom   string  `json:"accountNamePaidFrom"`
	AccountNumberPaidFrom string  `json:"accountNumberPaidFrom"`
	DatePaidFrom          string  `json:"datePaidFrom"`
	UserName              string  `json:"username"`
	CategoryIndex         int     `json:"categoryIndex"`
}

type CreateAccountRequest struct {
	basemodule.BaseDTO

	FullName      string  `json:"fullName"`
	Contribution  float64 `json:"contribution"`
	ParentId      string  `json:"parentId"`
	ContributorId string  `json:"contributorId"`
	AccountId     string  `json:"accountId"`
	Status        string  `json:"status"`
	IsComfirmed   bool    `json:"isComfirmed"`
	Base64String  string  `json:"base64String"`
	UserName      string  `json:"username"`
}

type UpdateAccountRequest struct {
	basemodule.BaseModel

	FullName      string  `json:"fullName"`
	Contribution  float64 `json:"contribution"`
	ParentId      string  `json:"parentId"`
	ContributorId string  `json:"contributorId"`
	AccountId     string  `json:"accountId"`
	Status        string  `json:"status"`
	IsComfirmed   bool    `json:"isComfirmed"`
	UserName      string  `json:"username"`
}

type AccountModelResponse struct {
	Id                    string    `json:"id"  bson:"_id"`
	FullName              string    `json:"fullName"`
	Contribution          float64   `json:"contribution"`
	RegistrationFee       float64   `json:"registrationFee"`
	ContributorId         string    `json:"contributorId"`
	AccountId             string    `json:"accountId"`
	Status                string    `json:"status"`
	Base64String          string    `json:"base64String"`
	EntryDate             string    `json:"entryDate"`
	IsComfirmed           bool      `json:"isComfirmed"`
	CreatedAt             time.Time `json:"createdAt"`
	BankNamePaidFrom      string    `json:"bankNamePaidFrom"`
	AccountNamePaidFrom   string    `json:"accountNamePaidFrom"`
	AccountNumberPaidFrom string    `json:"accountNumberPaidFrom"`
	UserName              string    `json:"username"`
}

type AccountResponse struct {
	AccountModels []AccountModelResponse `json:"accountModels"`
}

type MediaDto struct {
	StatusCode int                    `json:"statusCode"`
	Message    string                 `json:"message"`
	Data       map[string]interface{} `json:"data"`
}

type DashboardDTO struct {
	CategorysDTO []categoryDTOPackage.Category `json:"categorysDTO"`
	HasPaid      bool                          `json:"hasPaid"`
	CategoryId   string                        `json:"categoryId"`
}
