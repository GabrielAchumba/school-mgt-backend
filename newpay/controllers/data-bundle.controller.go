package controllers

import (
	"net/http"

	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/GabrielAchumba/school-mgt-backend/newpay/dtos"
	"github.com/GabrielAchumba/school-mgt-backend/newpay/services"
	"github.com/gin-gonic/gin"
)

type DataBundleController interface {
	GetVariationCodes(ctx *gin.Context) *rest.Response
	PurchaseDataBundle(ctx *gin.Context) *rest.Response
}
type DataBundleControllerImpl struct {
	dataBundleService services.DataBundleService
}

var dataBundleResponse rest.Response

func NewDataBundleController(dataBundleService services.DataBundleService) DataBundleController {
	return &DataBundleControllerImpl{
		dataBundleService: dataBundleService,
	}
}

func (ctrl *DataBundleControllerImpl) GetVariationCodes(ctx *gin.Context) *rest.Response {
	var model dtos.VTURequest

	if er := ctx.BindJSON(&model); er != nil {
		return dataBundleResponse.GetError(http.StatusBadRequest, er.Error())
	}

	if m, er := ctrl.dataBundleService.GetVariationCodes(model.ServiceID); er != nil {
		return dataBundleResponse.GetError(http.StatusOK, er.Error())
	} else {
		return dataBundleResponse.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *DataBundleControllerImpl) PurchaseDataBundle(ctx *gin.Context) *rest.Response {
	var model dtos.VTURequest

	if er := ctx.BindJSON(&model); er != nil {
		return dataBundleResponse.GetError(http.StatusBadRequest, er.Error())
	}

	if m, er := ctrl.dataBundleService.PurchaseDataBundle(model.Phone, model.NetworkID, model.Amount); er != nil {
		return dataBundleResponse.GetError(http.StatusOK, er.Error())
	} else {
		return dataBundleResponse.GetSuccess(http.StatusOK, m)
	}
}
