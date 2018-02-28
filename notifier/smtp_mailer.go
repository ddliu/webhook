package notifier

import (
	"crypto/tls"
	"github.com/ddliu/webhook/contact"
	"github.com/ddliu/webhook/context"
	// log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"net"
	"net/mail"
	"net/smtp"
	"strings"
)

type SmtpConfig struct {
	Server   string
	Port     uint
	SSL      bool
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

	host, port, _ := net.SplitHostPort(config.Server)
	config.Server = host
	config.Port = cast.ToUint(port)

	if config.Port == 0 {
		if config.SSL {
			config.Port = 465
		} else {
			config.Port = 25
		}
	} else if config.Port == 465 {
		config.SSL = true
	}

	e.config = config
}

func (e *SmtpMailer) Notify(c *contact.Contact, title, content string) error {
	email := c.GetProperty("Email").(string)
	auth := smtp.PlainAuth("", e.config.Username, e.config.Password, e.config.Server)
	addr := e.config.Server + ":" + cast.ToString(e.config.Port)
	var conn net.Conn
	var err error
	if e.config.SSL {
		conn, err = tls.Dial("tcp", addr, &tls.Config{
			InsecureSkipVerify: true,
			ServerName:         e.config.Server,
		})

		if err != nil {
			// log.Debug("smtp.Dial.ssl error")
			return err
		}
	} else {
		conn, err = net.Dial("tcp", addr)
		if err != nil {
			// log.Debug("smtp.Dial error")
			return err
		}
	}

	client, err := smtp.NewClient(conn, e.config.Server)
	if err != nil {
		// log.Debug("smtp.NewClient error")
		return err
	}

	from := e.config.From
	fromFormated := (&mail.Address{
		Address: from,
	}).String()
	to := []string{email}
	msg := "From: " + fromFormated + "\r\n" +
		"To: " + strings.Join(to, ";") + "\r\n" +
		"Subject: " + title + "\r\n" +
		"\r\n" +
		content +
		"\r\n"

	if err = client.Auth(auth); err != nil {
		// log.Debug("smtp.client.Auth error")
		return err
	}

	if err := client.Mail(from); err != nil {
		// log.Debug("smtp.client.Mail error")
		return err
	}

	if err = client.Rcpt(email); err != nil {
		// log.Debug("smtp.client.Rcpt error")
		return err
	}

	writer, err := client.Data()
	if err != nil {
		// log.Debug("smtp.client.Data error")
		return err
	}

	_, err = writer.Write([]byte(msg))
	if err != nil {
		// log.Debug("smtp.writer.write error")
		return err
	}

	err = writer.Close()
	if err != nil {
		// log.Debug("smtp.writer.close error")
		return err
	}

	return client.Quit()
}

func (e *SmtpMailer) IsMatch(c *contact.Contact) bool {
	return c.GetProperty("Email") != nil
}

func init() {
	RegisterNotifier(&SmtpMailer{})
}
