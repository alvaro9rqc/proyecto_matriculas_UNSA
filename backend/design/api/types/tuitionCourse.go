package types

import (
	. "goa.design/goa/v3/dsl"
)

var TuitionCoursePayload = Type("TuitionCoursePayload", func() {
	Description("Tuition details of the course")

	Attribute("total_places", Int, func() {
		Description("Total number of places available for the course")
		Example(30)
		Minimum(1)
	})

	Attribute("taken_places", Int, func() {
		Description("Number of places already taken in the course")
		Example(10)
		Minimum(0)
	})

	Attribute("events", ArrayOf(EventCoursePayload))

	Required("tuition", "currency")
})

var TuitionCourse = Type("TuitionCourse", func() {
	Description("Tuition details of the course")

	Extend(TuitionCoursePayload)

	Attribute("id", String, func() {
		Description("Unique identifier for the tuition course")
		Example("144")
	})

	Required("id")
})
