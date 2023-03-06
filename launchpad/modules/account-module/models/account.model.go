package models

type Account struct {
	Id           string `json:"id" bson:"_id"`
	CreatedDay   int    `json:"createdDay"`
	CreatedMonth int    `json:"createdMonth"`
	CreatedYear  int    `json:"createdYear"`

	ParentId              string  `json:"parentId"`
	ContributorId         string  `json:"contributorId"`
	Contribution          float64 `json:"contribution"`
	Message               string  `json:"message"`
	Reference             string  `json:"reference"`
	Status                string  `json:"status"`
	Trans                 string  `json:"trans"`
	TransferCode          string  `json:"transferCode"`
	TransferId            string  `json:"transferId"`
	Integration           string  `json:"integration"`
	Recipient             string  `json:"recipient"`
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
}
