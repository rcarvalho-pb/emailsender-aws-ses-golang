package routes

import (
	"net/http"

	"github.com/rcarvalho-pb/uber-email-sender-go/internal/controller"
)

var EmailSenderRoute = []Route{
	{
		Uri:      "/api/emails",
		Method:   http.MethodPost,
		Function: controller.SendEmail,
	},
}
