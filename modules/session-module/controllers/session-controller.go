package controllers

import (
	"net/http"

	"github.com/GabrielAchumba/school-mgt-backend/authentication/middleware"
	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/GabrielAchumba/school-mgt-backend/modules/session-module/dtos"
	"github.com/GabrielAchumba/school-mgt-backend/modules/session-module/services"
	"github.com/gin-gonic/gin"
)

var _response rest.Response

type SessionController interface {
	CreateSession(ctx *gin.Context) *rest.Response
	CreateSessions(ctx *gin.Context) *rest.Response
	DeleteSession(ctx *gin.Context) *rest.Response
	GetSession(ctx *gin.Context) *rest.Response
	GetSessions(ctx *gin.Context) *rest.Response
	PutSession(ctx *gin.Context) *rest.Response
}

type controllerImpl struct {
	SessionService services.SessionService
}

func New(SessionService services.SessionService) SessionController {
	return &controllerImpl{
		SessionService: SessionService,
	}
}

func (ctrl *controllerImpl) CreateSession(ctx *gin.Context) *rest.Response {
	var model dtos.CreateSessionRequest

	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	payload, _ := middleware.GetAuthorizationPayload(ctx)
	var userId string = payload.UserId
	if userId == "" {
		return _response.NotAuthorized()
	}

	if m, er := ctrl.SessionService.CreateSession(userId, model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) CreateSessions(ctx *gin.Context) *rest.Response {
	var models []dtos.CreateSessionRequest

	if er := ctx.BindJSON(&models); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	payload, _ := middleware.GetAuthorizationPayload(ctx)
	var userId string = payload.UserId
	if userId == "" {
		return _response.NotAuthorized()
	}

	if m, er := ctrl.SessionService.CreateSessions(userId, models); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) DeleteSession(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")
	schoolId := ctx.Param("schoolId")

	m, er := ctrl.SessionService.DeleteSession(id, schoolId)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) GetSession(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")
	schoolId := ctx.Param("schoolId")

	m, er := ctrl.SessionService.GetSession(id, schoolId)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) GetSessions(ctx *gin.Context) *rest.Response {

	schoolId := ctx.Param("schoolId")

	m, er := ctrl.SessionService.GetSessions(schoolId)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) PutSession(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")
	var model dtos.UpdateSessionRequest
	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	if m, er := ctrl.SessionService.PutSession(id, model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}
