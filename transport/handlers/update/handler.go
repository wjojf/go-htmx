package update

import (
	"github.com/labstack/echo/v4"
	"github.com/wjojf/go-htmx/types"
)

func Handler(c echo.Context) error {
	types.DefaultCounter.Count++
	return c.Render(200, "count", types.DefaultCounter)
}
