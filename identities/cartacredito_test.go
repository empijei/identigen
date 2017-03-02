package identities

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

/*
Here there are some test cases found on the internet:
5105105105105100 Mastercard
5555555555554444 Mastercard
4111111111111111 Visa
4012888888881881 Visa
378282246310005  AMEX
371449635398431  AMEX
*/

var ccTests = []struct {
	input    int64
	expected int
}{
	{input: 510510510510510,
		expected: 0},
	{input: 555555555555444,
		expected: 4},
	{input: 411111111111111,
		expected: 1},
	{input: 401288888888188,
		expected: 1},
	{input: 37828224631000,
		expected: 5},
	{input: 37144963539843,
		expected: 1},
}

func TestLuhn(t *testing.T) {
	for _, tc := range ccTests {
		if cc := luhn(tc.input); cc != tc.expected {
			t.Errorf("Failed test with %v, calculated: %v, expected: %v", tc.input, cc, tc.expected)
		}
	}
}

func TestCcBuilder(t *testing.T) {
	for _, cct := range ccData {
		cc := ccBuilder(cct)
		n := strings.Join(strings.Split(cc.Number, "-"), "")
		p, _ := strconv.Atoi(n[:len(fmt.Sprintf("%d", cct.base))])
		if p > cct.base+cct.delta || p < cct.base {
			t.Errorf("Failed prefix test: value out of bound. %d !<= %d !< %d", cct.base, p, cct.base+cct.delta)
		}
		if l := len(n); l != cct.bodylength {
			t.Errorf("Failed length test: expected %d, obtained %d", cct.bodylength, l)
		}
		if cvv := len(cc.Cvv); cvv != cct.cvvlength {
			t.Errorf("Failed cvv length test: expected %d, obtained %d", cct.cvvlength, cvv)
		}
	}
}
