package db

import (
	"github.com/enrollment/gen/db"
	"github.com/enrollment/src/db/ports"
	"github.com/jackc/pgx/v5"
)

type AccountRepository struct {
	*db.Queries
}

func NewAccountRepository(conn *pgx.Conn) ports.AccountRepositoryInterface {
	return &AccountRepository{
		Queries: db.New(conn),
	}
}
