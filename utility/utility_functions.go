package utility

import (
	"math"
	"strings"
)

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

func MergeE1E2(list1 *[4]string, list2 *[4]string) *[8]string {
	var resultList = &[8]string{}
	for i := 0; i < 4; i++ {
		resultList[i] = list1[i]
		resultList[i+4] = list2[i]
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

func Abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func TrimWhiteSpace(str string) string {
	return strings.Trim(str, " ")
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func ToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
