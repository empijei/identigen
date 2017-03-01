package lists

import "testing"

func TestAddresses(t *testing.T) {
	cities := make(map[string]struct{}, len(Cities))
	for _, city := range Cities {
		cities[city.Name] = struct{}{}
		_, ok := Addresses[city.Name]
		if !ok {
			t.Errorf("%s doesn't have an address!", city)
		}
	}
	addresses := make(map[string]struct{}, len(Addresses))
	for city, address := range Addresses {
		_, ok := cities[city]
		_, notOk := addresses[address]
		if !notOk {
			addresses[address] = struct{}{}
		} else {
			t.Errorf("Address %s is used more than once!", address)
		}
		if !ok {
			t.Errorf("%s has an address but is not in Cities!", city)
		}
	}
}
