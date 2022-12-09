package requests_to_server

import "log"

func GetSidebarSportsBlocking() map[string]interface{} {
	response := GetSidebarSports()
	for response == nil {
		log.Println("Merkurxtip: Stuck on GetSidebarSports()...")
		response = GetSidebarSports()
	}
	return response
}

func GetAllSubgamesBlocking() map[string]interface{} {
	response := GetAllSubgames()
	for response == nil {
		log.Println("Merkurxtip: Stuck on GetAllSubgames()...")
		response = GetAllSubgames()
	}
	return response
}

func GetSidebarSportGroupsBlocking(sportID string) map[string]interface{} {
	response := GetSidebarSportGroups(sportID)
	for response == nil {
		log.Println("Merkurxtip: Stuck on GetSidebarSportGroups(sportID=" + sportID + ")...")
		response = GetSidebarSportGroups(sportID)
	}
	return response
}

func GetSidebarSportGroupLeaguesBlocking(sportID string, groupID string) map[string]interface{} {
	response := GetSidebarSportGroupLeagues(sportID, groupID)
	for response == nil {
		log.Println("Merkurxtip: Stuck on GetSidebarSportGroupLeagues(sportID=" + sportID + ", groupID=" + groupID + ")...")
		response = GetSidebarSportGroupLeagues(sportID, groupID)
	}
	return response
}

func GetMatchIDsBlocking(sportID string, leagueID string) map[string]interface{} {
	response := GetMatchIDs(sportID, leagueID)
	for response == nil {
		log.Println("Merkurxtip: Stuck on GetMatchIDs(sportID=" + sportID + ", leagueID=" + leagueID + ")...")
		response = GetMatchIDs(sportID, leagueID)
	}
	return response
}
