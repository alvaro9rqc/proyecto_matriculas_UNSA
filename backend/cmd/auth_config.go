package main

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"os"
)

var googleOAuth = &oauth2.Config{
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
	RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
	Scopes:       []string{"email", "profile"},
	Endpoint:     google.Endpoint,
}
