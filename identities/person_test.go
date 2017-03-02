package identities

import (
	"testing"
)

func TestToMap(t *testing.T) {
	ppl, _ := RandomPeople(25, 50, 1)
	person := ppl[0]
	for _, field := range AllFields {
		err := SetFilter([]string{field})
		if err != nil {
			t.Errorf("%s", err)
		}
		m := person.toMap()
		if m[field] == "" {
			t.Errorf("Field %s was requested but not generated", field)
		}
	}
}
