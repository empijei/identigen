package identities

import (
	"github.com/empijei/identigen/identities/lists"
)

//Returns a person's address and sets the location for the partita iva field.
func (p *Person) Address() string {
	var location int
	var county string
	//Getting a random element from the map.
	for location, county = range lists.Cities {
		break
	}
	p.locationCode = location
	p.partitaIvaCounty = county
	p.residence = lists.Addresses[county]
	return p.residence
}
