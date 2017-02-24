package identities

import "errors"

type printer struct {
	fieldName string
	function  func(*Person) string
}

//Maps are random-accessed and we want this list to stay ordered as it is here
var printers = []printer{
	{"Nome", (*Person).FirstName},
	{"Cognome", (*Person).LastName},
	{"Gender", (*Person).Gender},
	{"PaeseDiNascita", (*Person).BirthTown},
	{"ProvinciaDiNascita", func(p *Person) string { return p.birthDistrict }},
	{"Indirizzo", (*Person).Address},
	{"NumeroDiTelefono", (*Person).Phone},
	{"DataDiNascita", func(p *Person) string { return p.birthDate.Format(LocalizDate.Format()) }},
	{"CodiceFiscale", (*Person).CodiceFiscale},
	{"PartitaIva", func(p *Person) string {
		pi, _ := p.PartitaIva()
		return pi
	}},
	{"ComunePartitaIva", func(p *Person) string {
		_, cpi := p.PartitaIva()
		return cpi
	}},
	{"Documento", (*Person).ID},
	{"Patente", (*Person).DrivingLicense},
	{"CartaDiCredito", func(p *Person) string {
		cc := p.CartaCredito()
		return cc.issuer + " " + cc.n + ", " + cc.cvv + ", " + cc.expDate
	}},
	{"Iban", (*Person).IBAN},
	{"Username", func(p *Person) string { return p.Credentials().username }},
	{"Password", func(p *Person) string { return p.Credentials().password }},
}
var AllFields []string
var printerMap map[string]func(*Person) string
var fields = generateFilters()

//This should be resolved at compile time, generates an ordered list of keys
//of printers and a map of printers to simplify printing rountines
func generateFilters() []string {
	AllFields = make([]string, 0, len(printers)-1)
	printerMap = make(map[string]func(*Person) string)
	for _, val := range printers {
		AllFields = append(AllFields, val.fieldName)
		printerMap[val.fieldName] = val.function
	}
	return AllFields
}

func SetFilter(newFields []string) error {
	set := make(map[string]struct{})
	for _, allowedField := range AllFields {
		set[allowedField] = struct{}{}
	}
	for _, field := range newFields {
		if _, ok := set[field]; !ok {
			return errors.New("Unknown Field: " + field)
		}
	}
	fields = newFields
	return nil
}
