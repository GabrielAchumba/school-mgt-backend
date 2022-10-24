package studentmodule

import (
	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/GabrielAchumba/school-mgt-backend/modules/student-module/controllers"
	"github.com/GabrielAchumba/school-mgt-backend/modules/student-module/services"

	"github.com/GabrielAchumba/school-mgt-backend/authentication/middleware"
	"github.com/GabrielAchumba/school-mgt-backend/authentication/token"
	"github.com/gin-gonic/gin"
)

type StudentModule struct {
	controller controllers.StudentController
}

func InjectService(service services.StudentService) *StudentModule {
	module := new(StudentModule)
	module.controller = controllers.New(service)
	return module
}

func (module *StudentModule) RegisterRoutes(rg *gin.RouterGroup, tokenMaker token.Maker) {
	moduleRoute := rg.Group("/student")
	serverHttp := rest.ServeHTTP

	moduleRoute.Use(middleware.AuthMiddleware(tokenMaker))
	{
		moduleRoute.POST("/create", serverHttp(module.controller.CreateStudent))
		moduleRoute.PUT("/:id", serverHttp(module.controller.PutStudent))
		moduleRoute.GET("/:schoolId", serverHttp(module.controller.GetStudents))
		//moduleRoute.GET("/:id", serverHttp(module.controller.GetStudent))
		moduleRoute.DELETE("/:id/:schoolId", serverHttp(module.controller.DeleteStudent))
	}
}
