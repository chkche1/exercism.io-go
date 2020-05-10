package scale

import "strings"

var (
	sharpChromatics = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
	flatChromatics  = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}
	//sharps = []string{"G", "D", "A", "E", "B", "F#", "e", "b", "f#", "c#", "g#", "d#", "C", "a"}
	flats = []string{"F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb"}
	skips = map[rune]int{'m': 1, 'M': 2, 'A': 3}
)

func Scale(tonic, interval string) []string {
	chromatic := getChromatic(tonic)
	if interval == "" {
		return chromatic
	}

	var scale []string
	scale = append(scale, chromatic[0])
	curr := 0
	for _, step := range interval {
		curr += skips[step]
		if curr < len(chromatic) {
			scale = append(scale, chromatic[curr])
		}
	}
	return scale
}

func getScale(tonic string) []string {
	for _, t := range flats {
		if t == tonic {
			return flatChromatics
		}
	}
	return sharpChromatics
}

func getChromatic(tonic string) []string {
	scale := getScale(tonic)

	for i, n := range scale {
		if strings.Title(tonic) == n {
			if i == 0 {
				return scale
			} else {
				return append(scale[i:], scale[0:i]...)
			}
		}
	}
	return scale
}
