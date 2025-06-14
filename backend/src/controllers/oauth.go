package controllers

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"

	oauth "github.com/enrollment/gen/oauth"
	"goa.design/clue/log"
	"golang.org/x/oauth2"
)

// oauth service example implementation.
// The example methods log the requests and return zero values.
type oauthsrvc struct {
	GoogleOAuthConfig oauth2.Config
}

// NewOauth returns the oauth service implementation.
func NewOauth(oauthConfig oauth2.Config) oauth.Service {
	return &oauthsrvc{
		GoogleOAuthConfig: oauthConfig,
	}
}

// Generate a redirection URL for the chosen OAuth provider
func (s *oauthsrvc) Login(ctx context.Context, p *oauth.LoginPayload) (res *oauth.OAuthRedirectResult, err error) {
	log.Printf(ctx, "oauth.redirect")

	//generate random state to prevent CSRF attacks
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return nil, fmt.Errorf("failed to generate random state: %w", err)
	}
	state := base64.URLEncoding.EncodeToString(b)

	//choose the redirect URL based on the provider
	var url string
	switch p.Provider {
	case "google":
		url = s.GoogleOAuthConfig.AuthCodeURL(state)
	case "microsoft":
		//url = os.Getenv("MICROSOFT_REDIRECT_URL")
		return nil, oauth.MakeInvalidProvider(fmt.Errorf("unsupported provider: %s", p.Provider))
	default:
		return nil, oauth.MakeInvalidProvider(fmt.Errorf("unsupported provider: %s", p.Provider))
	}
	res = &oauth.OAuthRedirectResult{
		RedirectURL: url,
	}
	return res, nil
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

// Returns the authenticated user's information
func (s *oauthsrvc) Me(ctx context.Context) (res *oauth.AccountUser, err error) {
	res = &oauth.AccountUser{}
	log.Printf(ctx, "oauth.me")
	return
}
