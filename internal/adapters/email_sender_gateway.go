package adapters

type EmailSenderGateway interface {
	SendEmail(toEmail, subject, body string)
}
