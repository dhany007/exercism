package prime

import (
	"errors"
)

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("n must be greater than 0")
	}

	if n == 1 {
		return 2, nil
	}

	prime := []int{}
	i := 2
	for {
		if isPrime(i) {
			prime = append(prime, i)
		}

		if len(prime) == n {
			break
		}

		i += 1
	}

	return prime[n-1], nil
}

func isPrime(n int) bool {
	counter := 0

	for i := 1; i < n; i++ {
		if n%i == 0 {
			counter += 1
		}

		if counter > 1 {
			return false
		}
	}

	return true
}
