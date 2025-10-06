package atbash

import (
	"strings"
	"unicode"
)

func Atbash(s string) string {
	s = strings.ToLower(s)
	var chiper string

	var result []string
	for _, str := range s {
		isDigit := unicode.IsDigit(str)
		isLetter := unicode.IsLetter(str)
		if !(isDigit || isLetter) {
			continue
		}

		if isDigit {
			chiper += string(str)
		}

		if isLetter {
			tmp := str - 'a'
			chiper += string('z' - tmp)
		}
	}

	word := ""
	for i, v := range chiper {
		word += string(v)
		if len(word) == 5 {
			result = append(result, word)
			word = ""
			continue
		}

		if i == len(chiper)-1 && word != "" {
			result = append(result, word)
		}
	}

	return strings.Join(result, " ")
}
