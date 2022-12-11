package merkurxtip

import (
	"OISA_2x_sistem/scrape/merkurxtip/odds_parsers"
	"OISA_2x_sistem/scrape/merkurxtip/requests_to_server"
	"OISA_2x_sistem/scrape/merkurxtip/server_response_parsers"
	"OISA_2x_sistem/scrape/merkurxtip/standardization"
	"OISA_2x_sistem/utility"
	"fmt"
	"time"
)

func GetSportsCurrentlyOffered() []string {
	response := requests_to_server.GetSidebarSportsBlocking()
	sidebarSportsIDsByName := server_response_parsers.ParseGetSidebarSports(response)

	var sports []string
	for sport := range sidebarSportsIDsByName {
		sports = append(sports, sport)
	}
	return sports
}

func getMatchIDs(sport string) []int {

	response := requests_to_server.GetSidebarSportsBlocking()
	sidebarSportIDByName := server_response_parsers.ParseGetSidebarSports(response)
	sportID := sidebarSportIDByName[sport]

	response = requests_to_server.GetSidebarSportGroupsBlocking(sportID)
	groupIDs := server_response_parsers.ParseGetSidebarGroups(response)

	var leagueIDs []string
	for _, groupID := range groupIDs {
		response = requests_to_server.GetSidebarSportGroupLeaguesBlocking(sportID, groupID)
		groupLeagueIDs := server_response_parsers.ParseGetSidebarSportGroupLeagues(response)
		leagueIDs = append(leagueIDs, groupLeagueIDs...)
	}

	var matchIDs []int
	for _, leagueID := range leagueIDs {
		response = requests_to_server.GetMatchIDsBlocking(sportID, leagueID)
		leagueMatchIDs := server_response_parsers.ParseGetMatchIDs(response)
		matchIDs = append(matchIDs, leagueMatchIDs...)
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
