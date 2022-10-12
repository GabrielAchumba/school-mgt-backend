package services

import (
	"context"
	"log"
	"time"

	"github.com/GabrielAchumba/school-mgt-backend/common/config"
	"github.com/GabrielAchumba/school-mgt-backend/common/conversion"
	errors "github.com/GabrielAchumba/school-mgt-backend/common/errors"
	"github.com/GabrielAchumba/school-mgt-backend/modules/payment-module/dtos"
	"github.com/GabrielAchumba/school-mgt-backend/modules/payment-module/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PaymentService interface {
	CreatePayment(userId string, requestModel dtos.CreatePaymentRequest) (interface{}, error)
	DeletePayment(id string, schoolId string) (int64, error)
	GetPayment(id string, schoolId string) (dtos.PaymentResponse, error)
	GetPayments(schoolId string) ([]dtos.PaymentResponse, error)
}

type serviceImpl struct {
	ctx        context.Context
	collection *mongo.Collection
}

func New(mongoClient *mongo.Client, config config.Settings, ctx context.Context) PaymentService {
	collection := mongoClient.Database(config.Database.DatabaseName).Collection(config.TableNames.Payment)

	return &serviceImpl{
		collection: collection,
		ctx:        ctx,
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

func (impl serviceImpl) GetPayment(id string, schoolId string) (dtos.PaymentResponse, error) {

	log.Print("Get Type of Payment called")
	objId := conversion.GetMongoId(id)
	var Payment dtos.PaymentResponse

	filter := bson.D{bson.E{Key: "_id", Value: objId},
		bson.E{Key: "schoolid", Value: schoolId}}

	err := impl.collection.FindOne(impl.ctx, filter).Decode(&Payment)
	if err != nil {
		return Payment, errors.Error("could not find type of Payment by id")
	}

	log.Print("Get type of Payment completed")
	return Payment, err

}

func (impl serviceImpl) GetPayments(schoolId string) ([]dtos.PaymentResponse, error) {

	log.Print("Call to get all types of Payment started.")

	var Payments []dtos.PaymentResponse
	filter := bson.D{bson.E{Key: "schoolid", Value: schoolId}}
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

	log.Print("Call to get all types of Payment completed.")
	return Payments, err
}

func (impl serviceImpl) CreatePayment(userId string, model dtos.CreatePaymentRequest) (interface{}, error) {

	log.Print("Call to create Payment started.")

	var modelObj []models.Payment
	var documents []interface{}
	conversion.Convert(model.CreatePayments, &modelObj)

	for i := 0; i < len(modelObj); i++ {
		modelObj[i].CreatedBy = userId
		modelObj[i].CreatedAt = time.Now()

		documents = append(documents, modelObj[i])
	}

	_, er := impl.collection.InsertMany(impl.ctx, documents)
	if er != nil {
		return nil, errors.Error("Error in creating Payments.")
	}
	log.Print("Call to create Payments completed.")
	return modelObj, er
}
