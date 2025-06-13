package strand

import "strings"

//BenchmarkRNATranscription-8      5645312               209.3 ns/op            56 B/op          6 allocs/op
func ToRNA(dna string) string {
	dnaLength := len(dna)
	rna := strings.Builder{}
	for i := 0; i < dnaLength; i++ {
		switch dna[i] {
		case 71:
			rna.WriteByte(67)
		case 67:
			rna.WriteByte(71)
		case 84:
			rna.WriteByte(65)
		case 65:
			rna.WriteByte(85)
		}
	}
	return rna.String()
}
