package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	result := make(map[string]int)

	for key, rows := range in {
		for _, row := range rows {
			row = strings.ToLower(row)
			result[row] = key
		}
	}

	return result
}
