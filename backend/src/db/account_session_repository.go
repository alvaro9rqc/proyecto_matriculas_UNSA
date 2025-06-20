package db

import (
	"github.com/enrollment/gen/db"
	"github.com/enrollment/src/db/ports"
	"github.com/jackc/pgx/v5"
)

type AccountSessionRepository struct {
	*db.Queries
}

func NewAccountSessionRepository(conn *pgx.Conn) ports.AccountSessionRepositoryInterface {
	return &AccountSessionRepository{
		Queries: db.New(conn),
	}
}
