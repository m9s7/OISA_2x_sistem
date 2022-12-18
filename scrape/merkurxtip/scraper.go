package merkurxtip

import (
	"OISA_2x_sistem/requests_to_server/merkurxtip"
	"OISA_2x_sistem/scrape/merkurxtip/odds_parsers"
	"OISA_2x_sistem/scrape/merkurxtip/server_response_parsers"
	"OISA_2x_sistem/scrape/merkurxtip/standardization"
	"OISA_2x_sistem/utility"
	"fmt"
	"time"
)

func GetSportsCurrentlyOffered() []string {

	sidebarSportsIDsByName := server_response_parsers.GetSidebarSportsIDsByName()

	var sports []string
	for sport := range sidebarSportsIDsByName {
		sports = append(sports, sport)
	}
	return sports
}

func getMatchIDs(sport string) []int {

	var matchIDs []int

	sidebarSportIDByName := server_response_parsers.GetSidebarSportsIDsByName()
	sportID := sidebarSportIDByName[sport]

	response, err := merkurxtip.GetSidebarSportGroups(sportID)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	groupIDs := server_response_parsers.ParseGetSidebarGroups(response)

	for _, groupID := range groupIDs {
		getMatchIDsResponse, err := merkurxtip.GetMatchIDs(sportID, groupID)
		if err != nil {
			fmt.Println(err)
			continue
		}

		groupMatchIDs := server_response_parsers.ParseGetMatchIDs(getMatchIDsResponse)
		matchIDs = append(matchIDs, groupMatchIDs...)
	}

	return matchIDs
}

func Scrape(sport string) []*[8]string {
	startTime := time.Now()
	fmt.Println("...scraping merkurxtip - ", sport)

	// Don't need it all subgames are hardcoded
	//allSubgames := requests_to_server.GetAllSubgamesBlocking()
	
	matchIDs := getMatchIDs(sport)

	var odds []*[8]string

	switch sport {
	case utility.Tennis:
		odds = odds_parsers.TennisOddsParser(matchIDs)
	case utility.Basketball:
		odds = odds_parsers.BasketballOddsParser(matchIDs)
	case utility.Soccer:
		odds = odds_parsers.SoccerOddsParser(matchIDs)
	default:
		panic("Sport offered at merkurxtip, but I dont offer it, why am I trying to scrape it?")
	}

	standardization.StandardizeData(odds, sport)

	fmt.Printf("--- %s seconds ---\n", time.Since(startTime))
	return odds
}
