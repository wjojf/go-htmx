package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/wjojf/go-htmx/transport/handlers/detail"
	"github.com/wjojf/go-htmx/transport/handlers/update"
	"github.com/wjojf/go-htmx/transport/template"
)

func New() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())

	e.Renderer = template.New()
	addRouters(e)

	return e
}

func addRouters(e *echo.Echo) {
	e.GET("/", detail.Handler)
	e.POST("/increment", update.Handler)
}
