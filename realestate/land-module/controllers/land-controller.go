package controllers

import (
	"net/http"
	"strconv"

	"github.com/GabrielAchumba/school-mgt-backend/authentication/middleware"
	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/GabrielAchumba/school-mgt-backend/realestate/land-module/dtos"
	"github.com/GabrielAchumba/school-mgt-backend/realestate/land-module/services"
	"github.com/gin-gonic/gin"
)

var _response rest.Response

type LandController interface {
	CreateLand(ctx *gin.Context) *rest.Response
	DeleteLand(ctx *gin.Context) *rest.Response
	GetLand(ctx *gin.Context) *rest.Response
	GetLands(ctx *gin.Context) *rest.Response
	GetPaginatedLands(ctx *gin.Context) *rest.Response
	PutLand(ctx *gin.Context) *rest.Response
}

type controllerImpl struct {
	LandService services.LandService
}

func New(LandService services.LandService) LandController {
	return &controllerImpl{
		LandService: LandService,
	}
}

func (ctrl *controllerImpl) CreateLand(ctx *gin.Context) *rest.Response {
	var model dtos.CreateLandRequest

	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	payload, _ := middleware.GetAuthorizationPayload(ctx)
	var userId string = payload.UserId
	if userId == "" {
		return _response.NotAuthorized()
	}

	if m, er := ctrl.LandService.CreateLand(userId, model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) DeleteLand(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")
	schoolId := ctx.Param("schoolId")

	m, er := ctrl.LandService.DeleteLand(id, schoolId)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) GetLand(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")

	m, er := ctrl.LandService.GetLand(id)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) GetLands(ctx *gin.Context) *rest.Response {

	filterModel := ctx.Param("filterModel")
	m, er := ctrl.LandService.GetLands(filterModel)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) GetPaginatedLands(ctx *gin.Context) *rest.Response {

	page, _ := strconv.Atoi(ctx.Param("page"))
	filterModel := ctx.Param("filterModel")
	m, er := ctrl.LandService.GetPaginatedLands(page, filterModel)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) PutLand(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")
	var model dtos.UpdateLandRequest
	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	if m, er := ctrl.LandService.PutLand(id, model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}
