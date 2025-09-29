package strain

// Implement the "Keep" and "Discard" function in this file.

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1

func Keep[T any](nums []T, predicate func(T) bool) []T {
	var result []T
	for _, num := range nums {
		if predicate(num) {
			result = append(result, num)
		}
	}

	return result
}

func Discard[T any](nums []T, predicate func(T) bool) []T {
	var result []T

	for _, num := range nums {
		if !predicate(num) {
			result = append(result, num)
		}
	}

	return result
}
