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

type CreateCompanyRequest struct {
	UserId      int64
	Name        string  `json:"name"`
	Url         *string `json:"url"`
	Description *string `json:"description"`
	LogoUrl     *string `json:"logo_url"`
}

func NewCreateCompanyRequest(r *http.Request) (*CreateCompanyRequest, error) {
	var req CreateCompanyRequest

	err := render.DecodeJSON(r.Body, &req)
	if err != nil {
		return nil, fmt.Errorf("failed to decode json body: %w", err)
	}
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user id")
	}
	req.UserId = int64(id)
	return &req, req.validate()
}

func (r *CreateCompanyRequest) validate() error {
	return validation.Errors{
		"name/": validation.Validate(
			&r.Name,
			validation.Required,
		),
	}.Filter()
}
