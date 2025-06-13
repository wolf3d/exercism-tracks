package protein

import (
	"bytes"
	"errors"
	"strings"
)

var (
	ErrStop        = errors.New("stop codon detected")
	ErrInvalidBase = errors.New("faulty protein")
)

const CodonLength int = 3

//BenchmarkProtein-8        557952              1935 ns/op             848 B/op         33 allocs/op
func FromRNA(rna string) (proteins []string, e error) {
	codonBuilder := strings.Builder{}
	buf := bytes.Buffer{}
	buf.WriteString(rna)
	codonSlice := buf.Next(CodonLength)
	for lengthSlice := len(codonSlice); lengthSlice > 0; {
		codonBuilder.Write(codonSlice)
		codon, err := FromCodon(codonBuilder.String())
		if err == nil {
			codonBuilder.Reset()
			proteins = append(proteins, codon)
		} else if err == ErrStop {
			return proteins, nil
		} else {
			return nil, ErrInvalidBase
		}
		codonSlice = buf.Next(CodonLength)
		lengthSlice = len(codonSlice)
	}
	return
}

//BenchmarkCodon-8         7037148               173.8 ns/op             0 B/op          0 allocs/op
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
		return "", ErrStop
	}
	return "", ErrInvalidBase
}
