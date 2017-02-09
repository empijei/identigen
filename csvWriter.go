package main

import (
	"encoding/csv"
	"io"

	identigen "github.com/empijei/identigen/identities"
)

var header = []string{"Nome", "Cognome", "Gender", "PaeseDiNascita", "Indirizzo", "NumeroDiTelefono", "DataDiNascita", "CodiceFiscale", "PartitaIva", "ComunePartitaIva", "Documento", "CartaDiCredito", "Iban", "Username"}

func MarshalCSV(people []identigen.Person, out io.Writer) (err error) {
	w := csv.NewWriter(out)

	if err := w.Write(header); err != nil {
		return err
	}

	for _, person := range people {
		if err := w.Write(person.MarshalCSV()); err != nil {
			return err
		}
	}

	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	if err := w.Error(); err != nil {
		return err
	}
	return nil
}
