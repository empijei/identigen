package identities

import "testing"

/*
Here there are some test cases found on the internet:
IT 30 R 03268 10001 100000000000
IT 14 D 06055 02100 000001234567
IT 60 X 05428 11101 000000123456
IT 02 D 03268 02801 052879623060
*/

var cinTests = []struct {
	abi         string
	cab         string
	cc          string
	expectedCIN string
	expectedCD  string
}{
	{abi: "03268",
		cab:         "10001",
		cc:          "100000000000",
		expectedCIN: "R",
		expectedCD:  "30",
	},
	{abi: "06055",
		cab:         "02100",
		cc:          "000001234567",
		expectedCIN: "D",
		expectedCD:  "14",
	},
	{abi: "05428",
		cab:         "11101",
		cc:          "000000123456",
		expectedCIN: "X",
		expectedCD:  "60",
	},
	{abi: "03268",
		cab:         "02801",
		cc:          "052879623060",
		expectedCIN: "D",
		expectedCD:  "02",
	},
}

func TestIban(t *testing.T) {
	for _, tc := range cinTests {
		tmpIban := &Iban{abi: tc.abi, cab: tc.cab, cc: tc.cc}
		if cin := tmpIban.cin(); cin != tc.expectedCIN {
			t.Errorf("Failed cin test, calculated: %s, expected: %s", cin, tc.expectedCIN)
		}
		if cd := tmpIban.cd(); cd != tc.expectedCD {
			t.Errorf("Failed cd test, calculated: %s, expected: %s", cd, tc.expectedCD)
		}
	}
}
