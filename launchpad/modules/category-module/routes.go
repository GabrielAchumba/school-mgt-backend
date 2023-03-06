package catergory

import (
	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/category-module/controllers"
	"github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/category-module/services"

	"github.com/GabrielAchumba/school-mgt-backend/authentication/middleware"
	"github.com/GabrielAchumba/school-mgt-backend/authentication/token"
	"github.com/gin-gonic/gin"
)

type CategoryModule struct {
	controller controllers.CategoryController
}

func InjectService(service services.CategoryService) *CategoryModule {
	module := new(CategoryModule)
	module.controller = controllers.New(service)
	return module
}

func (module *CategoryModule) RegisterRoutes(rg *gin.RouterGroup, tokenMaker token.Maker, relativePath string) {
	moduleRoute := rg.Group(relativePath)

	serverHttp := rest.ServeHTTP
	moduleRoute.Use(middleware.AuthMiddleware(tokenMaker))
	{
		moduleRoute.POST("/createcategory", serverHttp(module.controller.CreateCategory))
		moduleRoute.GET("/getcategories", serverHttp(module.controller.GetCategories))
		moduleRoute.GET("/getcompletedlevelxcategories/:levelIndex/:categoryIndex", serverHttp(module.controller.GetCompletedLevelXCategories))
		moduleRoute.GET("/getpersonaldataList", serverHttp(module.controller.GetPersonalDataList))
	}
}
