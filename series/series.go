package series

func All(n int, s string) (substr []string) {

	for i := 0; i+n <= len(s); i++ {
		substr = append(substr, s[i:i+n])
	}

	return

}

func UnsafeFirst(n int, s string) string {

	if n > len(s) {
		return ""
	}
	return s[:n]
}

func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		return "", false
	}
	return s[:n], true
}
