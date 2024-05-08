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
		helpers.Log(r).Error(errors.Wrap(err, "failed to insert user").Error())
		render.JSON(w, r, problems.BadRequest(errors.Wrap(err, "failed to insert user")))
		return
	}

	render.JSON(w, r, users)
	return
}
