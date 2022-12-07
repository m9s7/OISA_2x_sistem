package merge

import (
	"sort"
)

func mapToSlice(myMap map[string][]*[8]string) []bookie {
	vec := make([]bookie, len(myMap))
	i := 0
	for k, v := range myMap {
		vec[i].name = k
		vec[i].rows = v
		i++
	}
	return vec
}

func orderBooksByNumOfRecords(bookies []bookie) {
	sort.Slice(bookies, func(i, j int) bool {
		return len((bookies[i]).rows) > len((bookies[j]).rows)
	})
}

func getColumnIndexes(numOfBookies int) map[string]int {
	ColIndx := map[string]int{
		"kick_off": 0,
		"league":   1,
	}
	ColIndx["1"] = ColIndx["league"] + numOfBookies
	ColIndx["2"] = ColIndx["1"] + 1
	ColIndx["tip1_name"] = ColIndx["2"] + 1
	ColIndx["tip2_name"] = ColIndx["tip1_name"] + numOfBookies + 1

	return ColIndx
}

func getMergedRecordsColumnNames(bookies []bookie) []string {

	colsNames := []string{"kick_off"}
	for _, bookie := range bookies {
		colsNames = append(colsNames, "league_"+bookie.name)
	}
	colsNames = append(colsNames, "1", "2", "tip1")
	for _, bookie := range bookies {
		colsNames = append(colsNames, "tip1_"+bookie.name)
	}
	colsNames = append(colsNames, "tip2")
	for _, bookie := range bookies {
		colsNames = append(colsNames, "tip2_"+bookie.name)
	}

	return colsNames
}
