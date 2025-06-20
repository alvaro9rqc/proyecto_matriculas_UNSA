package repositories

import (
	"github.com/enrollment/gen/db"
	"github.com/enrollment/src/db/ports"
	"github.com/jackc/pgx/v5"
)

type StudentGroupRepository struct {
	*db.Queries
}

func NewStudentGroupRepository(conn *pgx.Conn) ports.StudentGroupRepositoryInterface {
	return &StudentGroupRepository{
		Queries: db.New(conn),
	}
}
