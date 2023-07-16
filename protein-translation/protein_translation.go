package protein

import "errors"

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
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

var (
	ErrInvalidBase = errors.New("invalid base")
	ErrStop        = errors.New("stop error")
)

func FromRNA(rna string) ([]string, error) {
	var stopped bool
	result := make([]string, 0, len(rna)/3)
	for i := 0; i < len(rna)-2; i += 3 {
		if v, ok := codons[rna[i:i+3]]; ok {
			if v == "STOP" {
				stopped = true
				break
			}
			result = append(result, v)
			continue
		}
		return nil, ErrInvalidBase
	}

	if stopped {
		if len(result) == 0 {
			return nil, ErrStop
		}
	}
	return result, nil
}

func FromCodon(codon string) (string, error) {
	if v, ok := codons[codon]; ok {
		if v == "STOP" {
			return "", ErrStop
		}
		return v, nil
	}

	return "", ErrInvalidBase
}
