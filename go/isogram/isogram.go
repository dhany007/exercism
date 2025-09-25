package isogram

import (
	"fmt"
	"strings"
)

func IsIsogram(word string) bool {
	word = strings.ToUpper(word)
	lib := map[string]bool{}

	for _, c := range word {
		char := fmt.Sprintf("%c", c)

		if (char == " ") || (char == "-") {
			continue
		}

		_, ok := lib[char]
		if ok {
			return false
		}

		lib[char] = true
	}

	return true
}
