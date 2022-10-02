package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/GabrielAchumba/school-mgt-backend/common/config"

	"time"

	usermodule "github.com/GabrielAchumba/school-mgt-backend/modules/user-module"
	userdtos "github.com/GabrielAchumba/school-mgt-backend/modules/user-module/dtos"
	userService "github.com/GabrielAchumba/school-mgt-backend/modules/user-module/services"

	staffmodule "github.com/GabrielAchumba/school-mgt-backend/modules/staff-module"
	staffService "github.com/GabrielAchumba/school-mgt-backend/modules/staff-module/services"

	studentmodule "github.com/GabrielAchumba/school-mgt-backend/modules/student-module"
	studentService "github.com/GabrielAchumba/school-mgt-backend/modules/student-module/services"

	subjectmodule "github.com/GabrielAchumba/school-mgt-backend/modules/subject-module"
	subjectService "github.com/GabrielAchumba/school-mgt-backend/modules/subject-module/services"

	classroommodule "github.com/GabrielAchumba/school-mgt-backend/modules/classroom-module"
	classroomService "github.com/GabrielAchumba/school-mgt-backend/modules/classroom-module/services"

	resultmodule "github.com/GabrielAchumba/school-mgt-backend/modules/result-module"
	resultService "github.com/GabrielAchumba/school-mgt-backend/modules/result-module/services"

	"github.com/GabrielAchumba/school-mgt-backend/authentication/token"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server         *gin.Engine
	configSettings config.Settings
	mongoClient    *mongo.Client
	ctx            context.Context
)

func isProduction() string {
	return os.Getenv("APP_ENV")
}

func init() {
	gin.SetMode(gin.ReleaseMode)
	server = gin.Default()

	if isProduction() != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	ctx = context.TODO()
	config.Setup()
	configSettings = *config.AppSettings
	mongoConn := options.Client().ApplyURI(config.AppSettings.Database.DatabaseConnection)
	client, err := mongo.Connect(ctx, mongoConn)
	if err != nil {
		log.Fatal(err)
	}
	mongoClient = client
	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongo connection established")
}

func main() {
	defer mongoClient.Disconnect(ctx)

	server.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "*",
		RequestHeaders:  "*",
		ExposedHeaders:  "Content-Length",
		MaxAge:          50 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	}))

	apiBaseName := server.Group("")

	var emailData = userdtos.EmailDto{
		Origin:    configSettings.EmailData.Origin,
		EmailFrom: configSettings.EmailData.EmailFrom,
		SMTPHost:  configSettings.EmailData.SMTPHost,
		SMTPPass:  configSettings.EmailData.SMTPPass,
		SMTPPort:  configSettings.EmailData.SMTPPort,
		SMTPUser:  configSettings.EmailData.SMTPUser,
	}

	tokenMaker, _ := token.NewJWTMaker(config.AppSettings.Token.TokenSecretKey, config.AppSettings.Token.RefreshTokenSecretKey, config.AppSettings.Token.AccessTokenDuration, config.AppSettings.Token.RefreshTokenDuration)

	_userService := userService.New(mongoClient, configSettings, ctx, tokenMaker, emailData)
	usermodule.InjectService(_userService).RegisterRoutes(apiBaseName, tokenMaker)

	_staffService := staffService.New(mongoClient, configSettings, ctx)
	staffmodule.InjectService(_staffService).RegisterRoutes(apiBaseName, tokenMaker)

	_studentService := studentService.New(mongoClient, configSettings, ctx)
	studentmodule.InjectService(_studentService).RegisterRoutes(apiBaseName, tokenMaker)

	_subjectService := subjectService.New(mongoClient, configSettings, ctx)
	subjectmodule.InjectService(_subjectService).RegisterRoutes(apiBaseName, tokenMaker)

	_classRoomService := classroomService.New(mongoClient, configSettings, ctx)
	classroommodule.InjectService(_classRoomService).RegisterRoutes(apiBaseName, tokenMaker)

	_resultService := resultService.New(mongoClient, configSettings, ctx, _userService,
		_studentService, _subjectService)
	resultmodule.InjectService(_resultService).RegisterRoutes(apiBaseName, tokenMaker)

	port := config.AppSettings.Server.Port

	networkingServer := &http.Server{
		Addr:         ":" + port,
		Handler:      server,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Networking service is running on port: " + port)
	log.Fatal(networkingServer.ListenAndServe())
}
