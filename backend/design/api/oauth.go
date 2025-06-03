package api

import (
	. "goa.design/goa/v3/dsl"
)

const (
	OAuthProviderGoogle    = "google"
	OAuthProviderMicrosoft = "microsoft"
)

var OAuthProviderType = Type("OAuthProvider", func() {
	Description("Supported OAuth provider")
	Attribute("name", String, "OAuth provider name", func() {
		Enum(OAuthProviderGoogle, OAuthProviderMicrosoft)
	})
	Required("name")
})

// Result type containing the URL to redirect the user to start OAuth login
var OAuthRedirectResult = Type("OAuthRedirectResult", func() {
	Description("Redirect URL for initiating OAuth login")
	Attribute("redirect_url", String, "OAuth authorization URL", func() {
		Format(FormatURI)
		Example("https://accounts.google.com/o/oauth2/auth?...code")
	})
	Required("redirect_url")
})

// Result type after successful OAuth login
var LoginResult = Type("LoginResult", func() {
	Description("Successful login result containing access token")
	Attribute("access_token", String, "Session access token")
	Attribute("expires_at", String, "Access token expiration timestamp", func() {
		Format(FormatDateTime)
	})
	Required("access_token", "expires_at")
})

