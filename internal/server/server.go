package server

import (
	"net/http"

	"github.com/FabioSebs/NotiService/internal/config"
	"github.com/labstack/echo/v4"
)

type Server struct {
	cfg config.HTTP
	// TODO: handlers
}

func NewServer(cfg config.Config) Server {
	return Server{
		cfg: cfg.HTTP,
	}
}

func (s *Server) StartServer() error {
	e := echo.New()

	e.Use(SetCORS())

	SetUpRouter(e)

	errChannel := make(chan error, 1)
	if err := e.Start(s.cfg.Port); err != nil && err != http.ErrServerClosed {
		return err
	}

	return <-errChannel
}
