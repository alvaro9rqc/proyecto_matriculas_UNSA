package controllers

import (
	"context"

	auth "github.com/enrollment/gen/auth"
	"goa.design/clue/log"
)

// Auth service example implementation.
// The example methods log the requests and return zero values.
type authsrvc struct{}

// NewAuth returns the Auth service implementation.
func NewAuth() auth.Service {
	return &authsrvc{}
}

// Returns information about the currently authenticated user.
func (s *authsrvc) Me(ctx context.Context) (res *auth.AccountUser, err error) {
	res = &auth.AccountUser{}
	log.Printf(ctx, "auth.Me")
	return
}

// Initiates the Google OAuth 2.0 login flow by redirecting to the Google
// authorization endpoint.
func (s *authsrvc) GoogleLogin(ctx context.Context) (err error) {
	log.Printf(ctx, "auth.Google login")
	return
}

// Logs the user out by clearing the session and redirecting to the appropriate
// post-logout URL.
func (s *authsrvc) GoogleLogout(ctx context.Context) (err error) {
	log.Printf(ctx, "auth.Google logout")
	return
}

// Google OAuth callback. Creates or reuses a session and redirects to the
// frontend.
func (s *authsrvc) GoogleCallback(ctx context.Context) (err error) {
	log.Printf(ctx, "auth.GoogleCallback")
	return
}
