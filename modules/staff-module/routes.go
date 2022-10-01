package Usermodule

import (
	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/GabrielAchumba/school-mgt-backend/modules/staff-module/controllers"
	"github.com/GabrielAchumba/school-mgt-backend/modules/staff-module/services"

	"github.com/GabrielAchumba/school-mgt-backend/authentication/middleware"
	"github.com/GabrielAchumba/school-mgt-backend/authentication/token"
	"github.com/gin-gonic/gin"
)

type staffModule struct {
	controller controllers.StaffController
}

func InjectService(service services.StaffService) *staffModule {
	module := new(staffModule)
	module.controller = controllers.New(service)
	return module
}

func (module *staffModule) RegisterRoutes(rg *gin.RouterGroup, tokenMaker token.Maker) {
	moduleRoute := rg.Group("/staff")
	serverHttp := rest.ServeHTTP

	moduleRoute.Use(middleware.AuthMiddleware(tokenMaker))
	{
		moduleRoute.POST("/create", serverHttp(module.controller.CreateStaff))
		moduleRoute.PUT("/:id", serverHttp(module.controller.PutStaff))
		moduleRoute.GET("", serverHttp(module.controller.GetStaffs))
		moduleRoute.GET("/:id", serverHttp(module.controller.GetStaff))
		moduleRoute.DELETE("/:id", serverHttp(module.controller.DeleteStaff))
	}
}
