package types

type Index struct {
	Contacts *ContactList
	Form     *FormData
}

func NewIndex(c *ContactList, f *FormData) *Index {
	return &Index{
		Contacts: c,
		Form:     f,
	}
}
