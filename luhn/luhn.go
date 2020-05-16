package luhn

import (
	"strings"
	"unicode"
)

/*
Valid returns true if the input is valid per the Luhn algorithm
*/
func Valid(s string) bool {
	s = strings.Replace(s, " ", "", -1)
	if len(s) <= 1 {
		return false
	}

	sum := 0
	even := false
	cs := []rune(s)
	for i := len(cs) - 1; i >= 0; i-- {
		r := cs[i]
		if !unicode.IsDigit(r) {
			return false
		}

		tmp := int(r - '0')
		if even {
			tmp *= 2
			if tmp > 9 {
				tmp -= 9
			}
		}
		sum += tmp
		even = !even
	}
	return sum%10 == 0
}
