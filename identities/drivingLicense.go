package identities

import (
	"strings"
	"time"
)

type DrivingLicense struct {
	Number  string
	Issuer  string
	ExpDate string
}

//Returns a Italian driving license number
func (p *Person) DrivingLicense() *DrivingLicense {
	if p.drv != nil {
		return p.drv
	}
	location := strings.Split(p.birthDistrict, "(")[1][:2]
	p.drv = &DrivingLicense{
		Number: location +
			randString([]rune("0123456789"), 7) +
			randString([]rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ"), 1),
		Issuer:  "MCTC-" + location,
		ExpDate: p.birthDate.Format("02/01/") + time.Now().AddDate(5, 0, 0).Format("2006"),
	}
	return p.drv
}
