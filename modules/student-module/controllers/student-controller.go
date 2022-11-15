package controllers

import (
	"net/http"
	"strconv"

	"github.com/GabrielAchumba/school-mgt-backend/authentication/middleware"
	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/GabrielAchumba/school-mgt-backend/modules/student-module/dtos"
	"github.com/GabrielAchumba/school-mgt-backend/modules/student-module/services"
	"github.com/gin-gonic/gin"
)

var _response rest.Response

type StudentController interface {
	CreateStudent(ctx *gin.Context) *rest.Response
	CreateStudents(ctx *gin.Context) *rest.Response
	DeleteStudent(ctx *gin.Context) *rest.Response
	GetStudent(ctx *gin.Context) *rest.Response
	GetStudents(ctx *gin.Context) *rest.Response
	PutStudent(ctx *gin.Context) *rest.Response
	GenerateTokens(ctx *gin.Context) *rest.Response
	GetStudentByToken(ctx *gin.Context) *rest.Response
	LogInStudent(ctx *gin.Context) *rest.Response
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

func (ctrl *controllerImpl) CreateStudents(ctx *gin.Context) *rest.Response {
	var model []dtos.CreateStudentRequest

	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	payload, _ := middleware.GetAuthorizationPayload(ctx)
	var userId string = payload.UserId
	if userId == "" {
		return _response.NotAuthorized()
	}

	if m, er := ctrl.StudentService.CreateStudents(userId, model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) DeleteStudent(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")
	schoolId := ctx.Param("schoolId")

	m, er := ctrl.StudentService.DeleteStudent(id, schoolId)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) GetStudent(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")
	schoolId := ctx.Param("schoolId")

	m, er := ctrl.StudentService.GetStudent(id, schoolId)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) LogInStudent(ctx *gin.Context) *rest.Response {
	token, _ := strconv.Atoi(ctx.Param("token"))
	schoolId := ctx.Param("schoolId")

	m, er := ctrl.StudentService.LogInStudent(token, schoolId)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) GetStudentByToken(ctx *gin.Context) *rest.Response {
	token, _ := strconv.Atoi(ctx.Param("token"))
	schoolId := ctx.Param("schoolId")

	m, er := ctrl.StudentService.GetStudentByToken(token, schoolId)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) GetStudents(ctx *gin.Context) *rest.Response {

	schoolId := ctx.Param("schoolId")

	m, er := ctrl.StudentService.GetStudents(schoolId)
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

func (ctrl *controllerImpl) GenerateTokens(ctx *gin.Context) *rest.Response {
	var model dtos.UpdateStudentRequest

	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	if m, er := ctrl.StudentService.GenerateTokens(model.StudentIds); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}
