package protein

import (
	"bytes"
	"errors"
	"strings"
)

func FromRNA(rna string) (proteins []string, e error) {
	builder := strings.Builder{}
	buf := bytes.Buffer{}
	buf.WriteString(rna)
	codonSlice := buf.Next(3)
	builder.Write(codonSlice)
	if codon, err := FromCodon(builder.String()); err == nil {
		proteins = append(proteins, codon)
	} else {
		return proteins, errors.New("faulty protein")
	}

	return
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
		// case "UAA", "UAG", "UGA":
		// 	return "STOP", errors.New("")
	}
	return "STOP", errors.New("stop codon detected")
}
