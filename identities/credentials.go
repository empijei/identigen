package identities

import "fmt"

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

	up := &Credentials{
		Username: fmt.Sprintf("%s%s%d", p.firstName[:3], p.lastName[:3], p.birthDate.Year()),
		Password: randString([]rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789.!-"), 8),
	}
	p.up = up
	return up
}
