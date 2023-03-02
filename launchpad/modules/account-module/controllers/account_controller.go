package controllers

import (
	"encoding/base64"
	"log"
	"net/http"
	"strconv"

	"github.com/GabrielAchumba/school-mgt-backend/common/rest"
	"github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/account-module/dtos"
	"github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/account-module/services"

	"github.com/gin-gonic/gin"
)

type AccountController interface {
	UploadPhoto(ctx *gin.Context) *rest.Response
	OffPlatformPayment(ctx *gin.Context) *rest.Response
	GetUnComfirmedAccounts(ctx *gin.Context) *rest.Response
	ComfirmPayment(ctx *gin.Context) *rest.Response
	RegisteredHaveNotContributed(ctx *gin.Context) *rest.Response
	GetDescendantsByLevel(ctx *gin.Context) *rest.Response
}

type ImageName struct {
	FileName     string `json:"fileName"`
	Base64String string `json:"base64String"`
}

type controllerImpl struct {
	accountService services.AccountService
}

var _response rest.Response

func New(accountService services.AccountService) AccountController {
	return &controllerImpl{
		accountService: accountService,
	}
}

func (ctrl *controllerImpl) toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func (ctrl *controllerImpl) UploadPhoto(ctx *gin.Context) *rest.Response {

	form, _ := ctx.MultipartForm()

	files := form.File["images[]"]

	checkError := false

	//var imageNames []ImageName
	imageName := ImageName{}

	for _, file := range files {

		f, err := file.Open()

		if err != nil {
			log.Println(err)
		}

		defer f.Close()

		size := file.Size
		buffer := make([]byte, size)

		f.Read(buffer)
		//fileBytes := bytes.NewReader(buffer)
		//log.Print("fileBytes: ", fileBytes)
		fileType := http.DetectContentType(buffer)
		//log.Print("fileType: ", fileType)
		//path := "/media/" + file.Filename
		//log.Print("path: ", path)

		// Prepend the appropriate URI scheme header depending
		// on the MIME type
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
		// Append the base64 encoded output
		imageName.Base64String += ctrl.toBase64(buffer)

		// Print the full base64 representation of the image
		log.Print(imageName.Base64String)
		/*             params := &s3.PutObjectInput{
		                   Bucket:        aws.String("article-s3-jpskgc"),
		                   Key:           aws.String(path),
		                   Body:          fileBytes,
		                   ContentLength: aws.Int64(size),
		                   ContentType:   aws.String(fileType),
		               }
		               resp, err := svc.PutObject(params)

		               fmt.Printf("response %s", awsutil.StringValue(resp)) */

		imageName.FileName = file.Filename

		//imageNames = append(imageNames, imageName)
	}

	return _response.GetSuccess(http.StatusOK, imageName)
}

func (ctrl *controllerImpl) OffPlatformPayment(ctx *gin.Context) *rest.Response {
	var model dtos.AccountDTO
	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	if m, er := ctrl.accountService.OffPlatformPayment(model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) GetUnComfirmedAccounts(ctx *gin.Context) *rest.Response {

	if m, er := ctrl.accountService.GetUnComfirmedAccounts(); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) ComfirmPayment(ctx *gin.Context) *rest.Response {
	id := ctx.Param("id")
	var model dtos.CreateAccountRequest
	if er := ctx.BindJSON(&model); er != nil {
		return _response.GetError(http.StatusBadRequest, er.Error())
	}

	if m, er := ctrl.accountService.ComfirmPayment(id, model); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)
	}
}

func (ctrl *controllerImpl) RegisteredHaveNotContributed(ctx *gin.Context) *rest.Response {

	if m, er := ctrl.accountService.RegisteredHaveNotContributed(); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)

	}
}

func (ctrl *controllerImpl) GetDescendantsByLevel(ctx *gin.Context) *rest.Response {
	levelIndex, _ := strconv.Atoi(ctx.Param("levelIndex"))
	parentId := ctx.Param("parentId")
	if m, er := ctrl.accountService.GetDescendantsByLevel(levelIndex, parentId); er != nil {
		return _response.GetError(http.StatusOK, er.Error())
	} else {
		return _response.GetSuccess(http.StatusOK, m)

	}
}
