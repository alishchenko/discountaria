package requests

import (
	"github.com/spf13/cast"
	"net/http"
)

type PaginationParams struct {
	Limit  uint64
	Number uint64
	Order  string
	Sort   string
}

func GetPaginationParams(r *http.Request) PaginationParams {
	var request PaginationParams

	request.Limit = cast.ToUint64(r.URL.Query().Get("[page]limit"))
	request.Number = cast.ToUint64(r.URL.Query().Get("[page]number"))
	request.Order = r.URL.Query().Get("[page]order")
	request.Sort = r.URL.Query().Get("sort")

	return request
}
