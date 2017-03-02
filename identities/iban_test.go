package identities

import (
	"testing"
)

/*
Here there are some test cases found on the internet:
IT 30 R 03268 10001 100000000000
IT 14 D 06055 02100 000001234567
IT 60 X 05428 11101 000000123456
IT 02 D 03268 02801 052879623060
*/

var cinTests = []struct {
	abi      string
	cab      string
	cc       string
	expected string
}{
	{abi: "03268",
		cab:      "10001",
		cc:       "100000000000",
		expected: "R"},
	{abi: "06055",
		cab:      "02100",
		cc:       "000001234567",
		expected: "D"},
	{abi: "05428",
		cab:      "11101",
		cc:       "000000123456",
		expected: "X"},
	{abi: "03268",
		cab:      "02801",
		cc:       "052879623060",
		expected: "D"},
}

var ibanTests = []struct {
	input    string
	expected string
}{
	{input: "R0326810001100000000000",
		expected: "30"},
	{input: "D0605502100000001234567",
		expected: "14"},
	{input: "X0542811101000000123456",
		expected: "60"},
	{input: "D0326802801052879623060",
		expected: "02"},
}

func TestCheckDigit(t *testing.T) {
	for _, tc := range ibanTests {
		if iban := checkDigit(tc.input); iban != tc.expected {
			t.Errorf("Failed test with %v, calculated: %v, expected: %v", tc.input, iban, tc.expected)
		}
	}
}

func TestCin(t *testing.T) {
	for _, bban := range cinTests {
		if cin := cin(bban.abi, bban.cab, bban.cc); cin != bban.expected {
			t.Errorf("Failed cin test, calculated: %s, expected: %s", cin, bban.expected)
		}
	}
}
