package dtos

import "time"

type CreatePaymentRequest struct {
	CreatePayments []CreatePayment `json:"createPayments" binding:"required"`
}

type CreatePayment struct {
	SchoolId     string `json:"schoolId" binding:"required"`
	Message      string `json:"message" binding:"required"`
	Reference    string `json:"reference" binding:"required"`
	Status       string `json:"status" binding:"required"`
	Trans        string `json:"trans" binding:"required"`
	Transactions string `json:"transactions" binding:"required"`
	Trxref       string `json:"trxref" binding:"required"`
	TotalAmount  string `json:"totalAmount" binding:"required"`
}

type UpdatePaymentRequest struct {
	UpdatePayments []CreatePayment `json:"createPayments" binding:"required"`
}

type PaymentResponse struct {
	Id           string    `json:"id"  bson:"_id"`
	CreatedAt    time.Time `json:"createdAt"`
	CreatedBy    string    `json:"createdBy"`
	SchoolId     string    `json:"schoolId" binding:"required"`
	Message      string    `json:"message" binding:"required"`
	Reference    string    `json:"reference" binding:"required"`
	Status       string    `json:"status" binding:"required"`
	Trans        string    `json:"trans" binding:"required"`
	Transactions string    `json:"transactions" binding:"required"`
	Trxref       string    `json:"trxref" binding:"required"`
	TotalAmount  string    `json:"totalAmount" binding:"required"`
}
