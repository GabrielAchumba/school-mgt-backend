package controllers

import (
	"net/http"

	"github.com/GabrielAchumba/school-mgt-backend/authentication/middleware"
	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/GabrielAchumba/school-mgt-backend/realestate/file-module/dtos"
	"github.com/GabrielAchumba/school-mgt-backend/realestate/file-module/services"
	"github.com/gin-gonic/gin"
)

var _response rest.Response

type FileController interface {
	CreateFile(ctx *gin.Context) *rest.Response
	DeleteFile(ctx *gin.Context) *rest.Response
	GetFile(ctx *gin.Context) *rest.Response
	GetFiles(ctx *gin.Context) *rest.Response
	PutFile(ctx *gin.Context) *rest.Response
	GetFileByParams(ctx *gin.Context) *rest.Response
}

type controllerImpl struct {
	FileService services.FileService
}

func New(FileService services.FileService) FileController {
	return &controllerImpl{
		FileService: FileService,
	}
}

func (ctrl *controllerImpl) CreateFile(ctx *gin.Context) *rest.Response {
	var model dtos.CreateFileRequest

	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	payload, _ := middleware.GetAuthorizationPayload(ctx)
	var userId string = payload.UserId
	if userId == "" {
		return _response.NotAuthorized()
	}

	if m, er := ctrl.FileService.CreateFile(userId, model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) DeleteFile(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")
	schoolId := ctx.Param("schoolId")

	m, er := ctrl.FileService.DeleteFile(id, schoolId)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) GetFile(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")

	m, er := ctrl.FileService.GetFile(id)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) GetFiles(ctx *gin.Context) *rest.Response {

	filterModel := ctx.Param("filterModel")
	m, er := ctrl.FileService.GetFiles(filterModel)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) GetFileByParams(ctx *gin.Context) *rest.Response {

	title := ctx.Param("title")
	userId := ctx.Param("userId")
	categoryId := ctx.Param("categoryId")
	m, er := ctrl.FileService.GetFileByParams(title, userId, categoryId)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) PutFile(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")
	var model dtos.UpdateFileRequest
	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	if m, er := ctrl.FileService.PutFile(id, model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}
