package handlers

//
//import (
//	"net/http"
//	"strings"
//
//	"github.com/go-ozzo/ozzo-validation/v4/is"
//
//	"gitlab.com/tokend/klon/verifier-svc/internal/service/api/helpers"
//	"gitlab.com/tokend/klon/verifier-svc/internal/types"
//
//	"gitlab.com/distributed_lab/ape"
//	"gitlab.com/distributed_lab/ape/problems"
//	"gitlab.com/tokend/klon/verifier-svc/internal/service/api/requests"
//)
//
//func SendVerifyPhoneNumber(w http.ResponseWriter, r *http.Request) {
//	log := helpers.Log(r)
//	request, err := requests.NewGenerateIdentifierRequest(r)
//	if err != nil {
//		log.WithError(err).Error("invalid request")
//		errObjects := problems.BadRequest(err)
//		if strings.Contains(err.Error(), is.ErrEmail.Error()) {
//			if len(errObjects) == 1 {
//				errObjects[0].Code = types.ErrInvalidEmail
//			}
//		}
//		ape.RenderErr(w, errObjects...)
//		return
//	}
//
//	if err = helpers.SendConfirmationSMS(request.Data.Attributes.IdentificationData, r); err != nil {
//		log.Error(err)
//		ape.RenderErr(w, problems.InternalError())
//		return
//	}
//
//	w.WriteHeader(http.StatusNoContent)
//	return
//}
