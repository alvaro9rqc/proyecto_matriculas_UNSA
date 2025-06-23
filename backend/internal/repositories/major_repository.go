package repositories

import (
	"github.com/enrollment/gen/db"
	"github.com/enrollment/internal/ports"
	"github.com/jackc/pgx/v5"
)

type MajorRepository struct {
	*db.Queries
}

func NewMajorRepository(conn *pgx.Conn) ports.MajorRepositoryInterface {
	return &MajorRepository{
		Queries: db.New(conn),
	}
}
