package config

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	goaLog "goa.design/clue/log"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	googleOauth2 "google.golang.org/api/oauth2/v2"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

const (
	PORT         = "PORT"
	DATABASE_URL = "DATABASE_URL"
	HOST_F       = "HOST"
	DOMAIN_F     = "DOMAIN"
	HTTP_PORT_F  = "HTTP_PORT"
	SECURE_F     = "SECURE"
	DBG_F        = "DEBUG"
)

type Config struct {
	DatabaseURL       string
	HostF             string
	DomainF           string
	HttpPortF         string
	SecureF           bool
	DbgF              bool
	Ctx               context.Context
	GoogleOAuthConfig oauth2.Config
}

func NewConfig() (*Config, error) {
	devMode := flag.Bool("development", false, "Usar configuración de desarrollo")
	flag.Parse()

	envFile := ".env"
	if *devMode {
		envFile = ".env.development"
	}

	err := godotenv.Load(envFile)
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	databaseURL := os.Getenv(DATABASE_URL)
	if databaseURL == "" {
		return nil, fmt.Errorf("environment variable %s not found", DATABASE_URL)
	}

	hostF := os.Getenv(HOST_F)
	if hostF == "" {
		hostF = "localhost"
		log.Printf("Environment variable %s not found, using default value: %s\n", HOST_F, hostF)
	}

	domainF := os.Getenv(DOMAIN_F)
	if domainF == "" {
		domainF = ""
		log.Printf("Environment variable %s not found, using default value: %s\n", DOMAIN_F, domainF)
	}

	httpPortF := os.Getenv(HTTP_PORT_F)
	if httpPortF == "" {
		httpPortF = "8080"
		log.Printf("Environment variable %s not found, using default value: %s\n", HTTP_PORT_F, httpPortF)
	}

	secureF, err := strconv.ParseBool(os.Getenv(SECURE_F))
	if err != nil {
		secureF = false
		log.Printf("Environment variable %s not found or invalid, using default value: %t\n", SECURE_F, secureF)
	}

	dbgF, err := strconv.ParseBool(os.Getenv(DBG_F))
	if err != nil {
		dbgF = false
		log.Printf("Environment variable %s not found or invalid, using default value: %t\n", DBG_F, dbgF)
	}

	format := goaLog.FormatJSON
	if goaLog.IsTerminal() {
		format = goaLog.FormatTerminal
	}

	ctx := goaLog.Context(context.Background(), goaLog.WithFormat(format))

	if dbgF {
		ctx = goaLog.Context(ctx, goaLog.WithDebug())
		goaLog.Debugf(ctx, "debug logs enabled")
	}
	goaLog.Print(ctx, goaLog.KV{K: "http-port", V: httpPortF})

	// configure Google OAuth
	googleOAuthConfig := oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		Endpoint:     google.Endpoint,
		Scopes:       []string{googleOauth2.UserinfoEmailScope, googleOauth2.UserinfoProfileScope},
	}
	return &Config{
		DatabaseURL:       databaseURL,
		HostF:             hostF,
		DomainF:           domainF,
		HttpPortF:         httpPortF,
		SecureF:           secureF,
		DbgF:              dbgF,
		Ctx:               ctx,
		GoogleOAuthConfig: googleOAuthConfig,
	}, nil
}
