package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/GabrielAchumba/school-mgt-backend/common/config"
	"github.com/GabrielAchumba/school-mgt-backend/common/conversion"
	errors "github.com/GabrielAchumba/school-mgt-backend/common/errors"
	"github.com/GabrielAchumba/school-mgt-backend/modules/subject-module/dtos"
	"github.com/GabrielAchumba/school-mgt-backend/modules/subject-module/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SubjectService interface {
	CreateSubject(userId string, requestModel dtos.CreateSubjectRequest) (interface{}, error)
	CreateSubjects(userId string, _models []dtos.CreateSubjectRequest) (interface{}, error)
	DeleteSubject(id string, schoolId string) (int64, error)
	GetSubject(id string, schoolId string) (dtos.SubjectResponse, error)
	GetSubjects(schoolId string) ([]dtos.SubjectResponse, error)
	GetSubjectsByIds(schoolId string, Ids []string) ([]dtos.SubjectResponse, error)
	PutSubject(id string, item dtos.UpdateSubjectRequest) (interface{}, error)
}

type serviceImpl struct {
	ctx        context.Context
	collection *mongo.Collection
}

func New(mongoClient *mongo.Client, config config.Settings, ctx context.Context) SubjectService {
	collection := mongoClient.Database(config.Database.DatabaseName).Collection(config.TableNames.Subject)

	return &serviceImpl{
		collection: collection,
		ctx:        ctx,
	}
}

func (impl serviceImpl) DeleteSubject(id string, schoolId string) (int64, error) {

	log.Print("Call to delete type of Subject by id started.")
	objId := conversion.GetMongoId(id)
	filter := bson.D{bson.E{Key: "_id", Value: objId},
		bson.E{Key: "schoolid", Value: schoolId}}

	result, err := impl.collection.DeleteOne(impl.ctx, filter)

	if err != nil {
		return result.DeletedCount, errors.Error("Error in deleting type of Subject.")
	}

	if result.DeletedCount < 1 {
		return result.DeletedCount, errors.Error("Type of Subject with specified ID not found!")
	}

	log.Print("Call to delete type of Subject by id completed.")
	return result.DeletedCount, nil
}

func (impl serviceImpl) GetSubject(id string, schoolId string) (dtos.SubjectResponse, error) {

	log.Print("Get Type of Subject called")
	objId := conversion.GetMongoId(id)
	var Subject dtos.SubjectResponse

	filter := bson.D{bson.E{Key: "_id", Value: objId},
		bson.E{Key: "schoolid", Value: schoolId}}

	err := impl.collection.FindOne(impl.ctx, filter).Decode(&Subject)
	if err != nil {
		return Subject, errors.Error("could not find type of Subject by id")
	}

	log.Print("Get type of Subject completed")
	return Subject, err

}

func (impl serviceImpl) GetSubjects(schoolId string) ([]dtos.SubjectResponse, error) {

	log.Print("Call to get all types of Subject started.")

	var Subjects []dtos.SubjectResponse
	filter := bson.D{bson.E{Key: "schoolid", Value: schoolId}}
	cur, err := impl.collection.Find(impl.ctx, filter)

	if err != nil {
		Subjects = make([]dtos.SubjectResponse, 0)
		return Subjects, errors.Error("Types of Subject not found!")
	}

	err = cur.All(impl.ctx, &Subjects)
	if err != nil {
		return Subjects, err
	}

	cur.Close(impl.ctx)
	if len(Subjects) == 0 {
		Subjects = make([]dtos.SubjectResponse, 0)
	}

	log.Print("Call to get all types of Subject completed.")
	return Subjects, err
}

func (impl serviceImpl) GetSubjectsByIds(schoolId string, Ids []string) ([]dtos.SubjectResponse, error) {

	log.Print("Call to get GetSubjectsByIds started.")

	var objIds = make([]primitive.ObjectID, 0)
	for _, id := range Ids {
		objIds = append(objIds, conversion.GetMongoId(id))
	}

	var subjects []dtos.SubjectResponse
	filter := bson.D{bson.E{Key: "schoolid", Value: schoolId},
		bson.E{Key: "_id", Value: bson.D{bson.E{Key: "$in", Value: objIds}}}}

	cur, err := impl.collection.Find(impl.ctx, filter)

	if err != nil {
		subjects = make([]dtos.SubjectResponse, 0)
		return subjects, errors.Error("Subjects not found!")
	}

	err = cur.All(impl.ctx, &subjects)
	if err != nil {
		return subjects, err
	}

	cur.Close(impl.ctx)
	if len(subjects) == 0 {
		subjects = make([]dtos.SubjectResponse, 0)
	}

	log.Print("Call to get subjects by Ids completed.")
	return subjects, err
}

func (impl serviceImpl) CreateSubject(userId string, model dtos.CreateSubjectRequest) (interface{}, error) {

	log.Print("Call to create Subject started.")

	var modelObj models.Subject
	conversion.Convert(model, &modelObj)

	modelObj.CreatedBy = userId
	modelObj.CreatedAt = time.Now()

	if modelObj.Type == "" {
		return nil, errors.Error("Type of Subject cannot be empty.")
	}

	filter := bson.D{bson.E{Key: "type", Value: modelObj.Type},
		bson.E{Key: "schoolid", Value: modelObj.SchoolId}}
	count, err := impl.collection.CountDocuments(impl.ctx, filter)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.Error(fmt.Sprintf("Type of Subject '%v'already exist.", model.Type))
	}
	_, er := impl.collection.InsertOne(impl.ctx, modelObj)
	if er != nil {
		return nil, errors.Error("Error in creating Subject.")
	}
	log.Print("Call to create Subject completed.")
	return modelObj, er
}

func (impl serviceImpl) CreateSubjects(userId string, _models []dtos.CreateSubjectRequest) (interface{}, error) {

	log.Print("Call to create subjects started.")

	types := make([]string, 0)
	var subjects []dtos.SubjectResponse
	for _, model := range _models {
		types = append(types, model.Type)
	}

	filter := bson.D{{Key: "type", Value: bson.D{
		bson.E{Key: "$in", Value: types}}}}

	cur, _ := impl.collection.Find(impl.ctx, filter)

	_ = cur.All(impl.ctx, &subjects)
	cur.Close(impl.ctx)

	modelObjs := make([]interface{}, 0)
	for _, model := range _models {
		var modelObj models.Subject
		modelObj.CreatedBy = userId
		modelObj.CreatedAt = time.Now()
		check := false
		for _, subject := range subjects {
			if model.Type == subject.Type {
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
		return nil, errors.Error("Error in creating subjects.")
	}
	log.Print("Call to create subjects completed.")
	return modelObjs, er
}

func (impl serviceImpl) PutSubject(id string, User dtos.UpdateSubjectRequest) (interface{}, error) {

	log.Print("PutSubject started")

	objId := conversion.GetMongoId(id)
	var updatedSubject dtos.UpdateSubjectRequest
	conversion.Convert(User, &updatedSubject)
	filter := bson.D{bson.E{Key: "_id", Value: objId}}
	var modelObj models.Subject

	update := bson.D{bson.E{Key: "type", Value: updatedSubject.Type},
		bson.E{Key: "schoolid", Value: updatedSubject.SchoolId}}
	_, err := impl.collection.UpdateOne(impl.ctx, filter, bson.D{bson.E{Key: "$set", Value: update}})

	if err != nil {
		return modelObj, errors.Error("Could not upadte type of subject")
	}

	log.Print("PutSubject completed")
	return modelObj, nil
}
