package pangram

import (
	"fmt"
	"strings"
)

func IsPangram(input string) bool {
	if input == "" {
		return false
	}
	input = strings.ToLower(input)
	fmt.Println("input", input)
	chars := make(map[rune]bool, 26)

	for _, in := range input {
		if isAlphabet(in) {
			_, ok := chars[in]
			if !ok {
				chars[in] = true
			}
		}
	}

	return len(chars) == 26
}

func isAlphabet(s rune) bool {
	return s >= 'a' && s <= 'z' // a-z
}
