package controllers

import (
	"net/http"

	"github.com/GabrielAchumba/school-mgt-backend/authentication/middleware"
	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/GabrielAchumba/school-mgt-backend/modules/classroom-module/dtos"
	"github.com/GabrielAchumba/school-mgt-backend/modules/classroom-module/services"
	"github.com/gin-gonic/gin"
)

var _response rest.Response

type ClassRoomController interface {
	CreateClassRoom(ctx *gin.Context) *rest.Response
	DeleteClassRoom(ctx *gin.Context) *rest.Response
	GetClassRoom(ctx *gin.Context) *rest.Response
	GetClassRooms(ctx *gin.Context) *rest.Response
	PutClassRoom(ctx *gin.Context) *rest.Response
}

type controllerImpl struct {
	ClassRoomService services.ClassRoomService
}

func New(ClassRoomService services.ClassRoomService) ClassRoomController {
	return &controllerImpl{
		ClassRoomService: ClassRoomService,
	}
}

func (ctrl *controllerImpl) CreateClassRoom(ctx *gin.Context) *rest.Response {
	var model dtos.CreateClassRoomRequest

	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	payload, _ := middleware.GetAuthorizationPayload(ctx)
	var userId string = payload.UserId
	if userId == "" {
		return _response.NotAuthorized()
	}

	if m, er := ctrl.ClassRoomService.CreateClassRoom(userId, model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) DeleteClassRoom(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")

	m, er := ctrl.ClassRoomService.DeleteClassRoom(id)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) GetClassRoom(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")

	m, er := ctrl.ClassRoomService.GetClassRoom(id)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) GetClassRooms(ctx *gin.Context) *rest.Response {

	m, er := ctrl.ClassRoomService.GetClassRooms()
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) PutClassRoom(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")
	var model dtos.UpdateClassRoomRequest
	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	if m, er := ctrl.ClassRoomService.PutClassRoom(id, model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}
