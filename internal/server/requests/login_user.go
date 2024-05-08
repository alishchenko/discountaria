package requests

import (
	"fmt"
	"github.com/go-chi/render"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"net/http"
)

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewLoginUserRequest(r *http.Request) (*LoginUserRequest, error) {
	var req LoginUserRequest

	err := render.DecodeJSON(r.Body, &req)
	if err != nil {
		return nil, fmt.Errorf("failed to decode json body: %w", err)
	}
	return &req, req.validate()
}

func (r *LoginUserRequest) validate() error {
	return validation.Errors{
		"data/email": validation.Validate(
			&r.Email,
			validation.Required,
		),
		"data/password": validation.Validate(
			&r.Password,
			validation.Required,
		),
	}.Filter()
}
