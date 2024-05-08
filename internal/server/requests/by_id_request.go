package requests

import (
	"github.com/go-chi/chi/v5"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

type ByIdRequest struct {
	Id int64 `json:"age,omitempty"`
}

func NewByIdRequest(r *http.Request) (int64, error) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return 0, errors.Wrap(err, "failed to get id")
	}
	return int64(id), nil
}
