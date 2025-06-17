package api

import (
	. "goa.design/goa/v3/dsl"
)

var OAuthProviderType = Type("OAuthProvider", String, func() {
	Description("OAuth provider options")
	Enum("google", "microsoft") // Or use variables if preferred
	Example("google")
})

// Result type containing the URL to redirect the user to start OAuth login
var OAuthRedirectResult = Type("OAuthRedirectResult", func() {
	Description("Redirect URL for initiating OAuth login")
	Attribute("Location", String, "OAuth authorization URL", func() {
		Format(FormatURI)
		Example("https://accounts.google.com/o/oauth2/auth?...code")
	})
	Required("Location")
})

// Result type after successful OAuth login
var LoginResult = Type("LoginResult", func() {
	Description("Successful login result containing access token")
	Attribute("access_token", String, "Session access token")
	Attribute("expires_at", String, "Access token expiration timestamp", func() {
		Format(FormatDateTime)
	})
	Attribute("session_token", String, "Cookie for session management")
	Required("access_token", "expires_at")
})

var _ = Service("oauth", func() {
	Description("OAuth-based authentication service for Google and Microsoft")

	// Initiates the login by generating a provider-specific OAuth authorization URL
	Method("login", func() {
		Description("Generate a redirection URL for the chosen OAuth provider")

		Payload(func() {
			Attribute("provider", OAuthProviderType, "OAuth provider name")
			Required("provider")
		})

		Result(OAuthRedirectResult)

		Error("invalid_provider", ErrorResult, "Unsupported OAuth provider")

		HTTP(func() {
			GET("/auth/{provider}/login")
			Response(StatusTemporaryRedirect, func () {
				Header("Location")
			})
		})
	})

	// OAuth callback handler, exchanges authorization code for access token, logs user in
	Method("callback", func() {
		Description("Handle OAuth callback and authenticate user")

		Payload(func() {
			Attribute("provider", OAuthProviderType, "OAuth provider name")
			Attribute("code", String, "Authorization code", func() {
				MinLength(1)
			})
			Attribute("state", String, "Anti-CSRF state token", func() {
				MinLength(10)
			})
			Attribute("ip_address", String, "Client IP address", func() {
				Format(FormatIP)
			})
			Attribute("user_agent", String, "User-Agent header value")
			Required("provider", "code", "state", "ip_address", "user_agent")
		})

		Result(LoginResult)

		Error("invalid_token", ErrorResult, "Invalid or expired OAuth token")
		Error("server_error", ErrorResult, "Internal server error")

		HTTP(func() {
			GET("/auth/{provider}/callback")
			Param("code")
			Param("state")
			Param("ip_address")
			Param("user_agent")
			Response(StatusOK, func() {
				Cookie("session_token:session_token", String, func() {
					Description("Session token set in cookie after successful login")
					Example("session_token=abc123xyz")
				})
				CookieHTTPOnly()
			})
			Response("invalid_token", StatusBadRequest)
			Response("server_error", StatusInternalServerError)
		})
	})

	// Logout endpoint to invalidate a session
	Method("logout", func() {
		Description("Terminate the current session and invalidate the token")

		Payload(func() {
			Attribute("token", String, "Session token to invalidate")
			Required("token")
		})

		Error("unauthorized", ErrorResult, "Missing or invalid token")

		HTTP(func() {
			POST("/auth/logout")
			Header("token:Authorization")
			Response(StatusNoContent)
			Response("unauthorized", StatusUnauthorized)
		})
	})

	Method("me", func() {
		Description("Returns the authenticated user's information")

		Result(AccountUser)

		Error("unauthorized", ErrorResult, "Unauthorized access")

		HTTP(func() {
			GET("/auth/me")
			Response(StatusOK)
			Response("unauthorized", StatusUnauthorized)
		})
	})
})
