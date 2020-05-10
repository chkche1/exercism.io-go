package strand

var rnaMap = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

func ToRNA(dna string) string {
	res := make([]rune, len(dna))
	for i, c := range dna {
		res[i] = rnaMap[c]
	}
	return string(res)
}
