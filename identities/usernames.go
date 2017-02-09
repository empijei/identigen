package identigen

import "fmt"

func (p *Person) Username() (username string, err error) {
	if p.username != "" {
		return p.username, nil
	}
	username = fmt.Sprintf("%s%s%d", p.firstName[:3], p.lastName[:3], p.birthDate.Year())
	//log.Print(username)
	return
}
