package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	/*  socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"

	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket" */

	"github.com/GabrielAchumba/school-mgt-backend/common/config"

	"time"

	usermodule "github.com/GabrielAchumba/school-mgt-backend/modules/user-module"
	userdtos "github.com/GabrielAchumba/school-mgt-backend/modules/user-module/dtos"
	userService "github.com/GabrielAchumba/school-mgt-backend/modules/user-module/services"

	simulationmodule "github.com/GabrielAchumba/school-mgt-backend/reservoir-simulation/modules/simulation-module"
	simulationservice "github.com/GabrielAchumba/school-mgt-backend/reservoir-simulation/modules/simulation-module/services"

	staffmodule "github.com/GabrielAchumba/school-mgt-backend/modules/staff-module"
	staffService "github.com/GabrielAchumba/school-mgt-backend/modules/staff-module/services"

	paymentgatewaymodule "github.com/GabrielAchumba/school-mgt-backend/payment-gateway"

	subjectmodule "github.com/GabrielAchumba/school-mgt-backend/modules/subject-module"
	subjectService "github.com/GabrielAchumba/school-mgt-backend/modules/subject-module/services"

	classroommodule "github.com/GabrielAchumba/school-mgt-backend/modules/classroom-module"
	classroomService "github.com/GabrielAchumba/school-mgt-backend/modules/classroom-module/services"

	assessmentmodule "github.com/GabrielAchumba/school-mgt-backend/modules/assessment-module"
	assessmentService "github.com/GabrielAchumba/school-mgt-backend/modules/assessment-module/services"

	schoolmodule "github.com/GabrielAchumba/school-mgt-backend/modules/school-module"
	schoolService "github.com/GabrielAchumba/school-mgt-backend/modules/school-module/services"

	resultmodule "github.com/GabrielAchumba/school-mgt-backend/modules/result-module"
	resultService "github.com/GabrielAchumba/school-mgt-backend/modules/result-module/services"

	paymentmodule "github.com/GabrielAchumba/school-mgt-backend/modules/payment-module"
	paymentService "github.com/GabrielAchumba/school-mgt-backend/modules/payment-module/services"

	sessionmodule "github.com/GabrielAchumba/school-mgt-backend/modules/session-module"
	sessionService "github.com/GabrielAchumba/school-mgt-backend/modules/session-module/services"

	grademodule "github.com/GabrielAchumba/school-mgt-backend/modules/grade-module"
	gradeService "github.com/GabrielAchumba/school-mgt-backend/modules/grade-module/services"

	levelmodule "github.com/GabrielAchumba/school-mgt-backend/modules/level-module"
	levelService "github.com/GabrielAchumba/school-mgt-backend/modules/level-module/services"

	competitionResultmodule "github.com/GabrielAchumba/school-mgt-backend/modules/competition-result-module"
	competitionResultService "github.com/GabrielAchumba/school-mgt-backend/modules/competition-result-module/services"

	//=============LaunchPad Packages===============================================================//

	launchpadusermodule "github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/user-module"
	launchpaduserService "github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/user-module/services"

	accountService "github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/account-module/services"
	cashoutService "github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/cashout-module/services"
	categoryService "github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/category-module/services"
	cycleService "github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/cycle-module/services"

	cyclemodule "github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/cycle-module"

	categoryN10000module "github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/category-module"
	categoryN1000module "github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/category-module"
	categoryN2000module "github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/category-module"
	categoryN5000module "github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/category-module"
	categoryN500module "github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/category-module"

	accountN10000module "github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/account-module"
	accountN1000module "github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/account-module"
	accountN2000module "github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/account-module"
	accountN5000module "github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/account-module"
	accountN500module "github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/account-module"

	cashoutN10000module "github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/cashout-module"
	cashoutN1000module "github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/cashout-module"
	cashoutN2000module "github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/cashout-module"
	cashoutN5000module "github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/cashout-module"
	cashoutN500module "github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/cashout-module"

	//=============================================================================================//

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

