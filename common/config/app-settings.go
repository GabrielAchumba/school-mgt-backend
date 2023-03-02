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
		AccountN500       string
		AccountN1000      string
		AccountN2000      string
		AccountN5000      string
		AccountN10000     string
		LaunchpadUser     string
		Contributor       string
		CashOutN500       string
		CashOutN1000      string
		CashOutN2000      string
		CashOutN5000      string
		CashOutN10000     string
		CategoryN500      string
		CategoryN1000     string
		CategoryN2000     string
		CategoryN5000     string
		CategoryN10000    string
	}
	EmailData struct {
		Origin    string
		EmailFrom string
		SMTPHost  string
		SMTPPass  string
		SMTPPort  int
		SMTPUser  string
	}
	PayStackKey struct {
		TestKey   string
		ActualKey string
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

	AppSettings.TableNames.AccountN500 = os.Getenv("ACCOUNTN500")
	AppSettings.TableNames.AccountN1000 = os.Getenv("ACCOUNTN1000")
	AppSettings.TableNames.AccountN2000 = os.Getenv("ACCOUNTN2000")
	AppSettings.TableNames.AccountN5000 = os.Getenv("ACCOUNTN5000")
	AppSettings.TableNames.AccountN10000 = os.Getenv("ACCOUNTN10000")
	AppSettings.TableNames.LaunchpadUser = os.Getenv("LAUNCHPADUSER")
	AppSettings.TableNames.Contributor = os.Getenv("CONTRIBUTOR")
	AppSettings.TableNames.CashOutN500 = os.Getenv("CASHOUTN500")
	AppSettings.TableNames.CashOutN1000 = os.Getenv("CASHOUTN1000")
	AppSettings.TableNames.CashOutN2000 = os.Getenv("CASHOUTN2000")
	AppSettings.TableNames.CashOutN5000 = os.Getenv("CASHOUTN5000")
	AppSettings.TableNames.CashOutN10000 = os.Getenv("CASHOUTN10000")
	AppSettings.TableNames.CategoryN500 = os.Getenv("CATEGORYN500")
	AppSettings.TableNames.CategoryN1000 = os.Getenv("CATEGORYN1000")
	AppSettings.TableNames.CategoryN2000 = os.Getenv("CATEGORYN2000")
	AppSettings.TableNames.CategoryN5000 = os.Getenv("CATEGORYN5000")
	AppSettings.TableNames.CategoryN10000 = os.Getenv("CATEGORYN10000")

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

	AppSettings.PayStackKey.TestKey = "sk_test_574ef246a0c1d74f8b9e0b8b10214d3959a00b01"

	fmt.Println("App settings was successfully loaded.")
}
