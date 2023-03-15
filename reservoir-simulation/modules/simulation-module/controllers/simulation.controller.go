package controllers

import (
	"net/http"

	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/GabrielAchumba/school-mgt-backend/reservoir-simulation/modules/simulation-module/dtos"
	"github.com/GabrielAchumba/school-mgt-backend/reservoir-simulation/modules/simulation-module/services"

	"github.com/gin-gonic/gin"
)

type SimulationController interface {
	Run(ctx *gin.Context) *rest.Response
}

type controllerImpl struct {
	simulationService services.SimulationService
}

var _response rest.Response

func New(simulationService services.SimulationService) SimulationController {
	return &controllerImpl{
		simulationService: simulationService,
	}
}

func (ctrl *controllerImpl) Run(ctx *gin.Context) *rest.Response {
	var model dtos.SimulationInputDTO
	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	if m, er := ctrl.simulationService.Run(model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}
