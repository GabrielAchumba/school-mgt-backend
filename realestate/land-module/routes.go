package Landmodule

import (
	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/GabrielAchumba/school-mgt-backend/realestate/land-module/controllers"
	"github.com/GabrielAchumba/school-mgt-backend/realestate/land-module/services"

	"github.com/GabrielAchumba/school-mgt-backend/authentication/middleware"
	"github.com/GabrielAchumba/school-mgt-backend/authentication/token"
	"github.com/gin-gonic/gin"
)

type LandModule struct {
	controller controllers.LandController
}

func InjectService(service services.LandService) *LandModule {
	module := new(LandModule)
	module.controller = controllers.New(service)
	return module
}

func (module *LandModule) RegisterRoutes(rg *gin.RouterGroup, tokenMaker token.Maker, relativePath string) {
	moduleRoute := rg.Group(relativePath)
	serverHttp := rest.ServeHTTP

	moduleRoute.Use(middleware.AuthMiddleware(tokenMaker))
	{
		moduleRoute.POST("/create", serverHttp(module.controller.CreateLand))
		moduleRoute.PUT("/:id", serverHttp(module.controller.PutLand))
		moduleRoute.GET("/:page/:filterModel", serverHttp(module.controller.GetPaginatedLands))
		//moduleRoute.GET("/:filterModel", serverHttp(module.controller.GetLands))
		moduleRoute.DELETE("/:id", serverHttp(module.controller.DeleteLand))
	}
}
