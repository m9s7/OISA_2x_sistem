package maxbet

import (
	"OISA_2x_sistem/requests_to_server/maxbet"
	"strings"
)

func createSidebar(response []maxbet.SidebarSport) map[string][]int {

	sports := make(map[string][]int)
	for _, sport := range response {

		var leagueBetIds []int
		for _, league := range sport.Leagues {

			if strings.HasPrefix(league.Name, "Max Bonus Tip") {
				continue
			}
			if league.Active == false || league.Blocked == true {
				continue
			}
			leagueBetIds = append(leagueBetIds, league.BetLeagueId)

		}

		if len(leagueBetIds) != 0 {
			sports[sport.Name] = leagueBetIds
		}

	}
	
	return sports
}
