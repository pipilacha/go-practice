package strand

var dna_to_rna2 map[rune]string = map[rune]string{
	'G': "C",
	'C': "G",
	'T': "A",
	'A': "U",
}

func ToRNA2(dna string) (rna string) {
	for _, v := range dna {
		rna += dna_to_rna2[v]
	}
	return
}
