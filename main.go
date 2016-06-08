package main

import (
	"fmt"
	"time"

	"github.com/empijei/gotrials/CodiceFiscale/identigen"
)

func main() {
	test := identigen.Person{FirstName: "Roberto", Surname: "Clapis", Town: "CARATE BRIANZA (MB)", TownCode: "B729"}
	test.BirthDate = time.Date(1992, time.January, 31, 12, 0, 0, 0, time.UTC)
	cf, err := test.CodiceFiscale()
	if err != nil {
	}
	fmt.Println(cf)
}
