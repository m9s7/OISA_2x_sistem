package utility

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

func IsElInSliceINT(el int, list []int) bool {
	for _, e := range list {
		if e == el {
			return true
		}
	}
	return false
}

func IsElInSliceSTR(el string, list []string) bool {
	for _, e := range list {
		if el == e {
			return true
		}
	}
	return false
}
