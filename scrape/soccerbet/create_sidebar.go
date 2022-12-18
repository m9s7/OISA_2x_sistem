package soccerbet

import (
	"OISA_2x_sistem/requests_to_server/soccerbet"
	"OISA_2x_sistem/scrape/soccerbet/server_response_parsers"
	"OISA_2x_sistem/utility"
	"fmt"
)

func createSidebar(masterData *soccerbet.MasterData, sportNameByIDMap map[int]string) map[string][]soccerbet.CompetitionMasterData {
	response, err := soccerbet.GetSidebarLeagueIDs()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	leagueIDs := server_response_parsers.ParseGetSidebarLeagueIDs(response)

	sidebar := map[string][]soccerbet.CompetitionMasterData{}

	for _, league := range masterData.CompetitionsData.Competitions {

		if utility.IsElInSliceINT(league.Id, leagueIDs) {

			sportName := sportNameByIDMap[league.SportId]

			sidebar[sportName] = append(sidebar[sportName], league)
		}
	}
	return sidebar
}
