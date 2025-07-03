package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ConnectDBConfig interface {
	GetConnectDBConfig() (string, context.Context)
}

func ConnectDB(cfg ConnectDBConfig) (*pgxpool.Pool, error) {
	databaseURL, ctx := cfg.GetConnectDBConfig()

	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return pool, nil
}
