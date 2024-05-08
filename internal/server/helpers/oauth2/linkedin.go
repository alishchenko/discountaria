package oauth2

import (
	"encoding/json"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"net/http"
)

type LinkedinUser struct {
	Sub   string `json:"sub"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetLinkedinUserInfo(token *oauth2.Token) (*User, error) {
	var user LinkedinUser

	request, err := http.NewRequest("GET", "https://api.linkedin.com/v2/userinfo", nil)

	if err != nil {
		return nil, errors.Wrap(err, "Error occurred while creating request for Linkedin")
	}
	request.Header.Add("Authorization", "Bearer "+token.AccessToken)
	client := &http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, "Error occurred while getting information from Linkedin")
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&user)

	if err != nil {
		return nil, errors.Wrap(err, "Error occurred while decoding information from Linkedin")
	}
	return &User{
		Id:    user.Sub,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
