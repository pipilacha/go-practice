package strand

var dna_to_rna = map[rune]byte{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

func ToRNA(dna string) (rna string) {
	rna_b := make([]byte, len(dna))
	for i, v := range dna {
		rna_b[i] = dna_to_rna[v]
	}
	return string(rna_b)
}
