package handlers

//
//import (
//	"gitlab.com/tokend/klon/verifier-svc/resources"
//	"net/http"
//	"time"
//
//	"github.com/google/jsonapi"
//	"gitlab.com/tokend/klon/verifier-svc/internal/data"
//	"gitlab.com/tokend/klon/verifier-svc/internal/types"
//
//	"gitlab.com/tokend/klon/verifier-svc/internal/service/api/helpers"
//
//	"gitlab.com/distributed_lab/ape"
//	"gitlab.com/distributed_lab/ape/problems"
//	"gitlab.com/tokend/klon/verifier-svc/internal/service/api/ctx"
//	"gitlab.com/tokend/klon/verifier-svc/internal/service/api/requests"
//)
//
//func VerifyEmail(w http.ResponseWriter, r *http.Request) {
//	log := helpers.Log(r)
//	request, err := requests.NewIdentifier(r)
//	if err != nil {
//		log.WithError(err).Error("invalid request")
//		ape.RenderErr(w, problems.BadRequest(err)...)
//		return
//	}
//	verification, errObject := validateVerificationToken(request.Data.Attributes.Identifier, r)
//	if errObject != nil {
//		ape.RenderErr(w, errObject)
//		return
//	}
//	w.WriteHeader(http.StatusOK)
//	ape.Render(w, resources.Key{
//		ID:   verification.IdentificationData,
//		Type: "email",
//	}.AsRelation())
//}
//
//func validateVerificationToken(token string, r *http.Request) (*data.UniqueIdentifier, *jsonapi.ErrorObject) {
//	log := helpers.Log(r)
//
//	verification, err := ctx.UniqueIdentifiersQ(r).New().FilterByIdentifier(token).Get()
//	if err != nil {
//		log.Error(err)
//		return nil, problems.InternalError()
//	}
//	if verification == nil {
//		log.Error("email verification token not found")
//		return nil, problems.NotFound(types.ErrNotFoundIdentifier)
//	}
//	// Needed to ensure that the email verification token in db has not been forged:
//	err = helpers.VerifyToken(verification.Identifier, ctx.EmailVerificationConfig(r).Secret)
//	if err != nil {
//		log.WithError(err).Error("invalid email verification token")
//		return nil, problems.Unauthorized(types.ErrInvalidIdentifier)
//	}
//	if verification.IdentifierExpiresAt.UTC().Before(time.Now().UTC()) {
//		log.Error("email verification token expired")
//		return nil, problems.Unauthorized(types.ErrExpiredIdentifier)
//	}
//	return verification, nil
//}
