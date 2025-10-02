package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	var result = make(Frequency)
	phrase = cleanParent(phrase)

	phrases := strings.Split(phrase, " ")
	for _, ph := range phrases {
		ph = clean(ph)
		if ph == "" {
			continue
		}
		val, found := result[ph]
		if found {
			result[ph] = val + 1
		} else {
			result[ph] = 1
		}
	}
	return result
}

func cleanParent(phrase string) string {
	if strings.HasPrefix(phrase, "'") && strings.HasSuffix(phrase, "'") {
		phrase = strings.TrimLeft(phrase, "'")
		phrase = strings.TrimRight(phrase, "'")
	}

	phrase = strings.ReplaceAll(phrase, ",", " ")
	phrase = strings.ReplaceAll(phrase, "\n", "")
	phrase = strings.ReplaceAll(phrase, ":", "")
	phrase = strings.TrimSpace(phrase)
	return phrase

}

func clean(str string) string {
	str = strings.ToLower(str)
	str = strings.TrimSpace(str)
	str = strings.ReplaceAll(str, " ", "")
	if str == "" {
		return ""
	}
	result := ""
	if strings.HasPrefix(str, "'") && strings.HasSuffix(str, "'") {
		str = strings.ReplaceAll(str, "'", "")
	}

	for _, s := range str {
		if unicode.IsLetter(s) {
			result += string(s)
		}

		if unicode.IsDigit(s) {
			result += string(s)
		}

		if s == '\'' {
			result += string(s)
		}
	}

	return result
}
