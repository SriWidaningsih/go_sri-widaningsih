package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func LogMiddlewares(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time.rfc3339}] ${status} ${method} ${host}${path} ${laatency_human}` + "\n",
	}))
}
