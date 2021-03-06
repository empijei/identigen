package lists

// CCI maps the numbers to compute the check digit algorithm for Italian IBANs
var CCI = map[int]int{
	0: 1,
	1: 0,
	2: 5,
	3: 7,
	4: 9,
	5: 13,
	6: 15,
	7: 17,
	8: 19,
	9: 21,
}
