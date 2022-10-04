package Usermodule

import (
	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/GabrielAchumba/school-mgt-backend/modules/user-module/controllers"
	"github.com/GabrielAchumba/school-mgt-backend/modules/user-module/services"

	"github.com/GabrielAchumba/school-mgt-backend/authentication/middleware"
	"github.com/GabrielAchumba/school-mgt-backend/authentication/token"
	"github.com/gin-gonic/gin"
)

type userModule struct {
	controller controllers.UserController
}

func InjectService(service services.UserService) *userModule {
	module := new(userModule)
	module.controller = controllers.New(service)

	service.SeedAdmin()
	return module
}

func (module *userModule) RegisterRoutes(rg *gin.RouterGroup, tokenMaker token.Maker) {
	moduleRoute := rg.Group("/user")
	serverHttp := rest.ServeHTTP

	moduleRoute.POST("/login", serverHttp(module.controller.Login))
	moduleRoute.POST("/forgotpassword", serverHttp(module.controller.ForgotPassword))
	moduleRoute.POST("/resetpassword", serverHttp(module.controller.ResetPassword))

	moduleRoute.Use(middleware.AuthMiddleware(tokenMaker))
	{
		moduleRoute.POST("/register", serverHttp(module.controller.RegisterUser))
		moduleRoute.PUT("/:id", serverHttp(module.controller.PutUser))
		moduleRoute.GET("", serverHttp(module.controller.GetUsers))
		moduleRoute.GET("/:id", serverHttp(module.controller.GetUser))
		moduleRoute.GET("category/:category", serverHttp(module.controller.GetUsersByCategory))
		moduleRoute.DELETE("/:id", serverHttp(module.controller.DeleteUser))
		moduleRoute.POST("/registeradministrator", serverHttp(module.controller.RegisterUser))
		moduleRoute.PUT("/updateadministratordto/:id", serverHttp(module.controller.UpdateAdminDTO))
		moduleRoute.POST("/uploadphoto", serverHttp(module.controller.UploadPhoto))
	}
}
