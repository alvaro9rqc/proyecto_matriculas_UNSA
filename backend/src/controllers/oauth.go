package controllers

import (
	"context"
	"fmt"
	"os"

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
func NewOauth(oathConfig oauth2.Config) oauth.Service {
	return &oauthsrvc{
		GoogleOAuthConfig: oathConfig,
	}
}

// Generate a redirection URL for the chosen OAuth provider
func (s *oauthsrvc) Redirect(ctx context.Context, p *oauth.RedirectPayload) (res *oauth.OAuthRedirectResult, err error) {
	//===============================//
	// Mira asi puedes importar la configuración de OAuth
	// s.GoogleOAuthConfig
	// ================================//

	var (
		GOOGLE_REDIRECT_URL = s.GoogleOAuthConfig.RedirectURL
	)

	log.Printf(ctx, "oauth.redirect")
	var url string
	switch p.Provider {
	case "google":
		url = GOOGLE_REDIRECT_URL
	case "microsoft":
		url = os.Getenv("MICROSOFT_REDIRECT_URL")
	default:
		return nil, oauth.MakeInvalidProvider(fmt.Errorf("unsupported provider: %s", p.Provider))
	}

	return &oauth.OAuthRedirectResult{
		RedirectURL: url,
	}, nil
}

// Handle OAuth callback and authenticate user
func (s *oauthsrvc) Callback(ctx context.Context, p *oauth.CallbackPayload) (res *oauth.LoginResult, err error) {
	log.Printf(ctx, "oauth.callback: provider=%s, code=%s", p.Provider, p.Code)
	res = &oauth.LoginResult{}
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
