package Usermodule

import (
	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/GabrielAchumba/school-mgt-backend/modules/result-module/controllers"
	"github.com/GabrielAchumba/school-mgt-backend/modules/result-module/services"

	"github.com/GabrielAchumba/school-mgt-backend/authentication/middleware"
	"github.com/GabrielAchumba/school-mgt-backend/authentication/token"
	"github.com/gin-gonic/gin"
)

type ResultModule struct {
	controller controllers.ResultController
}

func InjectService(service services.ResultService) *ResultModule {
	module := new(ResultModule)
	module.controller = controllers.New(service)
	return module
}

func (module *ResultModule) RegisterRoutes(rg *gin.RouterGroup, tokenMaker token.Maker) {
	moduleRoute := rg.Group("/result")
	serverHttp := rest.ServeHTTP

	moduleRoute.Use(middleware.AuthMiddleware(tokenMaker))
	{
		moduleRoute.POST("/create", serverHttp(module.controller.CreateResult))
		moduleRoute.PUT("/:id", serverHttp(module.controller.PutResult))
		moduleRoute.GET("", serverHttp(module.controller.GetResults))
		moduleRoute.GET("/:id", serverHttp(module.controller.GetResult))
		moduleRoute.DELETE("/:id", serverHttp(module.controller.DeleteResult))
	}
}
