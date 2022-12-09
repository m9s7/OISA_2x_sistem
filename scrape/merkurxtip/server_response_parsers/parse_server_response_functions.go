package server_response_parsers

func ParseGetSidebarSports(response map[string]interface{}) map[string]string {

	sidebarSportIDByName := map[string]string{}
	for _, sport := range response["categories"].([]interface{}) {
		sport := sport.(map[string]interface{})
		sidebarSportIDByName[sport["name"].(string)] = sport["id"].(string)
	}
	return sidebarSportIDByName
}

func ParseGetSidebarGroups(response map[string]interface{}) []string {
	var groupIDs []string
	for _, group := range response["categories"].([]interface{}) {
		group := group.(map[string]interface{})
		if group["type"].(string) != "GROUP" || int(group["count"].(float64)) < 1 {
			continue
		}
		groupIDs = append(groupIDs, group["id"].(string))
	}
	return groupIDs
}

func ParseGetSidebarSportGroupLeagues(response map[string]interface{}) []string {
	var leagueIDs []string
	for _, group := range response["categories"].([]interface{}) {
		group := group.(map[string]interface{})
		if group["type"].(string) != "LEAGUE" || int(group["count"].(float64)) < 1 {
			continue
		}
		leagueIDs = append(leagueIDs, group["id"].(string))
	}
	return leagueIDs
}

func ParseGetMatchIDs(response map[string]interface{}) []int {
	var matchIDs []int

	for _, match := range response["esMatches"].([]interface{}) {
		match := match.(map[string]interface{})
		if match["blocked"].(bool) == true {
			continue
		}
		matchIDs = append(matchIDs, int(match["id"].(float64)))
	}
	return matchIDs
}
