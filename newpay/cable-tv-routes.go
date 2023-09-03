package newpay

import (
	"github.com/GabrielAchumba/school-mgt-backend/common/rest"

	"github.com/GabrielAchumba/school-mgt-backend/authentication/token"
	"github.com/GabrielAchumba/school-mgt-backend/newpay/controllers"
	"github.com/GabrielAchumba/school-mgt-backend/newpay/services"
	"github.com/gin-gonic/gin"
)

type CableTvModule struct {
	controller controllers.CableTvController
}

func InjectService(service services.CableTVService) *CableTvModule {
	module := new(CableTvModule)
	module.controller = controllers.NewCableTvController(service)
	return module
}

func (module *CableTvModule) RegisterRoutes(rg *gin.RouterGroup, tokenMaker token.Maker, relativePath string) {
	moduleRoute := rg.Group(relativePath)

	serverHttp := rest.ServeHTTP

	/* moduleRoute.Use(middleware.AuthMiddleware(tokenMaker))
	{ */

	moduleRoute.POST("/fetch-variation-codes", serverHttp(module.controller.GetVariationCodes))
	moduleRoute.POST("/verify-smartcard-number", serverHttp(module.controller.SmartCardNumberVerification))
	moduleRoute.POST("/purchase", serverHttp(module.controller.PurchaseCableTV))
	//}
}
