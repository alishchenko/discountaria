package oauth2

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

const (
	FacebookOAuth2Provider string = "facebook"
	GoogleOAuth2Provider   string = "google"
	LinkedinOAuth2Provider string = "linkedin"
)
