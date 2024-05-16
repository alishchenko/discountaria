package handlers

import (
	"github.com/alishchenko/discountaria/internal/server/helpers"
	"github.com/alishchenko/discountaria/internal/server/requests"
	"github.com/alishchenko/discountaria/internal/server/responses"
	"github.com/alishchenko/discountaria/internal/server/responses/problems"
	"github.com/go-chi/render"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewLoginUserRequest(r)
	if err != nil {
		helpers.Log(r).Error(errors.Wrap(err, "failed to parse request").Error())
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, problems.BadRequest(errors.Wrap(err, "failed to parse request")))
		return
	}

	user, err := helpers.DB(r).NewUsers().FilterByEmail(request.Email).Get()
	if err != nil {
		helpers.Log(r).Error(errors.Wrap(err, "failed to get user").Error())
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, problems.InternalError())
		return
	}
	if user == nil {
		helpers.Log(r).Error("user not found")
		render.JSON(w, r, problems.NotFound())
		return
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		helpers.Log(r).Error("invalid password")
		render.Status(r, http.StatusUnauthorized)
		render.JSON(w, r, problems.Unauthorized())
		return
	}

	accessToken, err := helpers.CreateToken(r, user.Id, true)
	if err != nil {
		helpers.Log(r).Error(errors.Wrap(err, "failed to create access token").Error())
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, problems.InternalError())
		return
	}
	refreshToken, err := helpers.CreateToken(r, user.Id, false)
	if err != nil {
		helpers.Log(r).Error(errors.Wrap(err, "failed to create refresh token").Error())
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, problems.InternalError())
		return
	}
	render.JSON(w, r, responses.AuthTokensResponse{Id: user.Id, AccessToken: accessToken, RefreshToken: refreshToken})
	return
}
