package types

import "errors"

var (
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrNameRequired       = errors.New("name is required")
	ErrEmailRequired      = errors.New("email is required")
)

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
