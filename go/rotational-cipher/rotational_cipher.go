package rotationalcipher

import (
	"fmt"
	"unicode"
)

func RotationalCipher(plain string, shiftKey int) string {
	chiper := ""
	for _, v := range plain {
		if unicode.IsLetter(v) {
			start := 'a'
			end := 'z'
			if unicode.IsUpper(v) {
				start = 'A'
				end = 'Z'
			}

			shift := (shiftKey % 26) + int(v)
			if shift > int(end) {
				shift = shift - int(end) + int(start) - 1
			}

			chiper += fmt.Sprintf("%c", rune(shift))
			continue
		}

		chiper += string(v)
	}

	return chiper
}
