package handlers

import (
	"github.com/alishchenko/discountaria/internal/data"
	"github.com/alishchenko/discountaria/internal/server/helpers"
	"github.com/alishchenko/discountaria/internal/server/requests"
	"github.com/alishchenko/discountaria/internal/server/responses"
	"github.com/alishchenko/discountaria/internal/server/responses/problems"
	"github.com/go-chi/render"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

func CreateOffer(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateOfferRequest(r)
	if err != nil {
		helpers.Log(r).Error(errors.Wrap(err, "failed to parse request").Error())
		render.JSON(w, r, problems.BadRequest(errors.Wrap(err, "failed to parse request")))
		return
	}
	company, err := helpers.DB(r).NewCompanies().FilterById(request.CompanyId).Get()
	if err != nil {
		helpers.Log(r).Error(errors.Wrap(err, "failed to get company").Error())
		render.JSON(w, r, problems.InternalError())
		return
	}
	if company == nil {
		helpers.Log(r).Error("company not found")
		render.JSON(w, r, problems.NotFound())
		return
	}
	id, err := helpers.DB(r).NewOffers().Insert(data.Offer{
		CompanyId:  request.CompanyId,
		IsPersonal: len(request.Users) > 0,
		Sale:       request.Sale,
		CreatedAt:  time.Now(),
		ExpiredAt:  request.ExpiredAt,
	})
	if err != nil {
		helpers.Log(r).Error(errors.Wrap(err, "failed to insert offer").Error())
		render.JSON(w, r, problems.BadRequest(errors.Wrap(err, "failed to insert offer")))
		return
	}
	if len(request.Users) > 0 {
		users, err := helpers.DB(r).NewUsers().FilterByEmail(request.Users...).Select()
		if err != nil {
			helpers.Log(r).Error(errors.Wrap(err, "failed to get user").Error())
			render.JSON(w, r, problems.BadRequest(errors.Wrap(err, "failed to ger user")))
			return
		}
		if err = helpers.DB(r).NewOffers().InsertUsers(id, users...); err != nil {
			helpers.Log(r).Error(errors.Wrap(err, "failed to insert users").Error())
			render.JSON(w, r, problems.BadRequest(errors.Wrap(err, "failed to insert users")))
			return
		}
	}
	render.JSON(w, r, responses.IdResponse{Id: id})
	return
}
