package requests_to_server

import (
	"OISA_2x_sistem/settings"
	"log"
)

func GetAllSubgamesBlocking() map[string]interface{} {

	response := GetAllSubgames()

	for i := 0; i < settings.NumOfTries; i++ {

		if response != nil {
			break
		}

		log.Println("Merkurxtip: Stuck on GetAllSubgames()...")
		response = GetAllSubgames()

	}

	return response
}

func GetSidebarSportGroupsBlocking(sportID string) map[string]interface{} {

	response := GetSidebarSportGroups(sportID)

	for i := 0; i < settings.NumOfTries; i++ {

		if response != nil {
			break
		}

		log.Println("Merkurxtip: Stuck on GetSidebarSportGroups(sportID=" + sportID + ")...")
		response = GetSidebarSportGroups(sportID)
	}

	return response
}

func GetSidebarSportGroupLeaguesBlocking(sportID string, groupID string) map[string]interface{} {

	response := GetSidebarSportGroupLeagues(sportID, groupID)

	for i := 0; i < settings.NumOfTries; i++ {

		if response != nil {
			break
		}

		log.Println("Merkurxtip: Stuck on GetSidebarSportGroupLeagues(sportID=" + sportID + ", groupID=" + groupID + ")...")
		response = GetSidebarSportGroupLeagues(sportID, groupID)
	}

	return response
}

func GetMatchIDsBlocking(sportID string, leagueID string) map[string]interface{} {

	response := GetMatchIDs(sportID, leagueID)

	for i := 0; i < settings.NumOfTries; i++ {

		if response != nil {
			break
		}

		log.Println("Merkurxtip: Stuck on GetMatchIDs(sportID=" + sportID + ", leagueID=" + leagueID + ")...")
		response = GetMatchIDs(sportID, leagueID)
	}

	return response
}
