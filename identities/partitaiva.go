package identigen

import (
	"fmt"
	"math/rand"
	"strconv"
)

var cities = map[int]string{1: "Torino", 2: "Vercelli-Biella", 3: "Novara-Verbania", 4: "Cuneo", 5: "Asti", 6: "Alessandria", 7: "Aosta", 8: "Imperia", 9: "Savona", 10: "Genova", 11: "La Spezia", 12: "Varese", 13: "Como-Lecco", 14: "Sondrio", 15: "Milano-Lodi", 16: "Bergamo", 17: "Brescia", 18: "Pavia", 19: "Cremona", 20: "Mantova", 21: "Bolzano", 22: "Trento", 23: "Verona", 24: "Vicenza", 25: "Belluno", 26: "Treviso", 27: "Venezia", 28: "Padova", 29: "Rovigo", 30: "Udine", 31: "Gorizia", 32: "Trieste", 33: "Piacenza", 34: "Parma", 35: "Reggio Emilia", 36: "Modena", 37: "Bologna", 38: "Ferrara", 39: "Ravenna", 40: "Forlì-Rimini", 41: "Pesaro", 42: "Ancona", 43: "Macerata", 44: "Ascoli Piceno", 45: "Massa Carrara", 46: "Lucca", 47: "Pistoia", 48: "Firenze", 49: "Livorno", 50: "Pisa", 51: "Arezzo", 52: "Siena", 53: "Grosseto", 54: "Perugia", 55: "Terni", 56: "Viterbo", 57: "Rieti", 58: "Roma", 59: "Latina", 60: "Frosinone", 61: "Caserta", 62: "Benevento", 63: "Napoli", 64: "Avellino", 65: "Salerno", 66: "L'Aquila", 67: "Teramo", 68: "Pescara", 69: "Chieti", 70: "Campobasso", 71: "Foggia", 72: "Bari", 73: "Taranto", 74: "Brindisi", 75: "Lecce", 76: "Potenza", 77: "Matera", 78: "Cosenza", 79: "CZ-KR-VV", 80: "Reggio Calabria", 81: "Trapani", 82: "Palermo", 83: "Messina", 84: "Agrigento", 85: "Caltanissetta", 86: "Enna", 87: "Catania", 88: "Ragusa", 89: "Siracusa", 90: "Sassari", 91: "Nuoro", 92: "Cagliari", 93: "Pordenone", 94: "Isernia", 95: "Oristano", 96: "Milano 2 (Monza)", 97: "Firenze 2 (Prato)", 98: "Brescia 2", 99: "Genova 2", 100: "Roma 2", 120: "Bologna 2", 121: "Napoli 2"}

func (p *Person) PartitaIva() (pi string, county string, err error) {
	location := rand.Int()%100 + 1
	county = cities[location]
	pi = fmt.Sprintf("%07d%03d", rand.Int()%1000000, location)
	num, _ := strconv.Atoi(pi)
	lastDigit := transformation(num, 10)
	pi = fmt.Sprintf("%s%d", pi, lastDigit)
	return
}

func transformation(num, len int) int {
	var digit, evenSum, oddSum int
	for pos := 0; pos < len; pos++ {
		digit = nthdigit(num, pos)
		if pos%2 == 0 {
			tmp := digit * 2
			if tmp > 9 {
				tmp -= 9
			}
			evenSum += tmp
		} else {
			oddSum += digit
		}
	}
	T := (evenSum + oddSum) % 10
	return 10 - T
}
