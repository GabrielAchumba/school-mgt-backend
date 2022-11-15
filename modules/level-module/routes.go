package Levelmodule

import (
	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/GabrielAchumba/school-mgt-backend/modules/level-module/controllers"
	"github.com/GabrielAchumba/school-mgt-backend/modules/level-module/services"

	"github.com/GabrielAchumba/school-mgt-backend/authentication/middleware"
	"github.com/GabrielAchumba/school-mgt-backend/authentication/token"
	"github.com/gin-gonic/gin"
)

type LevelModule struct {
	controller controllers.LevelController
}

func InjectService(service services.LevelService) *LevelModule {
	module := new(LevelModule)
	module.controller = controllers.New(service)
	return module
}

func (module *LevelModule) RegisterRoutes(rg *gin.RouterGroup, tokenMaker token.Maker) {
	moduleRoute := rg.Group("/level")
	serverHttp := rest.ServeHTTP

	moduleRoute.Use(middleware.AuthMiddleware(tokenMaker))
	{
		moduleRoute.POST("/create", serverHttp(module.controller.CreateLevel))
		moduleRoute.POST("/createmany", serverHttp(module.controller.CreateLevels))
		moduleRoute.PUT("/:id", serverHttp(module.controller.PutLevel))
		moduleRoute.GET("/:schoolId", serverHttp(module.controller.GetLevels))
		//moduleRoute.GET("/:id", serverHttp(module.controller.GetLevel))
		moduleRoute.DELETE("/:id/:schoolId", serverHttp(module.controller.DeleteLevel))
	}
}
