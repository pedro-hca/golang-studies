package organization

import (
	"errors"
	"fmt"
	"strings"

	uuid "github.com/satori/go.uuid"
)

type Identifiable interface {
	ID() string
}

type TwitterHandler string

func (th TwitterHandler) RedirectUrl() string {
	cleanHandler := strings.TrimPrefix(string(th), "@")
	return fmt.Sprintf("https://www.twitter.com/%s", cleanHandler)
}

type Name struct {
	first string
	last  string
}

func (n *Name) FullName() string {
	return fmt.Sprintf("%s %s", n.first, n.last)
}

type Employee struct {
	name Name
}
type Person struct {
	name           Name
	TwitterHandler TwitterHandler
}

func NewPerson(firstName string, lastName string) Person {
	return Person{
		name: Name{
			first: firstName,
			last:  lastName,
		},
	}
}

func (p *Person) ID() string {
	return uuid.NewV4().String()
}

func (p *Person) SetTwitterHandler(handler TwitterHandler) error {
	if len(handler) == 0 {
		p.TwitterHandler = handler
	} else if !strings.HasPrefix(string(handler), "@") {
		return errors.New("twitter handler must start with an @ symbol")
	}

	p.TwitterHandler = handler
	return nil
}

func (p *Person) GetTwitterHandler() string {
	return string(p.TwitterHandler)
}
