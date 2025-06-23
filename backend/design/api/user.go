package api

import (
	. "goa.design/goa/v3/dsl"
)

var AccountUser = Type("AccountUser", func() {
	Description("AccountUser type")

	Attribute("id", Int, "Unique user ID")
	Attribute("email", String, "User email")
	Attribute("name", String, "User first name")
	Attribute("surname", String, "User last names")
	Attribute("avatar_url", String, "User remaining names")

	Required("id", "email", "name", "surname", "avatar_url")
})
