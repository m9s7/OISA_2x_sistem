package mozzart

import (
	"OISA_2x_sistem/requests_to_server/mozzart"
	"OISA_2x_sistem/scrape/mozzart/odds_parsers"
	"OISA_2x_sistem/scrape/mozzart/server_response_parsers"
	"OISA_2x_sistem/scrape/mozzart/standardization"
	"OISA_2x_sistem/utility"
	"fmt"
	"time"
)

func GetSportsCurrentlyOffered() []string {
	response, err := mozzart.GetSidebar()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var sports []string
	for _, sport := range response {
		sports = append(sports, sport.Name)
	}
	return sports
}

func Scrape(sport string) []*[8]string {
	startTime := time.Now()

	response, _ := mozzart.GetSidebar()

	sportIDByName := server_response_parsers.GetSportIDByNameMap(response)

	allSubgamesResponse, _ := mozzart.GetAllSubgames()
	for response == nil {
		fmt.Println("Mozzart: Stuck on GetAllSubgames()...")
		allSubgamesResponse, _ = mozzart.GetAllSubgames()
	}

	if _, ok := sportIDByName[sport]; !ok {
		fmt.Println(sport, " not currently offered at mozzart")
		return nil
	}

	fmt.Println("...scraping mozzart - ", sport)
	var odds []*[8]string

	switch sport {
	case utility.Tennis:
		odds = odds_parsers.TennisOddsParser(sportIDByName[sport], allSubgamesResponse)
	case utility.Basketball:
		odds = odds_parsers.BasketballOddsParser(sportIDByName[sport], allSubgamesResponse)
	case utility.Soccer:
		odds = odds_parsers.SoccerOddsParser(sportIDByName[sport], allSubgamesResponse)
	default:
		panic("Sport offered at mozzart, but I dont offer it, why am I trying to scrape it?")
	}

	standardization.StandardizeData(odds, sport)

	fmt.Printf("--- %s seconds ---\n", time.Since(startTime))
	return odds
}
