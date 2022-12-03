package odds_parsers

func mergeE1E2(list1 [4]string, list2 [4]string) [8]string {
	var resultList [8]string
	for i, el := range list1 {
		resultList[i] = el
	}
	for i, el := range list2 {
		resultList[4+i] = el
	}
	return resultList
}
