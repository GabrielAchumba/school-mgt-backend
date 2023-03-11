package cashoutmodule

import (
	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/cashout-module/controllers"
	"github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/cashout-module/services"

	"github.com/GabrielAchumba/school-mgt-backend/authentication/middleware"
	"github.com/GabrielAchumba/school-mgt-backend/authentication/token"
	"github.com/gin-gonic/gin"
)

type CashOutModule struct {
	controller controllers.CashOutController
}

func InjectService(service services.CashOutService) *CashOutModule {
	module := new(CashOutModule)
	module.controller = controllers.New(service)
	return module
}

func (module *CashOutModule) RegisterRoutes(rg *gin.RouterGroup, tokenMaker token.Maker, relativePath string) {
	moduleRoute := rg.Group(relativePath)

	serverHttp := rest.ServeHTTP
	moduleRoute.Use(middleware.AuthMiddleware(tokenMaker))
	{
		moduleRoute.POST("/createcashoutdto", serverHttp(module.controller.CreateCashOutDTO))
		moduleRoute.GET("/getcashouts/:categoryIndex", serverHttp(module.controller.GetCashOuts))
		moduleRoute.GET("/getcashout/:id/:categoryIndex", serverHttp(module.controller.GetCashOut))
		moduleRoute.GET("/getcashoutbycategoryid/:levelIndex/:categoryId/:categoryIndex", serverHttp(module.controller.GetCashOutByCategoryId))
		moduleRoute.POST("/uploadphoto", serverHttp(module.controller.UploadPhoto))
		moduleRoute.GET("/getcategorybankdetails/:categoryIndex", serverHttp(module.controller.GetCategoryBankDetails))
	}
}
