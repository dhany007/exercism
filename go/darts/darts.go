package darts

import "math"

func Score(x, y float64) int {
	// persamaan lingkaran
	// x^2 + y^2 = r^2
	r := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))

	var (
		inner  float64 = 1
		middle float64 = 5
		outer  float64 = 10
	)

	if r >= 0 && r <= inner {
		return 10
	}

	if r > inner && r <= middle {
		return 5
	}

	if r > middle && r <= outer {
		return 1
	}

	return 0
}
