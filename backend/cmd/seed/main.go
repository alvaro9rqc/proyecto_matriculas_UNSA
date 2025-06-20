package main

import (
	"github.com/enrollment/config"
	"github.com/enrollment/src/db"
	"github.com/enrollment/src/db/repositories"
)

func main() {
	cfg, err := config.NewSeedConfig()
	if err != nil {
		panic(err)
	}

	conn, err := db.ConnectDB(cfg)
	if err != nil {
		panic(err)
	}

	oauthRepo := repositories.NewOauthRepository(conn)

	seedOauthRepository(cfg.Ctx, oauthRepo)

}
