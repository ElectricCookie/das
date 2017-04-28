package user

import (
	"fmt"

	"github.com/ElectricCookie/das-cms/configLoader"
	"github.com/ElectricCookie/das-cms/db"
	"github.com/ElectricCookie/das-cms/i18n"
	gomail "gopkg.in/gomail.v2"
)

func sendEmail(targetEmail string, subject, content string) *error {

	message := gomail.NewMessage()
	message.SetHeader("From", configLoader.GetConfig().SMTPFrom)
	message.SetHeader("To", targetEmail)
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", content)

	d := gomail.NewDialer(
		configLoader.GetConfig().SMTPURL,
		configLoader.GetConfig().SMTPPort,
		configLoader.GetConfig().SMTPUser,
		configLoader.GetConfig().SMTPPassword,
	)

	if err := d.DialAndSend(message); err != nil {
		fmt.Println(err)
		return &err
	}

	return nil

}

// SendRegistrationEmail sends a registration email, with confirmation link
func SendRegistrationEmail(user db.User) *error {

	if user.EmailVerified {
		return nil
	}

	t := i18n.GetTranslator(user.Language)

	return sendEmail(user.Email, t("register:subject", user), t("register:emailBody", user))

}
