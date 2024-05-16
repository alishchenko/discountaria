package problems

import (
	"fmt"
	"github.com/spf13/cast"
	"net/http"

	"github.com/google/jsonapi"
)

func InternalError() *jsonapi.ErrorObject {
	return &jsonapi.ErrorObject{
		Title:  http.StatusText(http.StatusInternalServerError),
		Code:   cast.ToString(http.StatusInternalServerError),
		Status: fmt.Sprintf("%d", http.StatusInternalServerError),
	}
}
