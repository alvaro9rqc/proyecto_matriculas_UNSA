package repositories

import (
	"github.com/enrollment/gen/db"
	"github.com/enrollment/internal/ports"
	"github.com/jackc/pgx/v5"
)

type CourseRepository struct {
	*db.Queries
}

func NewCourseRepository(conn *pgx.Conn) ports.CourseRepositoryInterface {
	return &CourseRepository{
		Queries: db.New(conn),
	}
}
