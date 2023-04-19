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
	"github.com/GabrielAchumba/school-mgt-backend/realestate/land-module/dtos"
	"github.com/GabrielAchumba/school-mgt-backend/realestate/land-module/models"
	realEstateUserModule "github.com/GabrielAchumba/school-mgt-backend/realestate/user-module/services"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LandService interface {
	CreateLand(userId string, requestModel dtos.CreateLandRequest) (interface{}, error)
	DeleteLand(id string, schoolId string) (int64, error)
	CustomFilter(rows []dtos.LandResponse, filter string) []dtos.LandResponse
	GetLand(id string) (dtos.LandResponse, error)
	GetLands(filterModel string) ([]dtos.LandResponse, error)
	GetPaginatedLands(page int, filterModel string) (dtos.LandResponsePaginated, error)
	PutLand(id string, item dtos.UpdateLandRequest) (interface{}, error)
}

type serviceImpl struct {
	ctx         context.Context
	collection  *mongo.Collection
	userService realEstateUserModule.UserService
}

func New(collection *mongo.Collection, config config.Settings, ctx context.Context,
	userService realEstateUserModule.UserService) LandService {

	return &serviceImpl{
		collection:  collection,
		ctx:         ctx,
		userService: userService,
	}
}

func (impl serviceImpl) DeleteLand(id string, schoolId string) (int64, error) {

	log.Print("Call to delete type of Land by id started.")
	objId := conversion.GetMongoId(id)
	filter := bson.D{bson.E{Key: "_id", Value: objId},
		bson.E{Key: "schoolid", Value: schoolId}}

	result, err := impl.collection.DeleteOne(impl.ctx, filter)

	if err != nil {
		return result.DeletedCount, errors.Error("Error in deleting type of Land.")
	}

	if result.DeletedCount < 1 {
		return result.DeletedCount, errors.Error("Type of Land with specified ID not found!")
	}

	log.Print("Call to delete type of Land by id completed.")
	return result.DeletedCount, nil
}

func (impl serviceImpl) CustomFilter(rows []dtos.LandResponse, filter string) []dtos.LandResponse {
	lowerSearch := ""
	filteredRows := make([]dtos.LandResponse, 0)

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

func (impl serviceImpl) GetLand(id string) (dtos.LandResponse, error) {

	log.Print("Get Type of Land called")
	objId := conversion.GetMongoId(id)
	var Land dtos.LandResponse

	filter := bson.D{bson.E{Key: "_id", Value: objId}}

	err := impl.collection.FindOne(impl.ctx, filter).Decode(&Land)
	if err != nil {
		return Land, errors.Error("could not find type of Land by id")
	}

	log.Print("Get type of Land completed")
	return Land, err

}

func (impl serviceImpl) GetPaginatedLands(page int, filterModel string) (dtos.LandResponsePaginated, error) {

	log.Print("Call to get paginated lands started.")

	var lands []dtos.LandResponse
	filter := bson.D{}
	cur, err := impl.collection.Find(impl.ctx, filter)

	if err != nil {
		lands = make([]dtos.LandResponse, 0)
		return dtos.LandResponsePaginated{
			PaginatedLands:     lands,
			TotalNumberOfUsers: len(lands),
		}, errors.Error("Lands not found!")
	}

	err = cur.All(impl.ctx, &lands)
	if err != nil {
		lands = make([]dtos.LandResponse, 0)
		return dtos.LandResponsePaginated{
			PaginatedLands:     lands,
			TotalNumberOfUsers: len(lands),
		}, err
	}

	cur.Close(impl.ctx)
	length := len(lands)
	if length == 0 {
		lands = make([]dtos.LandResponse, 0)
	}

	filteredLands := impl.CustomFilter(lands, filterModel)

	paginatedLands := make([]dtos.LandResponse, 0)
	limit := 10
	skip := int64(page*limit - limit)

	counter := 0
	for i := skip; i < int64(len(filteredLands)); i++ {
		counter++
		if counter > limit {
			break
		}
		paginatedLands = append(filteredLands, filteredLands[i])
	}

	landResponsePaginated := dtos.LandResponsePaginated{
		PaginatedLands:     paginatedLands,
		TotalNumberOfUsers: length,
		Limit:              limit,
	}
	log.Print("Call to get paginated lands completed.")
	return landResponsePaginated, err
}

func (impl serviceImpl) GetLands(filterModel string) ([]dtos.LandResponse, error) {

	log.Print("Call to get all lands started.")

	var Users []dtos.LandResponse
	filter := bson.D{}
	cur, err := impl.collection.Find(impl.ctx, filter)

	if err != nil {
		Users = make([]dtos.LandResponse, 0)
		return Users, errors.Error("Lands not found!")
	}

	err = cur.All(impl.ctx, &Users)
	if err != nil {
		return Users, err
	}

	cur.Close(impl.ctx)
	length := len(Users)
	if length == 0 {
		Users = make([]dtos.LandResponse, 0)
	}

	filteredUsers := impl.CustomFilter(Users, filterModel)

	log.Print("Call to get lands completed.")
	return filteredUsers, err
}

func (impl serviceImpl) CreateLand(userId string, model dtos.CreateLandRequest) (interface{}, error) {

	log.Print("Call to create Land started.")

	var modelObj models.Land
	conversion.Convert(model, &modelObj)

	modelObj.CreatedBy = userId
	modelObj.CreatedAt = time.Now()

	if modelObj.Title == "" {
		return nil, errors.Error("Title of Land cannot be empty.")
	}

	filter := bson.D{bson.E{Key: "title", Value: modelObj.Title},
		bson.E{Key: "createdby", Value: modelObj.CreatedBy}}
	count, err := impl.collection.CountDocuments(impl.ctx, filter)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.Error(fmt.Sprintf("Title of Land '%v'already exist.", model.Title))
	}
	_, er := impl.collection.InsertOne(impl.ctx, modelObj)
	if er != nil {
		return nil, errors.Error("Error in creating Land.")
	}
	log.Print("Call to create Land completed.")
	return modelObj, er
}

func (impl serviceImpl) PutLand(id string, User dtos.UpdateLandRequest) (interface{}, error) {

	log.Print("PutLand started")

	objId := conversion.GetMongoId(id)
	var updatedLand dtos.UpdateLandRequest
	conversion.Convert(User, &updatedLand)
	filter := bson.D{bson.E{Key: "_id", Value: objId}}
	var modelObj models.Land

	update := bson.D{bson.E{Key: "title", Value: updatedLand.Title},
		{Key: "wholeplot", Value: updatedLand.WholePlot},
		{Key: "fractionplot", Value: updatedLand.FractionPlot},
		{Key: "partialaddress", Value: updatedLand.PartialAddress}}
	_, err := impl.collection.UpdateOne(impl.ctx, filter, bson.D{bson.E{Key: "$set", Value: update}})

	if err != nil {
		return modelObj, errors.Error("Could not upadte type of Land")
	}

	log.Print("PutLand completed")
	return modelObj, nil
}
