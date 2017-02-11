package identities

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
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
	if p.id != "" {
		return p.id
	}
	p.id = fmt.Sprintf("A%s%d", string("QWERTYUIOPASDFGHJKLZXCVBNM"[rand.Int()%26]), rand.Int()%10000000)
	return p.id
}

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
	return json.Marshal(p.toMap())
}

func (p Person) MarshalCSV() []string {
	m := p.toMap()
	var out []string
	for _, f := range fields {
		out = append(out, m[f])
	}
	return out
}

func (p *Person) toMap() map[string]string {
	toret := make(map[string]string)
	for _, f := range fields {
		switch f {
		case "Nome":
			toret[f] = p.FirstName()
		case "Cognome":
			toret[f] = p.LastName()
		case "Gender":
			toret[f] = p.Gender()
		case "PaeseDiNascita":
			toret[f] = p.BirthTown()
		case "ProvinciaDiNascita":
			toret[f] = p.birthDistrict
		case "Indirizzo":
			toret[f] = p.Address()
		case "NumeroDiTelefono":
			toret[f] = p.Phone()
		case "DataDiNascita":
			toret[f] = fmt.Sprintf(p.birthDate.Format(LocalizDate.Format()))
		case "CodiceFiscale":
			toret[f] = logErr(p.CodiceFiscale)
		case "PartitaIva":
			toret[f] = logErr(func() (string, error) {
				pi, _, err := p.PartitaIva()
				return pi, err
			})
		case "ComunePartitaIva":
			toret[f] = logErr(func() (string, error) {
				_, county, err := p.PartitaIva()
				return county, err
			})
		case "Documento":
			toret[f] = p.ID()
		case "CartaDiCredito":
			toret[f] = logErr(p.CartaCredito)
		case "Iban":
			toret[f] = logErr(p.IBAN)
		case "Username":
			toret[f] = logErr(p.Username)
		}
	}
	return toret
}

func logErr(f func() (string, error)) string {
	toret, err := f()
	if err != nil {
		log.Println(err)
	}
	return toret
}