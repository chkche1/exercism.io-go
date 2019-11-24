package protein

import "errors"

var ErrInvalidBase = errors.New("invalid base")
var ErrStop = errors.New("stop")

// Convert codon into protein
func FromCodon(input string) (string, error) {
	switch input {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}

// Translate RNA sequences into proteins
func FromRNA(input string) ([]string, error) {
	length := len(input)
	var res []string
	for start := 0; start < length; start += 3 {
		codon := input[start : start+3]
		codon, e := FromCodon(codon)
		if e == ErrStop {
			return res, nil
		} else if e == ErrInvalidBase {
			return res, e
		}
		res = append(res, codon)
	}

	return res, nil
}
