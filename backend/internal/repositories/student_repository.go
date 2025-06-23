package repositories

import (
	"github.com/enrollment/gen/db"
	"github.com/enrollment/internal/ports"
	"github.com/jackc/pgx/v5"
)

type StudentRepository struct {
	*db.Queries
}

func NewStudentRepository(conn *pgx.Conn) ports.StudentRepositoryInterface {
	return &StudentRepository{
		Queries: db.New(conn),
	}
}
