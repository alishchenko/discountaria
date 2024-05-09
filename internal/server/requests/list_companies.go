package requests

import (
	"net/http"
)

type ListCompaniesRequest struct {
	PaginationParams
	Name string
}

func NewListCompaniesRequest(r *http.Request) ListCompaniesRequest {
	var request ListCompaniesRequest
	request.PaginationParams = GetPaginationParams(r)
	request.Name = r.URL.Query().Get("[filter]name")
	return request
}
