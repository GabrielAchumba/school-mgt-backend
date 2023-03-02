package models

type CashOut struct {
	Id                    string  `json:"id" bson:"_id"`
	CreatedDay            int     `json:"createdDay"`
	CreatedMonth          int     `json:"createdMonth"`
	CreatedYear           int     `json:"createdYear"`
	CreatedBy             string  `json:"createdBy"`
	FullName              string  `json:"fullName"`
	Username              string  `json:"username"`
	Category              string  `json:"category"`
	Level                 int     `json:"level"`
	CategoryBankName      string  `json:"categoryBankName"`
	CategoryAccountName   string  `json:"categoryAccountName"`
	CategoryAccountNumber string  `json:"categoryAccountNumber"`
	BankName              string  `json:"bankName"`
	AccountName           string  `json:"accountName"`
	AccountNumber         string  `json:"accountNumber"`
	Base64String          string  `json:"base64String"`
	FileName              string  `json:"fileName"`
	CategoryId            string  `json:"categoryId"`
	ContributorId         string  `json:"contributorId"`
	ReturnOnInvestment    float64 `json:"returnOnInvestment"`
}
