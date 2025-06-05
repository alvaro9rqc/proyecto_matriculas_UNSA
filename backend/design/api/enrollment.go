package api

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("enrollment", func() {
	Description("Manages attendee enrollment in courses")

	Error("not_found", ErrorResult, "The resource was not found")
	Error("bad_request", ErrorResult, "Invalid request")
	Error("un_authorized", ErrorResult, "Unauthorized access")

	Method("enroll", func() {
		Description("Enroll an student in selected courses")

		Payload(EnrollmentCourses)
		Result(Empty)

		HTTP(func() {
			POST("/enrollment/enroll")
			Response(StatusCreated)
			Response("un_authorized", StatusUnauthorized)
			Response("bad_request", StatusBadRequest)
		})
	})

	Method("get_enrollment_courses", func() {
		Description("Get all courses enrolled by an attendee")

		Result(EnrollmentCourses)

		HTTP(func() {
			POST("/enrollment/enrollement_courses")
			Response(StatusOK)
			Response("un_authorized", StatusUnauthorized)
			Response("bad_request", StatusBadRequest)
		})
	})
})

var EnrollCourse = Type("EnrollCourseType", func() {
	Description("Represents a course enrollment")

	Attribute("id", Int32, "Enrollment ID")
	Attribute("course_id", Int32, "Course ID")
	Attribute("program_id", Int32, "Program ID")

	Required("id", "course_id", "program_id")
})

var EnrollmentCourses = Type("EnrollmentPayload", func() {
	Description("Input data to enroll an attendee in a courses")

	Attribute("enrollCourses", ArrayOf(EnrollCourse), "Attendee ID")

	Required("enrollCourses")
})
