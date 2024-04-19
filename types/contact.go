package types

type Contact struct {
	Name  string
	Email string
}

func NewContact(name, email string) *Contact {
	return &Contact{
		Name:  name,
		Email: email,
	}
}
