package types

import (
	. "goa.design/goa/v3/dsl"
)

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

	Attribute("id", Int, func() {
		Description("Unique identifier for the course")
		Example(144)
	})

	Required("id")
})
