package services

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/GabrielAchumba/school-mgt-backend/common/config"
	"github.com/GabrielAchumba/school-mgt-backend/common/conversion"
	errors "github.com/GabrielAchumba/school-mgt-backend/common/errors"
	"github.com/GabrielAchumba/school-mgt-backend/realestate/file-module/dtos"
	"github.com/GabrielAchumba/school-mgt-backend/realestate/file-module/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type FileService interface {
	CreateFile(userId string, requestModel dtos.CreateFileRequest) (interface{}, error)
	DeleteFile(id string, schoolId string) (int64, error)
	CustomFilter(rows []dtos.FileResponse, filter string) []dtos.FileResponse
	GetFile(id string) (dtos.FileResponse, error)
	GetFiles(filterModel string) ([]dtos.FileResponse, error)
	PutFile(id string, item dtos.UpdateFileRequest) (interface{}, error)
	GetFileByParams(title string, userId string, categoryId string) (bool, error)
}

type serviceImpl struct {
	ctx        context.Context
	collection *mongo.Collection
}

func New(mongoClient *mongo.Client, config config.Settings, ctx context.Context) FileService {
	collection := mongoClient.Database(config.Database.DatabaseName).Collection(config.TableNames.File)

	return &serviceImpl{
		collection: collection,
		ctx:        ctx,
	}
}

func (impl serviceImpl) DeleteFile(id string, schoolId string) (int64, error) {

	log.Print("Call to delete type of File by id started.")
	objId := conversion.GetMongoId(id)
	filter := bson.D{bson.E{Key: "_id", Value: objId},
		bson.E{Key: "schoolid", Value: schoolId}}

	result, err := impl.collection.DeleteOne(impl.ctx, filter)

	if err != nil {
		return result.DeletedCount, errors.Error("Error in deleting type of File.")
	}

	if result.DeletedCount < 1 {
		return result.DeletedCount, errors.Error("Type of File with specified ID not found!")
	}

	log.Print("Call to delete type of File by id completed.")
	return result.DeletedCount, nil
}

func (impl serviceImpl) CustomFilter(rows []dtos.FileResponse, filter string) []dtos.FileResponse {
	lowerSearch := ""
	filteredRows := make([]dtos.FileResponse, 0)

	if filter != "@" {
		lowerSearch = strings.ToLower(filter)
	}

	s1 := true

	for _, row := range rows {
		if lowerSearch != "" && lowerSearch != "@" {
			s1 = false
			//Get the values
			v := reflect.ValueOf(row)
			s1_values := make([]interface{}, v.NumField())
			for i := 0; i < v.NumField(); i++ {
				s1_values[i] = v.Field(i).Interface()
			}
			//Convert to lowercase
			//let s1_lower = s1_values.map(x => x.toString().toLowerCase())
			s1_lower := make([]string, 0)
			for _, item := range s1_values {
				dataType := reflect.TypeOf(item).String()
				if dataType == "string" {
					txt := reflect.ValueOf(item).String()
					s1_lower = append(s1_lower, strings.ToLower(txt))
				}
			}

			for val := 0; val < len(s1_lower); val++ {
				check := strings.Contains(s1_lower[val], lowerSearch)
				if check {
					s1 = true
					break
				}
			}

			if s1 {
				filteredRows = append(filteredRows, row)
			}

		} else {
			filteredRows = append(filteredRows, row)
		}
	}

	return filteredRows
}

func (impl serviceImpl) GetFile(id string) (dtos.FileResponse, error) {

	log.Print("Get Type of File called")
	objId := conversion.GetMongoId(id)
	var File dtos.FileResponse

	filter := bson.D{bson.E{Key: "_id", Value: objId}}

	err := impl.collection.FindOne(impl.ctx, filter).Decode(&File)
	if err != nil {
		return File, errors.Error("could not find type of File by id")
	}

	log.Print("Get type of File completed")
	return File, err

}

func (impl serviceImpl) GetFileByParams(title string, userId string, categoryId string) (bool, error) {

	log.Print("Get File by params called")

	filter := bson.D{bson.E{Key: "title", Value: title},
		bson.E{Key: "categoryid", Value: categoryId},
		bson.E{Key: "createdby", Value: userId}}

	count, err := impl.collection.CountDocuments(impl.ctx, filter)
	if err != nil {
		return false, errors.Error("Error in fetching the file")
	}
	if count > 0 {
		return true, errors.Error(fmt.Sprintf("Title of File '%v'already exist.", title))
	}

	log.Print("Get File by params completed")
	return false, err

}

func (impl serviceImpl) GetFiles(filterModel string) ([]dtos.FileResponse, error) {

	log.Print("Call to get all Files started.")

	var Users []dtos.FileResponse
	filter := bson.D{}
	cur, err := impl.collection.Find(impl.ctx, filter)

	if err != nil {
		Users = make([]dtos.FileResponse, 0)
		return Users, errors.Error("Files not found!")
	}

	err = cur.All(impl.ctx, &Users)
	if err != nil {
		return Users, err
	}

	cur.Close(impl.ctx)
	length := len(Users)
	if length == 0 {
		Users = make([]dtos.FileResponse, 0)
	}

	filteredUsers := impl.CustomFilter(Users, filterModel)

	log.Print("Call to get Files completed.")
	return filteredUsers, err
}

func (impl serviceImpl) CreateFile(userId string, model dtos.CreateFileRequest) (interface{}, error) {

	log.Print("Call to create File started.")

	var modelObj models.File
	conversion.Convert(model, &modelObj)

	modelObj.CreatedBy = userId
	modelObj.CreatedAt = time.Now()

	if modelObj.Title == "" {
		return nil, errors.Error("Title of File cannot be empty.")
	}

	filter := bson.D{bson.E{Key: "title", Value: modelObj.Title},
		bson.E{Key: "createdby", Value: modelObj.CreatedBy}}
	count, err := impl.collection.CountDocuments(impl.ctx, filter)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.Error(fmt.Sprintf("Title of File '%v'already exist.", model.Title))
	}
	_, er := impl.collection.InsertOne(impl.ctx, modelObj)
	if er != nil {
		return nil, errors.Error("Error in creating File.")
	}
	log.Print("Call to create File completed.")
	return modelObj, er
}

func (impl serviceImpl) PutFile(id string, User dtos.UpdateFileRequest) (interface{}, error) {

	log.Print("PutFile started")

	objId := conversion.GetMongoId(id)
	var updatedFile dtos.UpdateFileRequest
	conversion.Convert(User, &updatedFile)
	filter := bson.D{bson.E{Key: "_id", Value: objId}}
	var modelObj models.File

	update := bson.D{bson.E{Key: "title", Value: updatedFile.Title},
		{Key: "description", Value: updatedFile.Description},
		{Key: "fileUrl", Value: updatedFile.FileUrl},
		{Key: "filename", Value: updatedFile.FileName},
		{Key: "originalfilename", Value: updatedFile.OriginalFileName},
		{Key: "filecategory", Value: updatedFile.FileCategory},
		{Key: "categoryid", Value: updatedFile.CategoryId}}
	_, err := impl.collection.UpdateOne(impl.ctx, filter, bson.D{bson.E{Key: "$set", Value: update}})

	if err != nil {
		return modelObj, errors.Error("Could not upadte type of File")
	}

	log.Print("PutFile completed")
	return modelObj, nil
}
