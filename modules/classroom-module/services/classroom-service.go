package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/GabrielAchumba/school-mgt-backend/common/config"
	"github.com/GabrielAchumba/school-mgt-backend/common/conversion"
	errors "github.com/GabrielAchumba/school-mgt-backend/common/errors"
	"github.com/GabrielAchumba/school-mgt-backend/modules/classroom-module/dtos"
	"github.com/GabrielAchumba/school-mgt-backend/modules/classroom-module/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ClassRoomService interface {
	CreateClassRoom(userId string, requestModel dtos.CreateClassRoomRequest) (interface{}, error)
	CreateClassRooms(userId string, _models []dtos.CreateClassRoomRequest) (interface{}, error)
	DeleteClassRoom(id string, schoolId string) (int64, error)
	GetClassRoom(id string, schoolId string) (dtos.ClassRoomResponse, error)
	GetClassRooms(schoolId string) ([]dtos.ClassRoomResponse, error)
	PutClassRoom(id string, item dtos.UpdateClassRoomRequest) (interface{}, error)
}

type serviceImpl struct {
	ctx        context.Context
	collection *mongo.Collection
}

func New(mongoClient *mongo.Client, config config.Settings, ctx context.Context) ClassRoomService {
	collection := mongoClient.Database(config.Database.DatabaseName).Collection(config.TableNames.ClassRoom)

	return &serviceImpl{
		collection: collection,
		ctx:        ctx,
	}
}

func (impl serviceImpl) DeleteClassRoom(id string, schoolId string) (int64, error) {

	log.Print("Call to delete type of ClassRoom by id started.")
	objId := conversion.GetMongoId(id)
	filter := bson.D{bson.E{Key: "_id", Value: objId},
		bson.E{Key: "schoolid", Value: schoolId}}

	result, err := impl.collection.DeleteOne(impl.ctx, filter)

	if err != nil {
		return result.DeletedCount, errors.Error("Error in deleting type of ClassRoom.")
	}

	if result.DeletedCount < 1 {
		return result.DeletedCount, errors.Error("Type of ClassRoom with specified ID not found!")
	}

	log.Print("Call to delete type of ClassRoom by id completed.")
	return result.DeletedCount, nil
}

func (impl serviceImpl) GetClassRoom(id string, schoolId string) (dtos.ClassRoomResponse, error) {

	log.Print("Get Type of ClassRoom called")
	objId := conversion.GetMongoId(id)
	var ClassRoom dtos.ClassRoomResponse

	filter := bson.D{bson.E{Key: "_id", Value: objId},
		bson.E{Key: "schoolid", Value: schoolId}}

	err := impl.collection.FindOne(impl.ctx, filter).Decode(&ClassRoom)
	if err != nil {
		return ClassRoom, errors.Error("could not find type of ClassRoom by id")
	}

	log.Print("Get type of ClassRoom completed")
	return ClassRoom, err

}

func (impl serviceImpl) GetClassRooms(schoolId string) ([]dtos.ClassRoomResponse, error) {

	log.Print("Call to get all types of ClassRoom started.")

	var ClassRooms []dtos.ClassRoomResponse
	filter := bson.D{bson.E{Key: "schoolid", Value: schoolId}}
	cur, err := impl.collection.Find(impl.ctx, filter)

	if err != nil {
		ClassRooms = make([]dtos.ClassRoomResponse, 0)
		return ClassRooms, errors.Error("Types of ClassRoom not found!")
	}

	err = cur.All(impl.ctx, &ClassRooms)
	if err != nil {
		return ClassRooms, err
	}

	cur.Close(impl.ctx)
	if len(ClassRooms) == 0 {
		ClassRooms = make([]dtos.ClassRoomResponse, 0)
	}

	log.Print("Call to get all types of ClassRoom completed.")
	return ClassRooms, err
}

func (impl serviceImpl) CreateClassRoom(userId string, model dtos.CreateClassRoomRequest) (interface{}, error) {

	log.Print("Call to create ClassRoom started.")

	var modelObj models.ClassRoom
	conversion.Convert(model, &modelObj)

	modelObj.CreatedBy = userId
	modelObj.CreatedAt = time.Now()

	if modelObj.Type == "" {
		return nil, errors.Error("Type of ClassRoom cannot be empty.")
	}

	filter := bson.D{bson.E{Key: "type", Value: modelObj.Type},
		bson.E{Key: "schoolid", Value: modelObj.SchoolId}}
	count, err := impl.collection.CountDocuments(impl.ctx, filter)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.Error(fmt.Sprintf("Type of ClassRoom '%v'already exist.", model.Type))
	}
	_, er := impl.collection.InsertOne(impl.ctx, modelObj)
	if er != nil {
		return nil, errors.Error("Error in creating ClassRoom.")
	}
	log.Print("Call to create ClassRoom completed.")
	return modelObj, er
}

func (impl serviceImpl) CreateClassRooms(userId string, _models []dtos.CreateClassRoomRequest) (interface{}, error) {

	log.Print("Call to create class rooms started.")

	types := make([]string, 0)
	var classRooms []dtos.ClassRoomResponse
	for _, model := range _models {
		types = append(types, model.Type)
	}

	filter := bson.D{{Key: "type", Value: bson.D{
		bson.E{Key: "$in", Value: types}}}}

	cur, _ := impl.collection.Find(impl.ctx, filter)

	_ = cur.All(impl.ctx, &classRooms)
	cur.Close(impl.ctx)

	modelObjs := make([]interface{}, 0)
	for _, model := range _models {
		var modelObj models.ClassRoom
		modelObj.CreatedBy = userId
		modelObj.CreatedAt = time.Now()
		check := false
		for _, classRoom := range classRooms {
			if model.Type == classRoom.Type {
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
		return nil, errors.Error("Error in creating class rooms.")
	}
	log.Print("Call to create class rooms completed.")
	return modelObjs, er
}

func (impl serviceImpl) PutClassRoom(id string, item dtos.UpdateClassRoomRequest) (interface{}, error) {

	log.Print("PutStaff started")

	objId := conversion.GetMongoId(id)
	var updatedClassRoom dtos.UpdateClassRoomRequest
	conversion.Convert(item, &updatedClassRoom)
	filter := bson.D{bson.E{Key: "_id", Value: objId}}
	var modelObj models.ClassRoom

	update := bson.D{bson.E{Key: "type", Value: updatedClassRoom.Type},
		bson.E{Key: "schoolid", Value: updatedClassRoom.SchoolId}}
	_, err := impl.collection.UpdateOne(impl.ctx, filter, bson.D{bson.E{Key: "$set", Value: update}})

	if err != nil {
		return modelObj, errors.Error("Could not upadte type of staff")
	}

	log.Print("PutStaff completed")
	return modelObj, nil
}
