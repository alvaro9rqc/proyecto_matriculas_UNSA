package repositories

import (
	"github.com/enrollment/gen/db"
	"github.com/enrollment/internal/ports"
	"github.com/jackc/pgx/v5"
)

type StudentMajorRepository struct {
	*db.Queries
}

func NewStudentMajorRepository(conn *pgx.Conn) ports.StudentMajorRepositoryInterface {
	return &StudentMajorRepository{
		Queries: db.New(conn),
	}
}
