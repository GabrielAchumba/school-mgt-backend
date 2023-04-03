package accountmodule

import (
	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/account-module/controllers"
	"github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/account-module/services"

	"github.com/GabrielAchumba/school-mgt-backend/authentication/middleware"
	"github.com/GabrielAchumba/school-mgt-backend/authentication/token"
	"github.com/gin-gonic/gin"
)

type AccountModule struct {
	controller controllers.AccountController
}

func InjectService(service services.AccountService) *AccountModule {
	module := new(AccountModule)
	module.controller = controllers.New(service)
	return module
}

func (module *AccountModule) RegisterRoutes(rg *gin.RouterGroup, tokenMaker token.Maker, relativePath string) {
	serverHttp := rest.ServeHTTP

	moduleRoute := rg.Group(relativePath)

	moduleRoute.Use(middleware.AuthMiddleware(tokenMaker))
	{
		moduleRoute.POST("/uploadphoto", serverHttp(module.controller.UploadPhoto))
		moduleRoute.POST("/offplatformpayment", serverHttp(module.controller.OffPlatformPayment))
		moduleRoute.GET("/getuncomfirmedaccounts", serverHttp(module.controller.GetUnComfirmedAccounts))
		moduleRoute.PUT("/comfirmpayment/:id", serverHttp(module.controller.ComfirmPayment))
		moduleRoute.GET("/registeredhavenotcontributed", serverHttp(module.controller.RegisteredHaveNotContributed))
		moduleRoute.GET("/getdescendantsbylevel/:levelIndex/:parentId", serverHttp(module.controller.GetDescendantsByLevel))
		moduleRoute.GET("/getcompletedlevelxcategories/:levelIndex/:categoryIndex", serverHttp(module.controller.GetCompletedLevelXCategories))
	}
}
