package controllers

import (
	"context"

	enrollment "github.com/enrollment/gen/enrollment"
	"goa.design/clue/log"
)

// enrollment service example implementation.
// The example methods log the requests and return zero values.
type enrollmentsrvc struct{}

// NewEnrollment returns the enrollment service implementation.
func NewEnrollment() enrollment.Service {
	return &enrollmentsrvc{}
}

// Enroll an student in selected courses
func (s *enrollmentsrvc) Enroll(ctx context.Context, p *enrollment.EnrollmentPayload) (err error) {
	log.Printf(ctx, "enrollment.enroll")
	return
}

// Get all courses enrolled by an attendee
func (s *enrollmentsrvc) GetEnrollmentCourses(ctx context.Context) (res *enrollment.EnrollmentPayload, err error) {
	res = &enrollment.EnrollmentPayload{}
	log.Printf(ctx, "enrollment.get_enrollment_courses")
	return
}
