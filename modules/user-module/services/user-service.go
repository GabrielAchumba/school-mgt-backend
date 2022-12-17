package services

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/GabrielAchumba/school-mgt-backend/common/config"
	"github.com/GabrielAchumba/school-mgt-backend/common/conversion"
	errors "github.com/GabrielAchumba/school-mgt-backend/common/errors"
	staffServicePackage "github.com/GabrielAchumba/school-mgt-backend/modules/staff-module/services"
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
	RegisterUsers(userId string, requestModel []dtos.CreateUserRequest) (interface{}, error)
	UserIsExist(requestModel dtos.CreateUserRequest) (interface{}, error)
	UserIsExist2(requestModel dtos.CreateUserRequest) (interface{}, error)
	RegisterAdminOrReferal(model dtos.CreateUserRequest) (interface{}, error)
	GetUsers(schoolId string) ([]dtos.UserResponse, error)
	GetStudents(schoolId string) ([]dtos.UserResponse, error)
	GetUser(id string, schoolId string) (dtos.UserResponse, error)
	GetStudent(id string, schoolId string) (dtos.UserResponse, error)
	GetUsersByCategory(category string, schoolId string) ([]dtos.UserResponse, error)
	GetRerals() ([]dtos.UserResponse, error)
	PutUser(id string, User dtos.UpdateUserRequest) (interface{}, error)
	PutReferal(id string, User dtos.UpdateUserRequest) (interface{}, error)
	PostUser(User dtos.CreateUserRequest) (interface{}, error)
	DeleteUser(id string, schoolId string) (int64, error)
	GetSelectedUser(filter primitive.D) (interface{}, interface{})
	UpdateAdminDTO(id string, adminDTO dtos.AdminDTO) (dtos.AdminDTO, error)
	ForgotPassword(forgotPasswordInput dtos.ForgotPasswordInput) (dtos.ForgotPasswordInput, error)
	ResetPassword(model dtos.ResetPasswordInput) (string, error)
	SeedAdmin()
	GenerateTokens(studentIds []string) (interface{}, error)
	GetSelecedStudents(Ids []string) ([]dtos.UserResponse, error)
	GetStudentsByClassRooms(schoolId string, levelId string,
		classRoomIds []string, sessionId string) ([]dtos.UserResponse, error)
	GetStudentByToken(token int, schoolId string) (dtos.UserResponse, error)
	LogInStudent(token int, schoolId string) (dtos.LoginUserResponse, error)
}

type serviceImpl struct {
	ctx          context.Context
	collection   *mongo.Collection
	tokenMaker   token.Maker
	emailDto     dtos.EmailDto
	staffService staffServicePackage.StaffService
	utils        utils.NumericTokenGenerator
}

func New(mongoClient *mongo.Client, config config.Settings, ctx context.Context,
	tokenMaker token.Maker, emailDto dtos.EmailDto,
	staffService staffServicePackage.StaffService) UserService {
	collection := mongoClient.Database(config.Database.DatabaseName).Collection(config.TableNames.User)

	return &serviceImpl{
		collection:   collection,
		ctx:          ctx,
		tokenMaker:   tokenMaker,
		emailDto:     emailDto,
		staffService: staffService,
	}
}

