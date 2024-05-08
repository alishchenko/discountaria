package requests

import (
	"encoding/json"
	"net/http"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"
)

type VerifyOAuth2Request struct {
	Code  string `json:"code"`
	State string `json:"state"`
}

func NewVerifyOAuth2Request(r *http.Request) (*VerifyOAuth2Request, error) {
	request := new(VerifyOAuth2Request)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return nil, ozzo.Errors{
			"/": errors.Wrap(err, "failed to decode request body"),
		}
	}
	return request, request.validate()
}

func (r *VerifyOAuth2Request) validate() error {
	return ozzo.Errors{
		"data/attributes/code":  ozzo.Validate(&r.Code, ozzo.Required),
		"data/attributes/state": ozzo.Validate(&r.State, ozzo.Required),
	}.Filter()
}
