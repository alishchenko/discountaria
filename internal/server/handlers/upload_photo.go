package handlers

import (
	"github.com/alishchenko/discountaria/internal/server/helpers"
	"github.com/alishchenko/discountaria/internal/server/requests"
	"github.com/alishchenko/discountaria/internal/server/responses"
	"github.com/alishchenko/discountaria/internal/server/responses/problems"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"net/http"
)

func UploadPhoto(w http.ResponseWriter, r *http.Request) {
	key, document, header, err := requests.NewUploadPhotoRequest(r)
	if err != nil {
		render.JSON(w, r, problems.BadRequest(err))
		return
	}

	ext, err := helpers.CheckDocumentMimeType(header.Header.Get("Content-Type"), r)
	if err != nil {
		render.JSON(w, r, problems.BadRequest(err))
		return
	}

	awsConfig := helpers.AwsConfig(r)

	if key == "" {
		key = uuid.New().String()
	} else {
		// checking if key exists (only in case of custom keys, not uuid-generated)
		// to not overwrite the existing document

		exists, err := helpers.IsKeyExists(key+"."+ext, awsConfig)
		if err != nil || exists {
			helpers.Log(r).Error(errors.Wrap(err, "failed to check key existence or key was found").Error())
			render.JSON(w, r, problems.BadRequest(
				errors.New("Document with such key already exists or it cannot be checked")))
			return
		}
	}
	key += "." + ext

	err = helpers.UploadFile(document, key, awsConfig)
	if err != nil {
		helpers.Log(r).Error(errors.Wrap(err, "failed to upload file").Error())
		render.JSON(w, r, problems.InternalError())
		return
	}

	render.JSON(w, r, responses.IdResponseString{Id: key})
}
