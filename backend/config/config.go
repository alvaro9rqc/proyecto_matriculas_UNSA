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
	HTTP_PORT    = "HTTP_PORT"
	DATABASE_URL = "DATABASE_URL"
	DBG          = "DEBUG"
)

type Config struct {
	DatabaseURL       string
	HttpPort          string
	Dbg               bool
	GoogleOAuthConfig oauth2.Config
	Ctx               context.Context
}

func NewConfig() (*Config, error) {
	devMode := flag.Bool("development", false, "Usar configuración de desarrollo")
	flag.Parse()

	envFile := ".env"
	if *devMode {
		log.Printf("Development mode: %t\n", *devMode)
		envFile = ".env.dev"
	}

	err := godotenv.Load(envFile)
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	databaseURL := os.Getenv(DATABASE_URL)
	if databaseURL == "" {
		return nil, fmt.Errorf("environment variable %s not found", DATABASE_URL)
	}

	httpPort := os.Getenv(HTTP_PORT)
	if httpPort == "" {
		httpPort = "8080"
		log.Printf("Environment variable %s not found, using default value: %s\n", HTTP_PORT, httpPort)
	}

	dbg, err := strconv.ParseBool(os.Getenv(DBG))
	if err != nil {
		dbg = false
		log.Printf("Environment variable %s not found or invalid, using default value: %t\n", DBG, dbg)
	}

	format := goaLog.FormatJSON
	if goaLog.IsTerminal() {
		format = goaLog.FormatTerminal
	}

	ctx := goaLog.Context(context.Background(), goaLog.WithFormat(format))

	if dbg {
		ctx = goaLog.Context(ctx, goaLog.WithDebug())
		goaLog.Debugf(ctx, "debug logs enabled")
	}
	goaLog.Print(ctx, goaLog.KV{K: "http-port", V: httpPort})

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
		HttpPort:          httpPort,
		Dbg:               dbg,
		GoogleOAuthConfig: googleOAuthConfig,
		Ctx:               ctx,
	}, nil
}
