package controllers

import (
	"net/http"

	"github.com/GabrielAchumba/school-mgt-backend/authentication/middleware"
	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/GabrielAchumba/school-mgt-backend/modules/student-module/dtos"
	"github.com/GabrielAchumba/school-mgt-backend/modules/student-module/services"
	"github.com/gin-gonic/gin"
)

var _response rest.Response

type StudentController interface {
	CreateStudent(ctx *gin.Context) *rest.Response
	DeleteStudent(ctx *gin.Context) *rest.Response
	GetStudent(ctx *gin.Context) *rest.Response
	GetStudents(ctx *gin.Context) *rest.Response
	PutStudent(ctx *gin.Context) *rest.Response
}

type controllerImpl struct {
	StudentService services.StudentService
}

func New(StudentService services.StudentService) StudentController {
	return &controllerImpl{
		StudentService: StudentService,
	}
}

func (ctrl *controllerImpl) CreateStudent(ctx *gin.Context) *rest.Response {
	var model dtos.CreateStudentRequest

	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	payload, _ := middleware.GetAuthorizationPayload(ctx)
	var userId string = payload.UserId
	if userId == "" {
		return _response.NotAuthorized()
	}

	if m, er := ctrl.StudentService.CreateStudent(userId, model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) DeleteStudent(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")

	m, er := ctrl.StudentService.DeleteStudent(id)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) GetStudent(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")

	m, er := ctrl.StudentService.GetStudent(id)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) GetStudents(ctx *gin.Context) *rest.Response {

	m, er := ctrl.StudentService.GetStudents()
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) PutStudent(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")
	var model dtos.UpdateStudentRequest
	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	if m, er := ctrl.StudentService.PutStudent(id, model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}
