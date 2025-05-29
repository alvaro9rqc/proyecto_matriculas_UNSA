package api

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("Auth", func() {
	Description("Service to handle authentication and authorization")

	Method("Me", func() {
		Description("Returns information about the currently authenticated user.")
		Result(AccountUser)

		HTTP(func() {
			GET("/auth/me")

			Response(StatusOK)
			Response(StatusUnauthorized)
		})
	})

	Method("Google login", func() {
		Description("Initiates the Google OAuth 2.0 login flow by redirecting to the Google authorization endpoint.")

		HTTP(func() {
			GET("/auth/google/login")

			Response(StatusTemporaryRedirect)
		})
	})

	Method("Google logout", func() {
		Description("Logs the user out by clearing the session and redirecting to the appropriate post-logout URL.")

		HTTP(func() {
			GET("/auth/google/logout")

			Response(StatusTemporaryRedirect)
		})
	})

	Method("GoogleCallback", func() {
		Description("Google OAuth callback. Creates or reuses a session and redirects to the frontend.")

		HTTP(func() {
			GET("/google/callback")

			Response(StatusTemporaryRedirect)
			Response(StatusUnauthorized)
		})
	})
})
