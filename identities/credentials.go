package identities

import (
	"fmt"
	"strings"
)

type Credentials struct {
	Username string
	Password string
}

//Returns ficticious credentials for an identity.
//The username is correlated at the person's name  and year of birth.
func (p *Person) Credentials() *Credentials {
	if p.up != nil {
		return p.up
	}
	name := p.firstName
	if len(name) < 3 {
		name = padding(p.firstName)
	}

	lastName := normalize(p.lastName)
	if len(lastName) < 3 {
		lastName = padding(p.lastName)
	}
	//This supposes no Names/Surnames are shorter than 3
	up := &Credentials{
		Username: fmt.Sprintf("%s%s%d", name[:3], lastName[:3], p.birthDate.Year()),
		Password: randString([]rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"), 6) +
			randString([]rune(".-!"), 2),
	}
	p.up = up
	return up
}

func normalize(name string) string {
	return strings.Join(strings.Split(name, "'"), "X")
}

func padding(name string) string {
	return fmt.Sprintf("%x3s", name)
}
