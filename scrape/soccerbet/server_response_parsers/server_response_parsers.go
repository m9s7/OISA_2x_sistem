package server_response_parsers

import "OISA_2x_sistem/requests_to_server/soccerbet"

func ParseGetSidebarLeagueIDs(response []soccerbet.SidebarLeagueId) []int {
	var leagueIDs []int
	for _, r := range response {
		leagueIDs = append(leagueIDs, r.CompetitionId)
	}
	return leagueIDs
}
