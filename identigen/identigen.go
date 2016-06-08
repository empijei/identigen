package identigen

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

type Person struct {
	FirstName, Surname string
	GenderIsFemale     bool
	BirthDate          time.Time
	Town, TownCode     string
}

func NewRandomPeople(minage, maxage int count int) (people []Person) {

	return
}

func randomLines(filepath string, count int) []string {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if count > len(lines) || count == 0 {
		count = len(lines)
	}

	rand.Seed(time.Now().UnixNano())
	indices := rand.Perm(len(lines))
	toreturn := make([]string, 0, count)
	for i := 0; i < count; i++ {
		toreturn = append(toreturn, lines[indices[i]])
	}
	return toreturn
}