// Easier to get running with CORS. Thanks for help @Vindexus and @erkie
var allowOriginFunc = func(r *http.Request) bool {
	return true
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

	_simulationservice := simulationservice.New(ctx)
	simulationmodule.InjectService(_simulationservice).RegisterRoutes(apiBaseName, tokenMaker)

	_paymentgatewayService := paymentgatewaymodule.New(ctx, configSettings)
	paymentgatewaymodule.InjectService(_paymentgatewayService).RegisterRoutes(apiBaseName, tokenMaker)

	_staffService := staffService.New(mongoClient, configSettings, ctx)
	staffmodule.InjectService(_staffService).RegisterRoutes(apiBaseName, tokenMaker)

	_schoolService := schoolService.New(mongoClient, configSettings, ctx)
	schoolmodule.InjectService(_schoolService).RegisterRoutes(apiBaseName, tokenMaker)

	_userService := userService.New(mongoClient, configSettings, ctx, tokenMaker, emailData,
		_staffService)
	usermodule.InjectService(_userService).RegisterRoutes(apiBaseName, tokenMaker)

	_subjectService := subjectService.New(mongoClient, configSettings, ctx)
	subjectmodule.InjectService(_subjectService).RegisterRoutes(apiBaseName, tokenMaker)

	_classRoomService := classroomService.New(mongoClient, configSettings, ctx)
	classroommodule.InjectService(_classRoomService).RegisterRoutes(apiBaseName, tokenMaker)

	_assessmentService := assessmentService.New(mongoClient, configSettings, ctx)
	assessmentmodule.InjectService(_assessmentService).RegisterRoutes(apiBaseName, tokenMaker)

	_paymentService := paymentService.New(mongoClient, configSettings, ctx)
	paymentmodule.InjectService(_paymentService).RegisterRoutes(apiBaseName, tokenMaker)

	_sessionService := sessionService.New(mongoClient, configSettings, ctx)
	sessionmodule.InjectService(_sessionService).RegisterRoutes(apiBaseName, tokenMaker)

	_gradeService := gradeService.New(mongoClient, configSettings, ctx)
	grademodule.InjectService(_gradeService).RegisterRoutes(apiBaseName, tokenMaker)

	_levelService := levelService.New(mongoClient, configSettings, ctx)
	levelmodule.InjectService(_levelService).RegisterRoutes(apiBaseName, tokenMaker)

	_resultService := resultService.New(mongoClient, configSettings, ctx, _userService,
		_subjectService, _classRoomService, _assessmentService,
		_staffService, _sessionService, _gradeService, _levelService)
	resultmodule.InjectService(_resultService).RegisterRoutes(apiBaseName, tokenMaker)

	_competitionResultService := competitionResultService.New(mongoClient, configSettings, ctx, _userService,
		_subjectService, _classRoomService, _assessmentService,
		_staffService, _sessionService, _gradeService, _levelService)
	competitionResultmodule.InjectService(_competitionResultService).RegisterRoutes(apiBaseName, tokenMaker)

	//==============Launch Pad Set Up========================================================//

	_launchpaduserService := launchpaduserService.New(mongoClient, configSettings, ctx, tokenMaker, emailData)
	launchpadusermodule.InjectService(_launchpaduserService).RegisterRoutes(apiBaseName, tokenMaker)

	categoryN500Collection := mongoClient.Database(configSettings.Database.DatabaseName).Collection(configSettings.TableNames.CategoryN500)
	_categoryN500Service := categoryService.New(categoryN500Collection, configSettings, ctx, _launchpaduserService)
	categoryN500module.InjectService(_categoryN500Service).RegisterRoutes(apiBaseName, tokenMaker, "/categoryn500")

	categoryN1000Collection := mongoClient.Database(configSettings.Database.DatabaseName).Collection(configSettings.TableNames.CategoryN1000)
	_categoryN1000Service := categoryService.New(categoryN1000Collection, configSettings, ctx, _launchpaduserService)
	categoryN1000module.InjectService(_categoryN1000Service).RegisterRoutes(apiBaseName, tokenMaker, "/categoryn1000")

	categoryN2000Collection := mongoClient.Database(configSettings.Database.DatabaseName).Collection(configSettings.TableNames.CategoryN2000)
	_categoryN2000Service := categoryService.New(categoryN2000Collection, configSettings, ctx, _launchpaduserService)
	categoryN2000module.InjectService(_categoryN2000Service).RegisterRoutes(apiBaseName, tokenMaker, "/categoryn2000")

	categoryN5000Collection := mongoClient.Database(configSettings.Database.DatabaseName).Collection(configSettings.TableNames.CategoryN5000)
	_categoryN5000Service := categoryService.New(categoryN5000Collection, configSettings, ctx, _launchpaduserService)
	categoryN5000module.InjectService(_categoryN5000Service).RegisterRoutes(apiBaseName, tokenMaker, "/categoryn5000")

	categoryN10000Collection := mongoClient.Database(configSettings.Database.DatabaseName).Collection(configSettings.TableNames.CategoryN10000)
	_categoryN10000Service := categoryService.New(categoryN10000Collection, configSettings, ctx, _launchpaduserService)
	categoryN10000module.InjectService(_categoryN10000Service).RegisterRoutes(apiBaseName, tokenMaker, "/categoryn10000")

	cyclemodule.InjectService(cycleService.New(ctx)).RegisterRoutes(apiBaseName, tokenMaker)

	accountN500Collection := mongoClient.Database(configSettings.Database.DatabaseName).Collection(configSettings.TableNames.AccountN500)
	accountN500module.InjectService(accountService.New(accountN500Collection, configSettings, ctx, _categoryN500Service, _launchpaduserService)).RegisterRoutes(apiBaseName, tokenMaker, "/accountn500")

	accountN1000Collection := mongoClient.Database(configSettings.Database.DatabaseName).Collection(configSettings.TableNames.AccountN1000)
	accountN1000module.InjectService(accountService.New(accountN1000Collection, configSettings, ctx, _categoryN1000Service, _launchpaduserService)).RegisterRoutes(apiBaseName, tokenMaker, "/accountn1000")

	accountN2000Collection := mongoClient.Database(configSettings.Database.DatabaseName).Collection(configSettings.TableNames.AccountN2000)
	accountN2000module.InjectService(accountService.New(accountN2000Collection, configSettings, ctx, _categoryN2000Service, _launchpaduserService)).RegisterRoutes(apiBaseName, tokenMaker, "/accountn2000")

	accountN5000Collection := mongoClient.Database(configSettings.Database.DatabaseName).Collection(configSettings.TableNames.AccountN5000)
	accountN5000module.InjectService(accountService.New(accountN5000Collection, configSettings, ctx, _categoryN5000Service, _launchpaduserService)).RegisterRoutes(apiBaseName, tokenMaker, "/accountn5000")

	accountN10000Collection := mongoClient.Database(configSettings.Database.DatabaseName).Collection(configSettings.TableNames.AccountN10000)
	accountN10000module.InjectService(accountService.New(accountN10000Collection, configSettings, ctx, _categoryN10000Service, _launchpaduserService)).RegisterRoutes(apiBaseName, tokenMaker, "/accountn10000")

	cashoutN500Collection := mongoClient.Database(configSettings.Database.DatabaseName).Collection(configSettings.TableNames.CashOutN500)
	cashoutN500module.InjectService(cashoutService.New(cashoutN500Collection, configSettings, ctx, _categoryN500Service)).RegisterRoutes(apiBaseName, tokenMaker, "/cashoutn500")

	cashoutN1000Collection := mongoClient.Database(configSettings.Database.DatabaseName).Collection(configSettings.TableNames.CashOutN1000)
	cashoutN1000module.InjectService(cashoutService.New(cashoutN1000Collection, configSettings, ctx, _categoryN1000Service)).RegisterRoutes(apiBaseName, tokenMaker, "/cashoutn1000")

	cashoutN2000Collection := mongoClient.Database(configSettings.Database.DatabaseName).Collection(configSettings.TableNames.CashOutN2000)
	cashoutN2000module.InjectService(cashoutService.New(cashoutN2000Collection, configSettings, ctx, _categoryN2000Service)).RegisterRoutes(apiBaseName, tokenMaker, "/cashoutn2000")

	cashoutN5000Collection := mongoClient.Database(configSettings.Database.DatabaseName).Collection(configSettings.TableNames.CashOutN5000)
	cashoutN5000module.InjectService(cashoutService.New(cashoutN5000Collection, configSettings, ctx, _categoryN1000Service)).RegisterRoutes(apiBaseName, tokenMaker, "/cashoutn5000")

	cashoutN10000Collection := mongoClient.Database(configSettings.Database.DatabaseName).Collection(configSettings.TableNames.CashOutN10000)
	cashoutN10000module.InjectService(cashoutService.New(cashoutN10000Collection, configSettings, ctx, _categoryN10000Service)).RegisterRoutes(apiBaseName, tokenMaker, "/cashoutn10000")

	//=======================================================================================//

	port := config.AppSettings.Server.Port

	//socketServer := socketio.NewServer(nil)

	/* socketServer := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			&websocket.Transport{
				CheckOrigin: allowOriginFunc,
			},
		},
	}) */

	/* socketServer.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		log.Println("connected:", s.ID())
		return nil
	})

	socketServer.OnEvent("/", "chat-message", func(s socketio.Conn, data map[string]interface{}) {
		log.Println("chat-message:", data["message"])
		s.Emit("chat-message", data)
	})

	socketServer.OnEvent("/", "typing", func(s socketio.Conn, data map[string]interface{}) {
		log.Println("typing:", data["message"])
		s.Emit("typing", data)
	})

	socketServer.OnEvent("/", "stopTyping", func(s socketio.Conn, data map[string]interface{}) {
		log.Println("stopTyping")
		s.Emit("stopTyping")
	}) */

	/* socketServer.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
	})

	socketServer.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})

	socketServer.OnError("/", func(s socketio.Conn, e error) {
		log.Println("meet error:", e)
	}) */

	/* socketServer.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("closed", reason, s.Context())
	}) */

	/* go func() {
		if err := socketServer.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}() */

	/* go socketServer.Serve()
	defer socketServer.Close() */

	//
	//http.Handle("/socket.io/", socketServer)

	networkingServer := &http.Server{
		Addr:         ":" + port,
		Handler:      server,
		ReadTimeout:  600 * time.Second,
		WriteTimeout: 1200 * time.Second,
	}

	fmt.Println("Networking service is running on port: " + port)
	log.Fatal(networkingServer.ListenAndServe())
}
