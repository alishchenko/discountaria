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

type UpdateCompanyRequest struct {
	Id      int64
	Name    *string `json:"name"`
	LogoUrl *string `json:"logo_url"`
}

func NewUpdateCompanyRequest(r *http.Request) (*UpdateCompanyRequest, error) {
	var req UpdateCompanyRequest

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

func (r *UpdateCompanyRequest) validate() error {
	return validation.Errors{
		"id/": validation.Validate(
			&r.Id,
			validation.Required,
		),
	}.Filter()
}
