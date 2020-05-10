package reverse

func Reverse(s string) string {
	if len(s) == 0 {
		return s
	}

	res := []rune(s)
	for l, r := 0, len(res)-1; l < r; l, r = l+1, r-1 {
		res[l], res[r] = res[r], res[l]
	}
	return string(res)
}
