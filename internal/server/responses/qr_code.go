package responses

import "github.com/alishchenko/discountaria/internal/server/helpers"

type QRCodeResponse struct {
	Data      helpers.SignatureData `json:"data"`
	Signature string                `json:"signature"`
}
