package models

import "time"

type Subscription struct {
	Name     string  `json:"name"`
	Variable string  `json:"variable"`
	Value    string  `json:"value"`
	Amount   float64 `json:"amount"`
}

type Payment struct {
	CreatedAt                    time.Time    `json:"createdAt"`
	CreatedBy                    string       `json:"createdBy"`
	CreatedTime                  string       `json:"createdTime"`
	SchoolId                     string       `json:"schoolId"`
	StudentIds                   []string     `json:"studentIds"`
	Message                      string       `json:"message"`
	Reference                    string       `json:"reference"`
	Status                       string       `json:"status"`
	Trans                        string       `json:"trans"`
	Transactions                 string       `json:"transactions"`
	Trxref                       string       `json:"trxref"`
	ResultSubscription           Subscription `json:"resultSubscription"`
	ExamsAndQuizSubscription     Subscription `json:"examsAndQuizSubscription"`
	FileManagementSubscription   Subscription `json:"fileManagementSubscription"`
	AppCustomizationSubscription Subscription `json:"appCustomizationSubscription"`
	OnLineLearningSubscription   Subscription `json:"onLineLearningSubscription"`
	AllSubscriptions             Subscription `json:"allSubscriptions"`
	PaymentStatus                Subscription `json:"paymentStatus"`
	PaymentMessage               Subscription `json:"paymentMessage"`
	ReceiptNo                    Subscription `json:"receiptNo"`
	BankNamePaidFrom             string       `json:"bankNamePaidFrom"`
	AccountNamePaidFrom          string       `json:"accountNamePaidFrom"`
	AccountNumberPaidFrom        string       `json:"accountNumberPaidFrom"`
	Base64String                 string       `json:"base64String"`
	FileName                     string       `json:"fileName"`
}
