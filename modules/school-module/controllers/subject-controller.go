package controllers

import (
	"net/http"

	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/GabrielAchumba/school-mgt-backend/modules/school-module/dtos"
	"github.com/GabrielAchumba/school-mgt-backend/modules/school-module/services"
	"github.com/gin-gonic/gin"
)

var _response rest.Response

type SchoolController interface {
	CreateSchool(ctx *gin.Context) *rest.Response
	DeleteSchool(ctx *gin.Context) *rest.Response
	GetSchool(ctx *gin.Context) *rest.Response
	GetSchools(ctx *gin.Context) *rest.Response
	GetSchoolByReferal(ctx *gin.Context) *rest.Response
	PutSchool(ctx *gin.Context) *rest.Response
}

type controllerImpl struct {
	SchoolService services.SchoolService
}

func New(SchoolService services.SchoolService) SchoolController {
	return &controllerImpl{
		SchoolService: SchoolService,
	}
}

func (ctrl *controllerImpl) CreateSchool(ctx *gin.Context) *rest.Response {
	var model dtos.CreateSchoolRequest

	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	if m, er := ctrl.SchoolService.CreateSchool(model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) DeleteSchool(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")

	m, er := ctrl.SchoolService.DeleteSchool(id)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) GetSchool(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")

	m, er := ctrl.SchoolService.GetSchool(id)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) GetSchools(ctx *gin.Context) *rest.Response {

	m, er := ctrl.SchoolService.GetSchools()
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) GetSchoolByReferal(ctx *gin.Context) *rest.Response {

	referalId := ctx.Param("referalId")
	m, er := ctrl.SchoolService.GetSchoolByReferal(referalId)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) PutSchool(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")
	var model dtos.UpdateSchoolRequest
	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	if m, er := ctrl.SchoolService.PutSchool(id, model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}
