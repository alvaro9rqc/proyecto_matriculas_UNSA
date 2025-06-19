package config

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	goaLog "goa.design/clue/log"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

const (
	HTTP_PORT    = "HTTP_PORT"
	DATABASE_URL = "DATABASE_URL"
	DBG          = "DEBUG"
)

type Config struct {
	DatabaseURL string
	HttpPort    string
	Dbg         bool
	Ctx         context.Context
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

	return &Config{
		DatabaseURL: databaseURL,
		HttpPort:    httpPort,
		Dbg:         dbg,
		Ctx:         ctx,
	}, nil
}
