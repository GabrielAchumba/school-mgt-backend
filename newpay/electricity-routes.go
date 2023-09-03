package newpay

import (
	"github.com/GabrielAchumba/school-mgt-backend/common/rest"

	"github.com/GabrielAchumba/school-mgt-backend/authentication/token"
	"github.com/GabrielAchumba/school-mgt-backend/newpay/controllers"
	"github.com/GabrielAchumba/school-mgt-backend/newpay/services"
	"github.com/gin-gonic/gin"
)

type ElectricityModule struct {
	controller controllers.ElectricityController
}

func InjectElectricityService(service services.ElectricityService) *ElectricityModule {
	module := new(ElectricityModule)
	module.controller = controllers.NewElectricityController(service)
	return module
}

func (module *ElectricityModule) RegisterElectricityRoutes(rg *gin.RouterGroup, tokenMaker token.Maker, relativePath string) {
	moduleRoute := rg.Group(relativePath)

	serverHttp := rest.ServeHTTP

	/* moduleRoute.Use(middleware.AuthMiddleware(tokenMaker))
	{ */

	moduleRoute.POST("/verify-meter-number", serverHttp(module.controller.MeterNumberVerification))
	moduleRoute.POST("/purchase", serverHttp(module.controller.PurchaseElectricity))
	//}
}
