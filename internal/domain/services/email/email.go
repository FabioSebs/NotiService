package email

import (
	"github.com/FabioSebs/NotiService/internal/config"
	"github.com/jordan-wright/email"
)

type Emailer interface {
	Send(subject string)
	SendHTML(subject string)
}

type Email struct {
	Client *email.Email
	Cfg    config.SMTP
}

func NewEmailer(cfg config.Config) Emailer {
	return &Email{
		Client: email.NewEmail(),
		Cfg:    cfg.SMTP,
	}
}

func (e *Email) Send(subject string) {
	// Logic goes here

}

func (e *Email) SendHTML(subject string) {}
