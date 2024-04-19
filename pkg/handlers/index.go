package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/wjojf/go-htmx/pkg/types"
)

func Index(c echo.Context) error {

	data := types.NewIndex(
		types.ContactsStorage,
		types.NewFormData(),
	)

	return c.Render(200, "index.html", data)
}
