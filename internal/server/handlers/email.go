package handlers

import (
	"github.com/FabioSebs/NotiService/internal/domain/services/email"
	"github.com/labstack/echo/v4"
)

type EmailHandler interface {
	SendEmail(ctx echo.Context) error
	// add more
}

type emailHandler struct {
	Service email.Emailer
}

func NewEmailHandler(svc email.Emailer) EmailHandler {
	return &emailHandler{
		Service: svc,
	}
}

func (e *emailHandler) SendEmail(ctx echo.Context) error {
	// request logic goes here
	e.Service.Send("") // TODO:
	return nil
}
