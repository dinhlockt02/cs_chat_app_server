package mailer

type Mailer interface {
	Send(subject string, receiverEmail string, receiverName string, htmlContent string) error
}
