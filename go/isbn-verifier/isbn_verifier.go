package isbn

import (
	"strconv"
	"strings"
)

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")

	if len(isbn) != 10 {
		return false
	}

	pointer := 10
	sum := 0
	for i := 1; i <= len(isbn); i++ {
		char := string(isbn[i-1])
		val, err := strconv.Atoi(char)
		if err != nil {
			if i != len(isbn) {
				return false
			}

			if i == len(isbn) && char != "X" {
				return false
			}

			val = 10
		}

		sum += val * pointer
		pointer -= 1
	}

	return sum%11 == 0
}
