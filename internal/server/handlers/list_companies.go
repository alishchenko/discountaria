package handlers

import (
	"github.com/alishchenko/discountaria/internal/server/helpers"
	"github.com/alishchenko/discountaria/internal/server/requests"
	"github.com/alishchenko/discountaria/internal/server/responses/problems"
	"github.com/go-chi/render"
	"github.com/pkg/errors"
	"net/http"
)

func ListCompanies(w http.ResponseWriter, r *http.Request) {
	request := requests.NewListCompaniesRequest(r)
	query := helpers.DB(r).NewCompanies().PageParams(request.PaginationParams)
	if request.Name != "" {
		query = query.FilterByName(request.Name)
	}

	companies, err := query.Select()
	if err != nil {
		helpers.Log(r).Error(errors.Wrap(err, "failed to get list companies").Error())
		render.JSON(w, r, problems.BadRequest(errors.Wrap(err, "failed to get list companies")))
		return
	}

	render.JSON(w, r, companies)
	return
}
