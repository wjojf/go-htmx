package types

var (
	ContactsStorage = NewContactList(
		NewContact("Alice", "alice@cooper.email.com"),
		NewContact("Bob", "bob@email.com"),
	)
)

type ContactList struct {
	Contacts []*Contact
}

func (c *ContactList) GetContacts() []*Contact {
	return c.Contacts
}

func (c *ContactList) AddContact(contact *Contact) error {

	if c.HasEmail(contact.Email) {
		return ErrEmailAlreadyExists
	}

	c.Contacts = append(c.Contacts, contact)
	return nil
}

func (c *ContactList) HasEmail(email string) bool {
	for _, contact := range c.Contacts {
		if contact.Email == email {
			return true
		}
	}
	return false
}

func NewContactList(contacts ...*Contact) *ContactList {
	return &ContactList{
		Contacts: contacts,
	}
}
