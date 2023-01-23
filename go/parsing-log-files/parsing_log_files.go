package parsinglogfiles

import (
	"fmt"
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`\A(\[)+(TRC|DBG|INF|WRN|ERR|FTL)+(\])*`)

	return re.Match([]byte(text))
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<([~|*|=|-])*>`)

	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	var result int = 0

	for _, line := range lines {
		newLine := strings.ToLower(line)
		re := regexp.MustCompile(`".*password.*"`)
		if re.Match([]byte(newLine)) {
			result += 1
		}
	}

	return result
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d*`)

	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	var result []string

	for _, line := range lines {
		tempLine := line
		re := regexp.MustCompile(`\s+`)
		newLine := re.ReplaceAllString(tempLine, " ")

		re = regexp.MustCompile(`User\s+[A-Za-z]*\d+`)
		user := re.FindString(newLine)

		if user != "" {
			re = regexp.MustCompile(`User`)
			newUser := re.ReplaceAllString(user, "[USR]")

			result = append(result, fmt.Sprintf("%s %s", newUser, line))
		} else {
			result = append(result, line)
		}

	}

	return result
}
