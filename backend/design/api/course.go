package api

import (
	"github.com/enrollment/design/api/types"
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

		Result(ArrayOf(types.Course))

		HTTP(func() {
			GET("/courses")
			Response(StatusOK)
			Response("bad_request", StatusBadRequest)
			Response("un_authorized", StatusUnauthorized)
		})
	})

	Method("get_user_available_courses", func() {
		Description("Get all courses available for the user, only user can use this method")

		Result(ArrayOf(types.Course))

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

	Attribute("courses", ArrayOf(types.CoursePayload))

	Required("courses")
})

var SpeakerCoursePayload = Type("SpeakerCoursePayload", func() {
	Extend(AccountUser)
})
