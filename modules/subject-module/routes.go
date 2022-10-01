package Usermodule

import (
	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/GabrielAchumba/school-mgt-backend/modules/subject-module/controllers"
	"github.com/GabrielAchumba/school-mgt-backend/modules/subject-module/services"

	"github.com/GabrielAchumba/school-mgt-backend/authentication/middleware"
	"github.com/GabrielAchumba/school-mgt-backend/authentication/token"
	"github.com/gin-gonic/gin"
)

type SubjectModule struct {
	controller controllers.SubjectController
}

func InjectService(service services.SubjectService) *SubjectModule {
	module := new(SubjectModule)
	module.controller = controllers.New(service)
	return module
}

func (module *SubjectModule) RegisterRoutes(rg *gin.RouterGroup, tokenMaker token.Maker) {
	moduleRoute := rg.Group("/subject")
	serverHttp := rest.ServeHTTP

	moduleRoute.Use(middleware.AuthMiddleware(tokenMaker))
	{
		moduleRoute.POST("/create", serverHttp(module.controller.CreateSubject))
		moduleRoute.PUT("/:id", serverHttp(module.controller.PutSubject))
		moduleRoute.GET("", serverHttp(module.controller.GetSubjects))
		moduleRoute.GET("/:id", serverHttp(module.controller.GetSubject))
		moduleRoute.DELETE("/:id", serverHttp(module.controller.DeleteSubject))
	}
}
