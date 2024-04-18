package detail

import (
	"github.com/labstack/echo/v4"
	"github.com/wjojf/go-htmx/types"
)

func Handler(c echo.Context) error {
	return c.Render(200, "index", types.DefaultCounter)
}
