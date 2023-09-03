package controllers

import (
	"net/http"

	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/GabrielAchumba/school-mgt-backend/newpay/dtos"
	"github.com/GabrielAchumba/school-mgt-backend/newpay/services"
	"github.com/gin-gonic/gin"
)

type CableTvController interface {
	GetVariationCodes(ctx *gin.Context) *rest.Response
	SmartCardNumberVerification(ctx *gin.Context) *rest.Response
	PurchaseCableTV(ctx *gin.Context) *rest.Response
}
type controllerImpl struct {
	cableTVService services.CableTVService
}

var cableTVresponse rest.Response

func NewCableTvController(cableTVService services.CableTVService) CableTvController {
	return &controllerImpl{
		cableTVService: cableTVService,
	}
}

func (ctrl *controllerImpl) GetVariationCodes(ctx *gin.Context) *rest.Response {
	var model dtos.VTURequest

	if er := ctx.BindJSON(&model); er != nil {
		return cableTVresponse.GetError(http.StatusBadRequest, er.Error())
	}

	if m, er := ctrl.cableTVService.GetVariationCodes(model.ServiceID); er != nil {
		return cableTVresponse.GetError(http.StatusOK, er.Error())
	} else {
		return cableTVresponse.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) SmartCardNumberVerification(ctx *gin.Context) *rest.Response {
	var model dtos.VTURequest

	if er := ctx.BindJSON(&model); er != nil {
		return cableTVresponse.GetError(http.StatusBadRequest, er.Error())
	}

	if m, er := ctrl.cableTVService.SmartCardNumberVerification(model.CustomerID, model.ServiceID); er != nil {
		return cableTVresponse.GetError(http.StatusOK, er.Error())
	} else {
		return cableTVresponse.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) PurchaseCableTV(ctx *gin.Context) *rest.Response {
	var model dtos.VTURequest

	if er := ctx.BindJSON(&model); er != nil {
		return cableTVresponse.GetError(http.StatusBadRequest, er.Error())
	}

	if m, er := ctrl.cableTVService.PurchaseCableTV(model.Phone,
		model.ServiceID, model.SmartCardNumber, model.VariationID); er != nil {
		return cableTVresponse.GetError(http.StatusOK, er.Error())
	} else {
		return cableTVresponse.GetSuccess(http.StatusOK, m)
	}
}
