package anagram

import (
	"strings"
)

func Detect(subject string, candidates []string) []string {
	var result []string

	subject = lower(subject)
	for _, candidate := range candidates {
		cndt := lower(candidate)
		if cndt == subject {
			continue
		}

		if isAnagram(subject, cndt) {
			result = append(result, candidate)
		}
	}

	return result
}

func lower(s string) string {
	return strings.ToLower(s)
}

func toMap(subject string) map[rune]int {
	var subjects = make(map[rune]int, 0)
	for _, s := range subject {
		val, found := subjects[s]
		if found {
			subjects[s] = val + 1
		} else {
			subjects[s] = 1
		}
	}

	return subjects
}

func isAnagram(subject, candidate string) bool {
	subjectMap := toMap(subject)
	candidateMap := toMap(candidate)

	if len(subjectMap) != len(candidateMap) {
		return false
	}

	for key, value := range candidateMap {
		val, found := subjectMap[key]
		if !found {
			return false
		}

		if val != value {
			return false
		}
	}

	return true
}
