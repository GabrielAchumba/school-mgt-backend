package cyclemodule

import (
	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/cycle-module/controllers"
	"github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/cycle-module/services"

	"github.com/GabrielAchumba/school-mgt-backend/authentication/middleware"
	"github.com/GabrielAchumba/school-mgt-backend/authentication/token"
	"github.com/gin-gonic/gin"
)

type CycleModule struct {
	controller controllers.CyleController
}

func InjectService(service services.CycleService) *CycleModule {
	module := new(CycleModule)
	module.controller = controllers.New(service)
	return module
}

func (module *CycleModule) RegisterRoutes(rg *gin.RouterGroup, tokenMaker token.Maker) {
	moduleRoute := rg.Group("/cycles")

	serverHttp := rest.ServeHTTP

	moduleRoute.GET("/getrois", serverHttp(module.controller.GetROIs))
	moduleRoute.Use(middleware.AuthMiddleware(tokenMaker))
	{
		moduleRoute.GET("/getcycleswithlevelsbyuserid", serverHttp(module.controller.GetCyclesWithLevelsByUserId))
	}
}
