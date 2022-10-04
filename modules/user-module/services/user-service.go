package services

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/GabrielAchumba/school-mgt-backend/common/config"
	"github.com/GabrielAchumba/school-mgt-backend/common/conversion"
	errors "github.com/GabrielAchumba/school-mgt-backend/common/errors"
	"github.com/GabrielAchumba/school-mgt-backend/modules/user-module/dtos"
	"github.com/GabrielAchumba/school-mgt-backend/modules/user-module/models"
	"github.com/GabrielAchumba/school-mgt-backend/modules/user-module/utils"

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
	GetUsersByCategory(category string) ([]dtos.UserResponse, error)
	PutUser(id string, User dtos.UpdateUserRequest) (interface{}, error)
	PostUser(User dtos.CreateUserRequest) (interface{}, error)
	DeleteUser(id string) (int64, error)
	GetSelectedUser(filter primitive.D) (interface{}, interface{})
	UpdateAdminDTO(id string, adminDTO dtos.AdminDTO) (dtos.AdminDTO, error)
	ForgotPassword(forgotPasswordInput dtos.ForgotPasswordInput) (dtos.ForgotPasswordInput, error)
	ResetPassword(model dtos.ResetPasswordInput) (string, error)
	SeedAdmin()
}

type serviceImpl struct {
	ctx        context.Context
	collection *mongo.Collection
	tokenMaker token.Maker
	emailDto   dtos.EmailDto
}

func New(mongoClient *mongo.Client, config config.Settings, ctx context.Context,
	tokenMaker token.Maker, emailDto dtos.EmailDto) UserService {
	collection := mongoClient.Database(config.Database.DatabaseName).Collection(config.TableNames.User)

	return &serviceImpl{
		collection: collection,
		ctx:        ctx,
		tokenMaker: tokenMaker,
		emailDto:   emailDto,
	}
}

