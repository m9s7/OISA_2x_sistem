package soccerbet

import (
	"OISA_2x_sistem/soccerbet/odds_parsers"
	"OISA_2x_sistem/soccerbet/requests_to_server"
	"OISA_2x_sistem/soccerbet/server_response_parsers"
	"OISA_2x_sistem/utility"
	"fmt"
	"time"
)

func GetSportsCurrentlyOffered() []string {
	masterData := requests_to_server.GetMasterData()
	for masterData == nil {
		fmt.Println("Stuck on soccerbet request: GetMasterData")
		masterData = requests_to_server.GetMasterData()
	}

	sportNameByIDMap := server_response_parsers.GetSportNameByIDMap(masterData)
	sidebar := createSidebar(masterData, sportNameByIDMap)

	var sidebarKeys []string
	for key := range sidebar {
		sidebarKeys = append(sidebarKeys, key)
	}
	return sidebarKeys
}

func Scrape(sport string) {
	startTime := time.Now()
	fmt.Println("...scraping soccerbet - ", sport)

	masterData := requests_to_server.GetMasterData()
	for masterData == nil {
		fmt.Println("Stuck on soccerbet request: GetMasterData")
		masterData = requests_to_server.GetMasterData()
	}

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
		return
	}

	var odds [][8]string

	switch sport {
	case "Tenis":
		odds = odds_parsers.TennisOddsParser(sidebar[sport], betgameByIdMap, betgameOutcomeByIdMap, betgameGroupByIdMap)
	case "Košarka":
		odds = odds_parsers.BasketballOddsParser(sidebar[sport], betgameByIdMap, betgameOutcomeByIdMap, betgameGroupByIdMap)
	case "Fudbal":
		odds = odds_parsers.SoccerOddsParser(sidebar[sport], betgameByIdMap, betgameOutcomeByIdMap, betgameGroupByIdMap)
	default:
		panic("Sport offered at maxbet, but I dont offer it, why am I trying to scrape it?")
	}
	
	for _, el := range odds {
		fmt.Println(el)
	}

	fmt.Printf("--- %s seconds ---", time.Since(startTime))
}
