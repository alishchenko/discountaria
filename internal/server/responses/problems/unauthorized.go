package problems

import (
	"fmt"
	"github.com/google/jsonapi"
	"github.com/spf13/cast"
	"net/http"
)

func Unauthorized() *jsonapi.ErrorObject {
	return &jsonapi.ErrorObject{
		Title:  http.StatusText(http.StatusUnauthorized),
		Code:   cast.ToString(http.StatusUnauthorized),
		Status: fmt.Sprintf("%d", http.StatusUnauthorized),
	}
}
