package problems

import (
	"fmt"
	"github.com/spf13/cast"
	"net/http"

	"github.com/google/jsonapi"
)

func NotFound() *jsonapi.ErrorObject {
	return &jsonapi.ErrorObject{
		Title:  http.StatusText(http.StatusNotFound),
		Code:   cast.ToString(http.StatusNotFound),
		Status: fmt.Sprintf("%d", http.StatusNotFound),
	}
}
