package repositories

import (
	"github.com/enrollment/gen/db"
	"github.com/enrollment/internal/ports"
	"github.com/jackc/pgx/v5"
)

type StudentProcessRepository struct {
	*db.Queries
}

func NewStudentProcessRepository(conn *pgx.Conn) ports.StudentProcessRepositoryInterface {
	return &StudentProcessRepository{
		Queries: db.New(conn),
	}
}
