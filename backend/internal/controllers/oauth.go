package controllers

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"

	"github.com/enrollment/config"
	"github.com/enrollment/gen/db"
	oauth "github.com/enrollment/gen/oauth"
	"github.com/enrollment/internal/ports"

	"goa.design/clue/log"
	"golang.org/x/oauth2"
	googleOauth2 "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

// oauth service example implementation.
// The example methods log the requests and return zero values.
type oauthsrvc struct {
	GoogleOAuthConfig *oauth2.Config
	OauthRep          ports.OauthRepositoryInterface
	FrontendURL       string
}

// NewOauth returns the oauth service implementation.
func NewOauth(cfg *config.MainConfig, oauthRep ports.OauthRepositoryInterface) oauth.Service {
	return &oauthsrvc{
		GoogleOAuthConfig: &cfg.GoogleOAuthConfig,
		OauthRep:          oauthRep,
		FrontendURL:       cfg.FrontendURL,
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
		Location: url,
	}
	return res, nil
}

func exchangesCode(ctx context.Context, p *oauth.CallbackPayload, s *oauthsrvc) (*googleOauth2.Userinfo, error) {
	token, err := s.GoogleOAuthConfig.Exchange(ctx, p.Code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}
	oauth2Service, err := googleOauth2.NewService(ctx, option.WithTokenSource(s.GoogleOAuthConfig.TokenSource(ctx, token)))
	if err != nil {
		return nil, fmt.Errorf("failed on creating oauth2 service: %w\n", err)
	}
	userinfo, err := oauth2Service.Userinfo.Get().Do()
	if err != nil {
		return nil, fmt.Errorf("failed retrieving user info %err\n", err)
	}
	return userinfo, nil
}

func generateSessionToken() (string, error) {
	// 32 bytes of cryptographically secure random data (~256 bits of entropy)
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("failed to read random bytes: %w", err)
	}

	// Directly encode the raw random bytes
	return base64.URLEncoding.EncodeToString(b), nil
}

// Handle OAuth callback and authenticate user
// Here we search for the user in the database, create a session, and return the access token
// or create a new user if not found
// TODO: use cookies to prevent CSRF attacks
// TODO: search a way to prevent unlimited login attempts ans unlimited sessions created
// TODO: erase access token for original result type

func createAccountSession(s *oauthsrvc, ctx *context.Context, p *oauth.CallbackPayload, account *db.Account) (res string, err error) {

	token, err := generateSessionToken()
	if err != nil {
		return res, fmt.Errorf("failed to generate session token: %w", err)
	}
	//get user agent
	var userAgent string
	if p.UserAgent != nil {
		userAgent = *p.UserAgent
	} else {
		userAgent = "unknown"
	}
	//get ip address
	var ipAddress string
	if p.IPAddress != nil {
		ipAddress = *p.IPAddress
	} else {
		ipAddress = "unknown"
	}
	//get timestamp for expiration date
	expirationDate := pgtype.Timestamptz{
		Time:  time.Now().Add(24 * time.Hour),
		Valid: true,
	}

	s.OauthRep.CreateAccountSession(*ctx, db.CreateAccountSessionParams{
		Token:          token,
		UserAgent:      userAgent,
		IpAddress:      ipAddress,
		ExpirationDate: expirationDate,
		AccountID:      account.ID,
	})
	return token, nil
}
func (s *oauthsrvc) Callback(ctx context.Context, p *oauth.CallbackPayload) (res *oauth.LoginResult, err error) {
	// returns the user info from the google's oauth service
	userinfo, err := exchangesCode(ctx, p, s)
	if err != nil {
		return nil, oauth.MakeServerError(fmt.Errorf("failed to exchange code: %w", err))
	}
	// search if account exists in the database
	account, err := s.OauthRep.GetAccountByEmail(ctx, userinfo.Email)

	if err != nil {
		return nil, oauth.MakeUnauthorized(fmt.Errorf("failed to get account by email: %w", err))
	}
	token, err := createAccountSession(s, &ctx, p, &account)
	if err != nil {
		return nil, oauth.MakeServerError(fmt.Errorf("failed to create account session: %w", err))
	}
	// put all data in the user.
	// create a new session
	res = &oauth.LoginResult{}
	//res.AccessToken = userinfo.Email
	//res.SessionToken = &userinfo.Email
	//res.ExpiresAt = "2025-06-12"
	res.SessionToken = token
	url := s.FrontendURL + "/dashboard"
	res.Location = &url
	log.Printf(ctx, "oauth.callback")
	return
}

// Terminate the current session and invalidate the token
// 1. Retrieeve the session token from the requests
// 2. Erase the session token from the database
// 3. Erase the session token from the cookies
// 4. Redirect the user to the frontend URL
func (s *oauthsrvc) Logout(ctx context.Context, p *oauth.LogoutPayload) (res *oauth.LogoutResult, err error) {
	log.Printf(ctx, "oauth.logout")
	// retrieve the session token from the payload
	token := p.SessionToken
	//erase the session token from the database
	err = s.OauthRep.DeleteAccountByToken(ctx, token)
	if err != nil {
		return nil, oauth.MakeUnauthorized(fmt.Errorf("failed to delete account by token: %w", err))
	}
	// erase the session token from the cookies (in design too)
	p.SessionToken = ""
	res = &oauth.LogoutResult{}
	res.SessionToken = ""
	res.Location = &s.FrontendURL
	return res, nil
}

// Returns the authenticated user's information
func (s *oauthsrvc) Me(ctx context.Context) (res *oauth.AccountUser, err error) {
	res = &oauth.AccountUser{}
	log.Printf(ctx, "oauth.me")
	return
}
