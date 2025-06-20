package repositories

import (
	"github.com/enrollment/gen/db"
	"github.com/enrollment/src/db/ports"
	"github.com/jackc/pgx/v5"
)

type ModalityRepository struct {
	*db.Queries
}

func NewModalityRepository(conn *pgx.Conn) ports.ModalityRepositoryInterface {
	return &ModalityRepository{
		Queries: db.New(conn),
	}
}
