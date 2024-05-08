package requests

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

type UpdateUserRequest struct {
	Id          int64
	Name        *string `json:"name" structs:"name"`
	PhotoURL    *string `json:"photo_url" structs:"photo_url"`
	PhoneNumber *string `json:"phone" structs:"phone"`
	Email       *string `json:"email" structs:"email"`
	Password    *string `json:"password" structs:"password"`
	OldPassword *string `json:"old_password"`
}

func NewUpdateUserRequest(r *http.Request) (*UpdateUserRequest, error) {
	var req UpdateUserRequest

	err := render.DecodeJSON(r.Body, &req)
	if err != nil {
		return nil, fmt.Errorf("failed to decode json body: %w", err)
	}
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user id")
	}
	req.Id = int64(id)
	return &req, req.validate()
}

func (r *UpdateUserRequest) validate() error {
	return validation.Errors{
		"id/": validation.Validate(
			&r.Id,
			validation.Required,
		),
	}.Filter()
}
