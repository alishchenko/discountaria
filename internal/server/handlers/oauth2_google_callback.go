package handlers

import (
	"github.com/alishchenko/discountaria/internal/data"
	"github.com/alishchenko/discountaria/internal/server/helpers"
	"github.com/alishchenko/discountaria/internal/server/helpers/oauth2"
	"github.com/alishchenko/discountaria/internal/server/requests"
	"github.com/alishchenko/discountaria/internal/server/responses"
	"github.com/alishchenko/discountaria/internal/server/responses/problems"
	"github.com/go-chi/render"
	"github.com/google/jsonapi"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

func OAuth2GoogleCallback(w http.ResponseWriter, r *http.Request) {
	log := helpers.Log(r)
	request, err := requests.NewVerifyOAuth2Request(r)
	if err != nil {
		log.Error(errors.Wrap(err, "invalid request").Error())
		render.JSON(w, r, problems.BadRequest(err))
		return
	}
	errObject := oauth2.ValidateOAuth2State(request.State, helpers.OAuth2StateConfig(r).StateSecret, r)
	if errObject != nil {
		render.JSON(w, r, errObject)
		return
	}
	user, errObject := getUserFromGoogle(request.Code, r)
	if errObject != nil {
		render.JSON(w, r, errObject)
		return
	}
	userDB, err := helpers.DB(r).NewUsers().FilterByEmail(user.Email).Get()
	if err != nil {
		helpers.Log(r).Error(errors.Wrap(err, "failed to get user").Error())
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
			CreatedAt:             time.Time{},
		})
		if err != nil {
			helpers.Log(r).Error(errors.Wrap(err, "failed to create user").Error())
			render.JSON(w, r, problems.InternalError())
			return
		}
	}
	accessToken, err := helpers.CreateToken(r, id, true)
	if err != nil {
		helpers.Log(r).Error(errors.Wrap(err, "failed to create access token").Error())
		render.JSON(w, r, problems.InternalError())
		return
	}
	refreshToken, err := helpers.CreateToken(r, id, false)
	if err != nil {
		helpers.Log(r).Error(errors.Wrap(err, "failed to create refresh token").Error())
		render.JSON(w, r, problems.InternalError())
		return
	}
	render.JSON(w, r, responses.AuthTokensResponse{Id: id, AccessToken: accessToken, RefreshToken: refreshToken})
	return
}

func getUserFromGoogle(code string, r *http.Request) (*oauth2.User, *jsonapi.ErrorObject) {
	log := helpers.Log(r)

	token, err := oauth2.GetUserToken(code, helpers.OAuth2GoogleConfig(r))
	if err != nil {
		log.Error(errors.Wrap(err, "oauth2 code is invalid").Error())
		return nil, problems.Unauthorized()
	}
	user, err := oauth2.GetGoogleUserInfo(token, helpers.OAuth2GoogleConfig(r))
	if err != nil {
		log.Error(errors.Wrap(err, "oauth2 code is invalid").Error())
		return nil, problems.Unauthorized()
	}
	return &oauth2.User{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
