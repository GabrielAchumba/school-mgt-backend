package controllers

import (
	"net/http"

	"github.com/GabrielAchumba/school-mgt-backend/authentication/middleware"
	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/GabrielAchumba/school-mgt-backend/modules/result-module/dtos"
	"github.com/GabrielAchumba/school-mgt-backend/modules/result-module/services"
	"github.com/gin-gonic/gin"
)

var _response rest.Response

type ResultController interface {
	CreateResult(ctx *gin.Context) *rest.Response
	DeleteResult(ctx *gin.Context) *rest.Response
	GetResult(ctx *gin.Context) *rest.Response
	GetResults(ctx *gin.Context) *rest.Response
	PutResult(ctx *gin.Context) *rest.Response
	ComputeSummaryResults(ctx *gin.Context) *rest.Response
	ComputeStudentsSummaryResults(ctx *gin.Context) *rest.Response
	SummaryStudentsPositions(ctx *gin.Context) *rest.Response
	ComputeStudentsResultsByDateRange(ctx *gin.Context) *rest.Response
}

type controllerImpl struct {
	ResultService services.ResultService
}

func New(ResultService services.ResultService) ResultController {
	return &controllerImpl{
		ResultService: ResultService,
	}
}

func (ctrl *controllerImpl) CreateResult(ctx *gin.Context) *rest.Response {
	var model dtos.CreateResultRequest

	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	payload, _ := middleware.GetAuthorizationPayload(ctx)
	var userId string = payload.UserId
	if userId == "" {
		return _response.NotAuthorized()
	}

	if m, er := ctrl.ResultService.CreateResult(userId, model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) DeleteResult(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")

	m, er := ctrl.ResultService.DeleteResult(id)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) GetResult(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")

	m, er := ctrl.ResultService.GetResult(id)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) GetResults(ctx *gin.Context) *rest.Response {

	m, er := ctrl.ResultService.GetResults()
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) PutResult(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")
	var model dtos.UpdateResultRequest
	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	if m, er := ctrl.ResultService.PutResult(id, model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) ComputeSummaryResults(ctx *gin.Context) *rest.Response {
	var model dtos.GetResultsRequest

	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	payload, _ := middleware.GetAuthorizationPayload(ctx)
	var userId string = payload.UserId
	if userId == "" {
		return _response.NotAuthorized()
	}

	if m, er := ctrl.ResultService.ComputeSummaryResults(model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) ComputeStudentsSummaryResults(ctx *gin.Context) *rest.Response {
	var model dtos.GetResultsRequest

	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	payload, _ := middleware.GetAuthorizationPayload(ctx)
	var userId string = payload.UserId
	if userId == "" {
		return _response.NotAuthorized()
	}

	if m, er := ctrl.ResultService.ComputeStudentsSummaryResults(model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) SummaryStudentsPositions(ctx *gin.Context) *rest.Response {
	var model dtos.GetResultsRequest

	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	payload, _ := middleware.GetAuthorizationPayload(ctx)
	var userId string = payload.UserId
	if userId == "" {
		return _response.NotAuthorized()
	}

	if m, er := ctrl.ResultService.SummaryStudentsPositions(model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) ComputeStudentsResultsByDateRange(ctx *gin.Context) *rest.Response {
	var model dtos.GetResultsRequest

	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	payload, _ := middleware.GetAuthorizationPayload(ctx)
	var userId string = payload.UserId
	if userId == "" {
		return _response.NotAuthorized()
	}

	if m, er := ctrl.ResultService.ComputeStudentsResultsByDateRange(model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}
