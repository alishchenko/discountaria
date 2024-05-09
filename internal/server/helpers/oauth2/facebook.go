package oauth2

import (
	"encoding/json"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"net/http"
)

func GetFacebookUserInfo(token *oauth2.Token) (*User, error) {
	var user User
	resp, err := http.Get("https://graph.facebook.com/me?fields=id,name,email&access_token=" + token.AccessToken)

	if err != nil {
		return nil, errors.Wrap(err, "Error occurred while getting information from Facebook")
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&user)

	if err != nil {
		return nil, errors.Wrap(err, "Error occurred while getting information from Facebook")
	}
	return &user, nil
}
