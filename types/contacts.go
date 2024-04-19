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

func NewContactList(contacts ...*Contact) *ContactList {
	return &ContactList{
		Contacts: contacts,
	}
}
