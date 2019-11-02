package bob

import (
	"strings"
	"unicode"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	if remark == "" {
		return "Fine. Be that way!"
	} else if isQuestion(remark) && allCaps(remark) {
		return "Calm down, I know what I'm doing!"
	} else if isQuestion(remark) {
		return "Sure."
	} else if allCaps(remark) {
		return "Whoa, chill out!"
	} else {
		return "Whatever."
	}
}

func allCaps(s string) bool {
	if s == "" || !hasLetter(s) {
		return false
	}

	for _, c := range s {
		if unicode.IsLower(c) {
			return false
		} else {
			continue
		}
	}
	return true
}

func isQuestion(s string) bool {
	return s[len(s) - 1] == '?'
}

func hasLetter(s string) bool {
	for _, c := range s {
		if unicode.IsLetter(c) {
			return true
		}
	}
	return false
}
