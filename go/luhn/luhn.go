package luhn

import (
	"fmt"
	"strconv"
	"strings"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	if len(id) <= 1 {
		return false
	}

	sum := 0
	for i := len(id) - 1; i >= 0; i-- {
		num, err := strconv.Atoi(fmt.Sprintf("%c", id[i]))
		if err != nil {
			return false
		}

		ptr := len(id) - 1 - i
		if (ptr+1)%2 == 0 {
			num *= 2
		}

		if num > 9 {
			num -= 9
		}

		sum += num
	}

	return sum%10 == 0
}
