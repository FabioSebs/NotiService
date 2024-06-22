package email_svc

import (
	"fmt"
	"net/smtp"

	"github.com/FabioSebs/NotiService/internal/config"
	"github.com/FabioSebs/NotiService/internal/constants"
	"github.com/FabioSebs/NotiService/internal/domain/controllers"
	"github.com/jordan-wright/email"
)

type EmailService interface {
	Send(subject string) (res constants.DEFAULT_RESPONSE, err error)
	SendHTML(subject string) (res constants.DEFAULT_RESPONSE, err error)
	SendNewScrape(recepients []string) (res constants.DEFAULT_RESPONSE, err error)
	SendNewEntry(recepients []string) (res constants.DEFAULT_RESPONSE, err error)
	SendWelcome(recepients []string) (res constants.DEFAULT_RESPONSE, err error)
}

type Email struct {
	Client *email.Email
	Cfg    config.SMTP
	Ctrl   controllers.Controllers
}

func NewEmailService(cfg config.Config, ctrl controllers.Controllers) EmailService {
	return &Email{
		Client: email.NewEmail(),
		Cfg:    cfg.SMTP,
		Ctrl:   ctrl, // make a master controller
	}
}

func (e *Email) Send(subject string) (res constants.DEFAULT_RESPONSE, err error) {
	// Logic goes here
	return
}

func (e *Email) SendHTML(subject string) (res constants.DEFAULT_RESPONSE, err error) {
	return
}

func (e *Email) SendNewScrape(recepients []string) (res constants.DEFAULT_RESPONSE, err error) {
	var (
		server string = e.Cfg.Server
		port   string = e.Cfg.Port
		sender string = e.Cfg.User
		pwd    string = e.Cfg.Password

		serverport string = fmt.Sprintf("%s:%s", server, port)

		subject  string = "ICCT New Source Added!"
		html_msg string = constants.HTML_NEW_SCRAPE
	)

	// setup client
	e.Client.From = sender
	e.Client.To = recepients
	e.Client.Subject = subject
	e.Client.HTML = []byte(html_msg)
	e.Client.AttachFile("stringer.opml")

	// send message
	if err = e.Client.Send(serverport, smtp.PlainAuth("scraper", sender, pwd, server)); err != nil {
		return
	}

	// send to db
	// e.Ctrl.CreateOne()

	//res
	res = constants.DEFAULT_RESPONSE{
		Message: constants.STATUS_SUCCESS_MSG,
		Data:    nil, // TODO: enhance
	}
	return
}

func (e *Email) SendNewEntry(recepients []string) (res constants.DEFAULT_RESPONSE, err error) {
	var (
		server string = e.Cfg.Server
		port   string = e.Cfg.Port
		sender string = e.Cfg.User
		pwd    string = e.Cfg.Password

		serverport string = fmt.Sprintf("%s:%s", server, port)

		subject  string = "ICCT New Entry Added!"
		html_msg string = constants.HTML_NEW_ENTRY
	)

	// setup client
	e.Client.From = sender
	e.Client.To = recepients
	e.Client.Subject = subject
	e.Client.HTML = []byte(html_msg)

	// attachment
	if _, err = e.Client.AttachFile(""); err != nil {
		return
	}

	// send message
	if err = e.Client.Send(serverport, smtp.PlainAuth("entry", sender, pwd, server)); err != nil {
		return
	}

	// send to db
	// e.Ctrl.CreateOne()

	//res
	res = constants.DEFAULT_RESPONSE{
		Message: constants.STATUS_SUCCESS_MSG,
		Data:    nil, // TODO: enhance
	}
	return
}

func (e *Email) SendWelcome(recepients []string) (res constants.DEFAULT_RESPONSE, err error) {
	var (
		server string = e.Cfg.Server
		port   string = e.Cfg.Port
		sender string = e.Cfg.User
		pwd    string = e.Cfg.Password

		serverport string = fmt.Sprintf("%s:%s", server, port)

		subject  string = "Welcome to the Fabrzy Email Subscription!"
		html_msg string = constants.HTML_NEW_EMAIL
	)

	// setup client
	e.Client.From = sender
	e.Client.To = recepients
	e.Client.Subject = subject
	e.Client.HTML = []byte(html_msg)

	// send message
	if err = e.Client.Send(serverport, smtp.PlainAuth("welcome", sender, pwd, server)); err != nil {
		return
	}

	// send to db
	// e.Ctrl.CreateOne()

	//res
	res = constants.DEFAULT_RESPONSE{
		Message: constants.STATUS_SUCCESS_MSG,
		Data:    nil, // TODO: enhance
	}
	return
}
