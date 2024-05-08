package handlers

import (
	"github.com/alishchenko/discountaria/internal/server/helpers"
	"github.com/alishchenko/discountaria/internal/server/requests"
	"github.com/alishchenko/discountaria/internal/server/responses/problems"
	"github.com/go-chi/render"
	"github.com/pkg/errors"
	"net/http"
)

func UseOffer(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewUseOfferRequest(r)
	if err != nil {
		helpers.Log(r).Error(errors.Wrap(err, "failed to parse request").Error())
		render.JSON(w, r, problems.BadRequest(errors.Wrap(err, "failed to parse request")))
		return
	}
	offer, err := helpers.DB(r).NewOffers().FilterById(request.OfferId).Get()
	if err != nil {
		helpers.Log(r).Error(errors.Wrap(err, "failed to get offer").Error())
		render.JSON(w, r, problems.BadRequest(errors.Wrap(err, "failed to get offer")))
		return
	}
	if offer == nil {
		helpers.Log(r).Error("offer with such id not found")
		render.JSON(w, r, problems.NotFound())
		return
	}
	user, err := helpers.DB(r).NewUsers().FilterById(request.UserId).Get()
	if err != nil {
		helpers.Log(r).Error(errors.Wrap(err, "failed to get user").Error())
		render.JSON(w, r, problems.BadRequest(errors.Wrap(err, "failed to get user")))
		return
	}
	if user == nil {
		helpers.Log(r).Error("user with such id not found")
		render.JSON(w, r, problems.NotFound())
		return
	}

	if offer.IsPersonal {

		helpers.Log(r).Error("offer with such id not found")
		render.JSON(w, r, problems.NotFound())
		return
	}

	render.JSON(w, r, offer)
	return
}
