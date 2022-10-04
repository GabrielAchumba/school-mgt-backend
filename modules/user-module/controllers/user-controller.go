package controllers

import (
	"encoding/base64"
	"log"
	"net/http"

	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/GabrielAchumba/school-mgt-backend/modules/user-module/dtos"
	"github.com/GabrielAchumba/school-mgt-backend/modules/user-module/services"

	"github.com/GabrielAchumba/school-mgt-backend/authentication/middleware"
	"github.com/gin-gonic/gin"
)

var _response rest.Response

type UserController interface {
	RegisterUser(ctx *gin.Context) *rest.Response
	Login(ctx *gin.Context) *rest.Response
	DeleteUser(ctx *gin.Context) *rest.Response
	GetUser(ctx *gin.Context) *rest.Response
	GetUsersByCategory(ctx *gin.Context) *rest.Response
	GetUsers(ctx *gin.Context) *rest.Response
	PutUser(ctx *gin.Context) *rest.Response
	UpdateAdminDTO(ctx *gin.Context) *rest.Response
	UploadPhoto(ctx *gin.Context) *rest.Response
	ForgotPassword(ctx *gin.Context) *rest.Response
	ResetPassword(ctx *gin.Context) *rest.Response
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

func (ctrl *controllerImpl) DeleteUser(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")

	m, er := ctrl.userService.DeleteUser(id)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) GetUser(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")

	m, er := ctrl.userService.GetUser(id)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) GetUsersByCategory(ctx *gin.Context) *rest.Response {

	category := ctx.Param("category")

	m, er := ctrl.userService.GetUsersByCategory(category)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) GetUsers(ctx *gin.Context) *rest.Response {

	m, er := ctrl.userService.GetUsers()
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
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
