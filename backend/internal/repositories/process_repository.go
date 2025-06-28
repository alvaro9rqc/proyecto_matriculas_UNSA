package repositories

import (
	"github.com/enrollment/gen/db"
	"github.com/enrollment/internal/ports"
	"github.com/jackc/pgx/v5"
)

type ProcessRepository struct {
	*db.Queries
}

func NewProcessRepository(conn *pgx.Conn) ports.ProcessRepositoryInterface {
	return &ProcessRepository{
		Queries: db.New(conn),
	}
}
