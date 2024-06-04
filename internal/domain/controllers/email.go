package controllers

import "github.com/jmoiron/sqlx"

type EmailController interface {
}

type emailController struct {
	DB *sqlx.DB
}

func NewEmailController(db *sqlx.DB) EmailController {
	return emailController{
		DB: db,
	}
}
