package competitionresultmodule

import (
	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/GabrielAchumba/school-mgt-backend/modules/competition-result-module/controllers"
	"github.com/GabrielAchumba/school-mgt-backend/modules/competition-result-module/services"

	"github.com/GabrielAchumba/school-mgt-backend/authentication/middleware"
	"github.com/GabrielAchumba/school-mgt-backend/authentication/token"
	"github.com/gin-gonic/gin"
)

type CompetitionResultModule struct {
	controller controllers.CompetitionResultController
}

func InjectService(service services.CompetitionResultService) *CompetitionResultModule {
	module := new(CompetitionResultModule)
	module.controller = controllers.New(service)
	return module
}

func (module *CompetitionResultModule) RegisterRoutes(rg *gin.RouterGroup, tokenMaker token.Maker) {
	moduleRoute := rg.Group("/competitionresult")
	serverHttp := rest.ServeHTTP

	moduleRoute.Use(middleware.AuthMiddleware(tokenMaker))
	{
		moduleRoute.POST("/create", serverHttp(module.controller.CreateCompetitionResult))
		moduleRoute.POST("/many", serverHttp(module.controller.CreateCompetitionResults))
		moduleRoute.PUT("/:id", serverHttp(module.controller.PutCompetitionResult))
		moduleRoute.GET("/:schoolId", serverHttp(module.controller.GetCompetitionResults))
		//moduleRoute.GET("/:id", serverHttp(module.controller.GetCompetitionResult))
		moduleRoute.DELETE("/:id/:schoolId", serverHttp(module.controller.DeleteCompetitionResult))
		moduleRoute.POST("/summarizedCompetitionResult2", serverHttp(module.controller.ComputeSummaryCompetitionResults))
		moduleRoute.POST("/summarizedstudentsCompetitionResult", serverHttp(module.controller.ComputeStudentsSummaryCompetitionResults))
		moduleRoute.POST("/summarizedstudentsCompetitionResultbydate", serverHttp(module.controller.ComputeStudentsCompetitionResultsByDateRange))
	}

}
