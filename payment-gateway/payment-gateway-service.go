package paymentgateway

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/GabrielAchumba/school-mgt-backend/common/config"
	errors "github.com/GabrielAchumba/school-mgt-backend/common/errors"
	paystack "github.com/rpip/paystack-go"
)

type PaymentGatewayService interface {
	GetBanks() []paystack.Bank
	InitiateTransfer(request PaymentGatewayRequest) (interface{}, error)
	FinalizeTransfer(request PaymentGatewayRequest) (interface{}, error)
	InitializePayment(request PaymentGatewayRequest) (interface{}, error)
}

type serviceImpl struct {
	ctx    context.Context
	client *paystack.Client
	apiKey string
}

func New(ctx context.Context, config config.Settings) PaymentGatewayService {
	client := paystack.NewClient(config.PayStackKey.TestKey, nil)

	return &serviceImpl{
		ctx:    ctx,
		client: client,
		apiKey: config.PayStackKey.TestKey,
	}
}

func (impl serviceImpl) GetBanks() []paystack.Bank {

	log.Print("GetBanks called")
	bankList, err := impl.client.Bank.List()
	if err != nil {
		return make([]paystack.Bank, 0)
	}

	log.Print("GetBanks completed")
	return bankList.Values

}

func (impl serviceImpl) InitiateTransfer(request PaymentGatewayRequest) (interface{}, error) {

	/* res, err := impl.client.Transfer.EnableOTP()
	fmt.Println(res)
	if err != nil {
		return nil, errors.Error(err.Error())
	} */

	recipient := &paystack.TransferRecipient{
		Type:          "Nuban",
		Name:          request.AccountName,
		Description:   request.Description,
		AccountNumber: request.AccountNumber,
		BankCode:      request.BankCode,
		Currency:      request.Currency,
		Metadata:      map[string]interface{}{"job": "Plumber"},
	}

	recipient1, err := impl.client.Transfer.CreateRecipient(recipient)
	if err != nil {
		return nil, errors.Error(err.Error())
	}

	amount, _ := strconv.ParseFloat(request.Amount, 32)

	req := &paystack.TransferRequest{
		Source:    "balance",
		Reason:    request.Reason,
		Amount:    float32(amount) * 100,
		Recipient: recipient1.RecipientCode,
		Currency:  request.Currency,
	}

	transfer, err := impl.client.Transfer.Initiate(req)

	if err != nil {
		return nil, errors.Error(err.Error())
	}

	if transfer.TransferCode == "" {
		return nil, errors.Error("Expected transfer code, got " + transfer.TransferCode)
	}

	// fetch transfer
	trf, err := impl.client.Transfer.Get(transfer.TransferCode)
	if err != nil {
		return nil, errors.Error("Could not fetch transfer details")
	}

	if trf.TransferCode == "" {
		return nil, errors.Error("Expected transfer code, got " + trf.TransferCode)
	}

	transfer.TransferCode = trf.TransferCode
	return transfer, nil
}

func (impl serviceImpl) FinalizeTransfer(request PaymentGatewayRequest) (interface{}, error) {

	response, err := impl.client.Transfer.Finalize(request.TransferCode, request.OTP)

	if err != nil {
		return nil, errors.Error(err.Error())
	}

	return response, nil
}

func (impl serviceImpl) makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func (impl serviceImpl) InitializePayment(request PaymentGatewayRequest) (interface{}, error) {

	txn := &paystack.TransactionRequest{
		Email:     "user123@gmail.com",
		Amount:    6000,
		Reference: "Txn-" + fmt.Sprintf("%d", impl.makeTimestamp()),
	}
	resp, err := impl.client.Transaction.Initialize(txn)
	if err != nil {
		return nil, errors.Error(err.Error())
	}

	if resp["authorization_code"] == "" {
		return nil, errors.Error("Missing transaction authorization code")
	}

	if resp["access_code"] == "" {
		return nil, errors.Error("Missing transaction access code")
	}

	if resp["reference"] == "" {
		return nil, errors.Error("Missing transaction reference")
	}

	txn1, err := impl.client.Transaction.Verify(resp["reference"].(string))

	if err != nil {
		return nil, errors.Error(err.Error())
	}

	if txn1.Amount != txn.Amount {
		x := fmt.Sprintf("Expected transaction amount %f, got %+v", txn.Amount, txn1.Amount)
		return nil, errors.Error(x)
	}

	if txn1.Reference == "" {
		return nil, errors.Error("Missing transaction reference")
	}

	reponse, err := impl.client.Transaction.Get(txn1.ID)

	if err != nil {
		return nil, errors.Error(err.Error())
	}

	return reponse, nil
}
