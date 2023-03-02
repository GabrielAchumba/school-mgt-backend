package dtos

type BankAccountDTO struct {
	FullName      string `json:"fullname"`
	BankName      string `json:"bankName"`
	AccountName   string `json:"accountName"`
	AccountNumber string `json:"accountNumber"`
	BVN           string `json:"bVN"`
	ContributorId string `json:"contributorId"`
}
