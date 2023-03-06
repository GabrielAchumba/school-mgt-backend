package paymentgateway

import (
	"github.com/GabrielAchumba/school-mgt-backend/common/rest"

	"github.com/GabrielAchumba/school-mgt-backend/authentication/middleware"
	"github.com/GabrielAchumba/school-mgt-backend/authentication/token"
	"github.com/gin-gonic/gin"
)

type PaymentGatewayModule struct {
	controller PaymentGatewayController
}

func InjectService(service PaymentGatewayService) *PaymentGatewayModule {
	module := new(PaymentGatewayModule)
	module.controller = NewPaymentGatewayController(service)
	return module
}

func (module *PaymentGatewayModule) RegisterRoutes(rg *gin.RouterGroup, tokenMaker token.Maker) {
	moduleRoute := rg.Group("/paymentgateway")

	serverHttp := rest.ServeHTTP

	moduleRoute.Use(middleware.AuthMiddleware(tokenMaker))
	{

		moduleRoute.GET("/getbanks", serverHttp(module.controller.GetBanks))
		moduleRoute.POST("/initiatetransfer", serverHttp(module.controller.InitiateTransfer))
		moduleRoute.POST("/finalizetransfer", serverHttp(module.controller.FinalizeTransfer))
	}
}
