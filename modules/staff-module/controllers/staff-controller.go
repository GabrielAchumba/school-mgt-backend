package controllers

import (
	"net/http"

	"github.com/GabrielAchumba/school-mgt-backend/authentication/middleware"
	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/GabrielAchumba/school-mgt-backend/modules/staff-module/dtos"
	"github.com/GabrielAchumba/school-mgt-backend/modules/staff-module/services"
	"github.com/gin-gonic/gin"
)

var _response rest.Response

type StaffController interface {
	CreateStaff(ctx *gin.Context) *rest.Response
	DeleteStaff(ctx *gin.Context) *rest.Response
	GetStaff(ctx *gin.Context) *rest.Response
	GetStaffs(ctx *gin.Context) *rest.Response
	PutStaff(ctx *gin.Context) *rest.Response
}

type controllerImpl struct {
	staffService services.StaffService
}

func New(staffService services.StaffService) StaffController {
	return &controllerImpl{
		staffService: staffService,
	}
}

func (ctrl *controllerImpl) CreateStaff(ctx *gin.Context) *rest.Response {
	var model dtos.CreateStaffRequest

	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	payload, _ := middleware.GetAuthorizationPayload(ctx)
	var userId string = payload.UserId
	if userId == "" {
		return _response.NotAuthorized()
	}

	if m, er := ctrl.staffService.CreateStaff(userId, model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) DeleteStaff(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")
	schoolId := ctx.Param("schoolId")

	m, er := ctrl.staffService.DeleteStaff(id, schoolId)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) GetStaff(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")
	schoolId := ctx.Param("schoolId")

	m, er := ctrl.staffService.GetStaff(id, schoolId)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) GetStaffs(ctx *gin.Context) *rest.Response {

	schoolId := ctx.Param("schoolId")
	m, er := ctrl.staffService.GetStaffs(schoolId)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) PutStaff(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")
	var model dtos.UpdateStaffRequest
	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	if m, er := ctrl.staffService.PutStaff(id, model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}
