package repositories

import (
	"github.com/enrollment/gen/db"
	"github.com/enrollment/internal/ports"
	"github.com/jackc/pgx/v5"
)

type SectionRepository struct {
	*db.Queries
}

func NewSectionRepository(conn *pgx.Conn) ports.SectionRepositoryInterface {
	return &SectionRepository{
		Queries: db.New(conn),
	}
}
