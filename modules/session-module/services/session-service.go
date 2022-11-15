package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/GabrielAchumba/school-mgt-backend/common/config"
	"github.com/GabrielAchumba/school-mgt-backend/common/conversion"
	errors "github.com/GabrielAchumba/school-mgt-backend/common/errors"
	"github.com/GabrielAchumba/school-mgt-backend/modules/session-module/dtos"
	"github.com/GabrielAchumba/school-mgt-backend/modules/session-module/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SessionService interface {
	CreateSession(userId string, requestModel dtos.CreateSessionRequest) (interface{}, error)
	CreateSessions(userId string, _models []dtos.CreateSessionRequest) (interface{}, error)
	DeleteSession(id string, schoolId string) (int64, error)
	GetSession(id string, schoolId string) (dtos.SessionResponse, error)
	GetSessions(schoolId string) ([]dtos.SessionResponse, error)
	PutSession(id string, item dtos.UpdateSessionRequest) (interface{}, error)
}

type serviceImpl struct {
	ctx        context.Context
	collection *mongo.Collection
}

func New(mongoClient *mongo.Client, config config.Settings, ctx context.Context) SessionService {
	collection := mongoClient.Database(config.Database.DatabaseName).Collection(config.TableNames.Session)

	return &serviceImpl{
		collection: collection,
		ctx:        ctx,
	}
}

func (impl serviceImpl) DeleteSession(id string, schoolId string) (int64, error) {

	log.Print("Call to delete type of Session by id started.")
	objId := conversion.GetMongoId(id)
	filter := bson.D{bson.E{Key: "_id", Value: objId},
		bson.E{Key: "schoolid", Value: schoolId}}

	result, err := impl.collection.DeleteOne(impl.ctx, filter)

	if err != nil {
		return result.DeletedCount, errors.Error("Error in deleting type of Session.")
	}

	if result.DeletedCount < 1 {
		return result.DeletedCount, errors.Error("Type of Session with specified ID not found!")
	}

	log.Print("Call to delete type of Session by id completed.")
	return result.DeletedCount, nil
}

func (impl serviceImpl) GetSession(id string, schoolId string) (dtos.SessionResponse, error) {

	log.Print("Get Type of Session called")
	objId := conversion.GetMongoId(id)
	var Session dtos.SessionResponse

	filter := bson.D{bson.E{Key: "_id", Value: objId},
		bson.E{Key: "schoolid", Value: schoolId}}

	err := impl.collection.FindOne(impl.ctx, filter).Decode(&Session)
	if err != nil {
		return Session, errors.Error("could not find type of Session by id")
	}

	log.Print("Get type of Session completed")
	return Session, err

}

func (impl serviceImpl) GetSessions(schoolId string) ([]dtos.SessionResponse, error) {

	log.Print("Call to get all types of Session started.")

	var Sessions []dtos.SessionResponse
	filter := bson.D{bson.E{Key: "schoolid", Value: schoolId}}
	cur, err := impl.collection.Find(impl.ctx, filter)

	if err != nil {
		Sessions = make([]dtos.SessionResponse, 0)
		return Sessions, errors.Error("Types of Session not found!")
	}

	err = cur.All(impl.ctx, &Sessions)
	if err != nil {
		return Sessions, err
	}

	cur.Close(impl.ctx)
	if len(Sessions) == 0 {
		Sessions = make([]dtos.SessionResponse, 0)
	}

	log.Print("Call to get all types of Session completed.")
	return Sessions, err
}

func (impl serviceImpl) CreateSession(userId string, model dtos.CreateSessionRequest) (interface{}, error) {

	log.Print("Call to create Session started.")

	var modelObj models.Session
	conversion.Convert(model, &modelObj)

	modelObj.CreatedBy = userId
	modelObj.CreatedAt = time.Now()

	if modelObj.Type == "" {
		return nil, errors.Error("Type of Session cannot be empty.")
	}

	filter := bson.D{bson.E{Key: "type", Value: modelObj.Type},
		bson.E{Key: "schoolid", Value: modelObj.SchoolId}}
	count, err := impl.collection.CountDocuments(impl.ctx, filter)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.Error(fmt.Sprintf("Type of Session '%v'already exist.", model.Type))
	}
	_, er := impl.collection.InsertOne(impl.ctx, modelObj)
	if er != nil {
		return nil, errors.Error("Error in creating Session.")
	}
	log.Print("Call to create Session completed.")
	return modelObj, er
}

func (impl serviceImpl) CreateSessions(userId string, _models []dtos.CreateSessionRequest) (interface{}, error) {

	log.Print("Call to create Sessions started.")

	modelObjs := make([]interface{}, 0)
	for _, model := range _models {
		var modelObj models.Session
		modelObj.CreatedBy = userId
		modelObj.CreatedAt = time.Now()
		conversion.Convert(model, &modelObj)
		modelObjs = append(modelObjs, modelObj)
	}

	_, er := impl.collection.InsertMany(impl.ctx, modelObjs)
	if er != nil {
		return nil, errors.Error("Error in creating Sessions.")
	}
	log.Print("Call to create Sessions completed.")
	return modelObjs, er
}

func (impl serviceImpl) PutSession(id string, User dtos.UpdateSessionRequest) (interface{}, error) {

	log.Print("PutSession started")

	objId := conversion.GetMongoId(id)
	var updatedSession dtos.UpdateSessionRequest
	conversion.Convert(User, &updatedSession)
	filter := bson.D{bson.E{Key: "_id", Value: objId}}
	var modelObj models.Session

	update := bson.D{bson.E{Key: "type", Value: updatedSession.Type},
		bson.E{Key: "schoolid", Value: updatedSession.SchoolId}}
	_, err := impl.collection.UpdateOne(impl.ctx, filter, bson.D{bson.E{Key: "$set", Value: update}})

	if err != nil {
		return modelObj, errors.Error("Could not upadte type of Session")
	}

	log.Print("PutSession completed")
	return modelObj, nil
}
