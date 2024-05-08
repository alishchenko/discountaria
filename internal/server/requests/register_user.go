package requests

import (
	"fmt"
	"github.com/go-chi/render"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"net/http"
)

type RegisterUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewRegisterUserRequest(r *http.Request) (*RegisterUserRequest, error) {
	var req RegisterUserRequest

	err := render.DecodeJSON(r.Body, &req)
	if err != nil {
		return nil, fmt.Errorf("failed to decode json body: %w", err)
	}
	return &req, req.validate()
}

func (r *RegisterUserRequest) validate() error {
	return validation.Errors{
		"name/": validation.Validate(
			&r.Name,
			validation.Required,
		),
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
