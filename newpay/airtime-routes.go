package newpay

import (
	"github.com/GabrielAchumba/school-mgt-backend/common/rest"

	"github.com/GabrielAchumba/school-mgt-backend/authentication/token"
	"github.com/GabrielAchumba/school-mgt-backend/newpay/controllers"
	"github.com/GabrielAchumba/school-mgt-backend/newpay/services"
	"github.com/gin-gonic/gin"
)

type AirtimeModule struct {
	controller controllers.AirtimeController
}

func InjectAirtimeService(service services.AirtimeService) *AirtimeModule {
	module := new(AirtimeModule)
	module.controller = controllers.NewAirtimeController(service)
	return module
}

func (module *AirtimeModule) RegisterAirtimeRoutes(rg *gin.RouterGroup, tokenMaker token.Maker, relativePath string) {
	moduleRoute := rg.Group(relativePath)

	serverHttp := rest.ServeHTTP

	/* moduleRoute.Use(middleware.AuthMiddleware(tokenMaker))
	{ */

	moduleRoute.POST("/purchase", serverHttp(module.controller.PurchaseAirtime))
	//}
}
