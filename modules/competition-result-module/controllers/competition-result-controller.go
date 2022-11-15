package controllers

import (
	"net/http"

	"github.com/GabrielAchumba/school-mgt-backend/authentication/middleware"
	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/GabrielAchumba/school-mgt-backend/modules/competition-result-module/dtos"
	"github.com/GabrielAchumba/school-mgt-backend/modules/competition-result-module/services"
	"github.com/gin-gonic/gin"
)

var _response rest.Response

type CompetitionResultController interface {
	CreateCompetitionResult(ctx *gin.Context) *rest.Response
	CreateCompetitionResults(ctx *gin.Context) *rest.Response
	DeleteCompetitionResult(ctx *gin.Context) *rest.Response
	GetCompetitionResult(ctx *gin.Context) *rest.Response
	GetCompetitionResults(ctx *gin.Context) *rest.Response
	PutCompetitionResult(ctx *gin.Context) *rest.Response
	ComputeSummaryCompetitionResults(ctx *gin.Context) *rest.Response
	ComputeSummaryCompetitionResults2(ctx *gin.Context) *rest.Response
	ComputeStudentsSummaryCompetitionResults(ctx *gin.Context) *rest.Response
	SummaryStudentsPositions(ctx *gin.Context) *rest.Response
	SummaryStudentsPositions2(ctx *gin.Context) *rest.Response
	ComputeStudentsCompetitionResultsByDateRange(ctx *gin.Context) *rest.Response
}

type controllerImpl struct {
	CompetitionResultService services.CompetitionResultService
}

func New(CompetitionResultService services.CompetitionResultService) CompetitionResultController {
	return &controllerImpl{
		CompetitionResultService: CompetitionResultService,
	}
}

func (ctrl *controllerImpl) CreateCompetitionResult(ctx *gin.Context) *rest.Response {
	var model dtos.CreateCompetitionResultRequest

	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	payload, _ := middleware.GetAuthorizationPayload(ctx)
	var userId string = payload.UserId
	if userId == "" {
		return _response.NotAuthorized()
	}

	if m, er := ctrl.CompetitionResultService.CreateCompetitionResult(userId, model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) CreateCompetitionResults(ctx *gin.Context) *rest.Response {
	var model []dtos.CreateCompetitionResultRequest

	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	payload, _ := middleware.GetAuthorizationPayload(ctx)
	var userId string = payload.UserId
	if userId == "" {
		return _response.NotAuthorized()
	}

	if m, er := ctrl.CompetitionResultService.CreateCompetitionResults(userId, model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) DeleteCompetitionResult(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")
	schoolId := ctx.Param("schoolId")

	m, er := ctrl.CompetitionResultService.DeleteCompetitionResult(id, schoolId)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) GetCompetitionResult(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")
	schoolId := ctx.Param("schoolId")

	m, er := ctrl.CompetitionResultService.GetCompetitionResult(id, schoolId)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) GetCompetitionResults(ctx *gin.Context) *rest.Response {

	schoolId := ctx.Param("schoolId")
	m, er := ctrl.CompetitionResultService.GetCompetitionResults(schoolId)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) PutCompetitionResult(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")
	var model dtos.UpdateCompetitionResultRequest
	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	if m, er := ctrl.CompetitionResultService.PutCompetitionResult(id, model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) ComputeSummaryCompetitionResults(ctx *gin.Context) *rest.Response {
	var model dtos.GetCompetitionResultsRequest

	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	payload, _ := middleware.GetAuthorizationPayload(ctx)
	var userId string = payload.UserId
	if userId == "" {
		return _response.NotAuthorized()
	}

	if m, er := ctrl.CompetitionResultService.ComputeSummaryCompetitionResults(model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) ComputeSummaryCompetitionResults2(ctx *gin.Context) *rest.Response {
	var model dtos.GetCompetitionResultsRequest

	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	payload, _ := middleware.GetAuthorizationPayload(ctx)
	var userId string = payload.UserId
	if userId == "" {
		return _response.NotAuthorized()
	}

	if m, er := ctrl.CompetitionResultService.ComputeSummaryCompetitionResults2(model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) ComputeStudentsSummaryCompetitionResults(ctx *gin.Context) *rest.Response {
	var model dtos.GetCompetitionResultsRequest

	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	payload, _ := middleware.GetAuthorizationPayload(ctx)
	var userId string = payload.UserId
	if userId == "" {
		return _response.NotAuthorized()
	}

	if m, er := ctrl.CompetitionResultService.ComputeStudentsSummaryCompetitionResults(model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) SummaryStudentsPositions(ctx *gin.Context) *rest.Response {
	var model dtos.GetCompetitionResultsRequest

	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	payload, _ := middleware.GetAuthorizationPayload(ctx)
	var userId string = payload.UserId
	if userId == "" {
		return _response.NotAuthorized()
	}

	if m, er := ctrl.CompetitionResultService.SummaryStudentsPositions(model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) SummaryStudentsPositions2(ctx *gin.Context) *rest.Response {
	var model dtos.GetCompetitionResultsRequest

	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	payload, _ := middleware.GetAuthorizationPayload(ctx)
	var userId string = payload.UserId
	if userId == "" {
		return _response.NotAuthorized()
	}

	if m, er := ctrl.CompetitionResultService.SummaryStudentsPositions2(model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) ComputeStudentsCompetitionResultsByDateRange(ctx *gin.Context) *rest.Response {
	var model dtos.GetCompetitionResultsRequest

	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	payload, _ := middleware.GetAuthorizationPayload(ctx)
	var userId string = payload.UserId
	if userId == "" {
		return _response.NotAuthorized()
	}

	if m, er := ctrl.CompetitionResultService.ComputeStudentsCompetitionResultsByDateRange(model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}
