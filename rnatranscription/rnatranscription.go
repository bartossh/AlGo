package strand

func mapToRNA(r rune) rune {
	switch r {
	case 'G':
		return 'C'
	case 'C':
		return 'G'
	case 'T':
		return 'A'
	case 'A':
		return 'U'
	default:
		return ' '
	}
}

// ToRNA changes DNA to RNA
func ToRNA(dna string) string {
	rna := ""
	for _, r := range dna {
		rna += string(mapToRNA(r))
	}
	return rna
}
