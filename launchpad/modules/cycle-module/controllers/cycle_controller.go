package controllers

import (
	"log"
	"net/http"

	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/cycle-module/services"

	"github.com/gin-gonic/gin"
)

type CyleController interface {
	GetCyclesWithLevelsByUserId(ctx *gin.Context) *rest.Response
	GetROIs(ctx *gin.Context) *rest.Response
}
type controllerImpl struct {
	cycleService services.CycleService
}

var _response rest.Response

func New(cycleService services.CycleService) CyleController {
	return &controllerImpl{
		cycleService: cycleService,
	}
}

func (ctrl *controllerImpl) GetCyclesWithLevelsByUserId(ctx *gin.Context) *rest.Response {
	log.Print("GetCyclesWithLevelsByUserId called")
	m := ctrl.cycleService.GetCyclesWithLevelsByUserId()
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) GetROIs(ctx *gin.Context) *rest.Response {
	m := ctrl.cycleService.GetROIs()
	return _response.GetSuccess(http.StatusOK, m)
}
