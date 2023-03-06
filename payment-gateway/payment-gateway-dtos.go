package paymentgateway

type PaymentGatewayRequest struct {
	AccountName   string                 `json:"accountName,omitempty"`
	Description   string                 `json:"description,omitempty"`
	AccountNumber string                 `json:"accountNumber,omitempty"`
	BankCode      string                 `json:"bankCode,omitempty"`
	Currency      string                 `json:"currency,omitempty"`
	Metadata      map[string]interface{} `json:"metadata,omitempty"`
	Amount        string                 `json:"amount,omitempty"`
	Reason        string                 `json:"reason,omitempty"`
	TransferCode  string                 `json:"transferCode,omitempty"`
	OTP           string                 `json:"oTP,omitempty"`
}

func NewPaymentGatewayRequest() PaymentGatewayRequest {
	return PaymentGatewayRequest{
		Metadata: make(map[string]interface{}),
	}
}
