package acronym

import (
	//"fmt"
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
	//fmt.Println(s)
	for _, v := range matches {
		//fmt.Println(v)

		if len(v) == 0 {
			continue
		}
		str := string(v[0])
		res.WriteString(strings.ToUpper(str))
	}
	return res.String()
}
