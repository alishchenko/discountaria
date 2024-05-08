package requests

import (
	"github.com/go-chi/chi/v5"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

type UseOfferRequest struct {
	OfferId int64
	UserId  int64
}

func NewUseOfferRequest(r *http.Request) (*UseOfferRequest, error) {
	var req UseOfferRequest
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user id")
	}
	req.UserId = int64(id)
	id, err = strconv.Atoi(chi.URLParam(r, "offer_id"))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get offer id")
	}
	req.OfferId = int64(id)
	return &req, req.validate()
}

func (r *UseOfferRequest) validate() error {
	return validation.Errors{
		"offer_id/": validation.Validate(
			&r.OfferId,
			validation.Required,
		),
		"user_id/": validation.Validate(
			&r.UserId,
			validation.Required,
		),
	}.Filter()
}
