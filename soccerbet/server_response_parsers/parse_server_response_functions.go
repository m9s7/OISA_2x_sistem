package server_response_parsers

func ParseGetSidebarLeagueIDs(response []map[string]int) []int {
	var leagueIDs []int
	for _, r := range response {
		leagueIDs = append(leagueIDs, r["CompetitionId"])
	}
	return leagueIDs
}

func ParseGetLeagueMatchesInfo(response []map[string]interface{}) []map[string]interface{} {
	var matches []map[string]interface{}
	for _, match := range response {
		matches = append(matches, map[string]interface{}{
			"match_id": match["Id"],
			"home":     match["HomeCompetitorName"],
			"away":     match["AwayCompetitorName"],
			"kickoff":  match["StartDate"],
		})
	}
	return matches
}
