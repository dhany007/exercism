package romannumerals

import (
	"errors"
)

var (
	arabic = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
)

func ToRomanNumeral(input int) (string, error) {
	var result string
	if input <= 0 || input >= 4000 {
		return "", errors.New("invalid input")
	}

	for {
		for i := 0; i < len(romans); i++ {
			temp := input - arabic[i]
			if temp >= 0 {
				result += romans[i]
				input = temp
				break
			}
		}

		if input == 0 {
			break
		}
	}

	return result, nil
}
