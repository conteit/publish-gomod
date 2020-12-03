package data

import (
	"fmt"
	"github.com/dropbox/godropbox/errors"
)

type Person struct {
	Name    string
	Surname string
}

func (p *Person) String() string {
	if p == nil {
		return "(none)"
	}
	return fmt.Sprintf("%s, %s", p.Surname, p.Name)
}

func NewPerson(name, surname string) (*Person, error) {
	if name == "" {
		return nil, errors.New("no name")
	}
	if surname == "" {
		return nil, errors.New("no surname")
	}
	return &Person{
		Name:    name,
		Surname: surname,
	}, nil
}
