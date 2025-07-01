package api

import (
	. "goa.design/goa/v3/dsl"
)

var InstitutionResult = Type("Institution", func() {
	Description("Institution represents an educational institution")

	Attribute("id", Int32, "Unique identifier for the institution", func() {
		Example(1)
	})
	Attribute("name", String, "Name of the institution", func() {
		Example("University of Example")
	})
	Attribute("logoUrl", String, "URL to the institution's logo", func() {
		Format(FormatURI)
		Example("https://example.edu/logo.png")
	})

	Required("id", "name")
})

var ProcessResult = Type("Process", func() {
	Description("Process represents an enrollment process for an institution")

	Attribute("id", Int32, "Unique identifier for the process", func() {
		Example(1)
	})
	Attribute("name", String, "Name of the process", func() {
		Example("Fall Semester Enrollment")
	})
	Attribute("startDay", String, "Start date of the process in YYYY-MM-DD format", func() {
		Example("2023-08-01")
	})
	Attribute("endDay", String, "End date of the process in YYYY-MM-DD format", func() {
		Example("2023-12-15")
	})
	Attribute("institutionId", Int32, "ID of the institution this process belongs to", func() {
		Example(1)
	})
})

var _ = Service("institution", func() {
	Description("Service for managing educational institutions")

	Error("not_authorized", ErrorResult, "User is not authorized to access this resource")

	Error("internal_server_error", ErrorResult, "An internal server error occurred")

	// Institutions service provides endpoints for managing educational institutions.
	Method("ListInstitutions", func() {
		Description("List all educational institutions avaliable for the user")

		Result(ArrayOf(InstitutionResult), "List of institutions available to the user")

		HTTP(func() {
			GET("/institutions")
			Response(StatusOK)
			Response("not_authorized", StatusForbidden)
		})
	})

	Method("ListProccesByInstitution", func() {
		Description("List all processes available for a specific institution")

		Payload(func() {
			Attribute("institutionId", Int32, "ID of the institution to list processes for", func() {
				Example(1)
			})
			Required("institutionId")
		})

		Result(ArrayOf(ProcessResult), "List of processes available for the specified institution")

		HTTP(func() {
			GET("/institutions/{institutionId}/processes")
			Param("institutionId", Int32, "ID of the institution")
			Response(StatusOK)
			Response("not_authorized", StatusForbidden)
			Response("internal_server_error", StatusInternalServerError)
		})
	})

})
