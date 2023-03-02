package services

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/GabrielAchumba/school-mgt-backend/common/config"
	"github.com/GabrielAchumba/school-mgt-backend/common/conversion"
	errors "github.com/GabrielAchumba/school-mgt-backend/common/errors"
	"github.com/GabrielAchumba/school-mgt-backend/modules/payment-module/dtos"
	"github.com/GabrielAchumba/school-mgt-backend/modules/payment-module/models"
	paystack "github.com/rpip/paystack-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PaymentService interface {
	CreatePayment(userId string, requestModel dtos.CreatePaymentRequest) (interface{}, error)
	DeletePayment(id string, schoolId string) (int64, error)
	GetPayment(schoolId string) (dtos.PaymentResponse, error)
	GetPendingPayments() ([]dtos.PaymentResponse, error)
	CheckSubscription(schoolId string) (dtos.CheckSubscription, error)
	PutPayment(id string) (interface{}, error)
	GetBanks() ([]paystack.Bank, error)
	InitiateTransfer(request dtos.CreatePaymentRequest) (interface{}, error)
	FinalizeTransfer(request dtos.CreatePaymentRequest) (interface{}, error)
}

type serviceImpl struct {
	ctx        context.Context
	collection *mongo.Collection
	client     *paystack.Client
	apiKey     string
}

func New(mongoClient *mongo.Client, config config.Settings, ctx context.Context) PaymentService {
	collection := mongoClient.Database(config.Database.DatabaseName).Collection(config.TableNames.Payment)
	client := paystack.NewClient(config.PayStackKey.TestKey, nil)

	return &serviceImpl{
		collection: collection,
		ctx:        ctx,
		client:     client,
		apiKey:     config.PayStackKey.TestKey,
	}
}

func (impl serviceImpl) DeletePayment(id string, schoolId string) (int64, error) {

	log.Print("Call to delete type of Payment by id started.")
	objId := conversion.GetMongoId(id)
	filter := bson.D{bson.E{Key: "_id", Value: objId},
		bson.E{Key: "schoolid", Value: schoolId}}

	result, err := impl.collection.DeleteOne(impl.ctx, filter)

	if err != nil {
		return result.DeletedCount, errors.Error("Error in deleting type of Payment.")
	}

	if result.DeletedCount < 1 {
		return result.DeletedCount, errors.Error("Type of Payment with specified ID not found!")
	}

	log.Print("Call to delete type of Payment by id completed.")
	return result.DeletedCount, nil
}

func (impl serviceImpl) GetPayment(schoolId string) (dtos.PaymentResponse, error) {

	log.Print("Get Type of Payment called")
	var Payment dtos.PaymentResponse

	filter := bson.D{bson.E{Key: "schoolid", Value: schoolId}}

	err := impl.collection.FindOne(impl.ctx, filter).Decode(&Payment)
	if err != nil {
		return Payment, errors.Error("could not find type of Payment by id")
	}

	log.Print("Get type of Payment completed")
	return Payment, err

}

