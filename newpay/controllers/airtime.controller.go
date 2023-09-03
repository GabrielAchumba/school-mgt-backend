package controllers

import (
	"net/http"

	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/GabrielAchumba/school-mgt-backend/newpay/dtos"
	"github.com/GabrielAchumba/school-mgt-backend/newpay/services"
	"github.com/gin-gonic/gin"
)

type AirtimeController interface {
	PurchaseAirtime(ctx *gin.Context) *rest.Response
}
type AirtimeControllerImpl struct {
	airtimeService services.AirtimeService
}

var airTimeResponse rest.Response

func NewAirtimeController(airtimeService services.AirtimeService) AirtimeController {
	return &AirtimeControllerImpl{
		airtimeService: airtimeService,
	}
}

func (ctrl *AirtimeControllerImpl) PurchaseAirtime(ctx *gin.Context) *rest.Response {
	var model dtos.VTURequest

	if er := ctx.BindJSON(&model); er != nil {
		return airTimeResponse.GetError(http.StatusBadRequest, er.Error())
	}

	if m, er := ctrl.airtimeService.PurchaseAirtime(model.Phone, model.NetworkID, model.Amount); er != nil {
		return airTimeResponse.GetError(http.StatusOK, er.Error())
	} else {
		return airTimeResponse.GetSuccess(http.StatusOK, m)
	}
}
