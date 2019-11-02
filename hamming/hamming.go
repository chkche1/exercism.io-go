package hamming

import "errors"

/*
Compute the hamming distance of the two input stirngs
 */
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("string lengths differ")
	}

	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}

	return count, nil
}
