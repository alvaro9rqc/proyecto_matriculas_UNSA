package controllers

import (
	"context"

	course "github.com/enrollment/gen/course"
	"goa.design/clue/log"
)

// course service example implementation.
// The example methods log the requests and return zero values.
type coursesrvc struct{}

// NewCourse returns the course service implementation.
func NewCourse() course.Service {
	return &coursesrvc{}
}

// Upload all courses from data of file, only admin can use this method
func (s *coursesrvc) UploadAllCourses(ctx context.Context, p *course.UploadAllCoursesPayload) (err error) {
	log.Printf(ctx, "course.upload_all_courses")
	return
}

// Get all courses, only admin can use this method
func (s *coursesrvc) GetAllCourses(ctx context.Context, p *course.GetAllCoursesPayload) (res []*course.Course, err error) {
	log.Printf(ctx, "course.get_all_courses")
	return
}

// Get all courses available for the user, only user can use this method
func (s *coursesrvc) GetUserAvailableCourses(ctx context.Context) (res []*course.Course, err error) {
	log.Printf(ctx, "course.get_user_available_courses")
	return
}
