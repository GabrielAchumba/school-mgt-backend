package services

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/GabrielAchumba/school-mgt-backend/common/config"
	"github.com/GabrielAchumba/school-mgt-backend/common/conversion"
	networkingerrors "github.com/GabrielAchumba/school-mgt-backend/common/errors"
	"github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/user-module/dtos"
	"github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/user-module/models"
	"github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/user-module/utils"
	userDtos "github.com/GabrielAchumba/school-mgt-backend/modules/user-module/dtos"

	"github.com/GabrielAchumba/school-mgt-backend/authentication/token"
	"github.com/thanhpk/randstr"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService interface {
	LoginUser(requestModel dtos.LoginUserRequest) (interface{}, error)
	RegisterUser(userId string, requestModel dtos.CreateUserRequest) (interface{}, error)
	GetUsers() ([]dtos.UserResponse, error)
	GetUser(id string) (dtos.UserResponse, error)
	PutUser(id string, User dtos.UpdateUserRequest) (interface{}, error)
	PostUser(User dtos.CreateUserRequest) (interface{}, error)
	DeleteUser(id string) (int64, error)
	GetSelectedUser(filter primitive.D) (interface{}, interface{})
	UpdateBioData(id string, bioDataDTO dtos.BioDataDTO) (dtos.BioDataDTO, error)
	UpdateContactDTO(id string, contactDTO dtos.ContactDTO) (dtos.ContactDTO, error)
	UpdateNextOfKinDTO(id string, nextOfKinDTO dtos.NextOfKinDTO) (dtos.NextOfKinDTO, error)
	UpdateBankAccountData(id string, bankAccountDTO dtos.BankAccountDTO) (dtos.BankAccountDTO, error)
	GetContacts() ([]dtos.ContactDTO, error)
	GetBioData() ([]dtos.BioDataDTO, error)
	GetBankDetails() ([]dtos.BankAccountDTO, error)
	GetNextOfKins() ([]dtos.NextOfKinDTO, error)
	GetContributor(id string) (models.User, error)
	GetSelectedContributor(filter primitive.D) (models.User, interface{})
	GetAllContributors() ([]dtos.UserResponse, error)
	GetPersonalDataList() ([]dtos.UserResponse, error)
	GetAdministrators() ([]dtos.UserResponse, error)
	UpdateAdminDTO(id string, adminDTO dtos.AdminDTO) (dtos.AdminDTO, error)
	ForgotPassword(forgotPasswordInput dtos.ForgotPasswordInput) (dtos.ForgotPasswordInput, error)
	ResetPassword(model dtos.ResetPasswordInput) (string, error)
	SeedAdmin()
}

type serviceImpl struct {
	ctx        context.Context
	collection *mongo.Collection
	tokenMaker token.Maker
	emailDto   userDtos.EmailDto
}

func New(mongoClient *mongo.Client, config config.Settings, ctx context.Context,
	tokenMaker token.Maker, emailDto userDtos.EmailDto) UserService {
	collection := mongoClient.Database(config.Database.DatabaseName).Collection(config.TableNames.LaunchpadUser)

	return &serviceImpl{
		collection: collection,
		ctx:        ctx,
		tokenMaker: tokenMaker,
		emailDto:   emailDto,
	}
}

func (impl serviceImpl) SeedAdmin() {
	admin := dtos.CreateUserRequest{
		Password:    "network",
		FirstName:   "admin",
		LastName:    "admin",
		UserName:    "admin@network.com",
		UserType:    "Admin",
		Designation: "CEO",
	}

	filter := bson.D{bson.E{Key: "username", Value: admin.UserName}}
	count, er := impl.collection.CountDocuments(impl.ctx, filter)
	if count == 0 && er == nil {
		impl.RegisterUser("admin", admin)
	}

}

