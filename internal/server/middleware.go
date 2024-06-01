package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetCORS() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			echo.GET,
			echo.HEAD,
			echo.PATCH,
			echo.PUT,
			echo.DELETE,
			echo.POST,
			echo.OPTIONS,
		},
		AllowHeaders: []string{
			"Origin",
			"Authorization",
			"Access-Control-Allow-Origin",
			"token",
			"Pv",
			echo.HeaderContentType,
			"Accept",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
		},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	})

}
