package hamming

import "errors"

func Distance(a, b string) (int, error) {
	result := 0
	tempA := []byte(a)
	tempB := []byte(b)

	if a == "" && b == "" {
		return 0, nil
	}

	if len(tempA) != len(tempB) {
		return 0, errors.New("disallow")
	}

	for i := range tempA {
		if tempA[i] != tempB[i] {
			result += 1
		}
	}

	return result, nil
}
