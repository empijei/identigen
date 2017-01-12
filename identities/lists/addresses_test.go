package lists

import "testing"

func TestAddresses(t *testing.T) {
	cities := make(map[string]struct{}, len(Cities))
	for _, city := range Cities {
		cities[city] = struct{}{}
		_, ok := Addresses[city]
		if !ok {
			t.Errorf("%s doesn't have an address!", city)
		}
	}
	for city, _ := range Addresses {
		_, ok := cities[city]
		if !ok {
			t.Errorf("%s has an address but is not in Cities!", city)
		}
	}
}
