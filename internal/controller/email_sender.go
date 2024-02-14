package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/rcarvalho-pb/uber-email-sender-go/internal/core"
	"github.com/rcarvalho-pb/uber-email-sender-go/internal/infra/ses"
	"github.com/rcarvalho-pb/uber-email-sender-go/internal/response"
)

func SendEmail(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	var emailRequest core.EmailRequest
	if err = json.Unmarshal(reqBody, &emailRequest); err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	amazonSES := ses.NewAmazonSES()
	amazonSES.SendEmailGateway(emailRequest.To, emailRequest.Subject, emailRequest.Body)

	response.JSON(w, http.StatusOK, "email sent successfully")
}