func (impl serviceImpl) LoginUser(requestModel dtos.LoginUserRequest) (interface{}, error) {

	log.Print("Call to login user started.")

	if requestModel.UserName == "" {
		return nil, networkingerrors.Error("UserName cannot be empty.")
	}

	var modelDto dtos.UserInternalOperation

	filter := bson.D{bson.E{Key: "username", Value: requestModel.UserName}}
	er := impl.collection.FindOne(impl.ctx, filter).Decode(&modelDto)
	if er != nil {
		return nil, er // exception.Error("Invalid credentials supplied.")
	}

	credentialError := models.CheckPassword(modelDto.Password, requestModel.Password)
	if credentialError != nil {
		return nil, networkingerrors.Error("Invalid credentials supplied.")
	}

	accessToken, accessPayload, accessError := impl.tokenMaker.CreateToken(modelDto.ID, modelDto.UserName)
	if accessError != nil {
		return nil, networkingerrors.Error("Internal server error.")
	}

	rsp := dtos.LoginUserResponse{
		Token:     accessToken,
		ExpiresAt: accessPayload.ExpiredAt,
		User: dtos.UserResponse{
			Id:           modelDto.ID,
			Title:        modelDto.Title,
			PhoneNumber:  modelDto.PhoneNumber,
			FirstName:    modelDto.FirstName,
			MiddleName:   modelDto.MiddleName,
			LastName:     modelDto.LastName,
			UserName:     modelDto.UserName,
			UserType:     modelDto.UserType,
			Designation:  modelDto.Designation,
			Region:       modelDto.Region,
			Description:  modelDto.Description,
			CreatedAt:    modelDto.CreatedAt,
			Base64String: modelDto.Base64String,
		},
	}

	log.Print("Call to login user completed.")
	return rsp, er
}

func (impl serviceImpl) RegisterUser(userId string, model dtos.CreateUserRequest) (interface{}, error) {

	log.Print("Call to register user started.")

	var modelObj models.User
	conversion.Convert(model, &modelObj)

	modelObj.CreatedBy = userId
	modelObj.CreatedAt = time.Now()

	if modelObj.UserName == "" {
		return nil, networkingerrors.Error("UserName cannot be empty.")
	}
	if modelObj.Password == "" {
		return nil, networkingerrors.Error("Password cannot be empty.")
	}
	if modelObj.FirstName == "" {
		return nil, networkingerrors.Error("FirstName cannot be empty.")
	}
	if modelObj.LastName == "" {
		return nil, networkingerrors.Error("LastName cannot be empty.")
	}

	er := modelObj.HashPassword()
	if er != nil {
		return nil, er
	}

	filter := bson.D{bson.E{Key: "username", Value: modelObj.UserName}}
	count, err := impl.collection.CountDocuments(impl.ctx, filter)
	if err != nil {
		return nil, err //exception.Error("Checking if title exist.")
	}
	if count > 0 {
		return nil, networkingerrors.Error(fmt.Sprintf("UserName '%v'already exist.", model.UserName))
	}
	_, er = impl.collection.InsertOne(impl.ctx, modelObj)
	if er != nil {
		return nil, networkingerrors.Error("Error in registering user.")
	}
	log.Print("Call to register user completed.")
	return modelObj, er
}

func (impl serviceImpl) DeleteUser(id string) (int64, error) {

	log.Print("Call to delete User by id started.")
	objId := conversion.GetMongoId(id)
	filter := bson.D{bson.E{Key: "_id", Value: objId}}

	result, err := impl.collection.DeleteOne(impl.ctx, filter)

	if err != nil {
		return result.DeletedCount, networkingerrors.Error("Error in deleting User.")
	}

	if result.DeletedCount < 1 {
		return result.DeletedCount, networkingerrors.Error("Adminstrator with specified ID not found!")
	}

	log.Print("Call to delete User by id completed.")
	return result.DeletedCount, nil
}

func (impl serviceImpl) GetUser(id string) (dtos.UserResponse, error) {

	log.Print("Get Adminstrator called")
	objId := conversion.GetMongoId(id)
	var User dtos.UserResponse

	filter := bson.D{bson.E{Key: "_id", Value: objId}}

	err := impl.collection.FindOne(impl.ctx, filter).Decode(&User)
	if err != nil {
		return User, networkingerrors.Error("could not find adminstrator by id")
	}

	log.Print("Get Adminstrator completed")
	return User, err

}

