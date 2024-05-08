package requests

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
	"time"
)

type CreateOfferRequest struct {
	CompanyId int64
	Sale      int64      `json:"sale"`
	Users     []string   `json:"users"`
	ExpiredAt *time.Time `json:"expired_at"`
}

func NewCreateOfferRequest(r *http.Request) (*CreateOfferRequest, error) {
	var req CreateOfferRequest

	err := render.DecodeJSON(r.Body, &req)
	if err != nil {
		return nil, fmt.Errorf("failed to decode json body: %w", err)
	}
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user id")
	}
	req.CompanyId = int64(id)
	return &req, req.validate()
}

func (r *CreateOfferRequest) validate() error {
	return validation.Errors{
		"company_id/": validation.Validate(
			&r.CompanyId,
			validation.Required,
		),
		"sale/": validation.Validate(
			&r.Sale,
			validation.Required,
		),
	}.Filter()
}
