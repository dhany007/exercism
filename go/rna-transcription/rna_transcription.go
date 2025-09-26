package strand

var rnai = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

func ToRNA(dna string) string {
	result := ""

	for _, d := range dna {
		val, ok := rnai[d]
		if !ok {
			result += string(d)
			continue
		}

		result += string(val)
	}

	return result
}