func (impl serviceImpl) GetUsers() ([]dtos.UserResponse, error) {

	log.Print("Call to get all Users started.")

	var Users []dtos.UserResponse
	filter := bson.D{}
	cur, err := impl.collection.Find(impl.ctx, filter)

	if err != nil {
		Users = make([]dtos.UserResponse, 0)
		return Users, networkingerrors.Error("Users not found!")
	}

	err = cur.All(impl.ctx, &Users)
	if err != nil {
		return Users, err
	}

	cur.Close(impl.ctx)
	if len(Users) == 0 {
		Users = make([]dtos.UserResponse, 0)
	}

	log.Print("Call to get all Users completed.")
	return Users, err
}

func (impl serviceImpl) PostUser(User dtos.CreateUserRequest) (interface{}, error) {

	log.Print("Call to create User started.")

	var _User models.User
	conversion.Convert(User, &_User)
	CreatedAt := time.Now()
	_User.CreatedDay = CreatedAt.Day()
	_User.CreatedMonth = int(CreatedAt.Month())
	_User.CreatedYear = CreatedAt.Year()

	filter := bson.D{bson.E{Key: "username", Value: _User.UserName}}
	count, err := impl.collection.CountDocuments(impl.ctx, filter)
	if err != nil {
		return nil, networkingerrors.Error("User exists!")
	}

	if count > 0 {
		return nil, networkingerrors.Error("UserName already exist.")
	}

	m, er := impl.collection.InsertOne(impl.ctx, _User)

	if er != nil {
		return nil, networkingerrors.Error("Error in creating User.")
	}
	log.Print("Call to create adminstrator completed.")
	return m.InsertedID, er
}

func (impl serviceImpl) PutUser(id string, User dtos.UpdateUserRequest) (interface{}, error) {

	objId := conversion.GetMongoId(id)
	var updatedUser dtos.UpdateUserRequest
	conversion.Convert(User, &updatedUser)
	filter := bson.D{bson.E{Key: "_id", Value: objId}}
	var oldUser models.User
	err := impl.collection.FindOne(impl.ctx, filter).Decode(&oldUser)
	if err == nil {
		return nil, networkingerrors.Error("User does not exist")
	}

	update := bson.D{bson.E{Key: "createdDay", Value: updatedUser.CreatedDay},
		bson.E{Key: "createdMonth", Value: updatedUser.CreatedMonth},
		bson.E{Key: "createdYear", Value: updatedUser.CreatedYear},
		bson.E{Key: "designation", Value: updatedUser.Designation},
		bson.E{Key: "firstName", Value: updatedUser.FirstName},
		bson.E{Key: "isPhotographUploaded", Value: updatedUser.IsPhotographUploaded},
		bson.E{Key: "lastName", Value: updatedUser.LastName},
		bson.E{Key: "middleName", Value: updatedUser.MiddleName},
		bson.E{Key: "password", Value: updatedUser.Password},
		bson.E{Key: "phoneNumber", Value: updatedUser.PhoneNumber},
		bson.E{Key: "region", Value: updatedUser.Region},
		bson.E{Key: "title", Value: updatedUser.Title},
		bson.E{Key: "username", Value: updatedUser.UserName},
		bson.E{Key: "userType", Value: updatedUser.UserType}}
	result, er := impl.collection.UpdateOne(impl.ctx, filter, bson.D{bson.E{Key: "$set", Value: update}})
	return result.UpsertedID, er
}

func (impl serviceImpl) GetSelectedUser(filter primitive.D) (interface{}, interface{}) {

	log.Print("GetSelectedUser started")
	var User models.User
	//var User dtos.UserResponse

	log.Print("filter: ", filter)
	//log.Print("UserServices: ", impl.UserServices)

	err := impl.collection.FindOne(impl.ctx, filter).Decode(&User)
	if err != nil {
		return User, "No User" //networkingerrors.Error("could not find selected User")
	}

	log.Print("GetSelectedUser completed")
	return User, err
}

