package romannumerals

import "errors"

var values = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
var symbols = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

/*
Convert from normal numbers to Roman Numerals
*/
func ToRomanNumeral(n int) (string, error) {
	if n <= 0 || n > 3000 {
		return "", errors.New("invalid input")
	}

	var res string
	for i, val := range values {
		for n >= val {
			n -= val
			res += symbols[i]
		}
	}
	return res, nil
}
