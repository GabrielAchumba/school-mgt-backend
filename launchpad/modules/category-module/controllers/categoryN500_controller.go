package controllers

import (
	"net/http"
	"strconv"

	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/category-module/services"
	contributorDTOPackage "github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/user-module/dtos"

	"github.com/gin-gonic/gin"
)

type CategoryController interface {
	CreateCategory(ctx *gin.Context) *rest.Response
	GetCategories(ctx *gin.Context) *rest.Response
	GetCompletedLevelXCategories(ctx *gin.Context) *rest.Response
	GetPersonalDataList(ctx *gin.Context) *rest.Response
}

type controllerImpl struct {
	categoryService services.CategoryService
}

var _response rest.Response

func New(categoryService services.CategoryService) CategoryController {
	return &controllerImpl{
		categoryService: categoryService,
	}
}

func (ctrl *controllerImpl) CreateCategory(ctx *gin.Context) *rest.Response {
	var model contributorDTOPackage.CreateUserRequest
	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	if m, er := ctrl.categoryService.CreateCategoryDTO(model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) GetCategories(ctx *gin.Context) *rest.Response {

	if m, er := ctrl.categoryService.GetCategorys(); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) GetCompletedLevelXCategories(ctx *gin.Context) *rest.Response {
	levelIndex, _ := strconv.Atoi(ctx.Param("levelIndex"))
	categoryIndex, _ := strconv.Atoi(ctx.Param("categoryIndex"))
	if m, er := ctrl.categoryService.GetCompletedLevelXCategorys(levelIndex, categoryIndex); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) GetPersonalDataList(ctx *gin.Context) *rest.Response {
	if m, er := ctrl.categoryService.GetPersonalDataList(); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}
