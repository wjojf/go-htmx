package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/wjojf/go-htmx/types"
)

func Index(c echo.Context) error {
	return c.Render(200, "index.html", types.ContactsStorage)
}
