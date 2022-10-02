package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/GabrielAchumba/school-mgt-backend/common/config"
	"github.com/GabrielAchumba/school-mgt-backend/common/conversion"
	errors "github.com/GabrielAchumba/school-mgt-backend/common/errors"
	"github.com/GabrielAchumba/school-mgt-backend/modules/staff-module/dtos"
	"github.com/GabrielAchumba/school-mgt-backend/modules/staff-module/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type StaffService interface {
	CreateStaff(userId string, requestModel dtos.CreateStaffRequest) (interface{}, error)
	DeleteStaff(id string) (int64, error)
	GetStaff(id string) (dtos.StaffResponse, error)
	GetStaffs() ([]dtos.StaffResponse, error)
	PutStaff(id string, User dtos.UpdateStaffRequest) (interface{}, error)
}

type serviceImpl struct {
	ctx        context.Context
	collection *mongo.Collection
}

func New(mongoClient *mongo.Client, config config.Settings, ctx context.Context) StaffService {
	collection := mongoClient.Database(config.Database.DatabaseName).Collection(config.TableNames.Staff)

	return &serviceImpl{
		collection: collection,
		ctx:        ctx,
	}
}

func (impl serviceImpl) DeleteStaff(id string) (int64, error) {

	log.Print("Call to delete type of staff by id started.")
	objId := conversion.GetMongoId(id)
	filter := bson.D{bson.E{Key: "_id", Value: objId}}

	result, err := impl.collection.DeleteOne(impl.ctx, filter)

	if err != nil {
		return result.DeletedCount, errors.Error("Error in deleting type of staff.")
	}

	if result.DeletedCount < 1 {
		return result.DeletedCount, errors.Error("Type of staff with specified ID not found!")
	}

	log.Print("Call to delete type of staff by id completed.")
	return result.DeletedCount, nil
}

func (impl serviceImpl) GetStaff(id string) (dtos.StaffResponse, error) {

	log.Print("Get Type of Staff called")
	objId := conversion.GetMongoId(id)
	var Staff dtos.StaffResponse

	filter := bson.D{bson.E{Key: "_id", Value: objId}}

	err := impl.collection.FindOne(impl.ctx, filter).Decode(&Staff)
	if err != nil {
		return Staff, errors.Error("could not find type of staff by id")
	}

	log.Print("Get type of staff completed")
	return Staff, err

}

func (impl serviceImpl) GetStaffs() ([]dtos.StaffResponse, error) {

	log.Print("Call to get all types of staff started.")

	var Staffs []dtos.StaffResponse
	filter := bson.D{}
	cur, err := impl.collection.Find(impl.ctx, filter)

	if err != nil {
		Staffs = make([]dtos.StaffResponse, 0)
		return Staffs, errors.Error("Types of staff not found!")
	}

	err = cur.All(impl.ctx, &Staffs)
	if err != nil {
		return Staffs, err
	}

	cur.Close(impl.ctx)
	if len(Staffs) == 0 {
		Staffs = make([]dtos.StaffResponse, 0)
	}

	log.Print("Call to get all types of staff completed.")
	return Staffs, err
}

func (impl serviceImpl) CreateStaff(userId string, model dtos.CreateStaffRequest) (interface{}, error) {

	log.Print("Call to create staff started.")

	var modelObj models.Staff
	conversion.Convert(model, &modelObj)

	modelObj.CreatedBy = userId
	modelObj.CreatedAt = time.Now()

	if modelObj.Type == "" {
		return nil, errors.Error("Type of staff cannot be empty.")
	}

	filter := bson.D{bson.E{Key: "type", Value: modelObj.Type}}

	err := impl.collection.FindOne(impl.ctx, filter).Decode(&modelObj)
	if err == nil {
		return nil, errors.Error(fmt.Sprintf("Type of staff ('%v') already exist.", model.Type))
	}

	_, er := impl.collection.InsertOne(impl.ctx, modelObj)
	if er != nil {
		return nil, errors.Error("Error in creating staff.")
	}
	log.Print("Call to create staff completed.")
	return modelObj, er
}

func (impl serviceImpl) PutStaff(id string, item dtos.UpdateStaffRequest) (interface{}, error) {

	log.Print("PutStaff started")
	objId := conversion.GetMongoId(id)
	var updatedStaff dtos.UpdateStaffRequest
	conversion.Convert(item, &updatedStaff)
	filter := bson.D{bson.E{Key: "_id", Value: objId}}
	var modelObj models.Staff

	update := bson.D{bson.E{Key: "type", Value: updatedStaff.Type}}
	_, err := impl.collection.UpdateOne(impl.ctx, filter, bson.D{bson.E{Key: "$set", Value: update}})

	if err != nil {
		return modelObj, errors.Error("Could not upadte type of staff")
	}

	log.Print("PutStaff completed")
	return modelObj, nil
}
