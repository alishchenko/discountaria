package responses

type OAuth2Response struct {
	Url string `json:"url"`
}
type OAuth2CallbackResponse struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func ComposeOAuth2(url string) OAuth2Response {
	return OAuth2Response{
		Url: url,
	}
}
