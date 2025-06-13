package controllers

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"

	oauth "github.com/enrollment/gen/oauth"
	goahttp "goa.design/goa/v3/http"
	"goa.design/clue/log"
	"golang.org/x/oauth2"
	"net/http"
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

	var (
		GOOGLE_REDIRECT_URL = s.GoogleOAuthConfig.RedirectURL
	)

	log.Printf(ctx, "oauth.redirect")
	//choose the redirect URL based on the provider
	var url string
	switch p.Provider {
	case "google":
		url = GOOGLE_REDIRECT_URL
	case "microsoft":
		//url = os.Getenv("MICROSOFT_REDIRECT_URL")
	default:
		return nil, oauth.MakeInvalidProvider(fmt.Errorf("unsupported provider: %s", p.Provider))
	}

	//generate random state to prevent CSRF attacks
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return nil, fmt.Errorf("failed to generate random state: %w", err)
	}
	state := base64.URLEncoding.EncodeToString(b)
	// add cokie http 
	httpRes = goahttp.
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