func (impl serviceImpl) GetContacts() ([]dtos.ContactDTO, error) {

	log.Print("GetContacts started")
	var contributors []dtos.UserResponse
	filter := bson.D{bson.E{Key: "usertype", Value: "Member"}}
	curr, err := impl.collection.Find(impl.ctx, filter)
	if err != nil {
		return make([]dtos.ContactDTO, 0),
			networkingerrors.Error("Could not get all contacts")
	}

	err = curr.All(impl.ctx, &contributors)
	if err != nil {
		return make([]dtos.ContactDTO, 0),
			networkingerrors.Error("Could not decode all contributors")
	}

	curr.Close(impl.ctx)
	if len(contributors) == 0 {
		return make([]dtos.ContactDTO, 0), nil
	}

	var contactsDTO []dtos.ContactDTO
	for _, val := range contributors {
		contactsDTO = append(contactsDTO, dtos.ContactDTO{
			FullName:         val.FirstName + " " + val.MiddleName + " " + val.LastName,
			Address:          val.Address,
			ResidentialCity:  val.ResidentialCity,
			ResidentialState: val.ResidentialState,
			Email:            val.Email,
			PhoneNumber:      val.PhoneNumber,
			ContributorId:    val.Id,
		})
	}

	log.Print("GetContacts completed")
	return contactsDTO, nil
}

func (impl serviceImpl) GetBioData() ([]dtos.BioDataDTO, error) {

	log.Print("GetBioData started")
	var contributors []dtos.UserResponse
	filter := bson.D{bson.E{Key: "usertype", Value: "Member"}}
	curr, err := impl.collection.Find(impl.ctx, filter)
	if err != nil {
		return make([]dtos.BioDataDTO, 0),
			networkingerrors.Error("Could not get all biodata")
	}

	err = curr.All(impl.ctx, &contributors)
	if err != nil {
		return make([]dtos.BioDataDTO, 0),
			networkingerrors.Error("Could not decode all biodata")
	}

	curr.Close(impl.ctx)
	if len(contributors) == 0 {
		return make([]dtos.BioDataDTO, 0), nil
	}

	var biodataDTO []dtos.BioDataDTO
	for _, val := range contributors {
		biodataDTO = append(biodataDTO, dtos.BioDataDTO{
			FullName:             val.FirstName + " " + val.MiddleName + " " + val.LastName,
			Base64String:         val.Base64String,
			IsPhotographUploaded: val.IsPhotographUploaded,
			BloodGroup:           val.BloodGroup,
			Genotype:             val.Genotype,
			MaritalStatus:        val.MaritalStatus,
			LGAOfOrigin:          val.LGAOfOrigin,
			StateOfOrigin:        val.StateOfOrigin,
			Country:              val.Country,
			ContributorId:        val.Id,
		})
	}

	log.Print("GetBioData completed")
	return biodataDTO, nil
}

func (impl serviceImpl) GetBankDetails() ([]dtos.BankAccountDTO, error) {

	log.Print("GetBankDetails started")
	var contributors []dtos.UserResponse
	filter := bson.D{bson.E{Key: "usertype", Value: "Member"}}
	curr, err := impl.collection.Find(impl.ctx, filter)
	if err != nil {
		return make([]dtos.BankAccountDTO, 0),
			networkingerrors.Error("Could not get all bank details")
	}

	err = curr.All(impl.ctx, &contributors)
	if err != nil {
		return make([]dtos.BankAccountDTO, 0),
			networkingerrors.Error("Could not decode all bank details")
	}

	curr.Close(impl.ctx)
	if len(contributors) == 0 {
		return make([]dtos.BankAccountDTO, 0), nil
	}

	var bankDetailsDTO []dtos.BankAccountDTO
	for _, val := range contributors {
		bankDetailsDTO = append(bankDetailsDTO, dtos.BankAccountDTO{
			FullName:      val.FirstName + " " + val.MiddleName + " " + val.LastName,
			BankName:      val.BankName,
			AccountName:   val.AccountName,
			AccountNumber: val.AccountNumber,
			BVN:           val.BVN,
			ContributorId: val.Id,
		})
	}

	log.Print("GetBankDetails completed")
	return bankDetailsDTO, nil
}

