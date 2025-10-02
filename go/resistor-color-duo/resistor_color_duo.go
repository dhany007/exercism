package resistorcolorduo

// Value should return the resistance value of a resistor with a given colors.
var resistors = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

func Value(colors []string) int {
	if len(colors) < 2 {
		return 0
	}

	result := 0
	first, ok := resistors[colors[0]]
	if !ok {
		return 0
	}
	result += first * 10

	second, ok := resistors[colors[1]]
	if !ok {
		return 0
	}
	result += second
	return result
}
