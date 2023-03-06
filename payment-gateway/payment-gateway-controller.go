package paymentgateway

import (
	"net/http"

	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/gin-gonic/gin"
)

type PaymentGatewayController interface {
	GetBanks(ctx *gin.Context) *rest.Response
	InitiateTransfer(ctx *gin.Context) *rest.Response
	FinalizeTransfer(ctx *gin.Context) *rest.Response
}
type controllerImpl struct {
	paymentGatewayService PaymentGatewayService
}

var _response rest.Response

func NewPaymentGatewayController(paymentGatewayService PaymentGatewayService) PaymentGatewayController {
	return &controllerImpl{
		paymentGatewayService: paymentGatewayService,
	}
}

func (ctrl *controllerImpl) GetBanks(ctx *gin.Context) *rest.Response {
	m := ctrl.paymentGatewayService.GetBanks()
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) InitiateTransfer(ctx *gin.Context) *rest.Response {

	var model PaymentGatewayRequest
	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	if m, er := ctrl.paymentGatewayService.InitiateTransfer(model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) FinalizeTransfer(ctx *gin.Context) *rest.Response {
	var model PaymentGatewayRequest
	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	if m, er := ctrl.paymentGatewayService.FinalizeTransfer(model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}
