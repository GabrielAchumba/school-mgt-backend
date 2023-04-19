package Filemodule

import (
	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/GabrielAchumba/school-mgt-backend/realestate/file-module/controllers"
	"github.com/GabrielAchumba/school-mgt-backend/realestate/file-module/services"

	"github.com/GabrielAchumba/school-mgt-backend/authentication/middleware"
	"github.com/GabrielAchumba/school-mgt-backend/authentication/token"
	"github.com/gin-gonic/gin"
)

type FileModule struct {
	controller controllers.FileController
}

func InjectService(service services.FileService) *FileModule {
	module := new(FileModule)
	module.controller = controllers.New(service)
	return module
}

func (module *FileModule) RegisterRoutes(rg *gin.RouterGroup, tokenMaker token.Maker) {
	moduleRoute := rg.Group("/file")
	serverHttp := rest.ServeHTTP

	moduleRoute.Use(middleware.AuthMiddleware(tokenMaker))
	{
		moduleRoute.POST("/create", serverHttp(module.controller.CreateFile))
		moduleRoute.PUT("/:id", serverHttp(module.controller.PutFile))
		moduleRoute.GET("/", serverHttp(module.controller.GetFiles))
		moduleRoute.GET("/:userId/:categoryId/:title", serverHttp(module.controller.GetFileByParams))
		moduleRoute.DELETE("/:id", serverHttp(module.controller.DeleteFile))
	}
}
