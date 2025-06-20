package main

import (
	"github.com/enrollment/config"
	"github.com/enrollment/src/db"
)

func main() {
	cfg, err := config.NewSeedConfig()
	if err != nil {
		panic(err)
	}

	conn, err := connectDB(cfg)
	if err != nil {
		panic(err)
	}

	accountRepo := db.NewAccountRepository(conn)
	accountSessionRepo := db.NewAccountSessionRepository(conn)
	courseRepo := db.NewCourseRepository(conn)

}
