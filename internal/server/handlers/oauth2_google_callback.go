package handlers

import (
	"github.com/alishchenko/discountaria/internal/data"
	"github.com/alishchenko/discountaria/internal/server/helpers"
	"github.com/alishchenko/discountaria/internal/server/helpers/oauth2"
	"github.com/alishchenko/discountaria/internal/server/requests"
	"github.com/alishchenko/discountaria/internal/server/responses"
	"github.com/alishchenko/discountaria/internal/server/responses/problems"
	"github.com/go-chi/render"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"net/http"
	"time"
)

func OAuth2GoogleCallback(w http.ResponseWriter, r *http.Request) {
	log := helpers.Log(r)
	request, err := requests.NewVerifyOAuth2Request(r)
	if err != nil {
		log.Error(errors.Wrap(err, "invalid request").Error())
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, problems.BadRequest(err))
		return
	}
	errObject := oauth2.ValidateOAuth2State(request.State, helpers.OAuth2StateConfig(r).StateSecret, r)
	if errObject != nil {
		render.Status(r, cast.ToInt(errObject.Code))
		render.JSON(w, r, errObject)
		return
	}
	user, err := getUserFromGoogle(request.Code, r)
	if errObject != nil {
		render.Status(r, http.StatusUnauthorized)
		render.JSON(w, r, problems.Unauthorized())
		return
	}
	userDB, err := helpers.DB(r).NewUsers().FilterByEmail(user.Email).Get()
	if err != nil {
		helpers.Log(r).Error(errors.Wrap(err, "failed to get user").Error())
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, problems.InternalError())
		return
	}
	var id int64
	if userDB == nil {
		googleOAuth2Provider := oauth2.GoogleOAuth2Provider
		id, err = helpers.DB(r).NewUsers().Insert(data.User{
			Name:                  user.Name,
			Email:                 user.Email,
			Oauth2AccountProvider: &googleOAuth2Provider,
			Password:              generatePassword(12, true, true),
			CreatedAt:             time.Now(),
		})
		if err != nil {
			helpers.Log(r).Error(errors.Wrap(err, "failed to create user").Error())
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, problems.InternalError())
			return
		}
	} else {
		id = userDB.Id
	}
	accessToken, err := helpers.CreateToken(r, id, true)
	if err != nil {
		helpers.Log(r).Error(errors.Wrap(err, "failed to create access token").Error())
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, problems.InternalError())
		return
	}
	refreshToken, err := helpers.CreateToken(r, id, false)
	if err != nil {
		helpers.Log(r).Error(errors.Wrap(err, "failed to create refresh token").Error())
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, problems.InternalError())
		return
	}
	render.JSON(w, r, responses.AuthTokensResponse{Id: id, AccessToken: accessToken, RefreshToken: refreshToken})
	return
}

func getUserFromGoogle(code string, r *http.Request) (*oauth2.User, error) {
	log := helpers.Log(r)

	token, err := oauth2.GetUserToken(code, helpers.OAuth2GoogleConfig(r))
	if err != nil {
		log.Error(errors.Wrap(err, "oauth2 code is invalid").Error())
		return nil, errors.Wrap(err, "oauth2 code is invalid")
	}
	user, err := oauth2.GetGoogleUserInfo(token, helpers.OAuth2GoogleConfig(r))
	if err != nil {
		log.Error(errors.Wrap(err, "oauth2 code is invalid").Error())
		return nil, errors.Wrap(err, "oauth2 code is invalid")
	}
	return &oauth2.User{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
