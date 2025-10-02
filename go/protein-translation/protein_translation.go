package protein

import (
	"errors"
)

var (
	ErrStop        = errors.New("ErrStop")
	ErrInvalidBase = errors.New("ErrInvalidBase")
)

var codons = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
}

var stopCodons = map[string]bool{
	"UAA": true,
	"UAG": true,
	"UGA": true,
}

func FromRNA(rna string) ([]string, error) {
	var result []string

	if len(rna)%3 != 0 {
		return result, ErrInvalidBase
	}

	for i := 0; i < len(rna); i += 3 {
		codon := rna[i : i+3]

		if _, stop := stopCodons[codon]; stop {
			return result, nil
		}

		if acid, ok := codons[codon]; ok {
			result = append(result, acid)
		}
	}
	return result, nil
}

func FromCodon(codon string) (string, error) {
	if _, stop := stopCodons[codon]; stop {
		return "", ErrStop
	}

	acid, ok := codons[codon]
	if !ok {
		return "", ErrInvalidBase
	}

	return acid, nil
}
