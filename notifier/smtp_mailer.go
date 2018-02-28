package notifier

import (
	"github.com/ddliu/webhook/contact"
	"github.com/ddliu/webhook/context"
	"net/smtp"
	"strings"
)

type SmtpConfig struct {
	Server   string
	Username string
	Password string
	From     string
}
type SmtpMailer struct {
	config SmtpConfig
}

func (e *SmtpMailer) GetId() string {
	return "smtp_mailer"
}

func (e *SmtpMailer) Config(c *context.Context) {
	var config SmtpConfig
	c.Unmarshal(&config)
}

func (e *SmtpMailer) Notify(c *contact.Contact, title, content string) error {
	email := c.GetProperty("Email").(string)
	auth := smtp.PlainAuth("", e.config.Username, e.config.Password, e.config.Server)

	from := e.config.From
	to := []string{email}
	msg := "To: " + strings.Join(to, ";") + "\r\n" +
		"Subject: " + title + "\r\n" +
		"\r\n" +
		content +
		"\r\n"

	return smtp.SendMail(e.config.Server, auth, from, to, []byte(msg))
}

func (e *SmtpMailer) IsMatch(c *contact.Contact) bool {
	return c.GetProperty("Email") != nil
}

func init() {
	RegisterNotifier(&SmtpMailer{})
}
