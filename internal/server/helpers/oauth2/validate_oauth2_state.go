package oauth2

import (
	"github.com/alishchenko/discountaria/internal/server/helpers"
	"github.com/alishchenko/discountaria/internal/server/responses/problems"
	"github.com/google/jsonapi"
	"net/http"
	"time"
)

func ValidateOAuth2State(state string, stateSecret string, r *http.Request) *jsonapi.ErrorObject {
	log := helpers.Log(r)

	oAuth2State, err := helpers.DB(r).NewOAuth2StatesQ().New().FilterByState(state).Get()
	if err != nil {
		log.Error(err.Error())
		return problems.InternalError()
	}
	if oAuth2State == nil {
		log.Error("not found state")
		return problems.NotFound()
	}
	// Needed to ensure that the state in db has not been forged:
	err = VerifyOAuthToken(oAuth2State.State, stateSecret)
	if err != nil {
		log.Error("invalid oauth2 state")
		return problems.Unauthorized()
	}
	if oAuth2State.ValidTill.UTC().Before(time.Now().UTC()) {
		log.Error("expired oauth2 state")
		return problems.Unauthorized()
	}
	return nil
}
