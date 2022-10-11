package Usermodule

import (
	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/GabrielAchumba/school-mgt-backend/modules/assessment-module/controllers"
	"github.com/GabrielAchumba/school-mgt-backend/modules/assessment-module/services"

	"github.com/GabrielAchumba/school-mgt-backend/authentication/middleware"
	"github.com/GabrielAchumba/school-mgt-backend/authentication/token"
	"github.com/gin-gonic/gin"
)

type AssessmentModule struct {
	controller controllers.AssessmentController
}

func InjectService(service services.AssessmentService) *AssessmentModule {
	module := new(AssessmentModule)
	module.controller = controllers.New(service)
	return module
}

func (module *AssessmentModule) RegisterRoutes(rg *gin.RouterGroup, tokenMaker token.Maker) {
	moduleRoute := rg.Group("/assessment")
	serverHttp := rest.ServeHTTP

	moduleRoute.Use(middleware.AuthMiddleware(tokenMaker))
	{
		moduleRoute.POST("/create", serverHttp(module.controller.CreateAssessment))
		moduleRoute.PUT("/:id", serverHttp(module.controller.PutAssessment))
		moduleRoute.GET("/:schoolId", serverHttp(module.controller.GetAssessments))
		//moduleRoute.GET("/:id", serverHttp(module.controller.GetAssessment))
		moduleRoute.DELETE("/:id", serverHttp(module.controller.DeleteAssessment))
	}
}
