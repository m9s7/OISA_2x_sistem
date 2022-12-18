package server_response_parsers

import (
	"OISA_2x_sistem/requests_to_server/merkurxtip"
	"fmt"
)

func ParseGetSidebarSports(response *merkurxtip.Sidebar) map[string]string {

	sidebarSportIDByName := map[string]string{}

	for _, sport := range response.Categories {
		sidebarSportIDByName[sport.Name] = sport.Id
	}

	return sidebarSportIDByName
}

func GetSidebarSportsIDsByName() map[string]string {
	response, err := merkurxtip.GetSidebarSports()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	sidebarSportsIDsByName := ParseGetSidebarSports(response)

	return sidebarSportsIDsByName
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
