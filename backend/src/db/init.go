package db

import (
	"context"
	"os"
	"sync"

	"github.com/jackc/pgx/v5"
)

var (
	conn *pgx.Conn
	once sync.Once
)

func InstanceDB() *pgx.Conn {
	DATABASE_URL := os.Getenv("DATABASE_URL")

	once.Do(func() {
		ctx := context.Background()
		var err error

		conn, err = pgx.Connect(ctx, DATABASE_URL)
		if err != nil {
			panic(err)
		}
		defer conn.Close(ctx)

		if err := conn.Ping(ctx); err != nil {
			panic(err)
		}
	})
	return conn
}
