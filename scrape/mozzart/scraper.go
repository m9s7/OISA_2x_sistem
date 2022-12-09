package mozzart

import (
	odds_parsers2 "OISA_2x_sistem/scrape/mozzart/odds_parsers"
	requests_to_server2 "OISA_2x_sistem/scrape/mozzart/requests_to_server"
	"OISA_2x_sistem/scrape/mozzart/server_response_parsers"
	"OISA_2x_sistem/scrape/mozzart/standardization"
	"OISA_2x_sistem/utility"
	"fmt"
	"time"
)

func GetSportsCurrentlyOffered() []string {
	response := requests_to_server2.GetSidebarSportsAndLeaguesBlocking()

	var sports []string
	for _, val := range response {
		sportName := val["name"].(string)
		sports = append(sports, sportName)
	}
	return sports
}

func Scrape(sport string) []*[8]string {
	startTime := time.Now()

	response := requests_to_server2.GetSidebarSportsAndLeaguesBlocking()

	getIDByNameMap := server_response_parsers.ParseGetSidebarSportsAndLeagues(response)

	allSubgamesResponse := requests_to_server2.GetAllSubgames()
	for response == nil {
		fmt.Println("Mozzart: Stuck on GetAllSubgames()...")
		allSubgamesResponse = requests_to_server2.GetAllSubgames()
	}

	if _, ok := getIDByNameMap[sport]; !ok {
		fmt.Println(sport, " not currently offered at mozzart")
		return nil
	}

	fmt.Println("...scraping mozzart - ", sport)
	var odds []*[8]string

	switch sport {
	case utility.Tennis:
		odds = odds_parsers2.TennisOddsParser(getIDByNameMap[sport], allSubgamesResponse)
	case utility.Basketball:
		odds = odds_parsers2.BasketballOddsParser(getIDByNameMap[sport], allSubgamesResponse)
	case utility.Soccer:
		odds = odds_parsers2.SoccerOddsParser(getIDByNameMap[sport], allSubgamesResponse)
	default:
		panic("Sport offered at maxbet, but I dont offer it, why am I trying to scrape it?")
	}

	standardization.StandardizeData(odds, sport)

	fmt.Printf("--- %s seconds ---\n", time.Since(startTime))
	return odds
}
