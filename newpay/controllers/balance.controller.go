package controllers

import (
	"net/http"

	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/GabrielAchumba/school-mgt-backend/newpay/services"
	"github.com/gin-gonic/gin"
)

type BalanceController interface {
	GetWalletBalance(ctx *gin.Context) *rest.Response
}
type BalanceControllerImpl struct {
	balanceService services.BalanceService
}

var balanceResponse rest.Response

func NewBalanceController(balanceService services.BalanceService) BalanceController {
	return &BalanceControllerImpl{
		balanceService: balanceService,
	}
}

func (ctrl *BalanceControllerImpl) GetWalletBalance(ctx *gin.Context) *rest.Response {

	if m, er := ctrl.balanceService.GetWalletBalance(); er != nil {
		return balanceResponse.GetError(http.StatusOK, er.Error())
	} else {
		return balanceResponse.GetSuccess(http.StatusOK, m)
	}
}
