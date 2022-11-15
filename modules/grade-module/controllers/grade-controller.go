package controllers

import (
	"net/http"

	"github.com/GabrielAchumba/school-mgt-backend/authentication/middleware"
	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/GabrielAchumba/school-mgt-backend/modules/grade-module/dtos"
	"github.com/GabrielAchumba/school-mgt-backend/modules/grade-module/services"
	"github.com/gin-gonic/gin"
)

var _response rest.Response

type GradeController interface {
	CreateGrade(ctx *gin.Context) *rest.Response
	CreateGrades(ctx *gin.Context) *rest.Response
	DeleteGrade(ctx *gin.Context) *rest.Response
	GetGrade(ctx *gin.Context) *rest.Response
	GetGrades(ctx *gin.Context) *rest.Response
	PutGrade(ctx *gin.Context) *rest.Response
}

type controllerImpl struct {
	GradeService services.GradeService
}

func New(GradeService services.GradeService) GradeController {
	return &controllerImpl{
		GradeService: GradeService,
	}
}

func (ctrl *controllerImpl) CreateGrade(ctx *gin.Context) *rest.Response {
	var model dtos.CreateGradeRequest

	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	payload, _ := middleware.GetAuthorizationPayload(ctx)
	var userId string = payload.UserId
	if userId == "" {
		return _response.NotAuthorized()
	}

	if m, er := ctrl.GradeService.CreateGrade(userId, model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) CreateGrades(ctx *gin.Context) *rest.Response {
	var models []dtos.CreateGradeRequest

	if er := ctx.BindJSON(&models); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	payload, _ := middleware.GetAuthorizationPayload(ctx)
	var userId string = payload.UserId
	if userId == "" {
		return _response.NotAuthorized()
	}

	if m, er := ctrl.GradeService.CreateGrades(userId, models); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) DeleteGrade(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")
	schoolId := ctx.Param("schoolId")

	m, er := ctrl.GradeService.DeleteGrade(id, schoolId)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) GetGrade(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")
	schoolId := ctx.Param("schoolId")

	m, er := ctrl.GradeService.GetGrade(id, schoolId)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) GetGrades(ctx *gin.Context) *rest.Response {

	schoolId := ctx.Param("schoolId")

	m, er := ctrl.GradeService.GetGrades(schoolId)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) PutGrade(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")
	var model dtos.UpdateGradeRequest
	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	if m, er := ctrl.GradeService.PutGrade(id, model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}
