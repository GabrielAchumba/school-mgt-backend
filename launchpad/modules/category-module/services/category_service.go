package services

import (
	"context"
	"log"
	"strconv"
	"strings"

	"github.com/GabrielAchumba/school-mgt-backend/common/config"
	"github.com/GabrielAchumba/school-mgt-backend/common/conversion"
	networkingerrors "github.com/GabrielAchumba/school-mgt-backend/common/errors"
	"github.com/GabrielAchumba/school-mgt-backend/launchpad/helpers"
	"github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/category-module/dtos"
	"github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/category-module/models"
	contributorDTOPackage "github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/user-module/dtos"
	contributorServicesPackage "github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/user-module/services"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryService interface {
	GetCategorys() ([]dtos.Category, error)
	GetSelectedCategorys(filter primitive.D) ([]models.Category, error)
	GetCategory(id string) (dtos.Category, error)
	GetSelectedCategory(filter primitive.D) (dtos.Category, interface{})
	CreateCategoryDTO(category contributorDTOPackage.CreateUserRequest) (dtos.Category, error)
	DeleteCategory(id string) (models.Category, error)
	UpdateCategory(id string, update primitive.D) error
	GetCompletedLevelXCategorys(levelIndex int, categoryIndex int) ([]dtos.Category, error)
	GetPersonalDataList() ([]contributorDTOPackage.UserResponse, error)
}

type serviceImpl struct {
	collection         *mongo.Collection
	ctx                context.Context
	contributorService contributorServicesPackage.UserService
	UserUtils          helpers.UserUtils
}

func New(collection *mongo.Collection, config config.Settings, ctx context.Context,
	UserService contributorServicesPackage.UserService) CategoryService {

	return &serviceImpl{
		collection:         collection,
		ctx:                ctx,
		contributorService: UserService,
	}
}

func (impl serviceImpl) CreateCategoryDTO(_category contributorDTOPackage.CreateUserRequest) (dtos.Category, error) {

	log.Print("CreateCategoryDTO started.")

	var category dtos.Category
	conversion.Convert(_category, &category)

	filter := bson.D{bson.E{Key: "username", Value: category.UserName}}

	err := impl.collection.FindOne(impl.ctx, filter).Decode(&category)
	if err == nil {
		return dtos.Category{}, networkingerrors.Error("category exists!")
	}

	filter = bson.D{bson.E{Key: "username", Value: category.ParentUserName}}
	parentCategory, err1 := impl.GetSelectedCategory(filter)

	if strings.ToLower(category.ParentUserName) == "admin" {
		category.ParentId = "admin"
		category.ParentUserName = "admin"

	} else if err1 == "No Category" {
		return dtos.Category{}, networkingerrors.Error("Referal username does not exist")
	} else {
		if parentCategory.NLevelOneRoomOneChildren == 3 {
			return dtos.Category{}, networkingerrors.Error("Referal user already has maximum (3) immediate downliners")
		}
		category.ParentId = parentCategory.CategoryId
		category.ParentUserName = parentCategory.UserName
	}

	ans, err := impl.collection.InsertOne(impl.ctx, category)

	if err != nil {
		return dtos.Category{}, networkingerrors.Error("Error in creating category.")
	}

	var stringObjectID = conversion.GetIdFromMongoId(ans.InsertedID)
	log.Print("CreatePersonalDataDTO completed.")
	return dtos.Category{
		CategoryId: stringObjectID,
		EntryDate: strconv.Itoa(category.CreatedDay) + "/" +
			strconv.Itoa(category.CreatedMonth) + "/" + strconv.Itoa(category.CreatedYear),
		ContributorId:  category.ContributorId,
		UserName:       category.UserName,
		ParentUserName: category.ParentUserName,
		ParentId:       category.ParentId,
		FullName: category.FirstName + "" +
			category.MiddleName + "" + category.LastName + "",
		Gender: category.Gender,
	}, err
}

func (impl serviceImpl) DeleteCategory(id string) (models.Category, error) {
	log.Print("DeleteCategory started.")
	objId := conversion.GetMongoId(id)
	filter := bson.D{bson.E{Key: "_id", Value: objId}}

	result, err := impl.collection.DeleteOne(impl.ctx, filter)

	if err != nil {
		return models.Category{}, networkingerrors.Error("Error in deleting category.")
	}

	if result.DeletedCount < 1 {
		return models.Category{}, networkingerrors.Error("Category with specified ID not found!")
	}

	log.Print("DeleteCategory completed.")
	return models.Category{}, nil
}

func (impl serviceImpl) GetCategory(id string) (dtos.Category, error) {

	log.Print("GetCategory started")
	objId := conversion.GetMongoId(id)
	var category models.Category

	filter := bson.D{bson.E{Key: "_id", Value: objId}}

	err := impl.collection.FindOne(impl.ctx, filter).Decode(&category)
	if err != nil {
		return dtos.Category{}, networkingerrors.Error("could not find category by id")
	}

	var parentCategory models.Category
	if category.ParentId == "admin" {
		parentCategory.UserName = "admin"

	} else {
		objId = conversion.GetMongoId(category.ParentId)
		filter = bson.D{bson.E{Key: "_id", Value: objId}}

		err = impl.collection.FindOne(impl.ctx, filter).Decode(&parentCategory)
		if err != nil {
			return dtos.Category{}, networkingerrors.Error("could not find category by id")
		}
	}

	categoryDTO := dtos.Category{
		FullName: category.FirstName + " " +
			category.MiddleName + " " +
			category.LastName + " ",
		CategoryId: category.Id,
		EntryDate: strconv.Itoa(category.CreatedDay) + "/" +
			strconv.Itoa(category.CreatedMonth) + "/" +
			strconv.Itoa(category.CreatedYear),
		ContributorId:              category.ContributorId,
		UserName:                   category.UserName,
		ParentUserName:             parentCategory.UserName,
		ParentId:                   category.ParentId,
		NLevelOneRoomOneChildren:   category.NLevelOneRoomOneChildren,
		NLevelTwoRoomOneChildren:   category.NLevelTwoRoomOneChildren,
		NLevelThreeRoomOneChildren: category.NLevelThreeRoomOneChildren,
		NLevelFourRoomOneChildren:  category.NLevelFourRoomOneChildren,
		NLevelFiveRoomOneChildren:  category.NLevelFiveRoomOneChildren,
		NLevelSixRoomOneChildren:   category.NLevelSixRoomOneChildren,
		NLevelSevenRoomOneChildren: category.NLevelSevenRoomOneChildren,
	}
	log.Print("GetCategory completed")
	return categoryDTO, err
}

func (impl serviceImpl) GetCategorys() ([]dtos.Category, error) {

	log.Print("GetCategorys started")
	var categorys []models.Category
	filter := bson.D{}
	curr, err := impl.collection.Find(impl.ctx, filter)
	if err != nil {
		return make([]dtos.Category, 0),
			networkingerrors.Error("Could not get all categorys")
	}

	err = curr.All(impl.ctx, &categorys)
	if err != nil {
		return make([]dtos.Category, 0),
			networkingerrors.Error("Could not decode all categorys")
	}

	curr.Close(impl.ctx)
	if len(categorys) == 0 {
		return make([]dtos.Category, 0), nil
	}

	var categorysDTO []dtos.Category
	for _, val := range categorys {
		categorysDTO = append(categorysDTO, dtos.Category{
			FullName:                 val.FirstName + " " + val.MiddleName + " " + val.LastName,
			Gender:                   val.Gender,
			ContributorId:            val.ContributorId,
			CategoryId:               val.Id,
			UserName:                 val.UserName,
			ParentId:                 val.ParentId,
			NLevelOneRoomOneChildren: val.NLevelOneRoomOneChildren,
		})
	}

	log.Print("GetContributors completed")
	return categorysDTO, nil

}

func (impl serviceImpl) GetCompletedLevelXCategorys(levelIndex int, categoryIndex int) ([]dtos.Category, error) {

	log.Print("GetCompletedLevelXCategorys started")
	var categorys []models.Category
	nLevel := 3
	nLessThan := nLevel - 1
	nGreaterThan := nLevel + 1
	returnOnInvestment := helpers.ROIs[categoryIndex-1][levelIndex-1]
	filter := bson.D{bson.E{Key: "nleveloneroomonechildren", Value: bson.D{bson.E{Key: "$gt", Value: nLessThan}}},
		bson.E{Key: "nleveloneroomonechildren", Value: bson.D{bson.E{Key: "$lt", Value: nGreaterThan}}},
		bson.E{Key: "isnleveloneroomonechildren", Value: false}}

	switch levelIndex {
	case 2:
		nLevel = 9
		filter = bson.D{bson.E{Key: "nleveltworoomonechildren", Value: bson.D{bson.E{Key: "$gt", Value: nLevel - 1}}},
			bson.E{Key: "nleveltworoomonechildren", Value: bson.D{bson.E{Key: "$lt", Value: nLevel + 1}}},
			bson.E{Key: "isnleveltworoomonechildren", Value: false}}

	case 3:
		nLevel = 27
		filter = bson.D{bson.E{Key: "nlevelthreeroomonechildren", Value: bson.D{bson.E{Key: "$gt", Value: nLevel - 1}}},
			bson.E{Key: "nlevelthreeroomonechildren", Value: bson.D{bson.E{Key: "$lt", Value: nLevel + 1}}},
			bson.E{Key: "isnlevelthreeroomonechildren", Value: false}}

	case 4:
		nLevel = 81
		filter = bson.D{bson.E{Key: "nlevelfourroomonechildren", Value: bson.D{bson.E{Key: "$gt", Value: nLevel - 1}}},
			bson.E{Key: "nlevelfourroomonechildren", Value: bson.D{bson.E{Key: "$lt", Value: nLevel + 1}}},
			bson.E{Key: "isnlevelfourroomonechildren", Value: false}}
	case 5:
		nLevel = 243
		filter = bson.D{bson.E{Key: "nlevelfiveroomonechildren", Value: bson.D{bson.E{Key: "$gt", Value: nLevel - 1}}},
			bson.E{Key: "nlevelfiveroomonechildren", Value: bson.D{bson.E{Key: "$lt", Value: nLevel + 1}}},
			bson.E{Key: "isnlevelfiveroomonechildren", Value: false}}
	case 6:
		nLevel = 729
		filter = bson.D{bson.E{Key: "nlevelsixroomonechildren", Value: bson.D{bson.E{Key: "$gt", Value: nLevel - 1}}},
			bson.E{Key: "nlevelsixroomonechildren", Value: bson.D{bson.E{Key: "$lt", Value: nLevel + 1}}},
			bson.E{Key: "isnlevelsixroomonechildren", Value: false}}
	case 7:
		nLevel = 2187
		filter = bson.D{bson.E{Key: "nlevelsevenroomonechildren", Value: bson.D{bson.E{Key: "$gt", Value: nLevel - 1}}},
			bson.E{Key: "nlevelsevenroomonechildren", Value: bson.D{bson.E{Key: "$lt", Value: nLevel + 1}}},
			bson.E{Key: "isnlevelsevenroomonechildren", Value: false}}
	}

	curr, err := impl.collection.Find(impl.ctx, filter)
	if err != nil {
		return make([]dtos.Category, 0),
			networkingerrors.Error("Could not get all completed level room-one categorys")
	}

	err = curr.All(impl.ctx, &categorys)
	if err != nil {
		return make([]dtos.Category, 0),
			networkingerrors.Error("Could not decode all level room-one categorys")
	}

	curr.Close(impl.ctx)
	if len(categorys) == 0 {
		return make([]dtos.Category, 0), nil
	}

	contributors, _ := impl.contributorService.GetAllContributors()

	var categorysDTO []dtos.Category
	for _, val := range categorys {
		NLevelXRoomOneChildren := 0
		switch levelIndex {
		case 1:
			NLevelXRoomOneChildren = val.NLevelOneRoomOneChildren
		case 2:
			NLevelXRoomOneChildren = val.NLevelTwoRoomOneChildren
		case 3:
			NLevelXRoomOneChildren = val.NLevelThreeRoomOneChildren
		case 4:
			NLevelXRoomOneChildren = val.NLevelFourRoomOneChildren
		case 5:
			NLevelXRoomOneChildren = val.NLevelFiveRoomOneChildren
		case 6:
			NLevelXRoomOneChildren = val.NLevelSixRoomOneChildren
		case 7:
			NLevelXRoomOneChildren = val.NLevelSevenRoomOneChildren
		}
		contributor := impl.UserUtils.FindUserById(contributors, val.ContributorId)
		categorysDTO = append(categorysDTO, dtos.Category{
			EntryDate:              strconv.Itoa(val.CreatedDay) + "/" + strconv.Itoa(val.CreatedMonth) + "/" + strconv.Itoa(val.CreatedYear),
			FullName:               val.FirstName + " " + val.MiddleName + " " + val.LastName,
			Gender:                 val.Gender,
			ContributorId:          val.Id,
			NLevelXRoomOneChildren: NLevelXRoomOneChildren,
			BankName:               contributor.BankName,
			AccountName:            contributor.AccountName,
			AccountNumber:          contributor.AccountNumber,
			UserName:               val.UserName,
			CategoryId:             val.Id,
			ReturnOnInvestment:     returnOnInvestment,
		})
	}

	log.Print("GetCompletedLevelXCategorys completed")
	return categorysDTO, nil
}

func (impl serviceImpl) GetSelectedCategory(filter primitive.D) (dtos.Category, interface{}) {

	log.Print("GetSelectedCategory started")
	var category models.Category

	err := impl.collection.FindOne(impl.ctx, filter).Decode(&category)
	if err != nil {
		return dtos.Category{}, "No Category" //networkingerrors.Error("could not find contributor by any key expect _id")
	}

	var parentCategory models.Category
	objId := conversion.GetMongoId(category.ParentId)
	filter = bson.D{bson.E{Key: "_id", Value: objId}}

	categoryDTO := dtos.Category{
		CategoryId: category.Id,
		EntryDate: strconv.Itoa(category.CreatedDay) + "/" +
			strconv.Itoa(category.CreatedMonth) + "/" +
			strconv.Itoa(category.CreatedYear),
		ContributorId:  category.ContributorId,
		UserName:       category.UserName,
		ParentUserName: parentCategory.UserName,
		ParentId:       category.ParentId,
		FullName: category.FirstName + " " +
			category.MiddleName + " " + category.LastName,
		NLevelOneRoomOneChildren: category.NLevelOneRoomOneChildren,
	}
	log.Print("GetSelectedCategory completed")
	return categoryDTO, err
}

func (impl serviceImpl) GetSelectedCategorys(filter primitive.D) ([]models.Category, error) {

	log.Print("GetSelectedCategorys started")
	var categorys []models.Category
	curr, err := impl.collection.Find(impl.ctx, filter)
	if err != nil {
		return make([]models.Category, 0),
			networkingerrors.Error("Could not get selected Categorys")
	}

	err = curr.All(impl.ctx, &categorys)
	if err != nil {
		return make([]models.Category, 0),
			networkingerrors.Error("Could not decode selected categorys")
	}

	curr.Close(impl.ctx)
	if len(categorys) == 0 {
		return make([]models.Category, 0), nil
	}

	log.Print("GetSelectedCategorys completed")
	return categorys, nil
}

func (impl serviceImpl) UpdateCategory(id string, update primitive.D) error {

	log.Print("UpdateCategory started")
	objId := conversion.GetMongoId(id)
	filter := bson.D{bson.E{Key: "_id", Value: objId}}
	_, err := impl.collection.UpdateOne(impl.ctx, filter, bson.D{bson.E{Key: "$set", Value: update}})

	if err != nil {
		return networkingerrors.Error("Could not upadte category")
	}

	log.Print("UpdateCategory completed")
	return nil
}

func (impl serviceImpl) GetPersonalDataList() ([]contributorDTOPackage.UserResponse, error) {

	log.Print("GetPersonalDataList Category started")
	contributors, _ := impl.contributorService.GetPersonalDataList()
	categories, _ := impl.GetCategorys()

	var result []contributorDTOPackage.UserResponse = make([]contributorDTOPackage.UserResponse, 0)
	for _, contributor := range contributors {
		check := false
		for _, category := range categories {
			if category.ContributorId == contributor.ContributorId {
				check = true
				break
			}
		}

		if !check {
			result = append(result, contributor)
		}

	}

	log.Print("GetPersonalDataList Category completed")
	return result, nil
}
