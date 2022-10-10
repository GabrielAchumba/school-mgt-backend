package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/GabrielAchumba/school-mgt-backend/common/config"
	"github.com/GabrielAchumba/school-mgt-backend/common/conversion"
	errors "github.com/GabrielAchumba/school-mgt-backend/common/errors"
	"github.com/GabrielAchumba/school-mgt-backend/modules/school-module/dtos"
	"github.com/GabrielAchumba/school-mgt-backend/modules/school-module/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SchoolService interface {
	CreateSchool(userId string, requestModel dtos.CreateSchoolRequest) (interface{}, error)
	DeleteSchool(id string) (int64, error)
	GetSchool(id string) (dtos.SchoolResponse, error)
	GetSchools() ([]dtos.SchoolResponse, error)
	PutSchool(id string, item dtos.UpdateSchoolRequest) (interface{}, error)
}

type serviceImpl struct {
	ctx        context.Context
	collection *mongo.Collection
}

func New(mongoClient *mongo.Client, config config.Settings, ctx context.Context) SchoolService {
	collection := mongoClient.Database(config.Database.DatabaseName).Collection(config.TableNames.School)

	return &serviceImpl{
		collection: collection,
		ctx:        ctx,
	}
}

func (impl serviceImpl) DeleteSchool(id string) (int64, error) {

	log.Print("Call to delete type of School by id started.")
	objId := conversion.GetMongoId(id)
	filter := bson.D{bson.E{Key: "_id", Value: objId}}

	result, err := impl.collection.DeleteOne(impl.ctx, filter)

	if err != nil {
		return result.DeletedCount, errors.Error("Error in deleting type of School.")
	}

	if result.DeletedCount < 1 {
		return result.DeletedCount, errors.Error("Type of School with specified ID not found!")
	}

	log.Print("Call to delete type of School by id completed.")
	return result.DeletedCount, nil
}

func (impl serviceImpl) GetSchool(id string) (dtos.SchoolResponse, error) {

	log.Print("Get Type of School called")
	objId := conversion.GetMongoId(id)
	var School dtos.SchoolResponse

	filter := bson.D{bson.E{Key: "_id", Value: objId}}

	err := impl.collection.FindOne(impl.ctx, filter).Decode(&School)
	if err != nil {
		return School, errors.Error("could not find type of School by id")
	}

	log.Print("Get type of School completed")
	return School, err

}

func (impl serviceImpl) GetSchools() ([]dtos.SchoolResponse, error) {

	log.Print("Call to get all types of School started.")

	var Schools []dtos.SchoolResponse
	filter := bson.D{}
	cur, err := impl.collection.Find(impl.ctx, filter)

	if err != nil {
		Schools = make([]dtos.SchoolResponse, 0)
		return Schools, errors.Error("Types of School not found!")
	}

	err = cur.All(impl.ctx, &Schools)
	if err != nil {
		return Schools, err
	}

	cur.Close(impl.ctx)
	if len(Schools) == 0 {
		Schools = make([]dtos.SchoolResponse, 0)
	}

	log.Print("Call to get all types of School completed.")
	return Schools, err
}

func (impl serviceImpl) CreateSchool(userId string, model dtos.CreateSchoolRequest) (interface{}, error) {

	log.Print("Call to create School started.")

	var modelObj models.School
	conversion.Convert(model, &modelObj)

	modelObj.CreatedBy = userId
	modelObj.CreatedAt = time.Now()

	if modelObj.SchoolName == "" {
		return nil, errors.Error("School cannot be empty.")
	}

	filter := bson.D{bson.E{Key: "type", Value: modelObj.SchoolName}}
	count, err := impl.collection.CountDocuments(impl.ctx, filter)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.Error(fmt.Sprintf("Type of School '%v'already exist.", model.SchoolName))
	}
	_, er := impl.collection.InsertOne(impl.ctx, modelObj)
	if er != nil {
		return nil, errors.Error("Error in creating School.")
	}
	log.Print("Call to create School completed.")
	return modelObj, er
}

func (impl serviceImpl) PutSchool(id string, User dtos.UpdateSchoolRequest) (interface{}, error) {

	log.Print("PutSchool started")

	objId := conversion.GetMongoId(id)
	var updatedSchool dtos.UpdateSchoolRequest
	conversion.Convert(User, &updatedSchool)
	filter := bson.D{bson.E{Key: "_id", Value: objId}}
	var modelObj models.School

	update := bson.D{bson.E{Key: "schoolname", Value: updatedSchool.SchoolName},
		bson.E{Key: "address", Value: updatedSchool.Address}}
	_, err := impl.collection.UpdateOne(impl.ctx, filter, bson.D{bson.E{Key: "$set", Value: update}})

	if err != nil {
		return modelObj, errors.Error("Could not upadte type of School")
	}

	log.Print("PutSchool completed")
	return modelObj, nil
}
