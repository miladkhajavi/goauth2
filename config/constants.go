package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func SetupConfig() *oauth2.Config {
	conf := &oauth2.Config{
		ClientID:     "632016499651-hf3ipco92fun3r8gkame0oe931c1vh3c.apps.googleusercontent.com",
		ClientSecret: "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",
		RedirectURL:  "http://localhost:3030/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
			"openid",
		},
		Endpoint: google.Endpoint,
	}
	return conf
}
