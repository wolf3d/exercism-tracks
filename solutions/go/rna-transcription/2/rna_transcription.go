package strand

var translate = map[rune]byte{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

// ToRNA - return DNA compliment
func ToRNA(dna string) string {
	rna := make([]byte, len(dna))
	for i, c := range dna {
		rna[i] = translate[c]
	}
	return string(rna)
}
