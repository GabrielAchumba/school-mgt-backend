package controllers

import (
	"net/http"

	"github.com/GabrielAchumba/school-mgt-backend/authentication/middleware"
	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/GabrielAchumba/school-mgt-backend/modules/level-module/dtos"
	"github.com/GabrielAchumba/school-mgt-backend/modules/level-module/services"
	"github.com/gin-gonic/gin"
)

var _response rest.Response

type LevelController interface {
	CreateLevel(ctx *gin.Context) *rest.Response
	CreateLevels(ctx *gin.Context) *rest.Response
	DeleteLevel(ctx *gin.Context) *rest.Response
	DeleteLevelMany(ctx *gin.Context) *rest.Response
	GetLevel(ctx *gin.Context) *rest.Response
	GetLevels(ctx *gin.Context) *rest.Response
	PutLevel(ctx *gin.Context) *rest.Response
}

type controllerImpl struct {
	LevelService services.LevelService
}

func New(LevelService services.LevelService) LevelController {
	return &controllerImpl{
		LevelService: LevelService,
	}
}

func (ctrl *controllerImpl) CreateLevel(ctx *gin.Context) *rest.Response {
	var model dtos.CreateLevelRequest

	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	payload, _ := middleware.GetAuthorizationPayload(ctx)
	var userId string = payload.UserId
	if userId == "" {
		return _response.NotAuthorized()
	}

	if m, er := ctrl.LevelService.CreateLevel(userId, model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) CreateLevels(ctx *gin.Context) *rest.Response {
	var model []dtos.CreateLevelRequest

	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	payload, _ := middleware.GetAuthorizationPayload(ctx)
	var userId string = payload.UserId
	if userId == "" {
		return _response.NotAuthorized()
	}

	if m, er := ctrl.LevelService.CreateLevels(userId, model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) DeleteLevel(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")
	schoolId := ctx.Param("schoolId")

	m, er := ctrl.LevelService.DeleteLevel(id, schoolId)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) DeleteLevelMany(ctx *gin.Context) *rest.Response {
	var model dtos.LevelIds

	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	m, er := ctrl.LevelService.DeleteLevelMany(model.Ids, model.SchoolId)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) GetLevel(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")
	schoolId := ctx.Param("schoolId")

	m, er := ctrl.LevelService.GetLevel(id, schoolId)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) GetLevels(ctx *gin.Context) *rest.Response {

	schoolId := ctx.Param("schoolId")
	m, er := ctrl.LevelService.GetLevels(schoolId)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) PutLevel(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")
	var model dtos.UpdateLevelRequest
	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	if m, er := ctrl.LevelService.PutLevel(id, model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}
