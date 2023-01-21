package maxbet

import (
	"OISA_2x_sistem/requests_to_server/maxbet"
	"OISA_2x_sistem/scrape/maxbet/odds_parsers"
	"OISA_2x_sistem/scrape/maxbet/server_response_parsers"
	"OISA_2x_sistem/scrape/maxbet/standardization"
	"OISA_2x_sistem/utility"
	"fmt"
	"time"
)

func GetSportsCurrentlyOffered() []string {
	response, err := maxbet.GetSidebar()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var sportsInSidebar []string
	for _, sport := range response {
		sportsInSidebar = append(sportsInSidebar, sport.Name)
	}
	return sportsInSidebar
}

func Scrape(sport string) []*[8]string {
	startTime := time.Now()

	sidebarResponse, err := maxbet.GetSidebar()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	sidebar := createSidebar(sidebarResponse)

	leagueIDs, ok := sidebar[sport]
	if !ok {
		fmt.Println(sport, " not currently offered at maxbet")
		return nil
	}

	fmt.Println("...scraping maxb - ", sport)

	response, err := maxbet.GetMatchIDs(leagueIDs)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	matchIDs := server_response_parsers.ParseGetMatchesIDsResponse(response)

	var odds []*[8]string

	switch sport {
	case utility.Tennis:
		odds = odds_parsers.Get2outcomeOdds(matchIDs, []string{"Konačan ishod", "Prvi set", "Drugi set", "Tie Break", "Tie Break prvi set", "Tie Break drugi set"})
	case utility.Basketball:
		odds = odds_parsers.Get2outcomeOdds(matchIDs, []string{"Konačan ishod sa produžecima"})
	case utility.Soccer:
		odds = odds_parsers.GetSoccerOdds(matchIDs)
	default:
		panic("Sport offered at maxbet, but I dont offer it, why am I trying to scrape it?")
	}

	standardization.StandardizeData(odds, sport)

	fmt.Printf("--- %s seconds ---\n", time.Since(startTime))
	return odds
}
