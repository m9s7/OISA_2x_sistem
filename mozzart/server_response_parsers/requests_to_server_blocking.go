package server_response_parsers

import (
	"OISA_2x_sistem/mozzart/requests_to_server"
	"fmt"
)

func GetSidebarSportsAndLeaguesBlocking() []map[string]interface{} {
	response := requests_to_server.GetSidebarSportsAndLeagues()
	for response == nil {
		fmt.Println("Mozzart: Stuck on GetSidebarSportsAndLeagues()...")
		response = requests_to_server.GetSidebarSportsAndLeagues()
	}
	return response
}

func GetMatchIDsBlocking(sportID int) map[string]interface{} {
	response := requests_to_server.GetMatchIDs(sportID)
	for response == nil {
		fmt.Println("Mozzart: Stuck on GetMatchIDs()...", "Match id: ", sportID)
		response = requests_to_server.GetMatchIDs(sportID)
	}
	return response
}

func GetOddsBlocking(matchIDs []int, subgameIDs []int) []map[string]interface{} {
	response := requests_to_server.GetOdds(matchIDs, subgameIDs)
	for response == nil {
		fmt.Println("Mozzart: Stuck on GetOdds(matchIDs, subgameIDs)...")
		fmt.Println("Match IDs: ", matchIDs)
		fmt.Println("Subgame IDs: ", subgameIDs)
		response = requests_to_server.GetOdds(matchIDs, subgameIDs)
	}
	return response
}
