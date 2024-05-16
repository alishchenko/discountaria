package handlers

import (
	"github.com/alishchenko/discountaria/internal/server/helpers"
	"github.com/alishchenko/discountaria/internal/server/requests"
	"github.com/alishchenko/discountaria/internal/server/responses/problems"
	"github.com/go-chi/render"
	"net/http"
)

func DeleteFile(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewGetDocumentByKeyRequest(r) // As only key is required
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, problems.BadRequest(err))
		return
	}

	awsConfig := helpers.AwsConfig(r)

	// Make sure the key exists as the DeleteFile method will not render the KeyNotFound error
	exists, err := helpers.IsKeyExists(req.Key, awsConfig)
	if err != nil || !exists {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, problems.NotFound())
		return
	}

	err = helpers.DeleteFile(req.Key, awsConfig)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, problems.InternalError())
		return
	}

	render.Status(r, http.StatusInternalServerError)
	return
}
