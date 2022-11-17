package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/GabrielAchumba/school-mgt-backend/common/config"
	"github.com/GabrielAchumba/school-mgt-backend/common/conversion"
	errors "github.com/GabrielAchumba/school-mgt-backend/common/errors"
	"github.com/GabrielAchumba/school-mgt-backend/modules/grade-module/dtos"
	"github.com/GabrielAchumba/school-mgt-backend/modules/grade-module/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type GradeService interface {
	CreateGrade(userId string, requestModel dtos.CreateGradeRequest) (interface{}, error)
	CreateGrades(userId string, _models []dtos.CreateGradeRequest) (interface{}, error)
	DeleteGrade(id string, schoolId string) (int64, error)
	GetGrade(id string, schoolId string) (dtos.GradeResponse, error)
	GetGrades(schoolId string) ([]dtos.GradeResponse, error)
	PutGrade(id string, item dtos.UpdateGradeRequest) (interface{}, error)
}

type serviceImpl struct {
	ctx        context.Context
	collection *mongo.Collection
}

func New(mongoClient *mongo.Client, config config.Settings, ctx context.Context) GradeService {
	collection := mongoClient.Database(config.Database.DatabaseName).Collection(config.TableNames.Grade)

	return &serviceImpl{
		collection: collection,
		ctx:        ctx,
	}
}

func (impl serviceImpl) DeleteGrade(id string, schoolId string) (int64, error) {

	log.Print("Call to delete type of Grade by id started.")
	objId := conversion.GetMongoId(id)
	filter := bson.D{bson.E{Key: "_id", Value: objId},
		bson.E{Key: "schoolid", Value: schoolId}}

	result, err := impl.collection.DeleteOne(impl.ctx, filter)

	if err != nil {
		return result.DeletedCount, errors.Error("Error in deleting type of Grade.")
	}

	if result.DeletedCount < 1 {
		return result.DeletedCount, errors.Error("Type of Grade with specified ID not found!")
	}

	log.Print("Call to delete type of Grade by id completed.")
	return result.DeletedCount, nil
}

func (impl serviceImpl) GetGrade(id string, schoolId string) (dtos.GradeResponse, error) {

	log.Print("Get Type of Grade called")
	objId := conversion.GetMongoId(id)
	var Grade dtos.GradeResponse

	filter := bson.D{bson.E{Key: "_id", Value: objId},
		bson.E{Key: "schoolid", Value: schoolId}}

	err := impl.collection.FindOne(impl.ctx, filter).Decode(&Grade)
	if err != nil {
		return Grade, errors.Error("could not find type of Grade by id")
	}

	log.Print("Get type of Grade completed")
	return Grade, err

}

func (impl serviceImpl) GetGrades(schoolId string) ([]dtos.GradeResponse, error) {

	log.Print("Call to get all types of Grade started.")

	var Grades []dtos.GradeResponse
	filter := bson.D{bson.E{Key: "schoolid", Value: schoolId}}
	cur, err := impl.collection.Find(impl.ctx, filter)

	if err != nil {
		Grades = make([]dtos.GradeResponse, 0)
		return Grades, errors.Error("Types of Grade not found!")
	}

	err = cur.All(impl.ctx, &Grades)
	if err != nil {
		return Grades, err
	}

	cur.Close(impl.ctx)
	if len(Grades) == 0 {
		Grades = make([]dtos.GradeResponse, 0)
	}

	log.Print("Call to get all types of Grade completed.")
	return Grades, err
}

func (impl serviceImpl) CreateGrade(userId string, model dtos.CreateGradeRequest) (interface{}, error) {

	log.Print("Call to create Grade started.")

	var modelObj models.Grade
	conversion.Convert(model, &modelObj)

	modelObj.CreatedBy = userId
	modelObj.CreatedAt = time.Now()

	if modelObj.Type == "" {
		return nil, errors.Error("Type of Grade cannot be empty.")
	}

	filter := bson.D{bson.E{Key: "type", Value: modelObj.Type},
		bson.E{Key: "schoolid", Value: modelObj.SchoolId}}
	count, err := impl.collection.CountDocuments(impl.ctx, filter)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.Error(fmt.Sprintf("Type of Grade '%v'already exist.", model.Type))
	}
	_, er := impl.collection.InsertOne(impl.ctx, modelObj)
	if er != nil {
		return nil, errors.Error("Error in creating Grade.")
	}
	log.Print("Call to create Grade completed.")
	return modelObj, er
}

func (impl serviceImpl) CreateGrades(userId string, _models []dtos.CreateGradeRequest) (interface{}, error) {

	log.Print("Call to create Grades started.")

	types := make([]string, 0)
	var grades []dtos.GradeResponse
	for _, model := range _models {
		types = append(types, model.Type)
	}

	filter := bson.D{{Key: "type", Value: bson.D{
		bson.E{Key: "$in", Value: types}}}}

	cur, _ := impl.collection.Find(impl.ctx, filter)

	_ = cur.All(impl.ctx, &grades)
	cur.Close(impl.ctx)

	modelObjs := make([]interface{}, 0)
	for _, model := range _models {
		var modelObj models.Grade
		modelObj.CreatedBy = userId
		modelObj.CreatedAt = time.Now()
		check := false
		for _, grade := range grades {
			if model.Type == grade.Type {
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
		return nil, errors.Error("Error in creating Grades.")
	}
	log.Print("Call to create Grades completed.")
	return modelObjs, er
}

func (impl serviceImpl) PutGrade(id string, User dtos.UpdateGradeRequest) (interface{}, error) {

	log.Print("PutGrade started")

	objId := conversion.GetMongoId(id)
	var updatedGrade dtos.UpdateGradeRequest
	conversion.Convert(User, &updatedGrade)
	filter := bson.D{bson.E{Key: "_id", Value: objId}}
	var modelObj models.Grade

	update := bson.D{bson.E{Key: "type", Value: updatedGrade.Type},
		bson.E{Key: "point", Value: updatedGrade.Point},
		bson.E{Key: "from", Value: updatedGrade.From},
		bson.E{Key: "to", Value: updatedGrade.To},
		bson.E{Key: "schoolid", Value: updatedGrade.SchoolId}}
	_, err := impl.collection.UpdateOne(impl.ctx, filter, bson.D{bson.E{Key: "$set", Value: update}})

	if err != nil {
		return modelObj, errors.Error("Could not upadte type of Grade")
	}

	log.Print("PutGrade completed")
	return modelObj, nil
}
