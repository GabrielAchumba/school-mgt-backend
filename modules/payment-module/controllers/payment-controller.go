package controllers

import (
	"net/http"

	"github.com/GabrielAchumba/school-mgt-backend/authentication/middleware"
	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/GabrielAchumba/school-mgt-backend/modules/payment-module/dtos"
	"github.com/GabrielAchumba/school-mgt-backend/modules/payment-module/services"
	"github.com/gin-gonic/gin"
)

var _response rest.Response

type PaymentController interface {
	CreatePayment(ctx *gin.Context) *rest.Response
	DeletePayment(ctx *gin.Context) *rest.Response
	GetPayment(ctx *gin.Context) *rest.Response
	GetPayments(ctx *gin.Context) *rest.Response
}

type controllerImpl struct {
	PaymentService services.PaymentService
}

func New(PaymentService services.PaymentService) PaymentController {
	return &controllerImpl{
		PaymentService: PaymentService,
	}
}

func (ctrl *controllerImpl) CreatePayment(ctx *gin.Context) *rest.Response {
	var model dtos.CreatePaymentRequest

	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	payload, _ := middleware.GetAuthorizationPayload(ctx)
	var userId string = payload.UserId
	if userId == "" {
		return _response.NotAuthorized()
	}

	if m, er := ctrl.PaymentService.CreatePayment(userId, model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) DeletePayment(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")
	schoolId := ctx.Param("schoolId")

	m, er := ctrl.PaymentService.DeletePayment(id, schoolId)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) GetPayment(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")
	schoolId := ctx.Param("schoolId")

	m, er := ctrl.PaymentService.GetPayment(id, schoolId)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) GetPayments(ctx *gin.Context) *rest.Response {

	schoolId := ctx.Param("schoolId")
	m, er := ctrl.PaymentService.GetPayments(schoolId)
	if er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	}
	return _response.GetSuccess(http.StatusOK, m)
}
