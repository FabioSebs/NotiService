package server

import "github.com/labstack/echo/v4"

func SetUpRouter(e *echo.Echo) {
	// get handlers

	//routes
	v1 := e.Group("/v1")
	{
		email := v1.Group("/email")
		{
			email.POST("", nil) //implement email sending
		}
	}
}
