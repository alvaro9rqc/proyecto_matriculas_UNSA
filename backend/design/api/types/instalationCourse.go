package types

import (
	. "goa.design/goa/v3/dsl"
)

var InstalatitonCoursePayload = Type("InstalatitonCoursePayload", func() {
	Description("Represents a course in the installation process")

	Attribute("name", String, func() {
		Description("Name of the place where the course is installed")
		Example("Sistem Ingenier School")
		MinLength(1)
		MaxLength(128)
	})

	Attribute("description", String, func() {
		Description("Description of the instalation of course")
		Example("This course is in Sistem Ingenier School")
		MinLength(1)
		MaxLength(512)
	})

	Required("name", "description")
})

var InstalatitonCourse = Type("InstalatitonCourse", func() {
	Description("Represents a course in the installation process")

	Extend(InstalatitonCoursePayload)

	Attribute("id", Int, func() {
		Description("Unique identifier for the installation of the course")
		Example(144)
	})

	Required("id")
})
