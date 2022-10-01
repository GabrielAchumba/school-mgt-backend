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
