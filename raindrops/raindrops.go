package raindrops

import (
	"strconv"
	"strings"
)

/*
Convert a number into a string that contains raindrop sounds corresponding to certain potential factors
 */
func Convert(input int) string {
	res := strings.Builder{}
	if input % 3 == 0 {
		res.WriteString("Pling")
	}
	if input % 5 == 0 {
		res.WriteString("Plang")
	}
	if input % 7 == 0 {
		res.WriteString("Plong")
	}
	if res.Len() == 0 {
		res.WriteString(strconv.Itoa(input))
	}
	return res.String()
}
