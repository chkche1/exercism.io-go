package isogram

import "unicode"

var hyp = []*unicode.RangeTable{
	unicode.Hyphen,
}

// IsIsogram determines if a word or phrase is an isogram
func IsIsogram(s string) bool {
	set := make(map[rune]bool)
	for _, c := range s {
		c = unicode.ToLower(c)
		if unicode.IsSpace(c) || unicode.IsOneOf(hyp, c) {
			continue
		}
		if set[c] {
			return false
		}
		set[c] = true
	}
	return true
}
