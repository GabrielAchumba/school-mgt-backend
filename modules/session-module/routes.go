package sessionmodule

import (
	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/GabrielAchumba/school-mgt-backend/modules/session-module/controllers"
	"github.com/GabrielAchumba/school-mgt-backend/modules/session-module/services"

	"github.com/GabrielAchumba/school-mgt-backend/authentication/middleware"
	"github.com/GabrielAchumba/school-mgt-backend/authentication/token"
	"github.com/gin-gonic/gin"
)

type SessionModule struct {
	controller controllers.SessionController
}

func InjectService(service services.SessionService) *SessionModule {
	module := new(SessionModule)
	module.controller = controllers.New(service)
	return module
}

func (module *SessionModule) RegisterRoutes(rg *gin.RouterGroup, tokenMaker token.Maker) {
	moduleRoute := rg.Group("/session")
	serverHttp := rest.ServeHTTP

	moduleRoute.Use(middleware.AuthMiddleware(tokenMaker))
	{
		moduleRoute.POST("/create", serverHttp(module.controller.CreateSession))
		moduleRoute.POST("/createmany", serverHttp(module.controller.CreateSessions))
		moduleRoute.PUT("/:id", serverHttp(module.controller.PutSession))
		moduleRoute.GET("/:schoolId", serverHttp(module.controller.GetSessions))
		//moduleRoute.GET("/:id", serverHttp(module.controller.GetSession))
		moduleRoute.DELETE("/:id/:schoolId", serverHttp(module.controller.DeleteSession))
	}
}
