package Usermodule

import (
	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/GabrielAchumba/school-mgt-backend/modules/payment-module/controllers"
	"github.com/GabrielAchumba/school-mgt-backend/modules/payment-module/services"

	"github.com/GabrielAchumba/school-mgt-backend/authentication/middleware"
	"github.com/GabrielAchumba/school-mgt-backend/authentication/token"
	"github.com/gin-gonic/gin"
)

type PaymentModule struct {
	controller controllers.PaymentController
}

func InjectService(service services.PaymentService) *PaymentModule {
	module := new(PaymentModule)
	module.controller = controllers.New(service)
	return module
}

func (module *PaymentModule) RegisterRoutes(rg *gin.RouterGroup, tokenMaker token.Maker) {
	moduleRoute := rg.Group("/payment")
	serverHttp := rest.ServeHTTP

	moduleRoute.Use(middleware.AuthMiddleware(tokenMaker))
	{
		moduleRoute.POST("/create", serverHttp(module.controller.CreatePayment))
		moduleRoute.GET("/:schoolId", serverHttp(module.controller.GetPayments))
		//moduleRoute.GET("/:id", serverHttp(module.controller.GetPayment))
		moduleRoute.DELETE("/:id/:schoolId", serverHttp(module.controller.DeletePayment))
	}
}
