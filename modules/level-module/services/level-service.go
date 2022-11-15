package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/GabrielAchumba/school-mgt-backend/common/config"
	"github.com/GabrielAchumba/school-mgt-backend/common/conversion"
	errors "github.com/GabrielAchumba/school-mgt-backend/common/errors"
	"github.com/GabrielAchumba/school-mgt-backend/modules/level-module/dtos"
	"github.com/GabrielAchumba/school-mgt-backend/modules/level-module/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LevelService interface {
	CreateLevel(userId string, requestModel dtos.CreateLevelRequest) (interface{}, error)
	CreateLevels(userId string, _models []dtos.CreateLevelRequest) (interface{}, error)
	DeleteLevel(id string, schoolId string) (int64, error)
	GetLevel(id string, schoolId string) (dtos.LevelResponse, error)
	GetLevels(schoolId string) ([]dtos.LevelResponse, error)
	PutLevel(id string, item dtos.UpdateLevelRequest) (interface{}, error)
}

type serviceImpl struct {
	ctx        context.Context
	collection *mongo.Collection
}

func New(mongoClient *mongo.Client, config config.Settings, ctx context.Context) LevelService {
	collection := mongoClient.Database(config.Database.DatabaseName).Collection(config.TableNames.Level)

	return &serviceImpl{
		collection: collection,
		ctx:        ctx,
	}
}

func (impl serviceImpl) DeleteLevel(id string, schoolId string) (int64, error) {

	log.Print("Call to delete type of Level by id started.")
	objId := conversion.GetMongoId(id)
	filter := bson.D{bson.E{Key: "_id", Value: objId},
		bson.E{Key: "schoolid", Value: schoolId}}

	result, err := impl.collection.DeleteOne(impl.ctx, filter)

	if err != nil {
		return result.DeletedCount, errors.Error("Error in deleting type of Level.")
	}

	if result.DeletedCount < 1 {
		return result.DeletedCount, errors.Error("Type of Level with specified ID not found!")
	}

	log.Print("Call to delete type of Level by id completed.")
	return result.DeletedCount, nil
}

func (impl serviceImpl) GetLevel(id string, schoolId string) (dtos.LevelResponse, error) {

	log.Print("Get Type of Level called")
	objId := conversion.GetMongoId(id)
	var Level dtos.LevelResponse

	filter := bson.D{bson.E{Key: "_id", Value: objId},
		bson.E{Key: "schoolid", Value: schoolId}}

	err := impl.collection.FindOne(impl.ctx, filter).Decode(&Level)
	if err != nil {
		return Level, errors.Error("could not find type of Level by id")
	}

	log.Print("Get type of Level completed")
	return Level, err

}

func (impl serviceImpl) GetLevels(schoolId string) ([]dtos.LevelResponse, error) {

	log.Print("Call to get all types of Level started.")

	var Levels []dtos.LevelResponse
	filter := bson.D{bson.E{Key: "schoolid", Value: schoolId}}
	cur, err := impl.collection.Find(impl.ctx, filter)

	if err != nil {
		Levels = make([]dtos.LevelResponse, 0)
		return Levels, errors.Error("Types of Level not found!")
	}

	err = cur.All(impl.ctx, &Levels)
	if err != nil {
		return Levels, err
	}

	cur.Close(impl.ctx)
	if len(Levels) == 0 {
		Levels = make([]dtos.LevelResponse, 0)
	}

	log.Print("Call to get all types of Level completed.")
	return Levels, err
}

func (impl serviceImpl) CreateLevel(userId string, model dtos.CreateLevelRequest) (interface{}, error) {

	log.Print("Call to create Level started.")

	var modelObj models.Level
	conversion.Convert(model, &modelObj)

	modelObj.CreatedBy = userId
	modelObj.CreatedAt = time.Now()

	if modelObj.Type == "" {
		return nil, errors.Error("Type of Level cannot be empty.")
	}

	filter := bson.D{bson.E{Key: "type", Value: modelObj.Type},
		bson.E{Key: "schoolid", Value: modelObj.SchoolId}}
	count, err := impl.collection.CountDocuments(impl.ctx, filter)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.Error(fmt.Sprintf("Type of Level '%v'already exist.", model.Type))
	}
	_, er := impl.collection.InsertOne(impl.ctx, modelObj)
	if er != nil {
		return nil, errors.Error("Error in creating Level.")
	}
	log.Print("Call to create Level completed.")
	return modelObj, er
}

func (impl serviceImpl) CreateLevels(userId string, _models []dtos.CreateLevelRequest) (interface{}, error) {

	log.Print("Call to create class rooms started.")

	modelObjs := make([]interface{}, 0)
	for _, model := range _models {
		var modelObj models.Level
		modelObj.CreatedBy = userId
		modelObj.CreatedAt = time.Now()
		conversion.Convert(model, &modelObj)
		modelObjs = append(modelObjs, modelObj)
	}

	_, er := impl.collection.InsertMany(impl.ctx, modelObjs)
	if er != nil {
		return nil, errors.Error("Error in creating class rooms.")
	}
	log.Print("Call to create class rooms completed.")
	return modelObjs, er
}

func (impl serviceImpl) PutLevel(id string, item dtos.UpdateLevelRequest) (interface{}, error) {

	log.Print("PutStaff started")

	objId := conversion.GetMongoId(id)
	var updatedLevel dtos.UpdateLevelRequest
	conversion.Convert(item, &updatedLevel)
	filter := bson.D{bson.E{Key: "_id", Value: objId}}
	var modelObj models.Level

	update := bson.D{bson.E{Key: "type", Value: updatedLevel.Type},
		bson.E{Key: "schoolid", Value: updatedLevel.SchoolId}}
	_, err := impl.collection.UpdateOne(impl.ctx, filter, bson.D{bson.E{Key: "$set", Value: update}})

	if err != nil {
		return modelObj, errors.Error("Could not upadte type of staff")
	}

	log.Print("PutStaff completed")
	return modelObj, nil
}