func (impl serviceImpl) GetNextOfKins() ([]dtos.NextOfKinDTO, error) {

	log.Print("GetBankDetails started")
	var contributors []dtos.UserResponse
	filter := bson.D{bson.E{Key: "usertype", Value: "Member"}}
	curr, err := impl.collection.Find(impl.ctx, filter)
	if err != nil {
		return make([]dtos.NextOfKinDTO, 0),
			networkingerrors.Error("Could not get all next of kins")
	}

	err = curr.All(impl.ctx, &contributors)
	if err != nil {
		return make([]dtos.NextOfKinDTO, 0),
			networkingerrors.Error("Could not decode all next of kins")
	}

	curr.Close(impl.ctx)
	if len(contributors) == 0 {
		return make([]dtos.NextOfKinDTO, 0), nil
	}

	var nextOfKinsDTO []dtos.NextOfKinDTO
	for _, val := range contributors {
		nextOfKinsDTO = append(nextOfKinsDTO, dtos.NextOfKinDTO{
			FullName:        val.FirstName + " " + val.MiddleName + " " + val.LastName,
			NOKNames:        val.NOKNames,
			NOKAddress:      val.NOKAddress,
			NOKPhoneNumber:  val.NOKPhoneNumber,
			NOKRelationship: val.NOKRelationship,
			ContributorId:   val.Id,
		})
	}

	log.Print("GetNextOfKins completed")
	return nextOfKinsDTO, nil
}

func (impl serviceImpl) UpdateBankAccountData(id string, bankAccountDTO dtos.BankAccountDTO) (dtos.BankAccountDTO, error) {

	log.Print("UpdateBankAccountData started")
	objId := conversion.GetMongoId(id)
	var modelObj dtos.BankAccountDTO
	conversion.Convert(bankAccountDTO, &modelObj)
	filter := bson.D{bson.E{Key: "_id", Value: objId}}

	update := bson.D{bson.E{Key: "accountName", Value: modelObj.AccountName},
		bson.E{Key: "accountNumber", Value: modelObj.AccountNumber},
		bson.E{Key: "bVN", Value: modelObj.BVN},
		bson.E{Key: "bankName", Value: modelObj.BankName}}
	_, err := impl.collection.UpdateOne(impl.ctx, filter, bson.D{bson.E{Key: "$set", Value: update}})

	if err != nil {
		return modelObj, networkingerrors.Error("Could not upadte contributor's account details")
	}

	log.Print("UpdateBankAccountData completed")
	return modelObj, nil
}

func (impl serviceImpl) UpdateBioData(id string, bioDataDTO dtos.BioDataDTO) (dtos.BioDataDTO, error) {

	//bson.E{Key: "isPhotographUploaded", Value: modelObj.IsPhotographUploaded}
	log.Print("UpdateBioData started")
	objId := conversion.GetMongoId(id)
	var modelObj dtos.BioDataDTO
	conversion.Convert(bioDataDTO, &modelObj)
	log.Print("modelObj: ", modelObj)
	log.Print("objId:", objId)
	filter := bson.D{bson.E{Key: "_id", Value: objId}}

	update := bson.D{bson.E{Key: "bloodGroup", Value: modelObj.BloodGroup},
		bson.E{Key: "country", Value: modelObj.Country},
		bson.E{Key: "genotype", Value: modelObj.Genotype},
		bson.E{Key: "lGAOfOrigin", Value: modelObj.LGAOfOrigin},
		bson.E{Key: "maritalStatus", Value: modelObj.MaritalStatus},
		bson.E{Key: "stateOfOrigin", Value: modelObj.StateOfOrigin},
		bson.E{Key: "base64String", Value: modelObj.Base64String}}

	_, err := impl.collection.UpdateOne(impl.ctx, filter, bson.D{bson.E{Key: "$set", Value: update}})

	if err != nil {
		return modelObj, networkingerrors.Error("Could not upadte contributor's bio-data details")
	}

	log.Print("UpdateBioData completed")
	return modelObj, nil
}

