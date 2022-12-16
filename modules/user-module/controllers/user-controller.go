package controllers

import (
	"encoding/base64"
	"log"
	"net/http"
	"strconv"

	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/GabrielAchumba/school-mgt-backend/modules/user-module/dtos"
	"github.com/GabrielAchumba/school-mgt-backend/modules/user-module/services"

	"github.com/GabrielAchumba/school-mgt-backend/authentication/middleware"
	"github.com/gin-gonic/gin"
)

var _response rest.Response

type UserController interface {
	RegisterUser(ctx *gin.Context) *rest.Response
	RegisterUsers(ctx *gin.Context) *rest.Response
	RegisterAdminOrReferal(ctx *gin.Context) *rest.Response
	UserIsExist(ctx *gin.Context) *rest.Response
	UserIsExist2(ctx *gin.Context) *rest.Response
	Login(ctx *gin.Context) *rest.Response
	DeleteUser(ctx *gin.Context) *rest.Response
	GetUser(ctx *gin.Context) *rest.Response
	GetUsersByCategory(ctx *gin.Context) *rest.Response
	GetRerals(ctx *gin.Context) *rest.Response
	GetUsers(ctx *gin.Context) *rest.Response
	GetStudents(ctx *gin.Context) *rest.Response
	PutUser(ctx *gin.Context) *rest.Response
	PutReferal(ctx *gin.Context) *rest.Response
	UpdateAdminDTO(ctx *gin.Context) *rest.Response
	UploadPhoto(ctx *gin.Context) *rest.Response
	ForgotPassword(ctx *gin.Context) *rest.Response
	ResetPassword(ctx *gin.Context) *rest.Response
	GenerateTokens(ctx *gin.Context) *rest.Response
	GetStudentByToken(ctx *gin.Context) *rest.Response
	GetStudentsByClassRooms(ctx *gin.Context) *rest.Response
	LogInStudent(ctx *gin.Context) *rest.Response
	toBase64(b []byte) string
}

type ImageName struct {
	FileName     string `json:"fileName"`
	Base64String string `json:"base64String"`
}

type controllerImpl struct {
	userService services.UserService
}

func New(userService services.UserService) UserController {
	return &controllerImpl{
		userService: userService,
	}
}

func (ctrl *controllerImpl) Login(ctx *gin.Context) *rest.Response {
	var model dtos.LoginUserRequest
	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	if m, er := ctrl.userService.LoginUser(model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) RegisterUser(ctx *gin.Context) *rest.Response {
	var model dtos.CreateUserRequest

	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	payload, _ := middleware.GetAuthorizationPayload(ctx)
	var userId string = payload.UserId
	if userId == "" {
		return _response.NotAuthorized()
	}

	if m, er := ctrl.userService.RegisterUser(userId, model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) RegisterUsers(ctx *gin.Context) *rest.Response {
	var model []dtos.CreateUserRequest

	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	payload, _ := middleware.GetAuthorizationPayload(ctx)
	var userId string = payload.UserId
	if userId == "" {
		return _response.NotAuthorized()
	}

	if m, er := ctrl.userService.RegisterUsers(userId, model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) RegisterAdminOrReferal(ctx *gin.Context) *rest.Response {
	var model dtos.CreateUserRequest

	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	if m, er := ctrl.userService.RegisterAdminOrReferal(model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) UserIsExist(ctx *gin.Context) *rest.Response {
	var model dtos.CreateUserRequest

	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	if m, er := ctrl.userService.UserIsExist(model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) UserIsExist2(ctx *gin.Context) *rest.Response {
	var model dtos.CreateUserRequest

	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	if m, er := ctrl.userService.UserIsExist2(model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) DeleteUser(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")
	schoolId := ctx.Param("schoolId")

	m, er := ctrl.userService.DeleteUser(id, schoolId)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) GetUser(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")
	schoolId := ctx.Param("schoolId")

	m, er := ctrl.userService.GetUser(id, schoolId)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) GetRerals(ctx *gin.Context) *rest.Response {

	m, er := ctrl.userService.GetRerals()
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) GetUsersByCategory(ctx *gin.Context) *rest.Response {

	category := ctx.Param("category")
	schoolId := ctx.Param("schoolId")

	m, er := ctrl.userService.GetUsersByCategory(category, schoolId)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) GetUsers(ctx *gin.Context) *rest.Response {

	schoolId := ctx.Param("schoolId")
	m, er := ctrl.userService.GetUsers(schoolId)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) GetStudents(ctx *gin.Context) *rest.Response {

	schoolId := ctx.Param("schoolId")
	m, er := ctrl.userService.GetStudents(schoolId)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) GetStudentsByClassRooms(ctx *gin.Context) *rest.Response {
	var model dtos.CreateUserRequest

	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	if m, er := ctrl.userService.GetStudentsByClassRooms(model.SchoolId, model.LevelId,
		model.ClassRoomIds, model.SessionId); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) PutUser(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")
	var model dtos.UpdateUserRequest
	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	if m, er := ctrl.userService.PutUser(id, model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) PutReferal(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")
	var model dtos.UpdateUserRequest
	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	if m, er := ctrl.userService.PutReferal(id, model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) UpdateAdminDTO(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")
	var model dtos.AdminDTO
	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	if m, er := ctrl.userService.UpdateAdminDTO(id, model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func (ctrl *controllerImpl) UploadPhoto(ctx *gin.Context) *rest.Response {

	form, _ := ctx.MultipartForm()

	files := form.File["images[]"]

	//var imageNames []ImageName
	checkError := false
	imageName := ImageName{}

	for _, file := range files {

		f, err := file.Open()

		if err != nil {
			log.Println(err)
		}

		defer f.Close()

		size := file.Size
		buffer := make([]byte, size)

		f.Read(buffer)
		fileType := http.DetectContentType(buffer)

		switch fileType {
		case "image/jpeg":
			imageName.Base64String += "data:image/jpeg;base64,"
		case "image/png":
			imageName.Base64String += "data:image/png;base64,"
		case "image/jpg":
			imageName.Base64String += "data:image/jpg;base64,"
		default:
			checkError = true
		}

		if checkError {
			return _response.GetError(http.StatusOK, "image must be: (.png, .jpg, or .jpeg)")
		}

		imageName.Base64String += ctrl.toBase64(buffer)

		log.Print(imageName.Base64String)

		imageName.FileName = file.Filename

	}

	return _response.GetSuccess(http.StatusOK, imageName)
}

func (ctrl *controllerImpl) ForgotPassword(ctx *gin.Context) *rest.Response {

	var model dtos.ForgotPasswordInput
	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	if m, er := ctrl.userService.ForgotPassword(model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) ResetPassword(ctx *gin.Context) *rest.Response {

	var model dtos.ResetPasswordInput
	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	if m, er := ctrl.userService.ResetPassword(model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) LogInStudent(ctx *gin.Context) *rest.Response {
	token, _ := strconv.Atoi(ctx.Param("token"))
	schoolId := ctx.Param("schoolId")

	m, er := ctrl.userService.LogInStudent(token, schoolId)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) GetStudentByToken(ctx *gin.Context) *rest.Response {
	token, _ := strconv.Atoi(ctx.Param("token"))
	schoolId := ctx.Param("schoolId")

	m, er := ctrl.userService.GetStudentByToken(token, schoolId)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) GenerateTokens(ctx *gin.Context) *rest.Response {
	var model dtos.UpdateUserRequest

	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	if m, er := ctrl.userService.GenerateTokens(model.StudentIds); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}
