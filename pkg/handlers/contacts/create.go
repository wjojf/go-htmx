package contacts

import (
	"github.com/labstack/echo/v4"
	"github.com/wjojf/go-htmx/pkg/mapper"
	"github.com/wjojf/go-htmx/pkg/types"
)

var (
	storage = types.ContactsStorage
)

func Create(c echo.Context) error {

	contact, err := mapper.GetContactFromContext(c)
	if err != nil {
		return handleError(c, err, 422)
	}

	if err := storage.AddContact(contact); err != nil {
		return handleError(c, err, 409)
	}

	c.Render(200, "contact-form.html", types.NewFormData())

	return c.Render(200, "contact-oob.html", contact)
}

func handleError(c echo.Context, err error, status int) error {
	formData := mapper.GetFormFromContext(c)
	mapper.AddErrorToForm(formData, err)

	return c.Render(status, "contact-form.html", formData)
}
