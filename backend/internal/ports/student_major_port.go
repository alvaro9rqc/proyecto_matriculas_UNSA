package ports

import (
	"context"

	"github.com/enrollment/gen/db"
)

type StudentMajorRepositoryInterface interface {
	CreateStudentMajor(ctx context.Context, arg db.CreateStudentMajorParams) error
	DeleteStudentMajor(ctx context.Context, arg db.DeleteStudentMajorParams) error
	ListMajorsByStudent(ctx context.Context, studentID int32) ([]db.ListMajorsByStudentRow, error)
	ListStudentByMajor(ctx context.Context, majorID int32) ([]db.ListStudentByMajorRow, error)
	ListStudentMajors(ctx context.Context, studentID int32) ([]db.ListStudentMajorsRow, error)
}
