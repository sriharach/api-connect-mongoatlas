package internal

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func Oauth() *oauth2.Config {
	var googleOauthConfig = &oauth2.Config{
		RedirectURL:  os.Getenv("REDIRECT_URL"),
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint: google.Endpoint,
	}

	return googleOauthConfig
}
