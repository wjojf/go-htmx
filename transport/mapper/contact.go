package mapper

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/wjojf/go-htmx/types"
)

func GetContactFromContext(c echo.Context) (*types.Contact, error) {

	name := c.FormValue("name")
	email := c.FormValue("email")

	if name == "" || email == "" {
		return nil, errors.New("name and email are required")
	}

	return types.NewContact(name, email), nil
}
