package controllers

import (
	"net/http"

	"github.com/GabrielAchumba/school-mgt-backend/authentication/middleware"
	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/GabrielAchumba/school-mgt-backend/modules/assessment-module/dtos"
	"github.com/GabrielAchumba/school-mgt-backend/modules/assessment-module/services"
	"github.com/gin-gonic/gin"
)

var _response rest.Response

type AssessmentController interface {
	CreateAssessment(ctx *gin.Context) *rest.Response
	DeleteAssessment(ctx *gin.Context) *rest.Response
	GetAssessment(ctx *gin.Context) *rest.Response
	GetAssessments(ctx *gin.Context) *rest.Response
	PutAssessment(ctx *gin.Context) *rest.Response
}

type controllerImpl struct {
	AssessmentService services.AssessmentService
}

func New(AssessmentService services.AssessmentService) AssessmentController {
	return &controllerImpl{
		AssessmentService: AssessmentService,
	}
}

func (ctrl *controllerImpl) CreateAssessment(ctx *gin.Context) *rest.Response {
	var model dtos.CreateAssessmentRequest

	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	payload, _ := middleware.GetAuthorizationPayload(ctx)
	var userId string = payload.UserId
	if userId == "" {
		return _response.NotAuthorized()
	}

	if m, er := ctrl.AssessmentService.CreateAssessment(userId, model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) DeleteAssessment(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")
	schoolId := ctx.Param("schoolId")

	m, er := ctrl.AssessmentService.DeleteAssessment(id, schoolId)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) GetAssessment(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")
	schoolId := ctx.Param("schoolId")

	m, er := ctrl.AssessmentService.GetAssessment(id, schoolId)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) GetAssessments(ctx *gin.Context) *rest.Response {

	schoolId := ctx.Param("schoolId")
	m, er := ctrl.AssessmentService.GetAssessments(schoolId)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) PutAssessment(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")
	var model dtos.UpdateAssessmentRequest
	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	if m, er := ctrl.AssessmentService.PutAssessment(id, model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}
