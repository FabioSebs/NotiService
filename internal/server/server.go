package server

import (
	"context"
	"net/http"

	"github.com/FabioSebs/NotiService/internal/config"
	"github.com/FabioSebs/NotiService/internal/server/handlers"
	"github.com/labstack/echo/v4"
)

type Server struct {
	cfg      config.HTTP
	handlers handlers.Handlers
	echo     *echo.Echo
}

func NewServer(cfg config.Config, handlers handlers.Handlers) Server {
	return Server{
		cfg:      cfg.HTTP,
		handlers: handlers,
	}
}

func (s *Server) StartServer() error {
	s.echo = echo.New()
	s.echo.HideBanner = true
	s.echo.Use(SetCORS())
	s.SetUpRouter(s.echo)

	errChannel := make(chan error, 1)
	if err := s.echo.Start(s.cfg.Port); err != nil && err != http.ErrServerClosed {
		return err
	}

	return <-errChannel
}

func (s *Server) ShutDownServer(ctx context.Context) error {
	if err := s.echo.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}
