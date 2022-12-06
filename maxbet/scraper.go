package maxbet

import (
	"OISA_2x_sistem/maxbet/odds_parsers"
	"OISA_2x_sistem/maxbet/requests_to_server"
	"OISA_2x_sistem/maxbet/server_response_parsers"
	"OISA_2x_sistem/maxbet/standardization"
	"fmt"
	"time"
)

func GetSportsCurrentlyOffered() []string {
	response := requests_to_server.GetSidebarBlocking()

	var sportsInSidebar []string
	for _, sport := range response {
		sportsInSidebar = append(sportsInSidebar, sport["name"].(string))
	}
	return sportsInSidebar
}

func Scrape(sport string) []*[8]string {
	startTime := time.Now()

	response := requests_to_server.GetSidebarBlocking()
	sidebar := createSidebar(response)

	leagueIDs, ok := sidebar[sport]
	if !ok {
		fmt.Println(sport, " not currently offered at maxbet")
		return nil
	}
	fmt.Println("...scraping maxb - ", sport)

	response = requests_to_server.GetMatchIDsBlocking(leagueIDs)
	matchIDs := server_response_parsers.ParseGetMatchesIDsResponse(response)

	var odds []*[8]string

	switch sport {
	case "Tenis":
		odds = odds_parsers.Get2outcomeOdds(matchIDs, []string{"Konačan ishod", "Prvi set", "Drugi set", "Tie Break", "Tie Break prvi set", "Tie Break drugi set"})
	case "Košarka":
		odds = odds_parsers.Get2outcomeOdds(matchIDs, []string{"Konačan ishod sa produžecima"})
	case "Fudbal":
		odds = odds_parsers.GetSoccerOdds(matchIDs)
	default:
		panic("Sport offered at maxbet, but I dont offer it, why am I trying to scrape it?")
	}

	standardization.StandardizeData(odds, sport)

	fmt.Printf("--- %s seconds ---\n", time.Since(startTime))
	return odds
}
