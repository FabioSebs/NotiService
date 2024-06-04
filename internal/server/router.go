package server

import (
	"github.com/FabioSebs/NotiService/internal/server/handlers"
	"github.com/labstack/echo/v4"
)

func (s *Server) SetUpRouter(e *echo.Echo) {
	// get handlers
	var (
		emailHandler handlers.EmailHandler = s.handlers.EmailHandler
	)

	//routes
	v1 := e.Group("/v1")
	{
		email := v1.Group("/email")
		{
			email.POST("", emailHandler.SendEmail)
		}
	}
}
