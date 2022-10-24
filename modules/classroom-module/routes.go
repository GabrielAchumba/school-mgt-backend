package classroommodule

import (
	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/GabrielAchumba/school-mgt-backend/modules/classroom-module/controllers"
	"github.com/GabrielAchumba/school-mgt-backend/modules/classroom-module/services"

	"github.com/GabrielAchumba/school-mgt-backend/authentication/middleware"
	"github.com/GabrielAchumba/school-mgt-backend/authentication/token"
	"github.com/gin-gonic/gin"
)

type ClassRoomModule struct {
	controller controllers.ClassRoomController
}

func InjectService(service services.ClassRoomService) *ClassRoomModule {
	module := new(ClassRoomModule)
	module.controller = controllers.New(service)
	return module
}

func (module *ClassRoomModule) RegisterRoutes(rg *gin.RouterGroup, tokenMaker token.Maker) {
	moduleRoute := rg.Group("/classroom")
	serverHttp := rest.ServeHTTP

	moduleRoute.Use(middleware.AuthMiddleware(tokenMaker))
	{
		moduleRoute.POST("/create", serverHttp(module.controller.CreateClassRoom))
		moduleRoute.PUT("/:id", serverHttp(module.controller.PutClassRoom))
		moduleRoute.GET("/:schoolId", serverHttp(module.controller.GetClassRooms))
		//moduleRoute.GET("/:id", serverHttp(module.controller.GetClassRoom))
		moduleRoute.DELETE("/:id/:schoolId", serverHttp(module.controller.DeleteClassRoom))
	}
}
