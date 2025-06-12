package controllers

import (
	"context"
	"fmt"

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
	log.Printf(ctx, "oauth.redirect")
	var url string
	switch p.Provider {
	case "google":
		url = "https://accounts.google.com/o/oauth2/auth?client_id=YOUR_CLIENT_ID&redirect_uri=YOUR_REDIRECT_URI&response_type=code&scope=email%20profile"
	case "microsoft":
		url = "https://login.microsoftonline.com/common/oauth2/v2.0/authorize?client_id=YOUR_CLIENT_ID&response_type=code&redirect_uri=YOUR_REDIRECT_URI&scope=openid%20profile%20email"
	default:
		return nil, oauth.MakeInvalidProvider(fmt.Errorf("unsupported provider: %s", p.Provider))
	}

	return &oauth.OAuthRedirectResult{
		RedirectURL: url,
	}, nil
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

func (s *oauthsrvc) Me(ctx context.Context) (res *oauth.AccountUser, err error) {
	res = &oauth.AccountUser{}

	return res, nil
}
