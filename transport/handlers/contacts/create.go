package contacts

import (
	"github.com/labstack/echo/v4"
	"github.com/wjojf/go-htmx/transport/mapper"
	"github.com/wjojf/go-htmx/types"
	"log"
)

func Create(c echo.Context) error {

	contact, err := mapper.GetContactFromContext(c)
	if err != nil {
		return c.String(400, err.Error())
	}

	types.ContactsStorage.Contacts = append(types.ContactsStorage.Contacts, contact)

	log.Println("Contact created", contact.Name, contact.Email)

	return c.Render(200, "contact-list.html", types.ContactsStorage)
}
