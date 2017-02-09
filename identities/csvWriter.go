package identigen

import (
	"encoding/csv"
	"io"
)

var header = []string{"Nome", "Cognome", "Gender", "PaeseDiNascita", "Indirizzo", "NumeroDiTelefono", "DataDiNascita", "CodiceFiscale", "PartitaIva", "ComunePartitaIva", "Documento", "CartaDiCredito", "Iban", "Username"}

func MarshalCSV(people []Person, out io.Writer) (err error) {
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
