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

func OAuth2FacebookCallback(w http.ResponseWriter, r *http.Request) {
	log := helpers.Log(r)
	request, err := requests.NewVerifyOAuth2Request(r)
	if err != nil {
		log.Error(errors.Wrap(err, "invalid request").Error())
		render.JSON(w, r, problems.BadRequest(err))
		return
	}
	errObject := oauth2.ValidateOAuth2State(request.State, helpers.OAuth2StateConfig(r).StateSecret, r)
	if errObject != nil {
		render.Status(r, cast.ToInt(errObject.Code))
		render.JSON(w, r, errObject)
		return
	}
	user, err := getUserFromFacebook(request.Code, r)
	if err != nil {
		log.Error(errors.Wrap(err, "oauth2 code is invalid").Error())
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
		facebookOAuth2Provider := oauth2.FacebookOAuth2Provider
		id, err = helpers.DB(r).NewUsers().Insert(data.User{
			Name:                  user.Name,
			Email:                 user.Email,
			Oauth2AccountProvider: &facebookOAuth2Provider,
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

func getUserFromFacebook(code string, r *http.Request) (*oauth2.User, error) {
	log := helpers.Log(r)

	token, err := oauth2.GetUserToken(code, helpers.OAuth2FacebookConfig(r))
	if err != nil {
		log.Error(errors.Wrap(err, "oauth2 code is invalid").Error())
		return nil, errors.Wrap(err, "oauth2 code is invalid")
	}
	user, err := oauth2.GetFacebookUserInfo(token)
	if err != nil {
		log.Error(errors.Wrap(err, "oauth2 code is invalid").Error())
		return nil, errors.Wrap(err, "oauth2 code is invalid")
	}
	return user, nil
}
