package usermodule

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
	moduleRoute.POST("/create-admin", serverHttp(module.controller.RegisterAdminOrReferal))
	moduleRoute.GET("/referals", serverHttp(module.controller.GetRerals))
	moduleRoute.POST("/user-is-exist", serverHttp(module.controller.UserIsExist))
	moduleRoute.POST("/user-is-exist2", serverHttp(module.controller.UserIsExist2))
	moduleRoute.GET("/loginstudent/:token/:schoolId", serverHttp(module.controller.LogInStudent))
	moduleRoute.POST("/create", serverHttp(module.controller.RegisterUser))

	moduleRoute.Use(middleware.AuthMiddleware(tokenMaker))
	{
		moduleRoute.POST("/createmany", serverHttp(module.controller.RegisterUsers))
		moduleRoute.PUT("/:id", serverHttp(module.controller.PutUser))
		moduleRoute.PUT("/update-referal/:id", serverHttp(module.controller.PutReferal))
		moduleRoute.PUT("/confirmuser/:id", serverHttp(module.controller.ConfirmUser))
		moduleRoute.PUT("/blockuser/:id", serverHttp(module.controller.BlockUser))
		moduleRoute.GET("/:schoolId", serverHttp(module.controller.GetUsers))
		moduleRoute.GET("/students/:schoolId", serverHttp(module.controller.GetStudents))
		moduleRoute.GET("/paginatedusers/:schoolId/:page/:filter", serverHttp(module.controller.GetPaginatedUnconfirmedUsers))
		moduleRoute.GET("/paginatedconfirmesusers/:schoolId/:page/:filter", serverHttp(module.controller.GetPaginatedConfirmedUsers))
		moduleRoute.GET("/unconfirmedusers/:schoolId/:filter", serverHttp(module.controller.GetUnconfirmedUsers))
		//moduleRoute.GET("/:id", serverHttp(module.controller.GetUser))
		moduleRoute.GET("category/:category/:schoolId", serverHttp(module.controller.GetUsersByCategory))
		moduleRoute.POST("selectedstudents", serverHttp(module.controller.GetStudentsByClassRooms))
		moduleRoute.DELETE("/:id/:schoolId", serverHttp(module.controller.DeleteUser))
		moduleRoute.POST("/registeradministrator", serverHttp(module.controller.RegisterUser))
		moduleRoute.PUT("/updateadministratordto/:id", serverHttp(module.controller.UpdateAdminDTO))
		moduleRoute.POST("/uploadphoto", serverHttp(module.controller.UploadPhoto))
		moduleRoute.POST("/generatetokens", serverHttp(module.controller.GenerateTokens))
		moduleRoute.GET("/get/:token/:schoolId", serverHttp(module.controller.GetStudentByToken))
	}
}
