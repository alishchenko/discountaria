package handlers

import (
	"github.com/alishchenko/discountaria/internal/data"
	"github.com/alishchenko/discountaria/internal/server/helpers"
	"github.com/alishchenko/discountaria/internal/server/requests"
	"github.com/alishchenko/discountaria/internal/server/responses"
	"github.com/alishchenko/discountaria/internal/server/responses/problems"
	"github.com/go-chi/render"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewRegisterUserRequest(r)
	if err != nil {
		helpers.Log(r).Error(errors.Wrap(err, "failed to parse request").Error())
		render.JSON(w, r, problems.BadRequest(errors.Wrap(err, "failed to parse request")))
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		helpers.Log(r).Error(errors.Wrap(err, "failed to hash password").Error())
		render.JSON(w, r, problems.InternalError())
		return
	}
	id, err := helpers.DB(r).NewUsers().Insert(data.User{
		Name:      request.Name,
		Email:     request.Email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
	})
	if err != nil {
		helpers.Log(r).Error(errors.Wrap(err, "failed to insert user").Error())
		render.JSON(w, r, problems.BadRequest(errors.Wrap(err, "failed to insert user")))
		return
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
