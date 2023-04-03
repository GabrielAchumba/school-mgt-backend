package services

import (
	"context"
	"log"
	"strconv"

	"github.com/GabrielAchumba/school-mgt-backend/common/config"
	"github.com/GabrielAchumba/school-mgt-backend/common/conversion"
	networkingerrors "github.com/GabrielAchumba/school-mgt-backend/common/errors"
	"github.com/GabrielAchumba/school-mgt-backend/launchpad/helpers"
	"github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/account-module/dtos"
	"github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/account-module/models"
	categoryDTOPackage "github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/category-module/dtos"
	categoryServicesPackage "github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/category-module/services"
	contributorDTOPackage "github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/user-module/dtos"
	contributorServicesPackage "github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/user-module/services"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AccountService interface {
	GetAccounts() ([]models.Account, error)
	GetSelectedAccounts(key string, value interface{}) ([]models.Account, error)
	GetAccount(id string) (models.Account, error)
	GetAccountByCategoryId(contributorId string) (models.Account, error)
	GetUnComfirmedAccounts() ([]dtos.AccountModelResponse, error)
	GetComfirmedAccounts() ([]models.Account, error)
	ComfirmPayment(id string, accountModel dtos.CreateAccountRequest) (interface{}, error)
	OffPlatformPayment(account dtos.AccountDTO) (interface{}, error)
	Payment(account dtos.AccountDTO) (interface{}, error)
	DeleteAccount(id string) (models.Account, error)
	RegisteredHaveNotContributed() ([]categoryDTOPackage.Category, error)
	GetDescendantsByLevel(levelIndex int, parentId string) (interface{}, error)
	GetCompletedLevelXCategorys(levelIndex int, categoryIndex int) ([]categoryDTOPackage.Category, error)
	GetDownliners(parentId string, contributors []contributorDTOPackage.UserResponse,
		accounts []models.Account, categories []categoryDTOPackage.Category, Counter int, levelCount int,
		nLevel int) bool
}

type serviceImpl struct {
	collection         *mongo.Collection
	ctx                context.Context
	categoryService    categoryServicesPackage.CategoryService
	contributorService contributorServicesPackage.UserService
	CategoryUtils      helpers.CategoryUtils
	AccountUtils       helpers.AccountUtils
	UserUtils          helpers.UserUtils
}

func New(collection *mongo.Collection, config config.Settings, ctx context.Context,
	categoryService categoryServicesPackage.CategoryService,
	userService contributorServicesPackage.UserService) AccountService {

	return &serviceImpl{
		collection:         collection,
		ctx:                ctx,
		categoryService:    categoryService,
		contributorService: userService,
	}
}

func (impl serviceImpl) ComfirmPayment(id string,
	modelObj dtos.CreateAccountRequest) (interface{}, error) {

	log.Print("ComfirmPayment started")
	objId := conversion.GetMongoId(id)

	var accountModel dtos.CreateAccountRequest
	conversion.Convert(modelObj, &accountModel)

	var account dtos.CreateAccountRequest

	filter := bson.D{bson.E{Key: "_id", Value: objId}}

	err := impl.collection.FindOne(impl.ctx, filter).Decode(&account)
	if err != nil {
		return account, networkingerrors.Error("Account not found.")
	}

	if accountModel.Status == "success" {
		account.IsComfirmed = true
	} else {
		account.IsComfirmed = false
	}

	filter = bson.D{bson.E{Key: "_id", Value: objId}}
	update := bson.D{
		{"$set", bson.D{bson.E{Key: "iscomfirmed", Value: account.IsComfirmed}}}}

	_, err = impl.collection.UpdateOne(
		impl.ctx,
		filter,
		update,
	)

	if err != nil {
		return account, networkingerrors.Error("Could not upadate IsComfirmed in the account")
	}

	log.Print("ComfirmPayment completed")
	return account, nil
}

func (impl serviceImpl) DeleteAccount(id string) (models.Account, error) {

	log.Print("DeleteAccount started.")
	objId := conversion.GetMongoId(id)
	filter := bson.D{bson.E{Key: "_id", Value: objId}}

	result, err := impl.collection.DeleteOne(impl.ctx, filter)

	if err != nil {
		return models.Account{}, networkingerrors.Error("Error in deleting accont.")
	}

	if result.DeletedCount < 1 {
		return models.Account{}, networkingerrors.Error("Account with specified ID not found!")
	}

	log.Print("DeleteAccount completed.")
	return models.Account{}, nil
}

func (impl serviceImpl) GetAccount(id string) (models.Account, error) {

	log.Print("GetAccount Started")
	objId := conversion.GetMongoId(id)
	filter := bson.D{bson.E{Key: "_id", Value: objId}}
	var modelObj models.Account
	err := impl.collection.FindOne(impl.ctx, filter).Decode(&modelObj)
	if err != nil {
		return models.Account{}, networkingerrors.Error("could not find account by id")
	}
	log.Print("GetAccount Completed")
	return modelObj, nil
}

func (impl serviceImpl) GetAccountByCategoryId(contributorId string) (models.Account, error) {

	log.Print("GetAccountByUserId Started")

	filter := bson.D{bson.E{Key: "contributorid", Value: contributorId}}
	var account models.Account
	err := impl.collection.FindOne(impl.ctx, filter).Decode(&account)

	if err != nil {
		return models.Account{},
			networkingerrors.Error("Could not get account by category Id")
	}

	log.Print("GetAccountByUserId completed")
	return account, nil
}

func (impl serviceImpl) GetAccounts() ([]models.Account, error) {
	log.Print("GetAccounts started")
	filter := bson.D{}
	var accounts []models.Account
	curr, err := impl.collection.Find(impl.ctx, filter)

	if err != nil {
		return make([]models.Account, 0),
			networkingerrors.Error("Could not get all accounts")
	}

	err = curr.All(impl.ctx, &accounts)
	if err != nil {
		return make([]models.Account, 0),
			networkingerrors.Error("Could not decode all accounts")
	}

	curr.Close(impl.ctx)
	if len(accounts) == 0 {
		return make([]models.Account, 0), networkingerrors.Error("empty accounts list found")
	}

	log.Print("GetAccounts completed")
	return accounts, nil
}

func (impl serviceImpl) GetSelectedAccounts(key string, value interface{}) ([]models.Account, error) {
	log.Print("GetAccounts started")
	filter := bson.D{bson.E{Key: key, Value: value}}
	var accounts []models.Account
	curr, err := impl.collection.Find(impl.ctx, filter)

	if err == nil {
		return make([]models.Account, 0),
			networkingerrors.Error("Could not get all accounts")
	}

	err = curr.All(impl.ctx, &accounts)
	if err == nil {
		return make([]models.Account, 0),
			networkingerrors.Error("Could not decode all accounts")
	}

	curr.Close(impl.ctx)
	if len(accounts) == 0 {
		return make([]models.Account, 0), nil
	}

	log.Print("GetAccounts completed")
	return accounts, nil
}

func (impl serviceImpl) GetUnComfirmedAccounts() ([]dtos.AccountModelResponse, error) {
	log.Print("GetUnComfirmedAccounts started")
	filter := bson.D{bson.E{Key: "iscomfirmed", Value: false}}
	var accounts []models.Account
	curr, err := impl.collection.Find(impl.ctx, filter)

	if err != nil {
		return make([]dtos.AccountModelResponse, 0),
			networkingerrors.Error("Could not get all accounts")
	}

	err = curr.All(impl.ctx, &accounts)
	if err != nil {
		return make([]dtos.AccountModelResponse, 0),
			networkingerrors.Error("Could not decode all accounts")
	}

	curr.Close(impl.ctx)
	if len(accounts) == 0 {
		return make([]dtos.AccountModelResponse, 0), nil
	}

	var accountDTO []dtos.AccountModelResponse
	for _, val := range accounts {
		//contributor, _ := impl.ContributorService.GetContributor()
		//filter := bson.D{bson.E{Key: "contributorid", Value: val.ContributorId}}
		category, _ := impl.categoryService.GetCategory(val.ContributorId)
		accountDTO = append(accountDTO, dtos.AccountModelResponse{
			FullName:              category.FullName,
			Contribution:          val.Contribution,
			ContributorId:         val.ContributorId,
			AccountId:             val.Id,
			Status:                val.Status,
			Base64String:          val.Base64String,
			EntryDate:             strconv.Itoa(val.CreatedDay) + "/" + strconv.Itoa(val.CreatedMonth) + "/" + strconv.Itoa(val.CreatedYear),
			BankNamePaidFrom:      val.BankNamePaidFrom,
			AccountNamePaidFrom:   val.AccountNamePaidFrom,
			AccountNumberPaidFrom: val.AccountNumberPaidFrom,
		})
	}

	log.Print("GetUnComfirmedAccounts completed")
	return accountDTO, nil
}

func (impl serviceImpl) GetComfirmedAccounts() ([]models.Account, error) {
	log.Print("GetComfirmedAccounts started")
	filter := bson.D{bson.E{Key: "isComfirmed", Value: true}}
	var accounts []models.Account
	curr, err := impl.collection.Find(impl.ctx, filter)

	if err == nil {
		return make([]models.Account, 0),
			networkingerrors.Error("Could not get comfirmed accounts")
	}

	err = curr.All(impl.ctx, &accounts)
	if err == nil {
		return make([]models.Account, 0),
			networkingerrors.Error("Could not decode comfirmed accounts")
	}

	curr.Close(impl.ctx)
	if len(accounts) == 0 {
		return make([]models.Account, 0), nil
	}

	log.Print("GetComfirmedAccounts completed")
	return accounts, nil
}

func (impl serviceImpl) OffPlatformPayment(modelObj dtos.AccountDTO) (interface{}, error) {

	/* bankName: context.bankNamePaidFrom,
	   accountName: context.bankNamePaidFrom,
	   accountNumber: context.bankNamePaidFrom,
	   contributorId: id,
	   base64String: context.Account.base64String,
	   fileName: context.Account.fileName,
	   createdYear: todayDate.getFullYear(),
	   createdMonth: todayDate.getMonth() + 1,
	   createdDay: todayDate.getDate(), */

	var _account dtos.AccountDTO
	conversion.Convert(modelObj, &_account)
	//account1, _ := impl.GetAccountByContributorId(_account.ContributorId)
	filter := bson.D{bson.E{Key: "contributorid", Value: _account.ContributorId}}
	category, err := impl.categoryService.GetSelectedCategory(filter)
	if err != nil {
		return _account, networkingerrors.Error("Category not found. Please create category")
	}
	var account dtos.AccountDTO
	account.Base64String = _account.Base64String
	account.ContributorId = category.CategoryId
	account.FileName = _account.FileName
	account.CreatedDay = _account.CreatedDay
	account.CreatedMonth = _account.CreatedMonth
	account.CreatedYear = _account.CreatedYear
	account.Contribution = helpers.CategoryAmount[_account.CategoryIndex-1]
	account.Status = helpers.PaymentSuccessful
	account.BankNamePaidFrom = _account.BankNamePaidFrom
	account.AccountNamePaidFrom = _account.AccountNamePaidFrom
	account.AccountNumberPaidFrom = _account.AccountNumberPaidFrom
	account.DatePaidFrom = _account.DatePaidFrom
	account.UserName = _account.UserName
	account.Recipient = _account.Recipient
	result, er := impl.Payment(account)
	return result, er
}

func (impl serviceImpl) Payment(_account dtos.AccountDTO) (interface{}, error) {

	//var _account models.Account
	//conversion.Convert(account, &_account)

	filter := bson.D{bson.E{Key: "contributorid", Value: _account.ContributorId}}

	err := impl.collection.FindOne(impl.ctx, filter).Decode(&_account)
	if err == nil {
		return _account, networkingerrors.Error("Payment transaction exists!")
	}

	m, er := impl.collection.InsertOne(impl.ctx, _account)

	log.Print("er: ", er)
	if er != nil {
		return nil, networkingerrors.Error("Error in creating payment transaction.")
	}
	log.Print("Call to create payment transaction.")
	return m.InsertedID, er
}

func (impl serviceImpl) RegisteredHaveNotContributed() ([]categoryDTOPackage.Category, error) {
	log.Print("RegisteredHaveNotContributed started")

	contributors, err := impl.categoryService.GetCategorys()
	if err != nil {
		return make([]categoryDTOPackage.Category, 0),
			networkingerrors.Error("Categorys not found")
	}

	accounts, err := impl.GetAccounts()
	if err != nil {
		return contributors, nil
	} else {
		var contributorsHavNotCoontriuted []categoryDTOPackage.Category = make([]categoryDTOPackage.Category, 0)
		for _, contributor := range contributors {
			isHasPaid := false
			for _, account := range accounts {
				if account.ContributorId == contributor.CategoryId {
					isHasPaid = true
					break
				}
			}

			if !isHasPaid {
				contributorsHavNotCoontriuted = append(contributorsHavNotCoontriuted, contributor)
			}
		}

		log.Print("RegisteredHaveNotContributed completed")
		return contributorsHavNotCoontriuted, nil
	}

}

func (impl serviceImpl) GetDescendantsByLevel(levelIndex int, parentId string) (interface{}, error) {

	filter := bson.D{bson.E{Key: "contributorid", Value: parentId}}
	parentCategory, _ := impl.categoryService.GetSelectedCategory(filter)
	parentAccount, _ := impl.GetAccountByCategoryId(parentCategory.CategoryId)

	var categorys []categoryDTOPackage.Category
	categorys = append(categorys, parentCategory)
	categories, _ := impl.categoryService.GetCategorys()
	accounts, _ := impl.GetAccounts()
	contributors, _ := impl.contributorService.GetAllContributors()

	var categorysDTO []categoryDTOPackage.Category
	for i := 0; i < levelIndex; i++ {

		var allCategorys []categoryDTOPackage.Category
		for _, item1 := range categorys {
			_categorys := impl.CategoryUtils.FindCategorys(categories, item1.CategoryId)
			allCategorys = append(allCategorys, _categorys...)
		}
		categorys = make([]categoryDTOPackage.Category, 0, len(categories))
		categorys = append(categorys, allCategorys...)

		if i == levelIndex-1 {
			for _, val := range categorys {

				account := impl.AccountUtils.FindAccountByCategoryIdId(accounts, val.CategoryId)
				contributor := impl.UserUtils.FindUserById(contributors, val.ContributorId)

				categorysDTO = append(categorysDTO, categoryDTOPackage.Category{
					EntryDate:              strconv.Itoa(val.CreatedDay) + "/" + strconv.Itoa(val.CreatedMonth) + "/" + strconv.Itoa(val.CreatedYear),
					FullName:               val.FullName,
					ContributorId:          val.CategoryId,
					NLevelXRoomOneChildren: val.NLevelOneRoomOneChildren,
					UserName:               val.UserName,
					CategoryId:             val.CategoryId,
					HasPaid:                account.IsComfirmed,
					PhoneNumber:            contributor.PhoneNumber,
				})
			}
		}
	}

	DashboardDTO := dtos.DashboardDTO{
		CategorysDTO: categorysDTO,
		HasPaid:      parentAccount.IsComfirmed,
		CategoryId:   parentCategory.CategoryId,
	}
	log.Print("GetDescendantsByLevel completed")
	return DashboardDTO, nil
}

func (impl serviceImpl) GetDownliners(parentId string, contributors []contributorDTOPackage.UserResponse,
	accounts []models.Account, categories []categoryDTOPackage.Category, Counter int, levelCount int,
	nLevel int) bool {

	confirmedCategories := make([]categoryDTOPackage.Category, 0)
	for _, chidCategory := range categories {
		if parentId == chidCategory.ParentId {
			for _, account := range accounts {
				if account.ContributorId == chidCategory.CategoryId &&
					account.IsComfirmed {
					confirmedCategories = append(confirmedCategories, chidCategory)
				}
			}
		}
	}

	Counter++
	ans := false
	if Counter > levelCount {
		if nLevel == len(confirmedCategories) {
			return true
		} else {
			return false
		}
	} else {
		for _, confirmedCategory := range confirmedCategories {
			ans = impl.GetDownliners(confirmedCategory.ParentId, contributors,
				accounts, categories, Counter, levelCount, nLevel)
		}
	}
	return ans
}

func (impl serviceImpl) GetCompletedLevelXCategorys(levelIndex int, categoryIndex int) ([]categoryDTOPackage.Category, error) {

	log.Print("GetCompletedLevelXCategorys started")
	nLevel := 3
	returnOnInvestment := helpers.ROIs[categoryIndex-1][levelIndex-1]

	switch levelIndex {
	case 2:
		nLevel = 9

	case 3:
		nLevel = 27

	case 4:
		nLevel = 81
	case 5:
		nLevel = 243
	case 6:
		nLevel = 729
	case 7:
		nLevel = 2187
	}

	categorys, _ := impl.categoryService.GetCompletedLevelXCategories(levelIndex, categoryIndex)
	contributors, _ := impl.contributorService.GetAllContributors()
	accounts, _ := impl.GetAccounts()
	categories, _ := impl.categoryService.GetCategorys()

	var categorysDTO []categoryDTOPackage.Category = make([]categoryDTOPackage.Category, 0)
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

		Counter := 1
		check := impl.GetDownliners(val.Id, contributors,
			accounts, categories, Counter, levelIndex,
			nLevel)

		if check {
			contributor := impl.UserUtils.FindUserById(contributors, val.ContributorId)
			categorysDTO = append(categorysDTO, categoryDTOPackage.Category{
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
	}

	log.Print("GetCompletedLevelXCategorys completed")
	return categorysDTO, nil
}
