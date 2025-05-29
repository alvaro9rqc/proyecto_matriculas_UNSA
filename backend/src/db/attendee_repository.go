package db

import (
	"context"

	"github.com/enrollment/gen/db"
	"github.com/jackc/pgx/v5"
)

type AttendeeRepository struct {
	querys *db.Queries
}

func NewAttendeeRepository(psqlDB *pgx.Conn) *AttendeeRepository {
	return &AttendeeRepository{
		querys: db.New(psqlDB),
	}
}

func (r *AttendeeRepository) EnrollmentCourse(ctx context.Context, arg db.EnrollmentAttendeeCourseParams) error {
	return r.querys.EnrollmentAttendeeCourse(ctx, arg)
}

func (r *AttendeeRepository) UpdateEnrollmentCourse(ctx context.Context, arg db.UpdateEnrollmentAttendeeCourseParams) error {
	return r.querys.UpdateEnrollmentAttendeeCourse(ctx, arg)
}

func (r *AttendeeRepository) DeleteEnrollmentCourse(ctx context.Context, arg db.DeleteAttendeeCourseParams) error {
	return r.querys.DeleteAttendeeCourse(ctx, arg)
}

func (r *AttendeeRepository) GetEnrolledUsersByCourse(ctx context.Context, courseID int32) ([]db.GetEnrolledUsersByCourseRow, error) {
	return r.querys.GetEnrolledUsersByCourse(ctx, courseID)
}
