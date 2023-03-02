package dtos

type CashOutDTO struct {
	CashOutId             string  `json:"cashOutId"`
	CreatedDay            int     `json:"createdDay"`
	CreatedMonth          int     `json:"createdMonth"`
	CreatedYear           int     `json:"createdYear"`
	CreatedBy             string  `json:"createdBy"`
	CreatedDate           string  `json:"createdDate"`
	FullName              string  `json:"fullName"`
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

type CategoryBankDetails struct {
	CategoryBankName      string `json:"categoryBankName"`
	CategoryAccountName   string `json:"categoryAccountName"`
	CategoryAccountNumber string `json:"categoryAccountNumber"`
}
