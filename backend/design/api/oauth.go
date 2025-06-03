package api

import (
	. "goa.design/goa/v3/dsl"
)

// Oauth provider recognized by the system. eg: Google, GitHub, etc.
var OAuthProvider = Type("OAuthProvider", func() {
	Description("registered provider in th system")
	Attribute("id", Int, "unique identifier", func() {
		Minimum(1)
		Example(1)
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

var UserOauthInfo = Type("UserOauthInfo", func() {
	Description("User information via Oauth")
	Attribute("oauth_provider_id", Int, "ID of the provider used", func() {
		Minimum(1)
		Example(1)
	})
	Attribute("provider_user_id", String, "Access token given by the provider", func() {
		MinLength(10)
		Example("ya29.a0AfH6SM...")
	})
	Attribute("profile_picture", String, "Profile picture URL", func() {
		Format(FormatURI)
		Example("https://avatars.githubusercontent.com/u/123456?v=4")
	})
})

var _ = Service("oauth", func() {
	Description("Oauth Authentication service")
	// list oauth providers
	Method("list_providers", func() {
		Description("List all available oauth providers")
		Result(ArrayOf(OAuthProvider))
	})
	// login
	Method("login", func() {
		Description("Login using OAuth with a specific provider")
		Payload(func() {
			Attribute("oauth_provider_id", Int, "ID of the OAuth provider", func() {
				Minimum(1)
			})
			Attribute("code", String, "OAuth code given by the provider", func() {
				MinLength(10)
			})
			Required("oauth_provider_id", "code")
		})
		Result(UserOauthInfo)
		Error("unauthorized", ErrorResult, "Invalid or expired code")
		Error("not_found", ErrorResult, "Not registered provider")
		HTTP(func() {
			POST("/oauth/login")
			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
			Response("not_found", StatusNotFound)
		})
	})
})
