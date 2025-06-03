package api

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("enrollment", func() {
	Description("Manages attendee enrollment in courses")

	Error("not_found", ErrorResult, "The resource was not found")
	Error("bad_request", ErrorResult, "Invalid request")

	Method("enroll", func() {
		Description("Enroll an attendee in a course")

		Payload(EnrollmentPayload)
		Result(Empty)

		HTTP(func() {
			POST("/enrollment")
			Response(StatusCreated)
			Response("bad_request", StatusBadRequest)
		})
	})

	Method("update_enrollment", func() {
		Description("Update the enrollment status of an attendee in a course")

		Payload(UpdateEnrollmentPayload)
		Result(Empty)

		HTTP(func() {
			PUT("/enrollment")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
	})

	Method("delete_enrollment", func() {
		Description("Delete an attendee's enrollment from a course")

		Payload(DeleteEnrollmentPayload)
		Result(Empty)

		HTTP(func() {
			DELETE("/enrollment/{attendee_id}/{course_id}")
			Response(StatusNoContent)
			Response("not_found", StatusNotFound)
		})
	})

	Method("list_enrolled_users", func() {
		Description("List enrolled users for a specific course")

		Payload(func() {
			Attribute("course_id", Int32, "Course ID")
			Required("course_id")
		})

		Result(ArrayOf(EnrolledUser))
		Error("not_found", ErrorResult)

		HTTP(func() {
			GET("/enrollment/course/{course_id}/users")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
	})
})

var EnrollmentPayload = Type("EnrollmentPayload", func() {
	Description("Input data to enroll an attendee in a course")

	Attribute("attendee_id", Int32, "Attendee ID")
	Attribute("course_id", Int32, "Course ID")
	Attribute("passed", Boolean, "Whether the attendee passed the course")

	Required("attendee_id", "course_id", "passed")
})

var UpdateEnrollmentPayload = Type("UpdateEnrollmentPayload", func() {
	Description("Input data to update the enrollment status")

	Attribute("attendee_id", Int32, "Attendee ID")
	Attribute("course_id", Int32, "Course ID")
	Attribute("passed", Boolean, "New passed status")

	Required("attendee_id", "course_id", "passed")
})

var DeleteEnrollmentPayload = Type("DeleteEnrollmentPayload", func() {
	Description("Input data to delete an enrollment")

	Attribute("attendee_id", Int32, "Attendee ID")
	Attribute("course_id", Int32, "Course ID")

	Required("attendee_id", "course_id")
})

var EnrolledUser = Type("EnrolledUser", func() {
	Description("Details of a user enrolled in a course")

	Attribute("first_name", String, "First name of the user")
	Attribute("remaining_names", String, "Remaining names (middle names) of the user")
	Attribute("last_names", String, "Last names of the user")
	Attribute("email", String, func() {
		Description("Email address")
		Format(FormatEmail)
	})

	Required("first_name", "last_names", "email")
})
