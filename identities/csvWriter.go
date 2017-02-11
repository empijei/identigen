package identities

import (
	"encoding/csv"
	"io"
)

func MarshalCSV(people []Person, out io.Writer) (err error) {
	w := csv.NewWriter(out)

	if err := w.Write(fields); err != nil {
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
