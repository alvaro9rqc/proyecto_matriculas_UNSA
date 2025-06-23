package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type ConnectDBConfig interface {
	GetConnectDBConfig() (string, context.Context)
}

func ConnectDB(cfg ConnectDBConfig) (*pgx.Conn, error) {
	databaseURL, ctx := cfg.GetConnectDBConfig()

	conn, err := pgx.Connect(ctx, databaseURL)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
