package ports

import (
	"context"

	"github.com/enrollment/gen/db"
)

type StudentRepositoryInterface interface {
	CreateStudent(ctx context.Context, arg db.CreateStudentParams) error
}
