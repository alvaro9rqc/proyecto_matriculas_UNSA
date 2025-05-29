package design

import (
	. "goa.design/goa/v3/dsl"
)

// Oauth provider recognized by the system. eg: Google, GitHub, etc.
var _ = Type("OauthProvider", func () {
	Description("registered provider in th system")
	Attribute("id", Int, "unique identifier", func() {
		Minimum(1)
		Example("1")
	})
	Attribute("name", String, "Intern provider name", func() {
		MinLength(2)
		Example("google")
	})
	Attribute("auth_url", String, "Provider authentication URL", func() {
		Format(FormatURI)
		Example("https://accounts.google.com/o/oauth2/auth")
	})
	Required("id", "name", "auth_url")
})
