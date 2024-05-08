package handlers

//
//import (
//	"github.com/alishchenko/discountaria/internal/server/helpers"
//	"github.com/google/jsonapi"
//	"gitlab.com/distributed_lab/ape"
//	"gitlab.com/distributed_lab/ape/problems"
//	"gitlab.com/tokend/klon/verifier-svc/internal/data"
//	"gitlab.com/tokend/klon/verifier-svc/internal/service/api/ctx"
//	"gitlab.com/tokend/klon/verifier-svc/internal/service/api/requests"
//	"gitlab.com/tokend/klon/verifier-svc/internal/types"
//	"gitlab.com/tokend/klon/verifier-svc/resources"
//	"net/http"
//	"time"
//)
//
//func VerifyPhoneNumber(w http.ResponseWriter, r *http.Request) {
//	log := helpers.Log(r)
//	request, err := requests.NewIdentifier(r)
//	if err != nil {
//		log.WithError(err).Error("invalid request")
//		ape.RenderErr(w, problems.BadRequest(err)...)
//		return
//	}
//	verification, errObject := validateVerificationCode(request.Data.Attributes, r)
//	if errObject != nil {
//		ape.RenderErr(w, errObject)
//		return
//	}
//	w.WriteHeader(http.StatusOK)
//	ape.Render(w, resources.Key{
//		ID:   verification.IdentificationData,
//		Type: "phone",
//	}.AsRelation())
//}
//
//func validateVerificationCode(data resources.CheckIdentifierAttributes, r *http.Request) (*data.UniqueIdentifier, *jsonapi.ErrorObject) {
//	log := helpers.Log(r)
//
//	if int64(len(data.Identifier)) != ctx.PhoneVerificationConfig(r).CodeLength {
//		log.Error("invalid phone code identifier")
//		return nil, problems.Unauthorized(types.ErrInvalidIdentifier)
//	}
//	verification, err := ctx.UniqueIdentifiersQ(r).New().FilterByIdentificationData(data.IdentificationData).Get()
//	if err != nil {
//		log.Error(err)
//		return nil, problems.InternalError()
//	}
//	if verification == nil {
//		log.Error("phone code identifier not found")
//		return nil, problems.NotFound(types.ErrNotFoundIdentifier)
//	}
//	if verification.Identifier != data.Identifier {
//		log.Error("invalid phone code identifier")
//		return nil, problems.Unauthorized(types.ErrInvalidIdentifier)
//	}
//	if verification.IdentifierExpiresAt.UTC().Before(time.Now().UTC()) {
//		log.Error("phone code identifier expired")
//		return nil, problems.Unauthorized(types.ErrExpiredIdentifier)
//	}
//	return verification, nil
//}
