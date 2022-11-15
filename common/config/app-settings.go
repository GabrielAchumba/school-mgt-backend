package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

//git push heroku main
type Settings struct {
	Token struct {
		AccessTokenDuration   time.Duration
		RefreshTokenDuration  time.Duration
		TokenSecretKey        string
		RefreshTokenSecretKey string
	}
	Server struct {
		Port     string
		UserName string
		Password string
	}
	Database struct {
		DatabaseConnection string
		DatabaseName       string
	}
	TableNames struct {
		User              string
		Staff             string
		Student           string
		Subject           string
		ClassRoom         string
		Result            string
		Assessment        string
		School            string
		Payment           string
		Session           string
		Grade             string
		Level             string
		CompetitionResult string
	}
	EmailData struct {
		Origin    string
		EmailFrom string
		SMTPHost  string
		SMTPPass  string
		SMTPPort  int
		SMTPUser  string
	}
}

var AppSettings = &Settings{}

func Setup() {
	AppSettings.Server.Port = os.Getenv("PORT")

	AppSettings.Database.DatabaseConnection = os.Getenv("DATABASECONNECTION")
	AppSettings.Database.DatabaseName = os.Getenv("DATABASENAME")

	AppSettings.TableNames.User = os.Getenv("USER")
	AppSettings.TableNames.Staff = os.Getenv("STAFF")
	AppSettings.TableNames.Student = os.Getenv("STUDENT")
	AppSettings.TableNames.Subject = os.Getenv("SUBJECT")
	AppSettings.TableNames.ClassRoom = os.Getenv("CLASSROOM")
	AppSettings.TableNames.Result = os.Getenv("RESULT")
	AppSettings.TableNames.Assessment = os.Getenv("ASSESSMENT")
	AppSettings.TableNames.School = os.Getenv("SCHOOL")
	AppSettings.TableNames.Payment = os.Getenv("PAYMENT")
	AppSettings.TableNames.Grade = os.Getenv("GRADE")
	AppSettings.TableNames.Session = os.Getenv("SESSION")
	AppSettings.TableNames.Level = os.Getenv("LEVEL")
	AppSettings.TableNames.CompetitionResult = os.Getenv("COMPETITIONRESULT")

	AppSettings.EmailData.EmailFrom = os.Getenv("EMAIL_FROM")
	AppSettings.EmailData.SMTPHost = os.Getenv("SMTP_HOST")
	AppSettings.EmailData.SMTPPass = os.Getenv("SMTP_PASS")
	AppSettings.EmailData.SMTPPort, _ = strconv.Atoi(os.Getenv("SMTP_PORT"))
	AppSettings.EmailData.SMTPUser = os.Getenv("SMTP_USER")

	accessTokenDuration, _ := time.ParseDuration(os.Getenv("ACCESS_TOKEN_DURATION"))
	refreshTokenDuration, _ := time.ParseDuration(os.Getenv("ACCESS_TOKEN_DURATION"))

	AppSettings.Token.AccessTokenDuration = accessTokenDuration
	AppSettings.Token.RefreshTokenDuration = refreshTokenDuration
	AppSettings.Token.TokenSecretKey = os.Getenv("TOKEN_SECRET_KEY")
	AppSettings.Token.RefreshTokenSecretKey = os.Getenv("REFRESH_TOKEN_SECRET_KEY")

	fmt.Println("App settings was successfully loaded.")
}
