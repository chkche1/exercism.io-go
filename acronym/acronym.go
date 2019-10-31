package acronym

import (
	"regexp"
	"strings"
)

/*
Returns an acronym for the given input string

e.g. Portable Network Graphics --> PNG
*/
func Abbreviate(s string) string {
	r, _ := regexp.Compile(`[a-zA-Z']*`)

	var res strings.Builder
	matches := r.FindAll([]byte(s), -1)
	for _, v := range matches {

		if len(v) == 0 {
			continue
		}
		str := string(v[0])
		res.WriteString(strings.ToUpper(str))
	}
	return res.String()
}
