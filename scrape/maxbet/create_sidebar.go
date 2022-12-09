package maxbet

import "strings"

func createSidebar(sidebarSportsJSON []map[string]interface{}) map[string][]int {
	sports := make(map[string][]int)
	for _, sport := range sidebarSportsJSON {

		var leagueBetIds []int

		for _, leagueDict := range sport["leagues"].([]interface{}) {
			leagueDict := leagueDict.(map[string]interface{})
			if strings.HasPrefix(leagueDict["name"].(string), "Max Bonus Tip") {
				continue
			}
			leagueBetIds = append(leagueBetIds, int(leagueDict["betLeagueId"].(float64)))
		}

		if len(leagueBetIds) != 0 {
			sports[sport["name"].(string)] = leagueBetIds
		}
	}
	return sports
}
