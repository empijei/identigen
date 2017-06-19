package identities

import (
	"math/rand"

	"github.com/empijei/identigen/identities/lists"
)

// Address returns a person's address and sets the location for the partita iva field.
func (p *Person) Address() string {
	randCity := lists.Cities[rand.Intn(len(lists.Cities))]
	p.locationCode = randCity.Code
	p.partitaIvaCounty = randCity.Name
	p.residence = lists.Addresses[p.partitaIvaCounty]
	return p.residence
}
