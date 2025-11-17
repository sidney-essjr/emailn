package mail

import (
	"emailn/internal/domain/campaign"
	"errors"
	"fmt"
	"os"

	"github.com/wneessen/go-mail"
)

func SendMail(campaign *campaign.Campaign) error {
	fmt.Println("Sending email...")

	emails := make([]string, 0, len(campaign.Contacts))
	for _, email := range campaign.Contacts {
		emails = append(emails, email.Email)
	}

	msg := mail.NewMsg()
	msg.From(os.Getenv("EMAIL_USER"))
	msg.To(emails...)
	msg.Subject(campaign.Name)
	msg.SetBodyString(mail.TypeTextHTML, "<h3>Olá</h3><p>"+campaign.Content+"</p>")

	client, err := mail.NewClient(
		os.Getenv("EMAIL_SMTP"),
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(os.Getenv("EMAIL_USER")),
		mail.WithPassword(os.Getenv("EMAIL_PASSWORD")),
		mail.WithPort(587),
		mail.WithTLSPolicy(mail.DefaultTLSPolicy),
	)

	if err != nil {
		return errors.New("it was not possible to create the client")
	}

	if err := client.DialAndSend(msg); err != nil {
		return errors.New("failed to send mail")
	}

	return nil
}
