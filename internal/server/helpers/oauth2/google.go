package oauth2

import (
	"context"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	googleoauth2 "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

func GetGoogleUserInfo(token *oauth2.Token, cfg *oauth2.Config) (*googleoauth2.Userinfo, error) {
	client := cfg.Client(context.Background(), token)
	service, err := googleoauth2.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create new google OAuth2 service")
	}
	return service.Userinfo.V2.Me.Get().Do()
}
