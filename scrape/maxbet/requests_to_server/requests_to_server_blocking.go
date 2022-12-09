package requests_to_server

import "fmt"

func GetSidebarBlocking() []map[string]interface{} {
	response := GetSidebar()
	for response == nil {
		fmt.Println("Maxbet: Stuck on GetSidebar()...")
		response = GetSidebar()
	}
	return response
}

func GetMatchIDsBlocking(leagueIDs []int) []map[string]interface{} {
	response := GetMatchIDs(leagueIDs)
	for response == nil {
		fmt.Println("Maxbet: Stuck on GetMatchIDs(leagueIDs)...", "league IDs: ", leagueIDs)
		response = GetMatchIDs(leagueIDs)
	}
	return response
}
