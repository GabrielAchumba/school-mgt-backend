package usermodule

import (
	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/GabrielAchumba/school-mgt-backend/realestate/user-module/controllers"
	"github.com/GabrielAchumba/school-mgt-backend/realestate/user-module/services"

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
	moduleRoute := rg.Group("/realestateuser")
	serverHttp := rest.ServeHTTP

	moduleRoute.POST("/login", serverHttp(module.controller.Login))
	moduleRoute.POST("/resetpassword", serverHttp(module.controller.ResetPassword))
	moduleRoute.POST("/user-is-exist", serverHttp(module.controller.UserIsExist))
	moduleRoute.POST("/user-is-exist2", serverHttp(module.controller.UserIsExist2))
	moduleRoute.POST("/create", serverHttp(module.controller.RegisterUser))

	moduleRoute.Use(middleware.AuthMiddleware(tokenMaker))
	{
		moduleRoute.PUT("/:id", serverHttp(module.controller.PutUser))
		moduleRoute.PUT("/confirmuser/:id", serverHttp(module.controller.ConfirmUser))
		moduleRoute.PUT("/blockuser/:id", serverHttp(module.controller.BlockUser))
		moduleRoute.DELETE("/:id/:schoolId", serverHttp(module.controller.DeleteUser))
		moduleRoute.POST("/uploadphoto", serverHttp(module.controller.UploadPhoto))
	}
}
