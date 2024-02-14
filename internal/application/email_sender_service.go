package application

import (
	"github.com/rcarvalho-pb/uber-email-sender-go/internal/adapters"
)

type EmailSenderUseCase struct {
	EmailSenderGateway adapters.EmailSenderGateway
}

func (es *EmailSenderUseCase) SendEmail(toEmail, subject, body string) {
	es.EmailSenderGateway.SendEmail(toEmail, subject, body)
}
