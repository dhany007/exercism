package logs

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	result := "default"

	table := map[string]string{
		"U+2757":  "recommendation",
		"U+1F50D": "search",
		"U+2600":  "weather",
	}

	for _, c := range log {
		indexRunes := fmt.Sprintf("%U", c)

		value, ok := table[indexRunes]
		if ok {
			result = value
			break
		}
	}

	return result
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	newLog := ""

	for _, c := range log {
		indexRune := fmt.Sprintf("%d", c)
		indexRune32, _ := strconv.Atoi(indexRune)
		if oldRune == rune(indexRune32) {
			newLog += fmt.Sprintf("%c", newRune)
		} else {
			newLog += fmt.Sprintf("%c", c)
		}
	}

	return newLog
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	result := false

	numberOfRunes := utf8.RuneCountInString(log)
	fmt.Println("numberOfRunes", numberOfRunes)
	fmt.Println("limit", limit)
	if limit >= numberOfRunes {
		result = true
	}

	return result
}
