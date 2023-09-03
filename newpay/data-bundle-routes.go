package newpay

import (
	"github.com/GabrielAchumba/school-mgt-backend/common/rest"

	"github.com/GabrielAchumba/school-mgt-backend/authentication/token"
	"github.com/GabrielAchumba/school-mgt-backend/newpay/controllers"
	"github.com/GabrielAchumba/school-mgt-backend/newpay/services"
	"github.com/gin-gonic/gin"
)

type DataBundleModule struct {
	controller controllers.DataBundleController
}

func InjectDataBundleService(service services.DataBundleService) *DataBundleModule {
	module := new(DataBundleModule)
	module.controller = controllers.NewDataBundleController(service)
	return module
}

func (module *DataBundleModule) RegisterDataBundleRoutes(rg *gin.RouterGroup, tokenMaker token.Maker, relativePath string) {
	moduleRoute := rg.Group(relativePath)

	serverHttp := rest.ServeHTTP

	/* moduleRoute.Use(middleware.AuthMiddleware(tokenMaker))
	{ */

	moduleRoute.POST("/fetch-variation-codes", serverHttp(module.controller.GetVariationCodes))
	moduleRoute.POST("/purchase", serverHttp(module.controller.PurchaseDataBundle))
	//}
}
