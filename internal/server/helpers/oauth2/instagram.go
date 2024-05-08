package oauth2

import (
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"net/http"
)

func GetInstagramUserInfo(token *oauth2.Token) (*User, error) {
	var user User
	resp, err := http.Get(`https://graph.instagram.com/v18.0/me?fields=id,name,email&access_token=` + token.AccessToken)

	if err != nil {
		return nil, errors.Wrap(err, "Error occurred while getting information from Instagram")
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&user)

	if err != nil {
		return nil, errors.Wrap(err, "Error occurred while getting information from Instagram")
	}
	spew.Dump(user)
	return &user, nil
}
