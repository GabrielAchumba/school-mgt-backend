package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/GabrielAchumba/school-mgt-backend/common/config"
	"github.com/GabrielAchumba/school-mgt-backend/common/conversion"
	errors "github.com/GabrielAchumba/school-mgt-backend/common/errors"
	"github.com/GabrielAchumba/school-mgt-backend/modules/assessment-module/dtos"
	"github.com/GabrielAchumba/school-mgt-backend/modules/assessment-module/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AssessmentService interface {
	CreateAssessment(userId string, requestModel dtos.CreateAssessmentRequest) (interface{}, error)
	DeleteAssessment(id string, schoolId string) (int64, error)
	GetAssessment(id string, schoolId string) (dtos.AssessmentResponse, error)
	GetAssessments(schoolId string) ([]dtos.AssessmentResponse, error)
	PutAssessment(id string, item dtos.UpdateAssessmentRequest) (interface{}, error)
}

type serviceImpl struct {
	ctx        context.Context
	collection *mongo.Collection
}

func New(mongoClient *mongo.Client, config config.Settings, ctx context.Context) AssessmentService {
	collection := mongoClient.Database(config.Database.DatabaseName).Collection(config.TableNames.Assessment)

	return &serviceImpl{
		collection: collection,
		ctx:        ctx,
	}
}

func (impl serviceImpl) DeleteAssessment(id string, schoolId string) (int64, error) {

	log.Print("Call to delete type of Assessment by id started.")
	objId := conversion.GetMongoId(id)
	filter := bson.D{bson.E{Key: "_id", Value: objId},
		bson.E{Key: "schoolid", Value: schoolId}}

	result, err := impl.collection.DeleteOne(impl.ctx, filter)

	if err != nil {
		return result.DeletedCount, errors.Error("Error in deleting type of Assessment.")
	}

	if result.DeletedCount < 1 {
		return result.DeletedCount, errors.Error("Type of Assessment with specified ID not found!")
	}

	log.Print("Call to delete type of Assessment by id completed.")
	return result.DeletedCount, nil
}

func (impl serviceImpl) GetAssessment(id string, schoolId string) (dtos.AssessmentResponse, error) {

	log.Print("Get Type of Assessment called")
	objId := conversion.GetMongoId(id)
	var Assessment dtos.AssessmentResponse

	filter := bson.D{bson.E{Key: "_id", Value: objId},
		bson.E{Key: "schoolid", Value: schoolId}}

	err := impl.collection.FindOne(impl.ctx, filter).Decode(&Assessment)
	if err != nil {
		return Assessment, errors.Error("could not find type of Assessment by id")
	}

	log.Print("Get type of Assessment completed")
	return Assessment, err

}

func (impl serviceImpl) GetAssessments(schoolId string) ([]dtos.AssessmentResponse, error) {

	log.Print("Call to get all types of Assessment started.")

	var Assessments []dtos.AssessmentResponse
	filter := bson.D{bson.E{Key: "schoolid", Value: schoolId}}
	cur, err := impl.collection.Find(impl.ctx, filter)

	if err != nil {
		Assessments = make([]dtos.AssessmentResponse, 0)
		return Assessments, errors.Error("Types of Assessment not found!")
	}

	err = cur.All(impl.ctx, &Assessments)
	if err != nil {
		return Assessments, err
	}

	cur.Close(impl.ctx)
	if len(Assessments) == 0 {
		Assessments = make([]dtos.AssessmentResponse, 0)
	}

	log.Print("Call to get all types of Assessment completed.")
	return Assessments, err
}

func (impl serviceImpl) CreateAssessment(userId string, model dtos.CreateAssessmentRequest) (interface{}, error) {

	log.Print("Call to create Assessment started.")

	var modelObj models.Assessment
	conversion.Convert(model, &modelObj)

	modelObj.CreatedBy = userId
	modelObj.CreatedAt = time.Now()

	if modelObj.Type == "" {
		return nil, errors.Error("Type of Assessment cannot be empty.")
	}

	filter := bson.D{bson.E{Key: "type", Value: modelObj.Type},
		bson.E{Key: "schoolid", Value: modelObj.SchoolId}}
	count, err := impl.collection.CountDocuments(impl.ctx, filter)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.Error(fmt.Sprintf("Type of Assessment '%v'already exist.", model.Type))
	}
	_, er := impl.collection.InsertOne(impl.ctx, modelObj)
	if er != nil {
		return nil, errors.Error("Error in creating Assessment.")
	}
	log.Print("Call to create Assessment completed.")
	return modelObj, er
}

func (impl serviceImpl) PutAssessment(id string, User dtos.UpdateAssessmentRequest) (interface{}, error) {

	log.Print("PutAssessment started")

	objId := conversion.GetMongoId(id)
	var updatedAssessment dtos.UpdateAssessmentRequest
	conversion.Convert(User, &updatedAssessment)
	filter := bson.D{bson.E{Key: "_id", Value: objId}}
	var modelObj models.Assessment

	update := bson.D{bson.E{Key: "type", Value: updatedAssessment.Type},
		bson.E{Key: "percentage", Value: updatedAssessment.Percentage},
		bson.E{Key: "schoolid", Value: updatedAssessment.SchoolId}}
	_, err := impl.collection.UpdateOne(impl.ctx, filter, bson.D{bson.E{Key: "$set", Value: update}})

	if err != nil {
		return modelObj, errors.Error("Could not upadte type of Assessment")
	}

	log.Print("PutAssessment completed")
	return modelObj, nil
}
