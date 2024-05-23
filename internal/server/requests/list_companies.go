package requests

import (
	"github.com/spf13/cast"
	"net/http"
)

type ListCompaniesRequest struct {
	PaginationParams
	Name    *string
	OwnerId *int64
}

func NewListCompaniesRequest(r *http.Request) ListCompaniesRequest {
	var request ListCompaniesRequest
	request.PaginationParams = GetPaginationParams(r)
	name := r.URL.Query().Get("filter[name]")
	if name != "" {
		request.Name = &name
	}
	id := r.URL.Query().Get("filter[owner_id]")
	if id != "" {
		idInt := cast.ToInt64(id)
		request.OwnerId = &idInt
	}

	return request
}