func (impl serviceImpl) UpdateContactDTO(id string, contactDTO dtos.ContactDTO) (dtos.ContactDTO, error) {

	log.Print("UpdateContactDTO started")
	objId := conversion.GetMongoId(id)
	var modelObj dtos.ContactDTO
	conversion.Convert(contactDTO, &modelObj)
	filter := bson.D{bson.E{Key: "_id", Value: objId}}
	update := bson.D{bson.E{Key: "address", Value: modelObj.Address},
		bson.E{Key: "email", Value: modelObj.Email},
		bson.E{Key: "phoneNumber", Value: modelObj.PhoneNumber},
		bson.E{Key: "residentialCity", Value: modelObj.ResidentialCity},
		bson.E{Key: "residentialState", Value: modelObj.ResidentialState},
		bson.E{Key: "base64String", Value: modelObj.Base64String}}
	_, err := impl.collection.UpdateOne(impl.ctx, filter, bson.D{bson.E{Key: "$set", Value: update}})

	if err != nil {
		return modelObj, networkingerrors.Error("Could not upadte contributor's contact details")
	}

	log.Print("UpdateContactDTO completed")
	return modelObj, nil
}

func (impl serviceImpl) UpdateNextOfKinDTO(id string, nextOfKinDTO dtos.NextOfKinDTO) (dtos.NextOfKinDTO, error) {

	log.Print("UpdateNextOfKinDTO started")
	objId := conversion.GetMongoId(id)
	var modelObj dtos.NextOfKinDTO
	conversion.Convert(nextOfKinDTO, &modelObj)
	filter := bson.D{bson.E{Key: "_id", Value: objId}}
	update := bson.D{bson.E{Key: "nOKAddress", Value: modelObj.NOKAddress},
		bson.E{Key: "nOKNames", Value: modelObj.NOKNames},
		bson.E{Key: "nOKPhoneNumber", Value: modelObj.NOKPhoneNumber},
		bson.E{Key: "nOKRelationship", Value: modelObj.NOKRelationship}}
	_, err := impl.collection.UpdateOne(impl.ctx, filter, bson.D{bson.E{Key: "$set", Value: update}})

	if err != nil {
		return modelObj, networkingerrors.Error("Could not upadte contributor's next of kin's details")
	}

	log.Print("UpdateNextOfKinDTO completed")
	return modelObj, nil
}

func (impl serviceImpl) GetContributor(id string) (models.User, error) {

	log.Print("GetContributor started")
	objId := conversion.GetMongoId(id)
	var contributor models.User

	filter := bson.D{bson.E{Key: "_id", Value: objId}}

	err := impl.collection.FindOne(impl.ctx, filter).Decode(&contributor)
	if err != nil {
		return contributor, networkingerrors.Error("could not find contributor by id")
	}

	log.Print("GetContributor completed")
	return contributor, err
}

func (impl serviceImpl) GetSelectedContributor(filter primitive.D) (models.User, interface{}) {

	log.Print("GetSelectedContributor started")
	var contributor models.User

	err := impl.collection.FindOne(impl.ctx, filter).Decode(&contributor)
	if err != nil {
		return contributor, "No Contributor" //networkingerrors.Error("could not find contributor by any key expect _id")
	}

	log.Print("GetSelectedContributor completed")
	return contributor, err
}

func (impl serviceImpl) GetAllContributors() ([]dtos.UserResponse, error) {

	log.Print("GetContributors started")
	var contributors []dtos.UserResponse
	filter := bson.D{bson.E{Key: "usertype", Value: "Member"}}
	curr, err := impl.collection.Find(impl.ctx, filter)
	if err != nil {
		return make([]dtos.UserResponse, 0),
			networkingerrors.Error("Could not get all contributors")
	}

	err = curr.All(impl.ctx, &contributors)
	if err != nil {
		return make([]dtos.UserResponse, 0),
			networkingerrors.Error("Could not decode all contributors")
	}

	curr.Close(impl.ctx)
	if len(contributors) == 0 {
		return make([]dtos.UserResponse, 0), nil
	}

	var contributorsDTO []dtos.UserResponse
	for _, val := range contributors {
		contributorsDTO = append(contributorsDTO, dtos.UserResponse{
			FullName:      val.FirstName + " " + val.MiddleName + " " + val.LastName,
			Gender:        val.Gender,
			ContributorId: val.Id,
			BankName:      val.BankName,
			AccountName:   val.AccountName,
			AccountNumber: val.AccountNumber,
			PhoneNumber:   val.PhoneNumber,
		})
	}

	log.Print("GetContributors completed")
	return contributorsDTO, nil
}

