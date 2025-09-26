package resistorcolor

import (
	"strings"
)

// Colors returns the list of all colors.
func Colors() []string {
	colors := []string{
		"black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white",
	}

	return colors
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	color = strings.ToLower(color)
	colors := Colors()

	code := 0
	for idx, c := range colors {
		if c == color {
			code = idx
			break
		}
	}

	return code
}
