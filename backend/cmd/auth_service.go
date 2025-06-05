package main

import (
	"context"
	"time"
	"log"
	auth "github.com/enrollment/gen/auth"
	"golang.org/x/oauth2"
	"goa.design/goa/v3/pkg"
)

type AuthService struct{}

func (s *AuthService) Redirect(ctx context.Context, p *auth.RedirectPayload) (*auth.OAuthRedirectResult, error) {
	var url string
	state := "test-state" // you should generate/store securely

	switch p.Provider {
	case "google":
		url = GoogleOAuth.AuthCodeURL(state)
	//case "microsoft":
	//	url = MicrosoftOAuth.AuthCodeURL(state)
	default:
		return nil, goa.PermanentError("invalid_provider", "Unsupported provider")
	}

	return &auth.OAuthRedirectResult{RedirectURL: url}, nil
}

func (s *AuthService) Callback(ctx context.Context, p *auth.CallbackPayload) (*auth.LoginResult, error) {
	var conf *oauth2.Config

	switch p.Provider {
	case "google":
		conf = GoogleOAuth
	//case "microsoft":
	//	conf = MicrosoftOAuth
	default:
		return nil, goa.PermanentError("invalid_provider", "Unsupported provider")
	}

	token, err := conf.Exchange(ctx, p.Code)
	if err != nil {
		log.Println("Token exchange failed:", err)
		return nil, goa.TemporaryError("invalid_token", "Failed to exchange code")
	}

	return &auth.LoginResult{
		AccessToken: token.AccessToken,
		ExpiresAt:   token.Expiry.Format(time.RFC3339),
	}, nil
}

func (s *AuthService) Logout(ctx context.Context, p *auth.LogoutPayload) error {
	log.Println("User logged out:", p.Token)
	return nil
}
