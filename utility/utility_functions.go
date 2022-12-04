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

func MergeE1E2(list1 [4]string, list2 [4]string) [8]string {
	var resultList [8]string
	for i, el := range list1 {
		resultList[i] = el
	}
	for i, el := range list2 {
		resultList[4+i] = el
	}
	return resultList
}

func RemoveDuplicates(intSlice *[]int) []int {
	keys := make(map[int]bool)
	var list []int

	for _, entry := range *intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
