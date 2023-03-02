package services

import (
	"context"
	"log"
	"strconv"

	"github.com/GabrielAchumba/school-mgt-backend/common/config"
	"github.com/GabrielAchumba/school-mgt-backend/common/conversion"
	networkingerrors "github.com/GabrielAchumba/school-mgt-backend/common/errors"
	"github.com/GabrielAchumba/school-mgt-backend/launchpad/helpers"
	"github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/cashout-module/dtos"
	"github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/cashout-module/models"
	"github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/category-module/services"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CashOutService interface {
	GetCashOuts(categoryIndex int) ([]dtos.CashOutDTO, error)
	GetCashOut(id string, categoryIndex int) (dtos.CashOutDTO, error)
	CreateCashOutDTO(cashOutDTO dtos.CashOutDTO) (dtos.CashOutDTO, error)
	DeleteCashOut(id string) (dtos.CashOutDTO, error)
	GetCategoryBankDetails(categoryIndex int) dtos.CategoryBankDetails
	GetCashOutByCategoryId(levelIndex int, categoryId string, categoryIndex int) (dtos.CashOutDTO, error)
}

type serviceImpl struct {
	collection      *mongo.Collection
	ctx             context.Context
	categoryService services.CategoryService
}

func New(collection *mongo.Collection, config config.Settings, ctx context.Context, categoryService services.CategoryService) CashOutService {

	return &serviceImpl{
		collection:      collection,
		ctx:             ctx,
		categoryService: categoryService,
	}
}

func (impl serviceImpl) CreateCashOutDTO(cashOutDTO dtos.CashOutDTO) (dtos.CashOutDTO, error) {

	log.Print("CreateCashOutDTO started.")

	var cashout dtos.CashOutDTO
	conversion.Convert(cashOutDTO, &cashout)

	filter := bson.D{bson.E{Key: "categoryid", Value: cashout.CategoryId},
		bson.E{Key: "category", Value: cashout.Category},
		bson.E{Key: "level", Value: cashout.Level}}

	err := impl.collection.FindOne(impl.ctx, filter).Decode(&cashout)
	if err == nil {
		return dtos.CashOutDTO{}, networkingerrors.Error("cashout exists!")
	}

	_, err = impl.collection.InsertOne(impl.ctx, cashout)

	if err != nil {
		return dtos.CashOutDTO{}, networkingerrors.Error("Error in creating cashout.")
	}

	update := bson.D{bson.E{Key: "isnleveloneroomonechildren", Value: true}}
	err = impl.categoryService.UpdateCategory(cashout.CategoryId, update)
	if err != nil {
		return dtos.CashOutDTO{}, networkingerrors.Error("Could not update category")
	}
	log.Print("CreateCashOutDTO completed.")
	return cashout, err
}

func (impl serviceImpl) DeleteCashOut(id string) (dtos.CashOutDTO, error) {

	log.Print("DeleteCashOut started.")
	objId := conversion.GetMongoId(id)
	filter := bson.D{bson.E{Key: "_id", Value: objId}}

	result, err := impl.collection.DeleteOne(impl.ctx, filter)

	if err != nil {
		return dtos.CashOutDTO{}, networkingerrors.Error("Error in deleting cashout.")
	}

	if result.DeletedCount < 1 {
		return dtos.CashOutDTO{}, networkingerrors.Error("Cashou with specified ID not found!")
	}

	log.Print("DeleteCashOut completed.")
	return dtos.CashOutDTO{}, nil
}

func (impl serviceImpl) GetCashOut(id string, categoryIndex int) (dtos.CashOutDTO, error) {

	log.Print("GetCashOut started")
	objId := conversion.GetMongoId(id)
	var cashout models.CashOut

	filter := bson.D{bson.E{Key: "_id", Value: objId}}

	err := impl.collection.FindOne(impl.ctx, filter).Decode(&cashout)
	if err != nil {
		return dtos.CashOutDTO{}, networkingerrors.Error("could not find cash by id")
	}

	log.Print("GetCashOut completed")
	return dtos.CashOutDTO{
		CashOutId:    cashout.Id,
		CreatedDay:   cashout.CreatedDay,
		CreatedMonth: cashout.CreatedMonth,
		CreatedYear:  cashout.CreatedYear,
		CreatedBy:    cashout.CreatedBy,
		CreatedDate: strconv.Itoa(cashout.CreatedDay) + "/" +
			strconv.Itoa(cashout.CreatedMonth) + "/" +
			strconv.Itoa(cashout.CreatedYear) + "/",
		FullName:              cashout.FullName,
		Category:              cashout.Category,
		Level:                 cashout.Level,
		CategoryBankName:      helpers.CategoryBankName[categoryIndex-1],
		CategoryAccountName:   helpers.CategoryAccountName[categoryIndex-1],
		CategoryAccountNumber: helpers.CategoryAccountNumber[categoryIndex-1],
		BankName:              cashout.BankName,
		AccountName:           cashout.AccountName,
		AccountNumber:         cashout.AccountName,
		Base64String:          cashout.Base64String,
	}, nil
}

func (impl serviceImpl) GetCashOutByCategoryId(levelIndex int, categoryId string, categoryIndex int) (dtos.CashOutDTO, error) {

	log.Print("GetCashOutByCategoryId started")
	var cashout models.CashOut

	filter := bson.D{bson.E{Key: "level", Value: levelIndex},
		bson.E{Key: "categoryid", Value: categoryId}}

	err := impl.collection.FindOne(impl.ctx, filter).Decode(&cashout)
	if err != nil {
		return dtos.CashOutDTO{}, networkingerrors.Error("could not find cash by categoryId and level")
	}

	log.Print("GetCashOutByCategoryId completed")
	return dtos.CashOutDTO{
		CashOutId:    cashout.Id,
		CreatedDay:   cashout.CreatedDay,
		CreatedMonth: cashout.CreatedMonth,
		CreatedYear:  cashout.CreatedYear,
		CreatedBy:    cashout.CreatedBy,
		CreatedDate: strconv.Itoa(cashout.CreatedDay) + "/" +
			strconv.Itoa(cashout.CreatedMonth) + "/" +
			strconv.Itoa(cashout.CreatedYear) + "/",
		FullName:              cashout.FullName,
		Category:              cashout.Category,
		Level:                 cashout.Level,
		CategoryBankName:      helpers.CategoryBankName[categoryIndex-1],
		CategoryAccountName:   helpers.CategoryAccountName[categoryIndex-1],
		CategoryAccountNumber: helpers.CategoryAccountNumber[categoryIndex-1],
		BankName:              cashout.BankName,
		AccountName:           cashout.AccountName,
		AccountNumber:         cashout.AccountName,
		Base64String:          cashout.Base64String,
		ReturnOnInvestment:    cashout.ReturnOnInvestment,
	}, nil
}

func (impl serviceImpl) GetCategoryBankDetails(categoryIndex int) dtos.CategoryBankDetails {

	return dtos.CategoryBankDetails{
		CategoryBankName:      helpers.CategoryBankName[categoryIndex-1],
		CategoryAccountName:   helpers.CategoryAccountName[categoryIndex-1],
		CategoryAccountNumber: helpers.CategoryAccountNumber[categoryIndex-1],
	}
}

func (impl serviceImpl) GetCashOuts(categoryIndex int) ([]dtos.CashOutDTO, error) {
	log.Print("GetCashOuts started")
	var cashouts []models.CashOut
	filter := bson.D{}
	curr, err := impl.collection.Find(impl.ctx, filter)
	if err != nil {
		return make([]dtos.CashOutDTO, 0),
			networkingerrors.Error("Could not get all cashouts")
	}

	err = curr.All(impl.ctx, &cashouts)
	if err != nil {
		return make([]dtos.CashOutDTO, 0),
			networkingerrors.Error("Could not decode all cashouts")
	}

	curr.Close(impl.ctx)
	if len(cashouts) == 0 {
		return make([]dtos.CashOutDTO, 0), nil
	}

	var cashoutDTO []dtos.CashOutDTO
	for _, val := range cashouts {
		cashoutDTO = append(cashoutDTO, dtos.CashOutDTO{
			FullName: val.FullName,
			CreatedDate: strconv.Itoa(val.CreatedDay) +
				"/" + strconv.Itoa(val.CreatedMonth) +
				"/" + strconv.Itoa(val.CreatedYear),
			CreatedBy:             val.CreatedBy,
			Category:              val.Category,
			Level:                 val.Level,
			BankName:              val.BankName,
			AccountName:           val.AccountName,
			AccountNumber:         val.AccountNumber,
			Base64String:          val.Base64String,
			CategoryBankName:      helpers.CategoryBankName[categoryIndex-1],
			CategoryAccountName:   helpers.CategoryAccountName[categoryIndex-1],
			CategoryAccountNumber: helpers.CategoryAccountNumber[categoryIndex-1],
			ReturnOnInvestment:    val.ReturnOnInvestment,
		})
	}

	log.Print("GetCashOuts completed")
	return cashoutDTO, nil
}
