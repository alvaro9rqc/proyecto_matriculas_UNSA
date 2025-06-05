package db

import (
	"context"

	"github.com/enrollment/gen/db"
	"github.com/jackc/pgx/v5"
)

type StudentRepository struct {
	querys *db.Queries
}

func NewStudentRepository(psqlDB *pgx.Conn) *StudentRepository {
	return &StudentRepository{
		querys: db.New(psqlDB),
	}
}

func (r *StudentRepository) EnrollmentCourse(ctx context.Context, arg db.EnrollmentStudentCourseParams) error {
	return r.querys.EnrollmentStudentCourse(ctx, arg)
}

func (r *StudentRepository) UpdateEnrollmentCourse(ctx context.Context, arg db.UpdateEnrollmentStudentCourseParams) error {
	return r.querys.UpdateEnrollmentStudentCourse(ctx, arg)
}

func (r *StudentRepository) DeleteEnrollmentCourse(ctx context.Context, arg db.DeleteStudentCourseParams) error {
	return r.querys.DeleteStudentCourse(ctx, arg)
}

func (r *StudentRepository) GetEnrolledUsersByCourse(ctx context.Context, courseID int32) ([]db.GetEnrolledUsersByCourseRow, error) {
	return r.querys.GetEnrolledUsersByCourse(ctx, courseID)
}
