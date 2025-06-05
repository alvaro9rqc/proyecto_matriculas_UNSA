package api

import (
	. "goa.design/goa/v3/dsl"
)

var AccountUser = Type("AccountUser", func() {
	Description("AccountUser type")

	Attribute("id", String, "Unique user ID")
	Attribute("email", String, "User email")
	Attribute("firstName", String, "User first name")
	Attribute("lastNames", String, "User last names")
	Attribute("remainingNames", String, "User remaining names")

	Required("id", "email", "firstName", "lastNames", "remainingNames")
})
