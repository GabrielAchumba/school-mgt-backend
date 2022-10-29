package dtos

import (
	"time"

	"github.com/GabrielAchumba/school-mgt-backend/modules/payment-module/models"
)

type CreatePaymentRequest struct {
	SchoolId                     string              `json:"schoolId"`
	StudentIds                   []string            `json:"studentIds"`
	Message                      string              `json:"message"`
	Reference                    string              `json:"reference"`
	Status                       string              `json:"status"`
	Trans                        string              `json:"trans"`
	Transactions                 string              `json:"transactions"`
	Trxref                       string              `json:"trxref"`
	ResultSubscription           models.Subscription `json:"resultSubscription"`
	ExamsAndQuizSubscription     models.Subscription `json:"examsAndQuizSubscription"`
	FileManagementSubscription   models.Subscription `json:"fileManagementSubscription"`
	AppCustomizationSubscription models.Subscription `json:"appCustomizationSubscription"`
	OnLineLearningSubscription   models.Subscription `json:"onLineLearningSubscription"`
	AllSubscriptions             models.Subscription `json:"allSubscriptions"`
	PaymentStatus                models.Subscription `json:"paymentStatus"`
	PaymentMessage               models.Subscription `json:"paymentMessage"`
	ReceiptNo                    models.Subscription `json:"receiptNo"`
	BankNamePaidFrom             string              `json:"bankNamePaidFrom"`
	AccountNamePaidFrom          string              `json:"accountNamePaidFrom"`
	AccountNumberPaidFrom        string              `json:"accountNumberPaidFrom"`
	Base64String                 string              `json:"base64String"`
	FileName                     string              `json:"fileName"`
}

type UpdatePaymentRequest struct {
	SchoolId                     string              `json:"schoolId"`
	StudentIds                   []string            `json:"studentIds"`
	Message                      string              `json:"message"`
	Reference                    string              `json:"reference"`
	Status                       string              `json:"status"`
	Trans                        string              `json:"trans"`
	Transactions                 string              `json:"transactions"`
	Trxref                       string              `json:"trxref"`
	ResultSubscription           models.Subscription `json:"resultSubscription"`
	ExamsAndQuizSubscription     models.Subscription `json:"examsAndQuizSubscription"`
	FileManagementSubscription   models.Subscription `json:"fileManagementSubscription"`
	AppCustomizationSubscription models.Subscription `json:"appCustomizationSubscription"`
	OnLineLearningSubscription   models.Subscription `json:"onLineLearningSubscription"`
	AllSubscriptions             models.Subscription `json:"allSubscriptions"`
	PaymentStatus                models.Subscription `json:"paymentStatus"`
	PaymentMessage               models.Subscription `json:"paymentMessage"`
	ReceiptNo                    models.Subscription `json:"receiptNo"`
	BankNamePaidFrom             string              `json:"bankNamePaidFrom"`
	AccountNamePaidFrom          string              `json:"accountNamePaidFrom"`
	AccountNumberPaidFrom        string              `json:"accountNumberPaidFrom"`
	Base64String                 string              `json:"base64String"`
	FileName                     string              `json:"fileName"`
}

type PaymentResponse struct {
	Id                           string              `json:"id"  bson:"_id"`
	CreatedAt                    time.Time           `json:"createdAt"`
	CreatedBy                    string              `json:"createdBy"`
	SchoolId                     string              `json:"schoolId"`
	StudentIds                   []string            `json:"studentIds"`
	Message                      string              `json:"message"`
	Reference                    string              `json:"reference"`
	Status                       string              `json:"status"`
	Trans                        string              `json:"trans"`
	Transactions                 string              `json:"transactions"`
	Trxref                       string              `json:"trxref"`
	ResultSubscription           models.Subscription `json:"resultSubscription"`
	ExamsAndQuizSubscription     models.Subscription `json:"examsAndQuizSubscription"`
	FileManagementSubscription   models.Subscription `json:"fileManagementSubscription"`
	AppCustomizationSubscription models.Subscription `json:"appCustomizationSubscription"`
	OnLineLearningSubscription   models.Subscription `json:"onLineLearningSubscription"`
	AllSubscriptions             models.Subscription `json:"allSubscriptions"`
	PaymentStatus                models.Subscription `json:"paymentStatus"`
	PaymentMessage               models.Subscription `json:"paymentMessage"`
	ReceiptNo                    models.Subscription `json:"receiptNo"`
	BankNamePaidFrom             string              `json:"bankNamePaidFrom"`
	AccountNamePaidFrom          string              `json:"accountNamePaidFrom"`
	AccountNumberPaidFrom        string              `json:"accountNumberPaidFrom"`
	Base64String                 string              `json:"base64String"`
	FileName                     string              `json:"fileName"`
}

type CheckSubscription struct {
	IsResultsAnalysis bool `json:"isResultsAnalysis"`
	IsFileManagement  bool `json:"isFileManagement"`
	IsAdvertizement   bool `json:"isAdvertizement"`
	IsExamsQuiz       bool `json:"isExamsQuiz"`
	IsOnlineLearning  bool `json:"isOnlineLearning"`
}
