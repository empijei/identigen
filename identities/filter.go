package identities

import "errors"

var AllFields = []string{
	"Nome",
	"Cognome",
	"Gender",
	"PaeseDiNascita",
	"ProvinciaDiNascita",
	"Indirizzo",
	"NumeroDiTelefono",
	"DataDiNascita",
	"CodiceFiscale",
	"PartitaIva",
	"ComunePartitaIva",
	"Documento",
	"Patente",
	"CartaDiCredito",
	"Iban",
	"Username",
	"Password",
}
var fields = AllFields

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
