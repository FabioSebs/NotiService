package controllers

import (
	"github.com/FabioSebs/NotiService/internal/infrastructure"
	"github.com/jmoiron/sqlx"
)

type EmailController interface {
}

type emailController struct {
	DB *sqlx.DB
}

func NewEmailController(infra infrastructure.Infrastructure) EmailController {
	return emailController{
		DB: infra.PostgresDB,
	}
}

func (e *emailController) SendWelcomingEmail() {

}
