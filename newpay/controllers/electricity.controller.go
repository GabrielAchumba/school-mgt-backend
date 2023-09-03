package controllers

import (
	"net/http"

	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/GabrielAchumba/school-mgt-backend/newpay/dtos"
	"github.com/GabrielAchumba/school-mgt-backend/newpay/services"
	"github.com/gin-gonic/gin"
)

type ElectricityController interface {
	MeterNumberVerification(ctx *gin.Context) *rest.Response
	PurchaseElectricity(ctx *gin.Context) *rest.Response
}
type ElectricityControllerImpl struct {
	electricityService services.ElectricityService
}

func NewElectricityController(electricityService services.ElectricityService) ElectricityController {
	return &ElectricityControllerImpl{
		electricityService: electricityService,
	}
}

var electricityResponse rest.Response

func (ctrl *ElectricityControllerImpl) MeterNumberVerification(ctx *gin.Context) *rest.Response {
	var model dtos.VTURequest

	if er := ctx.BindJSON(&model); er != nil {
		return electricityResponse.GetError(http.StatusBadRequest, er.Error())
	}

	if m, er := ctrl.electricityService.MeterNumberVerification(model.CustomerID,
		model.VariationID, model.ServiceID); er != nil {
		return electricityResponse.GetError(http.StatusOK, er.Error())
	} else {
		return electricityResponse.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *ElectricityControllerImpl) PurchaseElectricity(ctx *gin.Context) *rest.Response {
	var model dtos.VTURequest

	if er := ctx.BindJSON(&model); er != nil {
		return electricityResponse.GetError(http.StatusBadRequest, er.Error())
	}

	if m, er := ctrl.electricityService.PurchaseElectricity(model.Phone,
		model.MeterNumber, model.ServiceID, model.VariationID, model.Amount); er != nil {
		return electricityResponse.GetError(http.StatusOK, er.Error())
	} else {
		return electricityResponse.GetSuccess(http.StatusOK, m)
	}
}
