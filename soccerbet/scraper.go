package soccerbet

import (
	"OISA_2x_sistem/soccerbet/odds_parsers"
	"OISA_2x_sistem/soccerbet/requests_to_server"
	"OISA_2x_sistem/soccerbet/server_response_parsers"
	"OISA_2x_sistem/soccerbet/standardization"
	"OISA_2x_sistem/utility"
	"fmt"
	"time"
)

func GetSportsCurrentlyOffered() []string {
	masterData := requests_to_server.GetMasterDataBlocking()

	sportNameByIDMap := server_response_parsers.GetSportNameByIDMap(masterData)
	sidebar := createSidebar(masterData, sportNameByIDMap)

	var sidebarKeys []string
	for key := range sidebar {
		sidebarKeys = append(sidebarKeys, key)
	}
	return sidebarKeys
}

func Scrape(sport string) []*[8]string {
	startTime := time.Now()
	fmt.Println("...scraping soccerbet - ", sport)

	masterData := requests_to_server.GetMasterDataBlocking()

	betgameByIdMap := server_response_parsers.GetBetgameByIdMap(masterData)
	betgameOutcomeByIdMap := server_response_parsers.GetBetgameOutcomeByIdMap(masterData)
	betgameGroupByIdMap := server_response_parsers.GetBetgameGroupByIdMap(masterData)
	sportNameByIDMap := server_response_parsers.GetSportNameByIDMap(masterData)

	sidebar := createSidebar(masterData, sportNameByIDMap)

	var sidebarKeys []string
	for key := range sidebar {
		sidebarKeys = append(sidebarKeys, key)
	}
	if !utility.IsElInSliceSTR(sport, sidebarKeys) {
		fmt.Println(sport, "not currently offered at soccerbet")
		return nil
	}

	var odds []*[8]string

	switch sport {
	case utility.Tennis:
		odds = odds_parsers.TennisOddsParser(sidebar[sport], betgameByIdMap, betgameOutcomeByIdMap, betgameGroupByIdMap)
	case utility.Basketball:
		odds = odds_parsers.BasketballOddsParser(sidebar[sport], betgameByIdMap, betgameOutcomeByIdMap, betgameGroupByIdMap)
	case utility.Soccer:
		odds = odds_parsers.SoccerOddsParser(sidebar[sport], betgameByIdMap, betgameOutcomeByIdMap, betgameGroupByIdMap)
	default:
		panic("Sport offered at maxbet, but I dont offer it, why am I trying to scrape it?")
	}

	standardization.StandardizeData(odds, sport)

	fmt.Printf("--- %s seconds ---\n", time.Since(startTime))
	return odds
}
