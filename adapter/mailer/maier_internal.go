package mailer

import (
	"github.com/kamiazya/memomail/model/attachment"
	"github.com/kamiazya/memomail/model/message"
	gomail "gopkg.in/gomail.v2"
)

func New(c *Config) Mailer {
	m := new(mailer)
	m.c = c
	m.d = gomail.NewDialer(
		m.c.Host,
		m.c.Port,
		m.c.Username,
		m.c.Password,
	)
	return m
}

type mailer struct {
	c *Config
	d *gomail.Dialer
}

func (mlr *mailer) Send(msg *message.Message, ats ...*attachment.Attachment) (err error) {
	m := gomail.NewMessage()
	m.SetHeader("From", mlr.c.EmailAddress)
	m.SetHeader("To", mlr.c.EmailAddress)
	m.SetHeader("Subject", "memomail")
	m.SetBody("text/plain", msg.String())
	for _, at := range ats {
		m.Attach(at.Path)
	}
	err = mlr.d.DialAndSend(m)
	return
}
