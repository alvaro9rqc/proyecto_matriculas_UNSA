package api

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("course", func() {
	Description("Manage the courses to enrollements")

	Error("not_found", ErrorResult, "The resource was not found")
	Error("bad_request", ErrorResult, "Invalid request")
	Error("un_authorized", ErrorResult, "Unauthorized access")

	Method("upload_all_courses", func() {
		Description("Upload all courses from data of file, only admin can use this method")

		Payload(UploadAllCoursesPayload)

		HTTP(func() {
			POST("/courses/upload")
			Response(StatusCreated)
			Response("bad_request", StatusBadRequest)
			Response("un_authorized", StatusUnauthorized)
		})
	})

	Method("get_all_courses", func() {
		Description("Get all courses, only admin can use this method")

		Payload(func() {
			Attribute("page", Int, func() {
				Description("Page number for pagination")
				Default(1)
				Minimum(1)
			})
			Attribute("limit", Int, func() {
				Description("Number of items per page")
				Default(10)
				Minimum(1)
			})
			Required("page", "limit")
		})

		Result(ArrayOf(Course))

		HTTP(func() {
			GET("/courses")
			Response(StatusOK)
			Response("bad_request", StatusBadRequest)
			Response("un_authorized", StatusUnauthorized)
		})
	})

	Method("get_user_available_courses", func() {
		Description("Get all courses available for the user, only user can use this method")

		Result(ArrayOf(Course))

		HTTP(func() {
			GET("/courses/available")
			Response(StatusOK)
			Response("bad_request", StatusBadRequest)
			Response("un_authorized", StatusUnauthorized)
		})
	})
})

var UploadAllCoursesPayload = Type("UploadAllCoursesPayload", func() {
	Description("Payload for uploading all courses")

	Attribute("courses", ArrayOf(CoursePayload))

	Required("courses")
})

var CoursePayload = Type("CoursePayload", func() {
	Description("Payload for uploading a single course")

	Attribute("name", String, func() {
		Description("Name of the course")
		Example("Introduction to Programming")
		MinLength(1)
		MaxLength(128)
	})

	Attribute("credits", Int, func() {
		Description("Number of credits for the course")
		Example(3)
		Minimum(1)
	})

	Attribute("cicle_number", Int, func() {
		Description("Cicle number of the course")
		Example(1)
		Minimum(1)
	})

	Required("name", "credits", "cicle_number")
})

var Course = Type("Course", func() {
	Description("Course represents a course in the system")

	Extend(CoursePayload)

	Attribute("id", String, func() {
		Description("Unique identifier for the course")
		Example("144")
	})

	Required("id", "name", "credits", "cicle_number")
})
