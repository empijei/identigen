package identities

import (
	"testing"
	"time"
)

var credTests = []struct {
	name, lastName string
	expected       string
}{
	{"", "", "XXXXXX1992"},
	{"Jo'", "Cotton", "JoXCot1992"},
	{"J o", "Cotton", "JXoCot1992"},
	{"a", "b", "aXXbXX1992"},
}

func TestCredentials(t *testing.T) {

	p := NewPerson(25, 55)
	p.birthDate = time.Date(1992, 1, 1, 1, 1, 1, 1, time.UTC)
	for _, tc := range credTests {
		p.firstName = tc.name
		p.lastName = tc.lastName
		if p.Credentials().Username != tc.expected {

		}
	}
}
