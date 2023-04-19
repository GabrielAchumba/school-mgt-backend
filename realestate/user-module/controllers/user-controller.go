package controllers

import (
	"encoding/base64"
	"log"
	"net/http"

	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/GabrielAchumba/school-mgt-backend/realestate/user-module/dtos"
	"github.com/GabrielAchumba/school-mgt-backend/realestate/user-module/services"

	"github.com/gin-gonic/gin"
)

var _response rest.Response

type UserController interface {
	RegisterUser(ctx *gin.Context) *rest.Response
	UserIsExist(ctx *gin.Context) *rest.Response
	UserIsExist2(ctx *gin.Context) *rest.Response
	Login(ctx *gin.Context) *rest.Response
	DeleteUser(ctx *gin.Context) *rest.Response
	GetUser(ctx *gin.Context) *rest.Response
	PutUser(ctx *gin.Context) *rest.Response
	UploadPhoto(ctx *gin.Context) *rest.Response
	ResetPassword(ctx *gin.Context) *rest.Response
	toBase64(b []byte) string
	BlockUser(ctx *gin.Context) *rest.Response
	ConfirmUser(ctx *gin.Context) *rest.Response
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

	//payload, _ := middleware.GetAuthorizationPayload(ctx)
	var userId string = model.CreatedBy // payload.UserId
	/* if userId == "" {
		return _response.NotAuthorized()
	} */

	if m, er := ctrl.userService.RegisterUser(userId, model); er != nil {
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

func Int64(s string) {
	panic("unimplemented")
}

func (ctrl *controllerImpl) ConfirmUser(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")
	var model dtos.UpdateUserRequest
	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	if m, er := ctrl.userService.ConfirmUser(id, model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) BlockUser(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")
	var model dtos.UpdateUserRequest
	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	if m, er := ctrl.userService.BlockUser(id, model); er != nil {
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
