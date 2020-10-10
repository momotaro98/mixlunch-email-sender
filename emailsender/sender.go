package emailsender

import (
	"context"
	"log"
	"os"

	"github.com/mailgun/mailgun-go/v4"
)

var domain string
var privateAPIKey string
var mg *mailgun.MailgunImpl

func InitEmailSender() {
	// Your available domain names can be found here:
	// (https://app.mailgun.com/app/domains)
	domain = os.Getenv("MAIL_DOMAIN_NAME")
	// You can find the Private API Key in your Account Menu, under "Settings":
	// (https://app.mailgun.com/app/account/security)
	privateAPIKey = os.Getenv("MAILGUN_PRIVATE_API_KEY")

	mg = mailgun.NewMailgun(domain, privateAPIKey)
}

type EmailSender interface {
	Send(ctx context.Context, sender string, subject string, body string, recipient string) error
}

type mailGunEmailSender struct{}

func NewSender() EmailSender {
	return &mailGunEmailSender{}
}

func (s *mailGunEmailSender) Send(ctx context.Context, sender string, subject string, body string, recipient string) error {
	// The message object allows you to add attachments and Bcc recipients
	message := mg.NewMessage(sender, subject, "text", recipient)

	message.SetHtml(body)

	// Send the message	with a 10 second timeout
	resp, id, err := mg.Send(ctx, message)

	if err != nil {
		log.Println("Error in EmailSender.Send", err)
		return err
	}

	log.Printf("ID: %s Resp: %s\n", id, resp)

	return nil
}