func (impl serviceImpl) SeedAdmin() {
	admin := dtos.CreateUserRequest{
		Password:      "school",
		FirstName:     "admin",
		LastName:      "admin",
		PhoneNumber:   "07032488605",
		CountryCode:   "+234",
		UserName:      "admin@school.com",
		UserType:      "Admin",
		DesignationId: "CEO",
		SchoolId:      "CEO",
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
			Id:            modelDto.ID,
			PhoneNumber:   modelDto.PhoneNumber,
			FirstName:     modelDto.FirstName,
			LastName:      modelDto.LastName,
			UserName:      modelDto.UserName,
			UserType:      modelDto.UserType,
			DesignationId: modelDto.DesignationId,
			CreatedAt:     modelDto.CreatedAt,
			Base64String:  modelDto.Base64String,
			SchoolId:      modelDto.SchoolId,
			CountryCode:   modelDto.CountryCode,
			SessionId:     modelDto.SessionId,
			LevelId:       modelDto.LevelId,
			ClassRoomId:   modelDto.ClassRoomId,
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

func (impl serviceImpl) RegisterUsers(userId string, _models []dtos.CreateUserRequest) (interface{}, error) {

	log.Print("Call to create users started.")

	usernames := make([]string, 0)
	var Users []dtos.UserResponse
	for _, model := range _models {
		usernames = append(usernames, model.UserName)
	}

	filter := bson.D{{Key: "username", Value: bson.D{
		bson.E{Key: "$in", Value: usernames}}}}

	cur, _ := impl.collection.Find(impl.ctx, filter)

	_ = cur.All(impl.ctx, &Users)
	cur.Close(impl.ctx)

	modelObjs := make([]interface{}, 0)
	for _, model := range _models {
		var modelObj models.User
		modelObj.CreatedBy = userId
		modelObj.CreatedAt = time.Now()
		check := false
		for _, user := range Users {
			if model.UserName == user.UserName {
				check = true
				break
			}
		}

		if !check {
			if model.UserName == "" {
				continue
			}
			if model.Password == "" {
				continue
			}
			if model.FirstName == "" {
				continue
			}
			if model.LastName == "" {
				continue
			}

			conversion.Convert(model, &modelObj)
			er := modelObj.HashPassword()
			if er != nil {
				continue
			}
			modelObjs = append(modelObjs, modelObj)
		}
	}

	_, er := impl.collection.InsertMany(impl.ctx, modelObjs)
	if er != nil {
		return nil, errors.Error("Error in creating users.")
	}
	log.Print("Call to create users completed.")
	return modelObjs, er
}

func (impl serviceImpl) RegisterAdminOrReferal(model dtos.CreateUserRequest) (interface{}, error) {

	log.Print("Call to register admin started.")

	var modelObj models.User
	conversion.Convert(model, &modelObj)

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
	log.Print("Call to register admin completed.")
	return modelObj, er
}

func (impl serviceImpl) UserIsExist(requestModel dtos.CreateUserRequest) (interface{}, error) {

	log.Print("UserIsExist started.")

	var modelObj models.User
	conversion.Convert(requestModel, &modelObj)

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
		return true, nil
	}

	log.Print("UserIsExist completed.")
	return false, er
}

func (impl serviceImpl) UserIsExist2(requestModel dtos.CreateUserRequest) (interface{}, error) {

	log.Print("UserIsExist2 started.")

	var modelObj models.User
	conversion.Convert(requestModel, &modelObj)

	modelObj.CreatedAt = time.Now()

	if modelObj.UserName == "" {
		return nil, errors.Error("UserName cannot be empty.")
	}
	if modelObj.PhoneNumber == "" {
		return nil, errors.Error("PhoneNumber cannot be empty.")
	}
	if modelObj.CountryCode == "" {
		return nil, errors.Error("CountryCode cannot be empty.")
	}

	filter := bson.D{bson.E{Key: "username", Value: modelObj.UserName},
		bson.E{Key: "phonenumber", Value: modelObj.PhoneNumber},
		bson.E{Key: "countrycode", Value: modelObj.CountryCode}}
	count, err := impl.collection.CountDocuments(impl.ctx, filter)
	if err != nil {
		return nil, err //exception.Error("Checking if title exist.")
	}
	if count > 0 {
		return true, nil
	}

	log.Print("UserIsExist2 completed.")
	return false, err
}

func (impl serviceImpl) DeleteUser(id string, schoolId string) (int64, error) {

	log.Print("Call to delete User by id started.")
	objId := conversion.GetMongoId(id)
	filter := bson.D{bson.E{Key: "_id", Value: objId},
		bson.E{Key: "schoolid", Value: schoolId}}

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

func (impl serviceImpl) GetUser(id string, schoolId string) (dtos.UserResponse, error) {

	log.Print("Get GetUser called")
	objId := conversion.GetMongoId(id)
	var User dtos.UserResponse

	filter := bson.D{bson.E{Key: "_id", Value: objId},
		bson.E{Key: "schoolid", Value: schoolId}}

	err := impl.collection.FindOne(impl.ctx, filter).Decode(&User)
	if err != nil {
		return User, errors.Error("could not find user by id")
	}

	staff, _ := impl.staffService.GetStaff(User.DesignationId, schoolId)
	User.Designation = staff.Type

	log.Print("Call GetUser completed")
	return User, err

}

func (impl serviceImpl) GetStudent(id string, schoolId string) (dtos.UserResponse, error) {

	log.Print("Get GetUser called")
	objId := conversion.GetMongoId(id)
	var User dtos.UserResponse

	filter := bson.D{bson.E{Key: "_id", Value: objId},
		bson.E{Key: "schoolid", Value: schoolId}}

	err := impl.collection.FindOne(impl.ctx, filter).Decode(&User)
	if err != nil {
		return User, errors.Error("could not find user by id")
	}

	log.Print("Call GetUser completed")
	return User, err

}

func (impl serviceImpl) GetUsers(schoolId string) ([]dtos.UserResponse, error) {

	log.Print("Call to get all Users started.")

	var Users []dtos.UserResponse
	filter := bson.D{bson.E{Key: "schoolid", Value: schoolId},
		bson.E{Key: "usertype", Value: "Member"}}
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
	length := len(Users)
	if length == 0 {
		Users = make([]dtos.UserResponse, 0)
	}

	for i := 0; i < length; i++ {
		staff, _ := impl.staffService.GetStaff(Users[i].DesignationId, schoolId)
		Users[i].Designation = staff.Type
	}

	log.Print("Call to get all Users completed.")
	return Users, err
}

func (impl serviceImpl) GetStudents(schoolId string) ([]dtos.UserResponse, error) {

	log.Print("Call to get all students started.")

	var Users []dtos.UserResponse
	filter := bson.D{bson.E{Key: "schoolid", Value: schoolId},
		bson.E{Key: "usertype", Value: "Student"}}
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
	length := len(Users)
	if length == 0 {
		Users = make([]dtos.UserResponse, 0)
	}

	/* for i := 0; i < length; i++ {
		staff, _ := impl.staffService.GetStaff(Users[i].DesignationId, schoolId)
		Users[i].Designation = staff.Type
	} */

	log.Print("Call to get all students completed.")
	return Users, err
}

func (impl serviceImpl) GetRerals() ([]dtos.UserResponse, error) {

	log.Print("Call to get all Referals started.")

	var Users []dtos.UserResponse
	filter := bson.D{bson.E{Key: "usertype", Value: "Referal"}}
	cur, err := impl.collection.Find(impl.ctx, filter)

	if err != nil {
		Users = make([]dtos.UserResponse, 0)
		return Users, errors.Error("Referals not found!")
	}

	err = cur.All(impl.ctx, &Users)
	if err != nil {
		return Users, err
	}

	cur.Close(impl.ctx)
	length := len(Users)
	if length == 0 {
		Users = make([]dtos.UserResponse, 0)
	}

	log.Print("Call to get all Referals completed.")
	return Users, err
}

func (impl serviceImpl) GetUsersByCategory(category string, schoolId string) ([]dtos.UserResponse, error) {

	log.Print("Call to get Users by category started.")

	var Users []dtos.UserResponse
	filter := bson.D{bson.E{Key: "designationid", Value: category},
		bson.E{Key: "schoolid", Value: schoolId}}
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
	length := len(Users)
	if length == 0 {
		Users = make([]dtos.UserResponse, 0)
	}

	for i := 0; i < length; i++ {
		staff, _ := impl.staffService.GetStaff(Users[i].DesignationId, schoolId)
		Users[i].Designation = staff.Type
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

	update := bson.D{bson.E{Key: "designationyd", Value: updatedUser.DesignationId},
		bson.E{Key: "firstname", Value: updatedUser.FirstName},
		bson.E{Key: "isPhotographuploaded", Value: updatedUser.IsPhotographUploaded},
		bson.E{Key: "lastname", Value: updatedUser.LastName},
		bson.E{Key: "phonenumber", Value: updatedUser.PhoneNumber},
		bson.E{Key: "username", Value: updatedUser.UserName},
		bson.E{Key: "usertype", Value: updatedUser.UserType},
		bson.E{Key: "classroomid", Value: updatedUser.ClassRoomId},
		bson.E{Key: "levelid", Value: updatedUser.LevelId},
		bson.E{Key: "sessionid", Value: updatedUser.SessionId},
		bson.E{Key: "schoolid", Value: updatedUser.SchoolId}}

	_, err := impl.collection.UpdateOne(impl.ctx, filter, bson.D{bson.E{Key: "$set", Value: update}})
	if err != nil {
		return modelObj, errors.Error("Could not upadte user")
	}

	log.Print("PutUser completed")
	return modelObj, nil
}

func (impl serviceImpl) PutReferal(id string, User dtos.UpdateUserRequest) (interface{}, error) {

	log.Print("PutUser PutReferal")

	objId := conversion.GetMongoId(id)
	var updatedUser dtos.UpdateUserRequest
	conversion.Convert(User, &updatedUser)
	filter := bson.D{bson.E{Key: "_id", Value: objId}}
	var modelObj models.User

	update := bson.D{bson.E{Key: "firstname", Value: updatedUser.FirstName},
		bson.E{Key: "lastname", Value: updatedUser.LastName},
		bson.E{Key: "phonenumber", Value: updatedUser.PhoneNumber},
		bson.E{Key: "username", Value: updatedUser.UserName},
		bson.E{Key: "countrycode", Value: updatedUser.CountryCode}}

	_, err := impl.collection.UpdateOne(impl.ctx, filter, bson.D{bson.E{Key: "$set", Value: update}})
	if err != nil {
		return modelObj, errors.Error("Could not upadte referal")
	}

	log.Print("PutReferal completed")
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

	//resetToken := userCredential.ResetToken

	var user models.User
	filter := bson.D{bson.E{Key: "username", Value: userCredential.UserName}}
	err := impl.collection.FindOne(impl.ctx, filter).Decode(&user)

	if err != nil {
		return "Error", errors.Error("Could not upadte user's password")
	}

	/* 	passwordResetToken, err := utils.Decode(user.PasswordResetToken)

	   	if err != nil {
	   		return "Error", errors.Error("Invalid or expired token")
	   	}

	   	if passwordResetToken != resetToken {
	   		return "Error", errors.Error("Invalid or expired token")
	   	}

	   	now := time.Now()
	   	if now.Sub(user.PasswordResetAt).Minutes() > 10 {
	   		return "Error", errors.Error("Toke life expired. Please generate another one")
	   	} */

	// Update User in Database
	query := bson.D{{Key: "username", Value: userCredential.UserName}}
	update := bson.D{{Key: "$set", Value: bson.D{
		bson.E{Key: "password", Value: modelObj.Password}}}}

	/* bson.E{Key: "passwordresettoken", Value: ""},
	bson.E{Key: "passwordresetat", Value: now} */

	_, err = impl.collection.UpdateOne(impl.ctx, query, update)

	if err != nil {
		return "Error", errors.Error("Could not update password")
	}

	return "Password data updated successfully", nil
}

func (impl serviceImpl) LogInStudent(token int, schoolId string) (dtos.LoginUserResponse, error) {

	var Student dtos.UserResponse

	filter := bson.D{bson.E{Key: "token", Value: token},
		bson.E{Key: "schoolid", Value: schoolId}}

	err := impl.collection.FindOne(impl.ctx, filter).Decode(&Student)
	if err != nil {
		return dtos.LoginUserResponse{}, errors.Error("could not find type of student by token")
	}

	accessToken, accessPayload, accessError := impl.tokenMaker.CreateToken(Student.Id, strconv.Itoa(token))
	if accessError != nil {
		return dtos.LoginUserResponse{}, errors.Error("Internal server error.")
	}

	rsp := dtos.LoginUserResponse{
		Token:     accessToken,
		ExpiresAt: accessPayload.ExpiredAt,
		User: dtos.UserResponse{
			Id:        Student.Id,
			FirstName: Student.FirstName,
			LastName:  Student.LastName,
			UserType:  Student.UserType,
			CreatedAt: Student.CreatedAt,
			SchoolId:  Student.SchoolId,
			Token:     token,
		},
	}

	log.Print("Get type of student completed")
	return rsp, err
}

func (impl serviceImpl) GetSelecedStudents(Ids []string) ([]dtos.UserResponse, error) {

	objIds := make([]primitive.ObjectID, 0)
	log.Print("Call GetSelecedStudents started.")
	for _, id := range Ids {
		objIds = append(objIds, conversion.GetMongoId(id))
	}

	var Students []dtos.UserResponse
	filter := bson.D{bson.E{Key: "_id", Value: bson.D{bson.E{Key: "$in", Value: objIds}}}}
	cur, err := impl.collection.Find(impl.ctx, filter)

	if err != nil {
		Students = make([]dtos.UserResponse, 0)
		return Students, errors.Error("Types of student not found!")
	}

	err = cur.All(impl.ctx, &Students)
	if err != nil {
		return Students, err
	}

	cur.Close(impl.ctx)
	if len(Students) == 0 {
		Students = make([]dtos.UserResponse, 0)
	}

	log.Print("Call GetSelecedStudents completed.")
	return Students, err
}

func (impl serviceImpl) GetStudentsByClassRooms(schoolId string, levelId string,
	classRoomIds []string, sessionId string) ([]dtos.UserResponse, error) {

	var Students []dtos.UserResponse
	filter := bson.D{bson.E{Key: "schoolid", Value: schoolId},
		bson.E{Key: "classroomid", Value: bson.D{bson.E{Key: "$in", Value: classRoomIds}}},
		bson.E{Key: "levelid", Value: levelId},
		bson.E{Key: "sessionid", Value: sessionId}}
	cur, err := impl.collection.Find(impl.ctx, filter)

	if err != nil {
		Students = make([]dtos.UserResponse, 0)
		return Students, errors.Error("students not found!")
	}

	err = cur.All(impl.ctx, &Students)
	if err != nil {
		return Students, err
	}

	cur.Close(impl.ctx)
	if len(Students) == 0 {
		Students = make([]dtos.UserResponse, 0)
	}

	log.Print("Call GetSelecedStudents completed.")
	return Students, err
}

func (impl serviceImpl) GenerateTokens(studentIds []string) (interface{}, error) {

	log.Print("GenerateTokens started")

	objIds := make([]primitive.ObjectID, 0)
	for _, studentId := range studentIds {
		objIds = append(objIds, conversion.GetMongoId(studentId))
	}

	selectedStudents, _ := impl.GetSelecedStudents(studentIds)
	tokens := make([]int, 0)
	schoolIds := make([]string, 0)

	for _, selectedStudent := range selectedStudents {
		tokens = append(tokens, selectedStudent.Token)
		schoolIds = append(schoolIds, selectedStudent.SchoolId)
	}

	//filter := bson.D{bson.E{Key: "_id", Value: bson.D{bson.E{Key: "$in", Value: objIds}}}
	var modelObj models.User
	CreatedSubscriptionDate := time.Now()
	newTokens := make([]int, 0)
	for i := 0; i < len(selectedStudents); i++ {
		newTokens = append(newTokens, impl.utils.GenerateToken(tokens))
		filter := bson.D{bson.E{Key: "_id", Value: objIds[i]}}
		update := bson.D{bson.E{Key: "createdsubscriptiondate", Value: CreatedSubscriptionDate},
			bson.E{Key: "schoolid", Value: schoolIds[i]},
			bson.E{Key: "token", Value: newTokens[i]}}
		_, err := impl.collection.UpdateOne(impl.ctx, filter, bson.D{bson.E{Key: "$set", Value: update}})
		if err != nil {
			return modelObj, errors.Error("Could not upadte students")
		}
	}

	/* update := bson.D{bson.E{Key: "createdsubscriptiondate", Value: CreatedSubscriptionDate},
	bson.E{Key: "schoolid", Value: bson.D{bson.E{Key: "$in", Value: schoolIds}}},
	bson.E{Key: "token", Value: bson.D{bson.E{Key: "$in", Value: newTokens}}}} */

	log.Print("GenerateTokens completed")
	return modelObj, nil
}

func (impl serviceImpl) GetStudentByToken(token int, schoolId string) (dtos.UserResponse, error) {

	log.Print("GetStudentByToken called")
	var Student dtos.UserResponse

	filter := bson.D{bson.E{Key: "token", Value: token},
		bson.E{Key: "schoolid", Value: schoolId}}

	err := impl.collection.FindOne(impl.ctx, filter).Decode(&Student)
	if err != nil {
		return Student, errors.Error("could not find type of student by token")
	}

	log.Print("GetStudentByToken completed")
	return Student, err

}
