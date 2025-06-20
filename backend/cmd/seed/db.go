package main

import (
	"github.com/enrollment/config"
	"github.com/jackc/pgx/v5"
)

func connectDB(cfg *config.SeedConfig) (*pgx.Conn, error) {
	conn, err := pgx.Connect(cfg.Ctx, "user=pqgotest dbname=pqgotest sslmode=verify-full")
	if err != nil {
		return nil, err
	}
	defer conn.Close(cfg.Ctx)

	return conn, nil
}
