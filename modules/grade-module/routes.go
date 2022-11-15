package Grademodule

import (
	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/GabrielAchumba/school-mgt-backend/modules/grade-module/controllers"
	"github.com/GabrielAchumba/school-mgt-backend/modules/grade-module/services"

	"github.com/GabrielAchumba/school-mgt-backend/authentication/middleware"
	"github.com/GabrielAchumba/school-mgt-backend/authentication/token"
	"github.com/gin-gonic/gin"
)

type GradeModule struct {
	controller controllers.GradeController
}

func InjectService(service services.GradeService) *GradeModule {
	module := new(GradeModule)
	module.controller = controllers.New(service)
	return module
}

func (module *GradeModule) RegisterRoutes(rg *gin.RouterGroup, tokenMaker token.Maker) {
	moduleRoute := rg.Group("/grade")
	serverHttp := rest.ServeHTTP

	moduleRoute.Use(middleware.AuthMiddleware(tokenMaker))
	{
		moduleRoute.POST("/create", serverHttp(module.controller.CreateGrade))
		moduleRoute.POST("/createmany", serverHttp(module.controller.CreateGrades))
		moduleRoute.PUT("/:id", serverHttp(module.controller.PutGrade))
		moduleRoute.GET("/:schoolId", serverHttp(module.controller.GetGrades))
		//moduleRoute.GET("/:id", serverHttp(module.controller.GetGrade))
		moduleRoute.DELETE("/:id/:schoolId", serverHttp(module.controller.DeleteGrade))
	}
}