func (impl serviceImpl) CheckSubscription(schoolId string) (dtos.CheckSubscription, error) {

	log.Print("CheckSubscription called")
	var Payment dtos.PaymentResponse
	var CheckSubscription dtos.CheckSubscription = dtos.CheckSubscription{
		IsResultsAnalysis:   false,
		IsFileManagement:    false,
		IsAdvertizement:     false,
		IsExamsQuiz:         false,
		IsOnlineLearning:    false,
		IsLibraryManagement: false,
		IsSocialize:         false,
	}

	filter := bson.D{bson.E{Key: "schoolid", Value: schoolId}}

	err := impl.collection.FindOne(impl.ctx, filter).Decode(&Payment)
	if err != nil {
		return CheckSubscription, errors.Error("You have not subscribed for paid features")
	}

	if Payment.PaymentStatus.Value == "PENDING" {
		return CheckSubscription, errors.Error("Subscription is yet to be approved")
	}

	CheckSubscription.IsAdvertizement = true
	CheckSubscription.IsExamsQuiz = true
	CheckSubscription.IsFileManagement = true
	CheckSubscription.IsOnlineLearning = true
	CheckSubscription.IsResultsAnalysis = true
	CheckSubscription.IsLibraryManagement = true
	CheckSubscription.IsSocialize = false

	today := time.Now()
	days := today.Sub(Payment.CreatedAt).Hours() / 24

	switch Payment.ResultSubscription.Variable {
	case "None":
		CheckSubscription.IsResultsAnalysis = false
	case "Results Analysis (90 Days)":
		if days > 90 {
			CheckSubscription.IsResultsAnalysis = false
		}
	case "Results Analysis (180 Days)":
		if days > 180 {
			CheckSubscription.IsResultsAnalysis = false
		}
	case "Results Analysis (360 Days)":
		if days > 360 {
			CheckSubscription.IsResultsAnalysis = false
		}
	}

	switch Payment.FileManagementSubscription.Variable {
	case "None":
		CheckSubscription.IsFileManagement = false
	case "File Management 10GB (Quaterly)":
		if days > 90 {
			CheckSubscription.IsFileManagement = false
		}
	case "File Management 20GB (Quaterly)":
		if days > 90 {
			CheckSubscription.IsFileManagement = false
		}
	case "File Management 80GB (Quaterly)":
		if days > 90 {
			CheckSubscription.IsFileManagement = false
		}
	case "File Management 10GB (Semi-Annually)":
		if days > 180 {
			CheckSubscription.IsFileManagement = false
		}
	case "File Management 20GB (Semi-Annually)":
		if days > 180 {
			CheckSubscription.IsFileManagement = false
		}
	case "File Management 80GB (Semi-Annually)":
		if days > 180 {
			CheckSubscription.IsFileManagement = false
		}
	case "File Management 10GB (Annually)":
		if days > 360 {
			CheckSubscription.IsFileManagement = false
		}
	case "File Management 20GB (Annually)":
		if days > 360 {
			CheckSubscription.IsFileManagement = false
		}
	case "File Management 80GB (Annually)":
		if days > 360 {
			CheckSubscription.IsFileManagement = false
		}
	}

	switch Payment.AppCustomizationSubscription.Variable {
	case "None":
		CheckSubscription.IsAdvertizement = false
	case "Branding & Advertisement 10GB (Quaterly)":
		if days > 90 {
			CheckSubscription.IsAdvertizement = false
		}
	case "Branding & Advertisement 20GB (Quaterly)":
		if days > 90 {
			CheckSubscription.IsAdvertizement = false
		}
	case "Branding & Advertisement 80GB (Quaterly)":
		if days > 90 {
			CheckSubscription.IsAdvertizement = false
		}
	case "Branding & Advertisement 10GB (Semi-Annually)":
		if days > 180 {
			CheckSubscription.IsAdvertizement = false
		}
	case "Branding & Advertisement 20GB (Semi-Annually)":
		if days > 180 {
			CheckSubscription.IsAdvertizement = false
		}
	case "Branding & Advertisement 80GB (Semi-Annually)":
		if days > 180 {
			CheckSubscription.IsAdvertizement = false
		}
	case "Branding & Advertisement 10GB (Annually)":
		if days > 360 {
			CheckSubscription.IsAdvertizement = false
		}
	case "Branding & Advertisement 20GB (Annually)":
		if days > 360 {
			CheckSubscription.IsAdvertizement = false
		}
	case "Branding & Advertisement 80GB (Annually)":
		if days > 360 {
			CheckSubscription.IsAdvertizement = false
		}
	}

	switch Payment.LibraryManagementSubscription.Variable {
	case "None":
		CheckSubscription.IsLibraryManagement = false
	case "Library Management 10GB (Quaterly)":
		if days > 90 {
			CheckSubscription.IsLibraryManagement = false
		}
	case "Library Management 20GB (Quaterly)":
		if days > 90 {
			CheckSubscription.IsLibraryManagement = false
		}
	case "Library Management 80GB (Quaterly)":
		if days > 90 {
			CheckSubscription.IsLibraryManagement = false
		}
	case "Library Management 10GB (Semi-Annually)":
		if days > 180 {
			CheckSubscription.IsLibraryManagement = false
		}
	case "Library Management 20GB (Semi-Annually)":
		if days > 180 {
			CheckSubscription.IsLibraryManagement = false
		}
	case "Library Management 80GB (Semi-Annually)":
		if days > 180 {
			CheckSubscription.IsLibraryManagement = false
		}
	case "Library Management 10GB (Annually)":
		if days > 360 {
			CheckSubscription.IsLibraryManagement = false
		}
	case "Library Management 20GB (Annually)":
		if days > 360 {
			CheckSubscription.IsLibraryManagement = false
		}
	case "Library Management 80GB (Annually)":
		if days > 360 {
			CheckSubscription.IsLibraryManagement = false
		}
	}

	if CheckSubscription.IsAdvertizement {
		CheckSubscription.IsSocialize = true
	}

	if CheckSubscription.IsExamsQuiz {
		CheckSubscription.IsSocialize = true
	}

	if CheckSubscription.IsFileManagement {
		CheckSubscription.IsSocialize = true
	}

	if CheckSubscription.IsOnlineLearning {
		CheckSubscription.IsSocialize = true
	}

	if CheckSubscription.IsLibraryManagement {
		CheckSubscription.IsSocialize = true
	}

	log.Print("CheckSubscription completed")
	return CheckSubscription, err

}

