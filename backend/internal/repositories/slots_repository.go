package repositories

import (
	"github.com/enrollment/gen/db"
	"github.com/enrollment/internal/ports"
	"github.com/jackc/pgx/v5"
)

type SlotsRepository struct {
	*db.Queries
}

func NewSlotsRepository(conn *pgx.Conn) ports.SlotsRepositoryInterface {
	return &SlotsRepository{
		Queries: db.New(conn),
	}
}
