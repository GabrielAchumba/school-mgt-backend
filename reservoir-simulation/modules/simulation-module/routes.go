package simulationmodule

import (
	"github.com/GabrielAchumba/school-mgt-backend/authentication/token"
	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/GabrielAchumba/school-mgt-backend/reservoir-simulation/modules/simulation-module/controllers"
	"github.com/GabrielAchumba/school-mgt-backend/reservoir-simulation/modules/simulation-module/services"
	"github.com/gin-gonic/gin"
)

type SimulationModule struct {
	controller controllers.SimulationController
}

func InjectService(service services.SimulationService) *SimulationModule {
	module := new(SimulationModule)
	module.controller = controllers.New(service)
	return module
}

func (module *SimulationModule) RegisterRoutes(rg *gin.RouterGroup, tokenMaker token.Maker) {
	moduleRoute := rg.Group("/simulation")
	serverHttp := rest.ServeHTTP

	moduleRoute.POST("/run", serverHttp(module.controller.Run))

}
