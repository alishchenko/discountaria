package handlers

import (
	"github.com/alishchenko/discountaria/internal/server/helpers"
	"github.com/alishchenko/discountaria/internal/server/responses/problems"
	"github.com/go-chi/render"
	"github.com/pkg/errors"
	"net/http"
)

func ListUsers(w http.ResponseWriter, r *http.Request) {
	//TODO: add filters
	users, err := helpers.DB(r).NewUsers().Select()
	if err != nil {
		helpers.Log(r).Error(errors.Wrap(err, "failed to get list users").Error())
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, problems.InternalError())
		return
	}

	render.Status(r, http.StatusInternalServerError)
	render.JSON(w, r, users)
	return
}
