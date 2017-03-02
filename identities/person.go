package identities

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

//Represents a person object. It must be initialized by the generator.
type Person struct {
	firstName, lastName           string
	genderIsFemale                bool
	birthDate                     time.Time
	town, townCode, birthDistrict string
	residence                     string
	drvLicense                    string
	fiscalCode                    string
	partitaIva                    string
	locationCode                  int
	partitaIvaCounty              string
	cc                            *CartaCredito
	mobilePhone                   string
	id                            string
	iban                          string
	up                            *Credentials
}

//TODO comment ALLL the getters
//Returns the first name of the person
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
func (p *Person) BirthDate() string {
	return p.birthDate.Format(LocalizDate.Format())
}
func (p *Person) BirthTown() string {
	return p.town
}
func (p *Person) BirthDistrict() string {
	return p.birthDistrict
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
	m := p.toMap()
	re := regexp.MustCompile("([a-z])([A-Z]+)")
	var buf bytes.Buffer
	for _, field := range fields {
		_, _ = buf.WriteString(re.ReplaceAllString(field, "$1 $2"))
		_, _ = buf.WriteString(": ")
		_, _ = buf.WriteString(m[field])
		_, _ = buf.WriteString(",\n")
	}
	return buf.String()
}

func (p *Person) MarshalJSON() (b []byte, err error) {
	return json.Marshal(p.toMap())
}

//TODO add sanitization
func (p *Person) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("%v", r))
		}
	}()
	_panic := func(e error) {
		if err != nil {
			panic(err)
		}
	}

	_panic(e.EncodeToken(start))

	for key, value := range p.toMap() {
		_panic(e.EncodeToken(xml.StartElement{Name: xml.Name{Local: key}}))
		_panic(e.EncodeToken(xml.CharData(value)))
		_panic(e.EncodeToken(xml.EndElement{Name: xml.Name{Local: key}}))
	}

	_panic(e.EncodeToken(xml.EndElement{start.Name}))

	// flush to ensure tokens are written
	return e.Flush()
}

func (p Person) MarshalCSV() []string {
	m := p.toMap()
	var out []string
	for _, f := range fields {
		out = append(out, m[f])
	}
	return out
}

//TESTTHIS (try every field alone)
func (p *Person) toMap() map[string]string {
	toret := make(map[string]string)
	for _, f := range fields {
		toret[f] = printerMap[f](p)
	}
	return toret
}
