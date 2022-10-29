package paymentmodule

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
		moduleRoute.POST("/uploadphoto", serverHttp(module.controller.UploadPhoto))
		moduleRoute.POST("/create", serverHttp(module.controller.CreatePayment))
		moduleRoute.PUT("/:id", serverHttp(module.controller.PutPayment))
		moduleRoute.GET("pending-payments", serverHttp(module.controller.GetPendingPayments))
		moduleRoute.GET("/:schoolId", serverHttp(module.controller.GetPayment))
		moduleRoute.GET("check-results-subscription/:schoolId", serverHttp(module.controller.CheckSubscription))
		moduleRoute.DELETE("/:id/:schoolId", serverHttp(module.controller.DeletePayment))
	}
}
