package mailer

import (
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type sendGridMailer struct {
	senderName  string
	senderEmail string
	apiKey      string
}

func NewSendGridMailer(
	senderName string,
	senderEmail string,
	apiKey string,
) *sendGridMailer {
	return &sendGridMailer{
		senderName:  senderName,
		senderEmail: senderEmail,
		apiKey:      apiKey,
	}
}

func (s *sendGridMailer) Send(subject string, receiverEmail string, receiverName string, htmlContent string) error {
	from := mail.NewEmail(s.senderName, s.senderEmail)
	to := mail.NewEmail(receiverName, receiverEmail)
	message := mail.NewSingleEmail(from, subject, to, "", htmlContent)
	client := sendgrid.NewSendClient(s.apiKey)
	_, err := client.Send(message)
	if err != nil {
		return err
	}
	return nil
}