func (impl serviceImpl) GetPersonalDataList() ([]dtos.UserResponse, error) {

	log.Print("GetPersonalDataList started")
	var contributors []dtos.UserResponse
	filter := bson.D{bson.E{Key: "usertype", Value: "Member"}}
	curr, err := impl.collection.Find(impl.ctx, filter)
	if err != nil {
		return make([]dtos.UserResponse, 0),
			networkingerrors.Error("Could not get all personal data list")
	}

	err = curr.All(impl.ctx, &contributors)
	if err != nil {
		return make([]dtos.UserResponse, 0),
			networkingerrors.Error("Could not decode all contributors")
	}

	curr.Close(impl.ctx)
	if len(contributors) == 0 {
		return make([]dtos.UserResponse, 0), nil
	}

	var PersonalDataListDTO []dtos.UserResponse
	for _, contributor := range contributors {
		PersonalDataListDTO = append(PersonalDataListDTO, dtos.UserResponse{
			CreatedDay:    contributor.CreatedDay,
			CreatedMonth:  contributor.CreatedMonth,
			CreatedYear:   contributor.CreatedYear,
			ContributorId: contributor.Id,
			FirstName:     contributor.FirstName,
			MiddleName:    contributor.MiddleName,
			LastName:      contributor.LastName,
			Gender:        contributor.Gender,
			UserName:      contributor.UserName,
			Password:      contributor.Password,
			UserType:      contributor.UserType,
		})
	}

	log.Print("GetPersonalDataList completed")
	return PersonalDataListDTO, nil
}

func (impl serviceImpl) GetAdministrators() ([]dtos.UserResponse, error) {

	log.Print("GetAdministrators started")
	var contributors []dtos.UserResponse
	filter := bson.D{bson.E{Key: "usertype", Value: "Admin"}}
	curr, err := impl.collection.Find(impl.ctx, filter)
	if err != nil {
		return make([]dtos.UserResponse, 0),
			networkingerrors.Error("Could not get all administrators")
	}

	err = curr.All(impl.ctx, &contributors)
	if err != nil {
		return make([]dtos.UserResponse, 0),
			networkingerrors.Error("Could not decode all administrators")
	}

	curr.Close(impl.ctx)
	if len(contributors) == 0 {
		return make([]dtos.UserResponse, 0), nil
	}

	var contributorsDTO []dtos.UserResponse
	for _, val := range contributors {
		contributorsDTO = append(contributorsDTO, dtos.UserResponse{
			FirstName:     val.FirstName,
			LastName:      val.LastName,
			ContributorId: val.Id,
			Base64String:  val.Base64String,
			UserName:      val.UserName,
			UserType:      val.UserType,
			Designation:   val.Designation,
		})
	}

	log.Print("GetAdministrators completed")
	return contributorsDTO, nil
}

func (impl serviceImpl) UpdateAdminDTO(id string, adminDTO dtos.AdminDTO) (dtos.AdminDTO, error) {

	log.Print("UpdateAdminDTO started")
	objId := conversion.GetMongoId(id)
	var modelObj dtos.AdminDTO
	conversion.Convert(adminDTO, &modelObj)
	filter := bson.D{bson.E{Key: "_id", Value: objId}}
	update := bson.D{bson.E{Key: "firstname", Value: modelObj.FirstName},
		bson.E{Key: "lastname", Value: modelObj.LastName},
		bson.E{Key: "username", Value: modelObj.UserName},
		bson.E{Key: "usertype", Value: modelObj.UserType},
		bson.E{Key: "base64string", Value: modelObj.Base64String},
		bson.E{Key: "designation", Value: modelObj.Designation}}
	_, err := impl.collection.UpdateOne(impl.ctx, filter, bson.D{bson.E{Key: "$set", Value: update}})

	if err != nil {
		return modelObj, networkingerrors.Error("Could not upadte adminstrator's details")
	}

	log.Print("UpdateAdminDTO completed")
	return modelObj, nil
}

