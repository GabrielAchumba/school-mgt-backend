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
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type StaffService interface {
	CreateStaff(userId string, requestModel dtos.CreateStaffRequest) (interface{}, error)
	CreateManyStaff(userId string, _models []dtos.CreateStaffRequest) (interface{}, error)
	DeleteStaff(id string, schoolId string) (int64, error)
	GetStaff(id string, schoolId string) (dtos.StaffResponse, error)
	GetStaffs(schoolId string) ([]dtos.StaffResponse, error)
	GetStaffsByIds(schoolId string, Ids []string) ([]dtos.StaffResponse, error)
	PutStaff(id string, item dtos.UpdateStaffRequest) (interface{}, error)
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

func (impl serviceImpl) DeleteStaff(id string, schoolId string) (int64, error) {

	log.Print("Call to delete type of staff by id started.")
	objId := conversion.GetMongoId(id)
	filter := bson.D{bson.E{Key: "_id", Value: objId},
		bson.E{Key: "schoolid", Value: schoolId}}

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

func (impl serviceImpl) GetStaff(id string, schoolId string) (dtos.StaffResponse, error) {

	log.Print("Get Type of Staff called")
	objId := conversion.GetMongoId(id)
	var Staff dtos.StaffResponse

	filter := bson.D{bson.E{Key: "_id", Value: objId},
		bson.E{Key: "schoolid", Value: schoolId}}

	err := impl.collection.FindOne(impl.ctx, filter).Decode(&Staff)
	if err != nil {
		return Staff, errors.Error("could not find type of staff by id")
	}

	log.Print("Get type of staff completed")
	return Staff, err

}

func (impl serviceImpl) GetStaffs(schoolId string) ([]dtos.StaffResponse, error) {

	log.Print("Call to get all types of staff started.")

	var Staffs []dtos.StaffResponse
	filter := bson.D{bson.E{Key: "schoolid", Value: schoolId}}
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

func (impl serviceImpl) GetStaffsByIds(schoolId string, Ids []string) ([]dtos.StaffResponse, error) {

	log.Print("Call to get GetStaffsByIds started.")

	var objIds = make([]primitive.ObjectID, 0)
	for _, id := range Ids {
		objIds = append(objIds, conversion.GetMongoId(id))
	}

	var staffs []dtos.StaffResponse
	filter := bson.D{bson.E{Key: "schoolid", Value: schoolId},
		bson.E{Key: "_id", Value: bson.D{bson.E{Key: "$in", Value: objIds}}}}

	cur, err := impl.collection.Find(impl.ctx, filter)

	if err != nil {
		staffs = make([]dtos.StaffResponse, 0)
		return staffs, errors.Error("Staffs not found!")
	}

	err = cur.All(impl.ctx, &staffs)
	if err != nil {
		return staffs, err
	}

	cur.Close(impl.ctx)
	if len(staffs) == 0 {
		staffs = make([]dtos.StaffResponse, 0)
	}

	log.Print("Call to get staffs by Ids completed.")
	return staffs, err
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

	filter := bson.D{bson.E{Key: "type", Value: modelObj.Type},
		bson.E{Key: "schoolid", Value: modelObj.SchoolId}}

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

func (impl serviceImpl) CreateManyStaff(userId string, _models []dtos.CreateStaffRequest) (interface{}, error) {

	log.Print("Call to create many staff started.")

	types := make([]string, 0)
	var staffs []dtos.StaffResponse
	for _, model := range _models {
		types = append(types, model.Type)
	}

	filter := bson.D{{Key: "type", Value: bson.D{
		bson.E{Key: "$in", Value: types}}}}

	cur, _ := impl.collection.Find(impl.ctx, filter)

	_ = cur.All(impl.ctx, &staffs)
	cur.Close(impl.ctx)

	modelObjs := make([]interface{}, 0)
	for _, model := range _models {
		var modelObj models.Staff
		modelObj.CreatedBy = userId
		modelObj.CreatedAt = time.Now()
		check := false
		for _, staff := range staffs {
			if model.Type == staff.Type {
				check = true
				break
			}
		}

		if !check {
			conversion.Convert(model, &modelObj)
			modelObjs = append(modelObjs, modelObj)
		}
	}

	_, er := impl.collection.InsertMany(impl.ctx, modelObjs)
	if er != nil {
		return nil, errors.Error("Error in creating many staff.")
	}
	log.Print("Call to create many staff completed.")
	return modelObjs, er
}

func (impl serviceImpl) PutStaff(id string, item dtos.UpdateStaffRequest) (interface{}, error) {

	log.Print("PutStaff started")
	objId := conversion.GetMongoId(id)
	var updatedStaff dtos.UpdateStaffRequest
	conversion.Convert(item, &updatedStaff)
	filter := bson.D{bson.E{Key: "_id", Value: objId}}
	var modelObj models.Staff

	update := bson.D{bson.E{Key: "type", Value: updatedStaff.Type},
		bson.E{Key: "schoolid", Value: updatedStaff.SchoolId}}
	_, err := impl.collection.UpdateOne(impl.ctx, filter, bson.D{bson.E{Key: "$set", Value: update}})

	if err != nil {
		return modelObj, errors.Error("Could not upadte type of staff")
	}

	log.Print("PutStaff completed")
	return modelObj, nil
}
