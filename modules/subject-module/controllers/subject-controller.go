package controllers

import (
	"net/http"

	"github.com/GabrielAchumba/school-mgt-backend/authentication/middleware"
	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/GabrielAchumba/school-mgt-backend/modules/subject-module/dtos"
	"github.com/GabrielAchumba/school-mgt-backend/modules/subject-module/services"
	"github.com/gin-gonic/gin"
)

var _response rest.Response

type SubjectController interface {
	CreateSubject(ctx *gin.Context) *rest.Response
	DeleteSubject(ctx *gin.Context) *rest.Response
	GetSubject(ctx *gin.Context) *rest.Response
	GetSubjects(ctx *gin.Context) *rest.Response
	PutSubject(ctx *gin.Context) *rest.Response
}

type controllerImpl struct {
	SubjectService services.SubjectService
}

func New(SubjectService services.SubjectService) SubjectController {
	return &controllerImpl{
		SubjectService: SubjectService,
	}
}

func (ctrl *controllerImpl) CreateSubject(ctx *gin.Context) *rest.Response {
	var model dtos.CreateSubjectRequest

	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	payload, _ := middleware.GetAuthorizationPayload(ctx)
	var userId string = payload.UserId
	if userId == "" {
		return _response.NotAuthorized()
	}

	if m, er := ctrl.SubjectService.CreateSubject(userId, model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) DeleteSubject(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")
	schoolId := ctx.Param("schoolId")

	m, er := ctrl.SubjectService.DeleteSubject(id, schoolId)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) GetSubject(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")
	schoolId := ctx.Param("schoolId")

	m, er := ctrl.SubjectService.GetSubject(id, schoolId)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) GetSubjects(ctx *gin.Context) *rest.Response {

	schoolId := ctx.Param("schoolId")

	m, er := ctrl.SubjectService.GetSubjects(schoolId)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) PutSubject(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")
	var model dtos.UpdateSubjectRequest
	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	if m, er := ctrl.SubjectService.PutSubject(id, model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}
