package reverse

func Reverse(input string) string {
	reversed := ""

	for _, v := range input {
		reversed = string(v) + reversed
	}

	return reversed
}
