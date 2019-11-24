package isogram

import "unicode"

var hyp = []*unicode.RangeTable{
	unicode.Hyphen,
}

func IsIsogram(s string) bool {
	set := make(map[rune]bool)
	for _, c := range s {
		c = unicode.ToLower(c)
		if unicode.IsSpace(c) || unicode.IsOneOf(hyp, c) {
			continue
		}
		_, ok := set[c]
		if ok {
			return false
		}
		set[c] = true
	}
	return true
}
