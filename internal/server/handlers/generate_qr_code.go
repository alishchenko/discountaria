package handlers

import (
	"github.com/alishchenko/discountaria/internal/server/helpers"
	"github.com/alishchenko/discountaria/internal/server/requests"
	"github.com/alishchenko/discountaria/internal/server/responses"
	"github.com/alishchenko/discountaria/internal/server/responses/problems"
	"github.com/go-chi/render"
	"github.com/pkg/errors"
	"net/http"
)

func GenerateQRCode(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewUseOfferRequest(r)
	if err != nil {
		helpers.Log(r).Error(errors.Wrap(err, "failed to parse request").Error())
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, problems.BadRequest(errors.Wrap(err, "failed to parse request")))
		return
	}
	offer, err := helpers.DB(r).NewOffers().FilterById(request.OfferId).Get()
	if err != nil {
		helpers.Log(r).Error(errors.Wrap(err, "failed to get offer").Error())
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, problems.InternalError())
		return
	}
	if offer == nil {
		helpers.Log(r).Error("offer with such id not found")
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, problems.NotFound())
		return
	}
	user, err := helpers.DB(r).NewUsers().FilterById(request.UserId).Get()
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

	sk, _, err := helpers.GenerateKeys([]byte(user.Email), []byte(user.Password), []byte(helpers.SignatureConfig(r).Salt), helpers.SignatureConfig(r).N, helpers.SignatureConfig(r).R, helpers.SignatureConfig(r).P)
	if err != nil {
		helpers.Log(r).Error("failed to generate keys")
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, problems.InternalError())
		return
	}
	data := helpers.SignatureData{
		UserId:  user.Id,
		OfferId: offer.Id,
	}
	sig, err := helpers.Sign(sk, data)
	render.JSON(w, r, responses.QRCodeResponse{
		Data:      data,
		Signature: sig,
	})
	return
}
