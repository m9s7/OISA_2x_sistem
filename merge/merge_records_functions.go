package merge

import (
	"OISA_2x_sistem/utility"
	"strconv"
	"strings"
)

func initRecordWithEl1(el1 *[8]string, indxMap map[string]int, numOfBooks int) []string {

	mergedRecord := make([]string, 5+3*numOfBooks)

	mergedRecord[(indxMap)["kick_off"]] = (*el1)[utility.Kickoff]
	mergedRecord[(indxMap)["league"]] = (*el1)[utility.League]
	mergedRecord[(indxMap)["1"]] = (*el1)[utility.Team1]
	mergedRecord[(indxMap)["2"]] = (*el1)[utility.Team2]
	mergedRecord[(indxMap)["tip1_name"]] = (*el1)[utility.Tip1Name]
	mergedRecord[(indxMap)["tip1_name"]+1] = (*el1)[utility.Tip1Value]
	mergedRecord[(indxMap)["tip2_name"]] = (*el1)[utility.Tip2Name]
	mergedRecord[(indxMap)["tip2_name"]+1] = (*el1)[utility.Tip2Value]

	return mergedRecord
}

func addElToRecord(el *[8]string, bookieOrder int, record *[]string, indxMap map[string]int, isSwitched bool) {

	(*record)[indxMap["league"]+bookieOrder] = (*el)[utility.League]
	if !isSwitched {
		// keep second record as is
		(*record)[indxMap["tip1_name"]+1+bookieOrder] = (*el)[utility.Tip1Value]
		(*record)[indxMap["tip2_name"]+1+bookieOrder] = (*el)[utility.Tip2Value]
	} else {
		// switch second record
		(*record)[indxMap["tip1_name"]+1+bookieOrder] = (*el)[utility.Tip1Value]
		(*record)[indxMap["tip2_name"]+1+bookieOrder] = (*el)[utility.Tip2Value]
	}
}

func shouldSwitchTipVals(tipName string, sportName string) bool {

	var tipNamesNotToSwitch [4]string
	if sportName == "tennis" {
		tipNamesNotToSwitch = [4]string{"TIE_BREAK_YES", "TIE_BREAK_NO", "TIE_BREAK_FST_SET_YES", "TIE_BREAK_FST_SET_NO"}
	} else {
		return false
	}
	for _, el := range tipNamesNotToSwitch {
		if tipName == el {
			return false
		}
	}
	return true
}

func getLeagueNum(league string) int {
	for _, s := range strings.Split(league, " ") {
		leagueNum, err := strconv.Atoi(s)
		if err != nil {
			continue
		}
		return leagueNum
	}
	return -1
}

func isSameLeagueNum(l1 string, l2 string) bool {
	l1LeagueNum := getLeagueNum(l1)
	l2LeagueNum := getLeagueNum(l2)

	if l1LeagueNum >= 2 || l2LeagueNum >= 2 {
		return l1LeagueNum == l2LeagueNum
	}
	return true
}
