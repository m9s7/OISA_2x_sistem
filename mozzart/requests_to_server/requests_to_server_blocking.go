package requests_to_server

import (
	"fmt"
)

func GetSidebarSportsAndLeaguesBlocking() []map[string]interface{} {
	response := GetSidebarSportsAndLeagues()
	for response == nil {
		fmt.Println("Mozzart: Stuck on GetSidebarSportsAndLeagues()...")
		response = GetSidebarSportsAndLeagues()
	}
	return response
}

func GetMatchIDsBlocking(sportID int) map[string]interface{} {
	response := GetMatchIDs(sportID)
	for response == nil {
		fmt.Println("Mozzart: Stuck on GetMatchIDs()...", "Match id: ", sportID)
		response = GetMatchIDs(sportID)
	}
	return response
}

func GetOddsBlocking(matchIDs []int, subgameIDs []int) []map[string]interface{} {
	response := GetOdds(matchIDs, subgameIDs)
	for response == nil {
		fmt.Println("Mozzart: Stuck on GetOdds(matchIDs, subgameIDs)...")
		fmt.Println("Match IDs: ", matchIDs)
		fmt.Println("Subgame IDs: ", subgameIDs)
		response = GetOdds(matchIDs, subgameIDs)
	}
	return response
}
