package responses

type AuthTokensResponse struct {
	Id           int64  `json:"id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
