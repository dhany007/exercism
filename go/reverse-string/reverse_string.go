package reverse

func Reverse(input string) string {
	result := ""

	for _, char := range input {
		result = string(char) + result
	}

	return result
}
