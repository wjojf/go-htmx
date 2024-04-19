package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/wjojf/go-htmx/transport/handlers"
	"github.com/wjojf/go-htmx/transport/handlers/contacts"
	"github.com/wjojf/go-htmx/transport/handlers/counter"
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
	e.GET("/", handlers.Index)

	e.POST("/increment", counter.Increment)

	e.POST("/contacts", contacts.Create)
}
