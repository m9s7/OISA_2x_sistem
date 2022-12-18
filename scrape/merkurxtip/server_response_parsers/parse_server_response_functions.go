package server_response_parsers

func ParseGetSidebarSports(response map[string]interface{}) map[string]string {

	sidebarSportIDByName := map[string]string{}

	categories, ok := response["categories"].([]interface{})
	if !ok {
		return nil
	}

	for _, sport := range categories {
		sport, ok := sport.(map[string]interface{})
		if !ok {
			continue
		}

		sportName, ok := sport["name"].(string)
		if !ok {
			continue
		}

		sportID, ok := sport["id"].(string)
		if !ok {
			continue
		}

		sidebarSportIDByName[sportName] = sportID
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

// TODO: add checking for casting from interface and print those structures when they fail, EVERYWHERE

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
