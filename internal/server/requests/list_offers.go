package requests

import (
	"net/http"
)

type ListOffersRequest struct {
	PaginationParams
	CompanyName string
}

func NewListOffersRequest(r *http.Request) ListOffersRequest {
	var request ListOffersRequest
	request.PaginationParams = GetPaginationParams(r)
	request.CompanyName = r.URL.Query().Get("filter[company_name]")
	return request
}
