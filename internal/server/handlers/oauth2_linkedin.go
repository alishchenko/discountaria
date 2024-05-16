package handlers

import (
	"github.com/alishchenko/discountaria/internal/data"
	"github.com/alishchenko/discountaria/internal/server/helpers"
	"github.com/alishchenko/discountaria/internal/server/helpers/oauth2"
	"github.com/alishchenko/discountaria/internal/server/responses"
	"github.com/alishchenko/discountaria/internal/server/responses/problems"
	"github.com/go-chi/render"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

func OAuth2Linkedin(w http.ResponseWriter, r *http.Request) {
	log := helpers.Log(r)
	state, err := oauth2.GenerateToken(helpers.OAuth2StateConfig(r).StateSecret)
	if err != nil {
		log.Error(errors.Wrap(err, "failed to generate oauth2 state").Error())
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, problems.InternalError())
		return
	}
	validTill := time.Now().Add(helpers.OAuth2StateConfig(r).StateLife).UTC()

	if _, err = helpers.DB(r).NewOAuth2StatesQ().New().Create(data.OAuth2State{
		State:     state,
		ValidTill: &validTill,
	}); err != nil {
		log.Error(err.Error())
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, problems.InternalError())
		return
	}

	url := helpers.OAuth2LinkedinConfig(r).AuthCodeURL(state)
	render.JSON(w, r, responses.ComposeOAuth2(url))
	return
}
