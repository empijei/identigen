package identities

import (
	"testing"
)

/*
Here there are some test cases found on the internet:
0764352056 7
0636339100 1
0329704036 6
0000000000 0
4444444444 0
1234567890 3
*/

var piTests = []struct {
	input    string
	expected int
}{
	{input: "0764352056",
		expected: 7},
	{input: "0636339100",
		expected: 1},
	{input: "0329704036",
		expected: 6},
	{input: "0000000000",
		expected: 0},
	{input: "4444444444",
		expected: 0},
	{input: "1234567890",
		expected: 3},
}

func TestPiCheckDigit(t *testing.T) {
	for _, tc := range piTests {
		if pi := piCheckDigit(tc.input); pi != tc.expected {
			t.Errorf("Failed test with %v, calculated: %v, expected: %v", tc.input, pi, tc.expected)
		}
	}
}
