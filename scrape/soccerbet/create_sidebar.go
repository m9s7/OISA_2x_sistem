package soccerbet

import (
	"OISA_2x_sistem/scrape/soccerbet/requests_to_server"
	"OISA_2x_sistem/scrape/soccerbet/server_response_parsers"
	"OISA_2x_sistem/utility"
)

func createSidebar(masterData map[string]interface{}, sportNameByIDMap map[int]string) map[string][]interface{} {
	response := requests_to_server.GetSidebarLeagueIDsBlocking()
	leagueIDs := server_response_parsers.ParseGetSidebarLeagueIDs(response)

	sidebar := map[string][]interface{}{}

	competitionsData := masterData["CompetitionsData"].(map[string]interface{})
	competitions := competitionsData["Competitions"].([]interface{})

	for _, league := range competitions {
		league := league.(map[string]interface{})
		leagueID := int(league["Id"].(float64))
		if utility.IsElInSliceINT(leagueID, leagueIDs) {
			sportName := sportNameByIDMap[int(league["SportId"].(float64))]

			var sidebarKeys []string
			for key := range sidebar {
				sidebarKeys = append(sidebarKeys, key)
			}
			if !utility.IsElInSliceSTR(sportName, sidebarKeys) {
				sidebar[sportName] = []interface{}{}
			}
			sidebar[sportName] = append(sidebar[sportName], league)
		}
	}
	return sidebar
}