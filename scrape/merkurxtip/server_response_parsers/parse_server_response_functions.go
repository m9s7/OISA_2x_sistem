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

func ParseGetSidebarGroups(response *merkurxtip.SidebarSportGroupsResponse) []string {

	var groupIDs []string

	for _, group := range response.Categories {
		if group.Type != "GROUP" || group.Count < 1 {
			continue
		}
		groupIDs = append(groupIDs, group.Id)
	}
	return groupIDs
}

func ParseGetMatchIDs(response *merkurxtip.GetMatchIDsResponse) []int {

	var matchIDs []int

	for _, match := range response.EsMatches {
		if match.Blocked {
			continue
		}
		matchIDs = append(matchIDs, match.Id)
	}

	return matchIDs
}
