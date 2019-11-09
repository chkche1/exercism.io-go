package raindrops

import "strconv"

/*
Convert a number into a string that contains raindrop sounds corresponding to certain potential factors
 */
func Convert(input int) string {
	var res string
	if input % 3 == 0 {
		res += "Pling"
	}
	if input % 5 == 0 {
		res += "Plang"
	}
	if input % 7 == 0 {
		res += "Plong"
	}
	if len(res) == 0 {
		res += strconv.Itoa(input)
	}
	return res
}
