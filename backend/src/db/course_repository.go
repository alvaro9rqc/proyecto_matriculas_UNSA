package db

import (
	"github.com/enrollment/gen/db"
	"github.com/enrollment/src/db/ports"
	"github.com/jackc/pgx/v5"
)

type CourseRepository struct {
	*db.Queries
}

func NewCourseRepository(conn *pgx.Conn) ports.CourseRepositoryInterface {
	return &AccountSessionRepository{
		Queries: db.New(conn),
	}
}
