package Usermodule

import (
	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/user-module/controllers"
	"github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/user-module/services"

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
	moduleRoute := rg.Group("/launchpadusers")
	serverHttp := rest.ServeHTTP

	moduleRoute.POST("/login", serverHttp(module.controller.Login))
	moduleRoute.POST("/register", serverHttp(module.controller.RegisterUser))
	moduleRoute.POST("/forgotpassword", serverHttp(module.controller.ForgotPassword))
	moduleRoute.POST("/resetpassword", serverHttp(module.controller.ResetPassword))

	moduleRoute.Use(middleware.AuthMiddleware(tokenMaker))
	{
		moduleRoute.PUT("/:id", serverHttp(module.controller.PutUser))
		moduleRoute.GET("", serverHttp(module.controller.GetUsers))
		moduleRoute.GET("/:id", serverHttp(module.controller.GetUser))
		moduleRoute.DELETE("/:id", serverHttp(module.controller.DeleteUser))
		moduleRoute.PUT("/updatebiodata/:id", serverHttp(module.controller.UpdateBioData))
		moduleRoute.PUT("/updatebankaccountdata/:id", serverHttp(module.controller.UpdateBankAccountData))
		moduleRoute.PUT("/updatecontactdto/:id", serverHttp(module.controller.UpdateContactDTO))
		moduleRoute.PUT("/updatenextofkindto/:id", serverHttp(module.controller.UpdateNextOfKinDTO))
		moduleRoute.GET("/getallcontributors", serverHttp(module.controller.GetAllContributors))
		moduleRoute.GET("/getcontacts", serverHttp(module.controller.GetContacts))
		moduleRoute.GET("/getbiodata", serverHttp(module.controller.GetBioData))
		moduleRoute.GET("/getbankdetails", serverHttp(module.controller.GetBankDetails))
		moduleRoute.GET("/getnextofkins", serverHttp(module.controller.GetNextOfKins))
		moduleRoute.GET("/getpersonaldataList", serverHttp(module.controller.GetPersonalDataList))

		moduleRoute.GET("/getadministrators", serverHttp(module.controller.GetAdministrators))
		moduleRoute.POST("/registeradministrator", serverHttp(module.controller.RegisterUser))
		moduleRoute.PUT("/updateadministratordto/:id", serverHttp(module.controller.UpdateAdminDTO))
		moduleRoute.POST("/uploadphoto", serverHttp(module.controller.UploadPhoto))
	}
}
