package repositories

import (
	"github.com/enrollment/gen/db"
	"github.com/enrollment/internal/ports"
	"github.com/jackc/pgx/v5"
)

type InstitutionRepository struct {
	*db.Queries
}

func NewInstitutionRepository(conn *pgx.Conn) ports.InstitutionRepositoryInterface {
	return &InstitutionRepository{
		Queries: db.New(conn),
	}
}
