package identities

import (
	"fmt"
	"regexp"
)

// Credentials exports username and password
type Credentials struct {
	Username string
	Password string
}

// Credentials returns fictitious credentials for an identity.
// The username is correlated at the person's name  and year of birth.
func (p *Person) Credentials() *Credentials {
	if p.up != nil {
		return p.up
	}
	name := normalize(p.firstName)
	if len(name) < 3 {
		name = padding(p.firstName)
	}

	lastName := normalize(p.lastName)
	if len(lastName) < 3 {
		lastName = padding(p.lastName)
	}
	up := &Credentials{
		Username: fmt.Sprintf("%s%s%d", name[:3], lastName[:3], p.birthDate.Year()),
		Password: randString([]rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"), 6) +
			randString([]rune(".-!"), 2),
	}
	p.up = up
	return up
}

func normalize(name string) string {
	re := regexp.MustCompile("('| )")
	return string(re.ReplaceAll([]byte(name), []byte("X")))
}

func padding(name string) string {
	for len(name) < 3 {
		name += "X"
	}
	return name
}
