package handlers

import (
	"github.com/alishchenko/discountaria/internal/server/helpers"
	"github.com/alishchenko/discountaria/internal/server/requests"
	"github.com/alishchenko/discountaria/internal/server/responses/problems"
	"github.com/go-chi/render"
	"github.com/pkg/errors"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := requests.NewByIdRequest(r)
	if err != nil {
		helpers.Log(r).Error(errors.Wrap(err, "failed to parse request").Error())
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, problems.BadRequest(errors.Wrap(err, "failed to parse request")))
		return
	}
	user, err := helpers.DB(r).NewUsers().FilterById(id).Get()
	if err != nil {
		helpers.Log(r).Error(errors.Wrap(err, "failed to get user").Error())
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, problems.InternalError())
		return
	}
	if user == nil {
		helpers.Log(r).Error("user with such id not found")
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, problems.NotFound())
		return
	}

	render.JSON(w, r, user)
	return
}
