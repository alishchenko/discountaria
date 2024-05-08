package handlers

import (
	"github.com/alishchenko/discountaria/internal/server/helpers"
	"github.com/alishchenko/discountaria/internal/server/responses/problems"
	"github.com/go-chi/render"
	"github.com/pkg/errors"
	"net/http"
)

func ListOffers(w http.ResponseWriter, r *http.Request) {
	//TODO: add filters
	offers, err := helpers.DB(r).NewOffers().Select()
	if err != nil {
		helpers.Log(r).Error(errors.Wrap(err, "failed to insert offer").Error())
		render.JSON(w, r, problems.BadRequest(errors.Wrap(err, "failed to insert offer")))
		return
	}

	render.JSON(w, r, offers)
	return
}
