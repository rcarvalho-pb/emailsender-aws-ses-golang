package core

type EmailSenderUseCase interface {
	SendEmail(toEmail, subject, body string)
}
