package newpay

import (
	"github.com/GabrielAchumba/school-mgt-backend/common/rest"

	"github.com/GabrielAchumba/school-mgt-backend/authentication/token"
	"github.com/GabrielAchumba/school-mgt-backend/newpay/controllers"
	"github.com/GabrielAchumba/school-mgt-backend/newpay/services"
	"github.com/gin-gonic/gin"
)

type BalanceModule struct {
	controller controllers.BalanceController
}

func InjectBalanceService(service services.BalanceService) *BalanceModule {
	module := new(BalanceModule)
	module.controller = controllers.NewBalanceController(service)
	return module
}

func (module *BalanceModule) RegisterBalanceRoutes(rg *gin.RouterGroup, tokenMaker token.Maker, relativePath string) {
	moduleRoute := rg.Group(relativePath)

	serverHttp := rest.ServeHTTP

	/* moduleRoute.Use(middleware.AuthMiddleware(tokenMaker))
	{ */

	moduleRoute.GET("/balance", serverHttp(module.controller.GetWalletBalance))
	//}
}
