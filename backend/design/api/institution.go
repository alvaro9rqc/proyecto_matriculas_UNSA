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

})
