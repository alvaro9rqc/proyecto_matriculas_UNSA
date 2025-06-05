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

// List all available oauth providers
func (s *oauthsrvc) ListProviders(ctx context.Context) (res []*oauth.OAuthProvider, err error) {
	log.Printf(ctx, "oauth.list_providers")
	return
}

// Login using OAuth with a specific provider
func (s *oauthsrvc) Login(ctx context.Context, p *oauth.LoginPayload) (res *oauth.UserOauthInfo, err error) {
	res = &oauth.UserOauthInfo{}
	log.Printf(ctx, "oauth.login")
	return
}
