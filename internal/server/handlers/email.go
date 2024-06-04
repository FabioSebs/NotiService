package handlers

import (
	"github.com/FabioSebs/NotiService/internal/constants"
	"github.com/FabioSebs/NotiService/internal/domain/entity"
	"github.com/FabioSebs/NotiService/internal/domain/services/email"
	"github.com/labstack/echo/v4"
)

type EmailHandler interface {
	SendEmail(c echo.Context) (err error)
	// add more services here
}

type emailHandler struct {
	Service email.Emailer
}

func NewEmailHandler(svc email.Emailer) EmailHandler {
	return &emailHandler{
		Service: svc,
	}
}

func (e *emailHandler) SendEmail(c echo.Context) (err error) {
	var (
		res    constants.DEFAULT_RESPONSE
		status int = constants.STATUS_SUCCESS
		req    entity.Email

		typeParam string = c.QueryParam("type")
	)

	// binding
	if err = c.Bind(&req); err != nil {
		return
	}

	// type handling
	switch typeParam {
	case "scrape":
		{
			res, err = e.Service.SendNewScrape(req.Recipients)
			if err != nil {
				return
			}
		}
	case "entry":
		{
			res, err = e.Service.SendNewEntry(req.Recipients)
			if err != nil {
				return
			}
		}
	default:
		return c.JSON(status, res)
	}

	return c.JSON(status, res)
}
