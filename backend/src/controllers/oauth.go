package controllers

import (
	"context"

	oauth "github.com/enrollment/gen/oauth"
	"goa.design/clue/log"
)

// oauth service example implementation.
// The example methods log the requests and return zero values.
type oauthsrvc struct{}

// NewOauth returns the oauth service implementation.
func NewOauth() oauth.Service {
	return &oauthsrvc{}
}

// Generate a redirection URL for the chosen OAuth provider
func (s *oauthsrvc) Redirect(ctx context.Context, p *oauth.RedirectPayload) (res *oauth.OAuthRedirectResult, err error) {
	res = &oauth.OAuthRedirectResult{}
	log.Printf(ctx, "oauth.redirect")
	return
}

// Handle OAuth callback and authenticate user
func (s *oauthsrvc) Callback(ctx context.Context, p *oauth.CallbackPayload) (res *oauth.LoginResult, err error) {
	res = &oauth.LoginResult{}
	log.Printf(ctx, "oauth.callback")
	return
}

// Terminate the current session and invalidate the token
func (s *oauthsrvc) Logout(ctx context.Context, p *oauth.LogoutPayload) (err error) {
	log.Printf(ctx, "oauth.logout")
	return
}
