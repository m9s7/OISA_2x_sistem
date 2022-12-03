package utility

func IsInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func IndexOf(s string, list []string) int {
	index := -1
	for i, el := range list {
		if el == s {
			index = i
			break
		}
	}
	return index
}