func (impl serviceImpl) SeedAdmin() {
	admin := dtos.CreateUserRequest{
		Password:    "school",
		FirstName:   "admin",
		LastName:    "admin",
		PhoneNumber: "07032488605",
		CountryCode: "+234",
		UserName:    "admin@school.com",
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
		return nil, errors.Error("UserName cannot be empty.")
	}

	var modelDto dtos.UserInternalOperation

	filter := bson.D{bson.E{Key: "username", Value: requestModel.UserName}}
	er := impl.collection.FindOne(impl.ctx, filter).Decode(&modelDto)
	if er != nil {
		return nil, er // exception.Error("Invalid credentials supplied.")
	}

	credentialError := models.CheckPassword(modelDto.Password, requestModel.Password)
	if credentialError != nil {
		return nil, errors.Error("Invalid credentials supplied.")
	}

	accessToken, accessPayload, accessError := impl.tokenMaker.CreateToken(modelDto.ID, modelDto.UserName)
	if accessError != nil {
		return nil, errors.Error("Internal server error.")
	}

	rsp := dtos.LoginUserResponse{
		Token:     accessToken,
		ExpiresAt: accessPayload.ExpiredAt,
		User: dtos.UserResponse{
			Id:           modelDto.ID,
			PhoneNumber:  modelDto.PhoneNumber,
			FirstName:    modelDto.FirstName,
			LastName:     modelDto.LastName,
			UserName:     modelDto.UserName,
			UserType:     modelDto.UserType,
			Designation:  modelDto.Designation,
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
		return nil, errors.Error("UserName cannot be empty.")
	}
	if modelObj.Password == "" {
		return nil, errors.Error("Password cannot be empty.")
	}
	if modelObj.FirstName == "" {
		return nil, errors.Error("FirstName cannot be empty.")
	}
	if modelObj.LastName == "" {
		return nil, errors.Error("LastName cannot be empty.")
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
		return nil, errors.Error(fmt.Sprintf("UserName '%v'already exist.", model.UserName))
	}
	_, er = impl.collection.InsertOne(impl.ctx, modelObj)
	if er != nil {
		return nil, errors.Error("Error in registering user.")
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
		return result.DeletedCount, errors.Error("Error in deleting User.")
	}

	if result.DeletedCount < 1 {
		return result.DeletedCount, errors.Error("Adminstrator with specified ID not found!")
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
		return User, errors.Error("could not find adminstrator by id")
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
		return Users, errors.Error("Users not found!")
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

func (impl serviceImpl) GetUsersByCategory(category string) ([]dtos.UserResponse, error) {

	log.Print("Call to get Users by category started.")

	var Users []dtos.UserResponse
	filter := bson.D{bson.E{Key: "designation", Value: category}}
	cur, err := impl.collection.Find(impl.ctx, filter)

	if err != nil {
		Users = make([]dtos.UserResponse, 0)
		return Users, errors.Error("Users not found!")
	}

	err = cur.All(impl.ctx, &Users)
	if err != nil {
		return Users, err
	}

	cur.Close(impl.ctx)
	if len(Users) == 0 {
		Users = make([]dtos.UserResponse, 0)
	}

	log.Print("Call to get Users by category completed.")
	return Users, err
}

func (impl serviceImpl) PostUser(User dtos.CreateUserRequest) (interface{}, error) {

	log.Print("Call to create User started.")

	var _User models.User
	conversion.Convert(User, &_User)

	filter := bson.D{bson.E{Key: "username", Value: _User.UserName}}
	count, err := impl.collection.CountDocuments(impl.ctx, filter)
	if err != nil {
		return nil, errors.Error("User exists!")
	}

	if count > 0 {
		return nil, errors.Error("UserName already exist.")
	}

	m, er := impl.collection.InsertOne(impl.ctx, _User)

	if er != nil {
		return nil, errors.Error("Error in creating User.")
	}
	log.Print("Call to create adminstrator completed.")
	return m.InsertedID, er
}

func (impl serviceImpl) PutUser(id string, User dtos.UpdateUserRequest) (interface{}, error) {

	log.Print("PutUser started")

	objId := conversion.GetMongoId(id)
	var updatedUser dtos.UpdateUserRequest
	conversion.Convert(User, &updatedUser)
	filter := bson.D{bson.E{Key: "_id", Value: objId}}
	var modelObj models.User

	update := bson.D{bson.E{Key: "designation", Value: updatedUser.Designation},
		bson.E{Key: "firstName", Value: updatedUser.FirstName},
		bson.E{Key: "isPhotographUploaded", Value: updatedUser.IsPhotographUploaded},
		bson.E{Key: "lastName", Value: updatedUser.LastName},
		bson.E{Key: "password", Value: updatedUser.Password},
		bson.E{Key: "phoneNumber", Value: updatedUser.PhoneNumber},
		bson.E{Key: "username", Value: updatedUser.UserName},
		bson.E{Key: "userType", Value: updatedUser.UserType}}

	_, err := impl.collection.UpdateOne(impl.ctx, filter, bson.D{bson.E{Key: "$set", Value: update}})
	if err != nil {
		return modelObj, errors.Error("Could not upadte user")
	}

	log.Print("PutUser completed")
	return modelObj, nil
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
		return modelObj, errors.Error("Could not upadte adminstrator's details")
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
		return dtos.ForgotPasswordInput{}, errors.Error("User does not exist")
	}

	if forgotPasswordInput.Email == "" {
		return dtos.ForgotPasswordInput{}, errors.Error("Email address must exist")
	}

	if forgotPasswordInput.Email != user.Email {
		return dtos.ForgotPasswordInput{}, errors.Error("In-coming email address is different from already existing email address")
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
		return dtos.ForgotPasswordInput{}, errors.Error("There was an error sending email")
	}

	if err != nil {
		return dtos.ForgotPasswordInput{}, errors.Error("There was an error sending email")
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
		return dtos.ForgotPasswordInput{}, errors.Error("There was an error sending email")
	}

	return dtos.ForgotPasswordInput{
		ResetToken: resetToken,
		Message: `You will receive a reset email if your email address is valid. 
		Please check your email address' inbox/junk. Thanks`,
	}, nil
}

func (impl serviceImpl) ResetPassword(userCredential dtos.ResetPasswordInput) (string, error) {

	if userCredential.Password != userCredential.PasswordConfirm {
		return "Error", errors.Error("Passwords do not match")
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
		return "Error", errors.Error("Could not upadte adminstrator's details")
	}

	passwordResetToken, err := utils.Decode(user.PasswordResetToken)

	if err != nil {
		return "Error", errors.Error("Invalid or expired token")
	}

	if passwordResetToken != resetToken {
		return "Error", errors.Error("Invalid or expired token")
	}

	now := time.Now()
	if now.Sub(user.PasswordResetAt).Minutes() > 10 {
		return "Error", errors.Error("Toke life expired. Please generate another one")
	}

	// Update User in Database
	query := bson.D{{Key: "username", Value: userCredential.UserName}}
	update := bson.D{{Key: "$set", Value: bson.D{
		bson.E{Key: "password", Value: modelObj.Password},
		bson.E{Key: "passwordresettoken", Value: ""},
		bson.E{Key: "passwordresetat", Value: now}}}}

	_, err = impl.collection.UpdateOne(impl.ctx, query, update)

	if err != nil {
		return "Error", errors.Error("Could not update password")
	}

	return "Password data updated successfully", nil
}
