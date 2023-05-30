package services

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	oAuthGoogleCfg = oauth2.Config{
		Endpoint:    google.Endpoint,
		Scopes:      []string{"https://www.googleapis.com/auth/userinfo.email"},
		RedirectURL: "http://localhost:8080/callback",
	}
	oAuthGoogleState string
)
