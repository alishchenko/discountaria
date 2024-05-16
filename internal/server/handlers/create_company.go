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

func CreateCompany(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateCompanyRequest(r)
	if err != nil {
		helpers.Log(r).Error(errors.Wrap(err, "failed to parse request").Error())
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, problems.BadRequest(errors.Wrap(err, "failed to parse request")))
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
		helpers.Log(r).Error("user not found")
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, problems.NotFound())
		return
	}
	id, err := helpers.DB(r).NewCompanies().Insert(data.Company{
		Name:        request.Name,
		LogoURL:     request.LogoUrl,
		URL:         request.Url,
		Description: request.Description,
		UserId:      request.UserId,
		CreatedAt:   time.Now(),
	})
	if err != nil {
		helpers.Log(r).Error(errors.Wrap(err, "failed to insert user").Error())
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, problems.InternalError())
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, responses.IdResponse{Id: id})
	return
}
