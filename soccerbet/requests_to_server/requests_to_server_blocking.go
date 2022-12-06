package requests_to_server

import "fmt"

func GetSidebarLeagueIDsBlocking() []map[string]int {
	response := GetSidebarLeagueIDs()
	for response == nil {
		fmt.Println("Stuck on soccerbet request: GetSidebarLeagueIDs")
		response = GetSidebarLeagueIDs()
	}
	return response
}

func GetMasterDataBlocking() map[string]interface{} {
	response := GetMasterData()
	for response == nil {
		fmt.Println("Stuck on soccerbet request: GetMasterData")
		response = GetMasterData()
	}

	return response
}
