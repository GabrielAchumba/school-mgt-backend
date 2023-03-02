package controllers

import (
	"encoding/base64"
	"log"
	"net/http"
	"strconv"

	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/cashout-module/dtos"
	"github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/cashout-module/services"

	"github.com/gin-gonic/gin"
)

type ImageName struct {
	FileName     string `json:"fileName"`
	Base64String string `json:"base64String"`
}

type CashOutController interface {
	UploadPhoto(ctx *gin.Context) *rest.Response
	CreateCashOutDTO(ctx *gin.Context) *rest.Response
	GetCashOuts(ctx *gin.Context) *rest.Response
	GetCashOut(ctx *gin.Context) *rest.Response
	GetCategoryBankDetails(ctx *gin.Context) *rest.Response
	GetCashOutByCategoryId(ctx *gin.Context) *rest.Response
}
type controllerImpl struct {
	cashOutService services.CashOutService
}

var _response rest.Response

func New(cashOutService services.CashOutService) CashOutController {
	return &controllerImpl{
		cashOutService: cashOutService,
	}
}

func (ctrl *controllerImpl) toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func (ctrl *controllerImpl) UploadPhoto(ctx *gin.Context) *rest.Response {

	form, _ := ctx.MultipartForm()

	files := form.File["images[]"]
	imageName := ImageName{}
	checkError := false

	for _, file := range files {

		f, err := file.Open()

		if err != nil {
			log.Println(err)
		}

		defer f.Close()

		size := file.Size
		buffer := make([]byte, size)

		f.Read(buffer)
		fileType := http.DetectContentType(buffer)
		switch fileType {
		case "image/jpeg":
			imageName.Base64String += "data:image/jpeg;base64,"
		case "image/png":
			imageName.Base64String += "data:image/png;base64,"
		case "image/jpg":
			imageName.Base64String += "data:image/jpg;base64,"
		default:
			checkError = true
		}

		if checkError {
			return _response.GetError(http.StatusOK, "image must be: (.png, .jpg, or .jpeg)")
		}

		imageName.Base64String += ctrl.toBase64(buffer)

		imageName.FileName = file.Filename

	}

	return _response.GetSuccess(http.StatusOK, imageName)

}

func (ctrl *controllerImpl) CreateCashOutDTO(ctx *gin.Context) *rest.Response {
	var model dtos.CashOutDTO
	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	if m, er := ctrl.cashOutService.CreateCashOutDTO(model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) GetCashOuts(ctx *gin.Context) *rest.Response {

	categoryIndex, _ := strconv.Atoi(ctx.Param("categoryIndex"))
	if m, er := ctrl.cashOutService.GetCashOuts(categoryIndex); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) GetCashOut(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")
	categoryIndex, _ := strconv.Atoi(ctx.Param("categoryIndex"))

	if m, er := ctrl.cashOutService.GetCashOut(id, categoryIndex); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) GetCategoryBankDetails(ctx *gin.Context) *rest.Response {

	categoryIndex, _ := strconv.Atoi(ctx.Param("categoryIndex"))
	m := ctrl.cashOutService.GetCategoryBankDetails(categoryIndex)
	return _response.GetSuccess(http.StatusOK, m)
}

func (ctrl *controllerImpl) GetCashOutByCategoryId(ctx *gin.Context) *rest.Response {
	levelIndex, _ := strconv.Atoi(ctx.Param("levelIndex"))
	categoryId := ctx.Param("categoryId")

	categoryIndex, _ := strconv.Atoi(ctx.Param("categoryIndex"))
	if m, er := ctrl.cashOutService.GetCashOutByCategoryId(levelIndex, categoryId, categoryIndex); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}
