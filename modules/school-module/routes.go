package Usermodule

import (
	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/GabrielAchumba/school-mgt-backend/modules/school-module/controllers"
	"github.com/GabrielAchumba/school-mgt-backend/modules/school-module/services"

	"github.com/GabrielAchumba/school-mgt-backend/authentication/middleware"
	"github.com/GabrielAchumba/school-mgt-backend/authentication/token"
	"github.com/gin-gonic/gin"
)

type SchoolModule struct {
	controller controllers.SchoolController
}

func InjectService(service services.SchoolService) *SchoolModule {
	module := new(SchoolModule)
	module.controller = controllers.New(service)
	return module
}

func (module *SchoolModule) RegisterRoutes(rg *gin.RouterGroup, tokenMaker token.Maker) {
	moduleRoute := rg.Group("/school")
	serverHttp := rest.ServeHTTP

	moduleRoute.Use(middleware.AuthMiddleware(tokenMaker))
	{
		moduleRoute.POST("/create", serverHttp(module.controller.CreateSchool))
		moduleRoute.PUT("/:id", serverHttp(module.controller.PutSchool))
		moduleRoute.GET("", serverHttp(module.controller.GetSchools))
		moduleRoute.GET("/:id", serverHttp(module.controller.GetSchool))
		moduleRoute.DELETE("/:id", serverHttp(module.controller.DeleteSchool))
	}
}
