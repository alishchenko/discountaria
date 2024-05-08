package handlers

import (
	"github.com/alishchenko/discountaria/internal/server/helpers"
	"github.com/alishchenko/discountaria/internal/server/requests"
	"github.com/alishchenko/discountaria/internal/server/responses/problems"
	"github.com/go-chi/render"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewUpdateUserRequest(r)
	if err != nil {
		helpers.Log(r).Error(errors.Wrap(err, "failed to parse request").Error())
		render.JSON(w, r, problems.BadRequest(errors.Wrap(err, "failed to parse request")))
		return
	}
	user, err := helpers.DB(r).NewUsers().FilterById(request.Id).Get()
	if err != nil {
		helpers.Log(r).Error(errors.Wrap(err, "failed to get user").Error())
		render.JSON(w, r, problems.BadRequest(errors.Wrap(err, "failed to get user")))
		return
	}
	if user == nil {
		helpers.Log(r).Error("user not found")
		render.JSON(w, r, problems.NotFound())
		return
	}

	query := helpers.DB(r).NewUsers().FilterById(request.Id)
	if user.Oauth2AccountProvider == nil && request.Password != nil {
		if request.OldPassword == nil {
			helpers.Log(r).Error("old password is required")
			render.JSON(w, r, problems.BadRequest(errors.New("old password is required")))
			return
		}
		if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(*request.OldPassword)); err != nil {
			helpers.Log(r).Error("invalid password")
			render.JSON(w, r, problems.BadRequest(errors.New("invalid password")))
			return
		}
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*request.Password), bcrypt.DefaultCost)
		if err != nil {
			helpers.Log(r).Error(errors.Wrap(err, "failed to hash password").Error())
			render.JSON(w, r, problems.InternalError())
			return
		}
		query = query.UpdatePassword(string(hashedPassword))
	}
	if request.Name != nil {
		query = query.UpdateName(*request.Name)
	}
	if request.PhotoURL != nil {
		query = query.UpdatePhotoUrl(*request.PhotoURL)
	}
	if user.Oauth2AccountProvider == nil && request.Email != nil {
		query = query.UpdateEmail(*request.Email)
	}
	if request.PhoneNumber != nil {
		query = query.UpdatePhone(*request.PhoneNumber)
	}

	if err = query.Update(); err != nil {
		helpers.Log(r).Error(errors.Wrap(err, "failed to update user").Error())
		render.JSON(w, r, problems.BadRequest(errors.Wrap(err, "failed to update user")))
		return
	}

	render.Status(r, http.StatusNoContent)
	render.JSON(w, r, nil)
	return
}
