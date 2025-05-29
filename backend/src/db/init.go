package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func ConnectDB(DATABASE_URL string) (*pgx.Conn, error) {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, DATABASE_URL)
	if err != nil {
		return nil, err
	}
	defer conn.Close(ctx)

	if err := conn.Ping(ctx); err != nil {
		return nil, err
	}

	return conn, nil
}