func (impl serviceImpl) ForgotPassword(forgotPasswordInput dtos.ForgotPasswordInput) (dtos.ForgotPasswordInput, error) {

	//message := "You will receive a reset email if user with that email exist. So check your email inbox/junk. Thanks."
	var user models.User
	filter := bson.D{bson.E{Key: "username", Value: forgotPasswordInput.UserName}}
	err := impl.collection.FindOne(impl.ctx, filter).Decode(&user)

	if err != nil {
		return dtos.ForgotPasswordInput{}, networkingerrors.Error("User does not exist")
	}

	if forgotPasswordInput.Email == "" {
		return dtos.ForgotPasswordInput{}, networkingerrors.Error("Email address must exist")
	}

	if forgotPasswordInput.Email != user.Email {
		return dtos.ForgotPasswordInput{}, networkingerrors.Error("In-coming email address is different from already existing email address")
	}

	// Generate Verification Code
	resetToken := randstr.String(20)

	passwordResetToken := utils.Encode(resetToken)

	// Update User in Database
	query := bson.D{{Key: "username", Value: user.UserName}}
	update := bson.D{{Key: "$set", Value: bson.D{
		bson.E{Key: "passwordresettoken", Value: passwordResetToken},
		bson.E{Key: "passwordresetat", Value: time.Now().Add(time.Minute * 15)},
	}}}
	result, err := impl.collection.UpdateOne(impl.ctx, query, update)

	if result.MatchedCount == 0 {
		return dtos.ForgotPasswordInput{}, networkingerrors.Error("There was an error sending email")
	}

	if err != nil {
		return dtos.ForgotPasswordInput{}, networkingerrors.Error("There was an error sending email")
	}
	var firstName = user.FirstName

	if strings.Contains(firstName, " ") {
		firstName = strings.Split(firstName, " ")[1]
	}

	// 👇 Send Email
	emailData := utils.EmailData{
		URL:       impl.emailDto.Origin + "/forgotPassword/" + resetToken,
		FirstName: firstName,
		Subject:   "Your password reset token (valid for 10min)",
	}

	err = utils.SendEmail(forgotPasswordInput.Email, &emailData, "resetPassword.html",
		impl.emailDto)
	if err != nil {
		return dtos.ForgotPasswordInput{}, networkingerrors.Error("There was an error sending email")
	}

	return dtos.ForgotPasswordInput{
		ResetToken: resetToken,
		Message: `You will receive a reset email if your email address is valid. 
		Please check your email address' inbox/junk. Thanks`,
	}, nil
}

func (impl serviceImpl) ResetPassword(userCredential dtos.ResetPasswordInput) (string, error) {

	if userCredential.Password != userCredential.PasswordConfirm {
		return "Error", networkingerrors.Error("Passwords do not match")
	}

	var modelObj = models.User{
		Password: userCredential.Password,
	}

	er := modelObj.HashPassword()
	if er != nil {
		return "Error", er
	}

	resetToken := userCredential.ResetToken

	var user models.User
	filter := bson.D{bson.E{Key: "username", Value: userCredential.UserName}}
	err := impl.collection.FindOne(impl.ctx, filter).Decode(&user)

	if err != nil {
		return "Error", networkingerrors.Error("Could not upadte adminstrator's details")
	}

	passwordResetToken, err := utils.Decode(user.PasswordResetToken)

	if err != nil {
		return "Error", networkingerrors.Error("Invalid or expired token")
	}

	if passwordResetToken != resetToken {
		return "Error", networkingerrors.Error("Invalid or expired token")
	}

	now := time.Now()
	if now.Sub(user.PasswordResetAt).Minutes() > 10 {
		return "Error", networkingerrors.Error("Toke life expired. Please generate another one")
	}

	// Update User in Database
	query := bson.D{{Key: "username", Value: userCredential.UserName}}
	update := bson.D{{Key: "$set", Value: bson.D{
		bson.E{Key: "password", Value: modelObj.Password},
		bson.E{Key: "passwordresettoken", Value: ""},
		bson.E{Key: "passwordresetat", Value: now}}}}

	_, err = impl.collection.UpdateOne(impl.ctx, query, update)

	if err != nil {
		return "Error", networkingerrors.Error("Could not update password")
	}

	return "Password data updated successfully", nil
}
