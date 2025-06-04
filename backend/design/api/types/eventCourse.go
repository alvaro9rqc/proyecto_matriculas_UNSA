package types

import (
	. "goa.design/goa/v3/dsl"
)

var EventCoursePayload = Type("EventCoursePayload", func() {
	Description("Is the schedule of the course in the event")

	Attribute("start_date", Int, func() {
		Description("Start date of the course in the event in Unix timestamp format")
		Example(1748992994154)
	})

	Attribute("end_date", Int, func() {
		Description("End date of the course in the event in Unix timestamp format")
		Example(1748992994154)
	})

	Attribute("instalation", InstalatitonCoursePayload, func() {
		Description("Installation details of the course")
	})

	Required("start_date", "end_date", "instalation")
})

var EventCourse = Type("EventCourse", func() {
	Description("Is the schedule of the course in the event")

	Extend(EventCoursePayload)

	Attribute("id", String, func() {
		Description("Unique identifier for the event course")
		Example("144")
	})

	Required("id")
})
