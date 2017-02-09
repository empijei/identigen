package identities

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Person struct {
	//TODO protect these fields, use accessor and cache data
	firstName, lastName           string
	genderIsFemale                bool
	birthDate                     time.Time
	town, townCode, birthDistrict string
	residence                     string
	fiscalCode                    string
	partitaIva                    string
	partitaIvaCounty              string
	cc                            string
	mobilePhone                   string
	id                            string
	iban                          string
	username                      string
}

func (p *Person) FirstName() string {
	return p.firstName
}
func (p *Person) LastName() string {
	return p.lastName
}
func (p *Person) Gender() string {
	if p.genderIsFemale {
		return "Donna"
	} else {
		return "Uomo"
	}
}
func (p *Person) Address() string {
	return p.residence
}
func (p *Person) BirthDate() time.Time {
	return p.birthDate
}
func (p *Person) BirthTown() string {
	return p.town
}
func (p *Person) Phone() string {
	return p.mobilePhone
}
func (p *Person) ID() string {
	//TODO
	if p.id != "" {
		return p.id
	}
	p.id = fmt.Sprintf("A%s%d", string("QWERTYUIOPASDFGHJKLZXCVBNM"[rand.Int()%26]), rand.Int()%10000000)
	return p.id
}

//TODO test
func (p Person) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString(p.firstName)
	_, _ = buf.WriteString(" ")
	_, _ = buf.WriteString(p.lastName)
	_, _ = buf.WriteString(", ")
	_, _ = buf.WriteString(p.Gender())
	_, _ = buf.WriteString(p.birthDate.Format(LocalizDate.Format()))
	_, _ = buf.WriteString(p.town)
	return buf.String()
}

func (p *Person) MarshalJSON() (b []byte, err error) {
	cf, err := p.CodiceFiscale()
	pi, county, err := p.PartitaIva()
	cc, err := p.CartaCredito()
	iban, err := p.IBAN()
	user, err := p.Username()
	bd := fmt.Sprintf(p.birthDate.Format(LocalizDate.Format()))
	//bd := fmt.Sprintf("%02d/%02d/%04d", p.BirthDate().Day, p.BirthDate().Month, p.BirthDate().Year)
	if err != nil {
		return
	}

	var wrapper = struct {
		Nome               string
		Cognome            string
		Gender             string
		PaeseDiNascita     string
		ProvinciaDiNascita string
		Indirizzo          string
		NumeroDiTelefono   string
		DataDiNascita      string
		CodiceFiscale      string
		PartitaIva         string
		ComunePartitaIva   string
		Documento          string
		CartaDiCredito     string
		Iban               string
		Username           string
	}{
		p.FirstName(),
		p.LastName(),
		p.Gender(),
		p.BirthTown(),
		p.birthDistrict,
		p.Address(),
		p.Phone(),
		bd,
		cf,
		pi,
		county,
		p.ID(),
		cc,
		iban,
		user,
	}
	return json.Marshal(wrapper)
}

func (p Person) MarshalCSV() []string {

	cf, err := p.CodiceFiscale()
	pi, county, err := p.PartitaIva()
	cc, err := p.CartaCredito()
	iban, err := p.IBAN()
	user, err := p.Username()
	bd := fmt.Sprintf(p.birthDate.Format(LocalizDate.Format()))
	//bd := fmt.Sprintf("%02d/%02d/%04d", p.BirthDate().Day, p.BirthDate().Month, p.BirthDate().Year)
	if err != nil {
		return nil
	}

	//FIXME: make this dynamic
	var tmp = []string{
		p.FirstName(),
		p.LastName(),
		p.Gender(),
		p.BirthTown(),
		p.birthDistrict,
		p.Address(),
		p.Phone(),
		bd,
		cf,
		pi,
		county,
		p.ID(),
		cc,
		iban,
		user,
	}
	return tmp
}
