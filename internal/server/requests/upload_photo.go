package requests

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"
	"mime/multipart"
	"net/http"
	"regexp"
)

const DocumentKey = "Document"
const FilenameKey = "Key"

var keyRegexp = regexp.MustCompile("^[^<>:;(),.?\"*|/]+$")

func NewUploadPhotoRequest(r *http.Request) (string, multipart.File, *multipart.FileHeader, error) {
	err := r.ParseMultipartForm(1 << 32)
	if err != nil {
		return "", nil, nil, errors.Wrap(err, "failed to parse document")
	}

	key := r.FormValue(FilenameKey)
	if key != "" {

		err = validation.Errors{
			"key": validation.Validate(key, validation.Required, validation.Match(keyRegexp)),
		}.Filter()

		if err != nil {
			return "", nil, nil, errors.Wrap(err, "failed to parse key")
		}
	}

	f, h, err := r.FormFile(DocumentKey)
	return key, f, h, err
}
