package cipher

import (
	"math"
	"strings"
	"unicode"
)

// Define the shift and vigenere types here.
type shift struct {
	distance int
}

// Both types should satisfy the Cipher interface.

func NewCaesar() Cipher {
	// default 3. not sure this
	return shift{
		distance: 3,
	}
}

func NewShift(distance int) Cipher {
	isDistanceValid := math.Abs(float64(distance)) >= 1 && math.Abs(float64(distance)) <= 25
	if !isDistanceValid {
		return nil
	}

	return shift{
		distance: distance,
	}
}

func (c shift) Encode(input string) string {
	plaintext := cleanInput(input)

	var chipper string
	for _, plain := range plaintext {
		tmp := plain + rune(c.distance)
		if tmp > 'z' {
			dif := tmp - 'z'
			tmp = 'a' + dif - 1
		}

		if tmp < 'a' {
			dif := 'a' - tmp
			tmp = 'z' - dif + 1
		}
		chipper += string(tmp)
	}

	return chipper
}

func (c shift) Decode(input string) string {
	plaintext := cleanInput(input)
	var chipper string
	for _, plain := range plaintext {
		tmp := plain - rune(c.distance)
		if tmp > 'z' {
			dif := tmp - 'z'
			tmp = 'a' + dif - 1
		}

		if tmp < 'a' {
			dif := 'a' - tmp
			tmp = 'z' - dif + 1
		}
		chipper += string(tmp)
	}

	return chipper
}

type vigenere struct {
	key string
}

func NewVigenere(key string) Cipher {
	if !isKeyValid(key) {
		return nil
	}
	return vigenere{
		key: key,
	}
}

func (v vigenere) Encode(input string) string {
	input = cleanInput(input)
	table := constructTable()
	chiper := ""

	for i, in := range input {
		key := v.key[i%len(v.key)]
		x := in - rune('a')
		y := rune(key) - rune('a')
		chip := table[x][y]
		chiper += string(rune(chip) + 'a')
	}

	return chiper
}

func (v vigenere) Decode(input string) string {
	input = cleanInput(input)
	table := constructTable()

	chiper := ""

	for i, in := range input {
		key := v.key[i%len(v.key)]

		row := key - 'a'
		chip := rune(in) - 'a'

		for i, v := range table[row] {
			if rune(v) == chip {
				chiper += string(rune(i) + 'a')
			}
		}
	}

	return chiper
}

func cleanInput(input string) string {
	input = strings.ToLower(input)
	input = strings.TrimSpace(input)

	var plaintext string
	for _, in := range input {
		if unicode.IsLetter(in) {
			plaintext += string(in)
		}
	}

	return plaintext
}

func isKeyValid(key string) bool {
	if len(key) == 0 {
		return false
	}
	if key == "a" {
		return false
	}

	var prev rune
	for _, k := range key {
		if k == 'a' && prev == 'a' {
			return false
		}
		if k < 'a' && k > 'z' {
			return false
		}

		if !unicode.IsLetter(k) {
			return false
		}

		if !unicode.IsLower(k) {
			return false
		}

		prev = k
	}

	return true
}

func constructTable() [][]int {
	result := [][]int{}
	for i := 0; i < 26; i++ {

		child := []int{}
		for j := 0; j < 26; j++ {
			val := i + j
			val = val % 26
			child = append(child, val)
		}

		result = append(result, child)
	}

	return result
}
