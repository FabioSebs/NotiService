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

func NewEmailer(cfg config.Config) (e Email) {
	e = Email{
		Client: email.NewEmail(),
		Cfg:    cfg.SMTP,
	}
	return
}

func (e *Email) Send(subject string) {

}

func (e *Email) SendHTML() {}
