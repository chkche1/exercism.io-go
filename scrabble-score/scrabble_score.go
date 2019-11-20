package scrabble

import "unicode"

/*
Given a word, compute the Scrabble score for that word.
*/
func Score(s string) int {
	res := 0
	for _, c := range s {
		switch unicode.ToUpper(c) {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			res += 1
		case 'D', 'G':
			res += 2
		case 'B', 'C', 'M', 'P':
			res += 3
		case 'F', 'H', 'V', 'W', 'Y':
			res += 4
		case 'K':
			res += 5
		case 'J', 'X':
			res += 8
		case 'Q', 'Z':
			res += 10
		}
	}
	return res
}
