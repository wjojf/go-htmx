package mapper

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/wjojf/go-htmx/pkg/types"
	"log"
)

func GetContactFromContext(c echo.Context) (*types.Contact, error) {

	name := c.FormValue("name")
	email := c.FormValue("email")

	if name == "" || name == "bob" {
		return nil, types.ErrNameRequired
	}

	if email == "" {
		return nil, types.ErrEmailRequired
	}

	return types.NewContact(name, email), nil
}

func GetFormFromContext(c echo.Context) *types.FormData {

	name := c.FormValue("name")
	email := c.FormValue("email")

	return &types.FormData{
		Values: map[string]string{
			"name":  name,
			"email": email,
		},
		Errors: map[string]string{},
	}
}

func AddErrorToForm(form *types.FormData, err error) {
	switch {
	case errors.Is(err, types.ErrEmailAlreadyExists):
		form.Errors["email"] = err.Error()
	case errors.Is(err, types.ErrEmailRequired):
		form.Errors["email"] = err.Error()
	case errors.Is(err, types.ErrNameRequired):
		form.Errors["name"] = err.Error()
	default:
		log.Println("unknown")
		form.Errors["unknown"] = err.Error()
	}
}