func (impl serviceImpl) GetPendingPayments() ([]dtos.PaymentResponse, error) {

	log.Print("GetPendingPayments started.")

	var Payments []dtos.PaymentResponse
	filter := bson.D{}
	cur, err := impl.collection.Find(impl.ctx, filter)

	if err != nil {
		Payments = make([]dtos.PaymentResponse, 0)
		return Payments, errors.Error("Types of Payment not found!")
	}

	err = cur.All(impl.ctx, &Payments)
	if err != nil {
		return Payments, err
	}

	cur.Close(impl.ctx)
	if len(Payments) == 0 {
		Payments = make([]dtos.PaymentResponse, 0)
	}

	pendingPayments := make([]dtos.PaymentResponse, 0)
	for _, payment := range Payments {
		if payment.PaymentStatus.Value == "PENDING" {
			pendingPayments = append(pendingPayments, payment)
		}
	}

	log.Print("GetPendingPayments completed.")
	return pendingPayments, err
}

func (impl serviceImpl) CreatePayment(userId string, model dtos.CreatePaymentRequest) (interface{}, error) {

	log.Print("Call to create Payment started.")

	var modelObj models.Payment
	conversion.Convert(model, &modelObj)

	modelObj.CreatedBy = userId
	modelObj.CreatedAt = time.Now()
	hours := modelObj.CreatedAt.Hour()
	minutes := modelObj.CreatedAt.Minute()
	seconds := modelObj.CreatedAt.Second()
	modelObj.CreatedTime = strconv.Itoa(hours) + ":" + strconv.Itoa(minutes) + ":" + strconv.Itoa(seconds)
	modelObj.PaymentStatus.Value = "PENDING"
	modelObj.PaymentMessage.Value = "Please wait for our administartion team to verify and approve your payment. It takes less than 24 hours. Thanks"
	min := 1234567
	max := 123456789
	modelObj.ReceiptNo.Value = strconv.Itoa(rand.Intn(max-min) + min)

	filter := bson.D{bson.E{Key: "schoolid", Value: modelObj.SchoolId}}
	count, err := impl.collection.CountDocuments(impl.ctx, filter)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.Error(fmt.Sprintf("Subscription already exist."))
	}

	_, er := impl.collection.InsertOne(impl.ctx, modelObj)
	if er != nil {
		return nil, errors.Error("Error in creating Payment.")
	}
	log.Print("Call to create Payment completed.")
	return modelObj, er
}

func (impl serviceImpl) PutPayment(id string) (interface{}, error) {

	log.Print("PutPayment started")
	objId := conversion.GetMongoId(id)
	filter := bson.D{bson.E{Key: "_id", Value: objId}}
	var modelObj models.Payment
	modelObj.CreatedAt = time.Now()

	update := bson.D{bson.E{Key: "paymentstatus.value", Value: "APPROVED"},
		bson.E{Key: "createdat", Value: modelObj.CreatedAt}}
	_, err := impl.collection.UpdateOne(impl.ctx, filter, bson.D{bson.E{Key: "$set", Value: update}})

	if err != nil {
		return modelObj, errors.Error("Could not upadte type of staff")
	}

	log.Print("PutStaff completed")
	return modelObj, nil
}

func (impl serviceImpl) GetBanks() ([]paystack.Bank, error) {

	log.Print("GetBanks called")
	bankList, err := impl.client.Bank.List()
	if err != nil {
		return make([]paystack.Bank, 0), errors.Error("could not fetch list of banks")
	}

	log.Print("GetBanks completed")
	return bankList.Values, err

}

func (impl serviceImpl) InitiateTransfer(request dtos.CreatePaymentRequest) (interface{}, error) {

	impl.client.Transfer.EnableOTP()

	recipient := &paystack.TransferRecipient{
		Type:          "Nuban",
		Name:          request.AccountNamePaidFrom,
		Description:   request.Description,
		AccountNumber: request.AccountNumberPaidFrom,
		BankCode:      request.BankCode,
		Currency:      request.CurrencyCode,
		Metadata:      map[string]interface{}{"Designation": request.Designation},
	}

	recipient1, _ := impl.client.Transfer.CreateRecipient(recipient)

	amount, _ := strconv.ParseFloat(request.AccountNamePaidFrom, 32)

	req := &paystack.TransferRequest{
		Source:    "balance",
		Reason:    request.Reason,
		Amount:    float32(amount),
		Recipient: recipient1.RecipientCode,
	}

	transfer, err := impl.client.Transfer.Initiate(req)

	if err != nil {
		return nil, errors.Error("Could not initiate transfer")
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

	return transfer, nil
}

func (impl serviceImpl) FinalizeTransfer(request dtos.CreatePaymentRequest) (interface{}, error) {

	response, err := impl.client.Transfer.Finalize(request.TransferCode, request.OTP)

	if err != nil {
		return nil, errors.Error("Could not initiate transfer")
	}

	return response, nil
}
