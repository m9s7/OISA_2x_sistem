package server_response_parsers

import "OISA_2x_sistem/requests_to_server/maxbet"

func ParseGetMatchesIDsResponse(response []maxbet.LeagueMatches) []int {

	var matchIDs []int

	for _, league := range response {

		for _, match := range league.MatchList {
			matchIDs = append(matchIDs, match.Id)
		}

	}
	return matchIDs
}
