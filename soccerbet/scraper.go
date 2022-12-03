package soccerbet

import (
	"OISA_2x_sistem/soccerbet/odds_parsers"
	"OISA_2x_sistem/soccerbet/requests_to_server"
	"OISA_2x_sistem/soccerbet/server_response_parsers"
	"OISA_2x_sistem/utility"
	"fmt"
	"log"
	"time"
)

func GetSportsCurrentlyOffered() []string {
	masterData := requests_to_server.GetMasterData()
	for masterData == nil {
		fmt.Println("Stuck on soccerbet request: GetMasterData")
		masterData = requests_to_server.GetMasterData()
	}

	sportNameByIDMap := server_response_parsers.GetSportNameByIDMap(masterData)
	sidebar := CreateSidebar(masterData, sportNameByIDMap)

	var sidebarKeys []string
	for key := range sidebar {
		sidebarKeys = append(sidebarKeys, key)
	}
	return sidebarKeys
}

func Scrape(sportName string) {
	startTime := time.Now()
	fmt.Println("...scraping soccerbet - ", sportName)

	masterData := requests_to_server.GetMasterData()
	for masterData == nil {
		fmt.Println("Stuck on soccerbet request: GetMasterData")
		masterData = requests_to_server.GetMasterData()
	}

	betgameByIdMap := server_response_parsers.GetBetgameByIdMap(masterData)
	betgameOutcomeByIdMap := server_response_parsers.GetBetgameOutcomeByIdMap(masterData)
	betgameGroupByIdMap := server_response_parsers.GetBetgameGroupByIdMap(masterData)

	sportNameByIDMap := server_response_parsers.GetSportNameByIDMap(masterData)
	sidebar := CreateSidebar(masterData, sportNameByIDMap)

	var sidebarKeys []string
	for key := range sidebar {
		sidebarKeys = append(sidebarKeys, key)
	}
	if !utility.IsElInSliceSTR(sportName, sidebarKeys) {
		fmt.Println(sportName, " not currently offered at soccerbet")
		return
	}

	if sportName == "Fudbal" {
		odds := odds_parsers.SoccerOddsParser(sidebar[sportName], betgameByIdMap, betgameOutcomeByIdMap, betgameGroupByIdMap)
		for _, el := range odds {
			fmt.Println(el)
		}
	}

	log.Printf("--- %s seconds ---", time.Since(startTime))
}
