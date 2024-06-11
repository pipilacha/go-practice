package protein

import (
	"errors"
)

var ErrStop = errors.New("error. stop codon encountered")

var ErrInvalidBase = errors.New("error. invalid base ")

func FromRNA(rna string) (proteins []string, err error) {
	if len(rna)%3 != 0 {
		return nil, ErrInvalidBase
	}

	for i := 0; i < len(rna); i += 3 {
		prot, errors := FromCodon(rna[i : i+3])

		if errors != nil {
			if errors == ErrInvalidBase {
				err = errors
			}
			break
		}

		proteins = append(proteins, prot)
	}

	return proteins, err
}

func FromCodon(codon string) (string, error) {
	switch codon {
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
		return "STOP", ErrStop
	default:
		return "", ErrInvalidBase
	}
}
