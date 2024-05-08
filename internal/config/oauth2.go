package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/linkedin"
	goauth2 "google.golang.org/api/oauth2/v2"
	"time"
)

type OAuth2Config struct {
	RedirectURL  string `yaml:"redirect_url" env-required:"true"`
	ClientID     string `yaml:"client_id" env-required:"true"`
	ClientSecret string `yaml:"client_secret" env-required:"true"`
}

func (o *OAuth2Config) ToFacebookOauth2() *oauth2.Config {
	return &oauth2.Config{
		RedirectURL:  o.RedirectURL,
		ClientID:     o.ClientID,
		ClientSecret: o.ClientSecret,
		Scopes: []string{
			"email",
		},
		Endpoint: facebook.Endpoint,
	}

}

func (o *OAuth2Config) ToGoogleOauth2() *oauth2.Config {
	return &oauth2.Config{
		RedirectURL:  o.RedirectURL,
		ClientID:     o.ClientID,
		ClientSecret: o.ClientSecret,
		Scopes: []string{
			goauth2.UserinfoEmailScope,
			goauth2.UserinfoProfileScope,
		},
		Endpoint: google.Endpoint,
	}
}

func (o *OAuth2Config) ToLinkedinOauth2() *oauth2.Config {
	return &oauth2.Config{
		RedirectURL:  o.RedirectURL,
		ClientID:     o.ClientID,
		ClientSecret: o.ClientSecret,
		Scopes: []string{
			"openid", "profile", "email",
		},
		Endpoint: linkedin.Endpoint,
	}
}

type OAuth2StateConfig struct {
	StateSecret string        `yaml:"state_secret" env-required:"true"`
	StateLife   time.Duration `yaml:"state_life" env-required:"true"`
}
