package counter

import (
	"github.com/labstack/echo/v4"
	"github.com/wjojf/go-htmx/pkg/types"
)

func Increment(c echo.Context) error {
	types.DefaultCounter.Count++
	return c.Render(200, "count", types.DefaultCounter)
}
