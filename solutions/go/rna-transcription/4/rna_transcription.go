package strand

//BenchmarkRNATranscription-8     10588022               111.1 ns/op            16 B/op          5 allocs/op
func ToRNA(dna string) string {
	dnaLength := len(dna)
	rna := make([]byte, len(dna))
	for i := 0; i < dnaLength; i++ {
		switch dna[i] {
		case 71: // G
			rna[i] = 67 // C
		case 67: // C
			rna[i] = 71 // G
		case 84: // T
			rna[i] = 65 // A
		case 65: // A
			rna[i] = 85 // U
		}
	}
	return string(rna)
}
