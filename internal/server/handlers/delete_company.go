package handlers

import (
	"github.com/alishchenko/discountaria/internal/server/helpers"
	"github.com/alishchenko/discountaria/internal/server/requests"
	"github.com/alishchenko/discountaria/internal/server/responses/problems"
	"github.com/go-chi/render"
	"github.com/pkg/errors"
	"net/http"
)

func DeleteCompany(w http.ResponseWriter, r *http.Request) {
	id, err := requests.NewByIdRequest(r)
	if err != nil {
		helpers.Log(r).Error(errors.Wrap(err, "failed to parse request").Error())
		render.JSON(w, r, problems.BadRequest(errors.Wrap(err, "failed to parse request")))
		return
	}
	company, err := helpers.DB(r).NewCompanies().FilterById(id).Get()
	if err != nil {
		helpers.Log(r).Error(errors.Wrap(err, "failed to get company").Error())
		render.JSON(w, r, problems.BadRequest(errors.Wrap(err, "failed to get company")))
		return
	}
	if company == nil {
		helpers.Log(r).Error("company not found")
		render.JSON(w, r, problems.NotFound())
		return
	}

	if err = helpers.DB(r).NewCompanies().Delete(id); err != nil {
		helpers.Log(r).Error(errors.Wrap(err, "failed to delete company").Error())
		render.JSON(w, r, problems.BadRequest(errors.Wrap(err, "failed to delete company")))
		return
	}

	render.Status(r, http.StatusNoContent)
	render.JSON(w, r, nil)
	return
}
