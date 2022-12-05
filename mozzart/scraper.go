package mozzart

import (
	"OISA_2x_sistem/mozzart/odds_parsers"
	"OISA_2x_sistem/mozzart/requests_to_server"
	"OISA_2x_sistem/mozzart/server_response_parsers"
	"fmt"
	"time"
)

func GetSportsCurrentlyOffered() []string {
	response := requests_to_server.GetSidebarSportsAndLeaguesBlocking()

	var sports []string
	for _, val := range response {
		sportName := val["name"].(string)
		sports = append(sports, sportName)
	}
	return sports
}

func Scrape(sport string) [][8]string {
	startTime := time.Now()

	response := requests_to_server.GetSidebarSportsAndLeaguesBlocking()

	getIDByNameMap := server_response_parsers.ParseGetSidebarSportsAndLeagues(response)

	allSubgamesResponse := requests_to_server.GetAllSubgames()
	for response == nil {
		fmt.Println("Mozzart: Stuck on GetAllSubgames()...")
		allSubgamesResponse = requests_to_server.GetAllSubgames()
	}

	if _, ok := getIDByNameMap[sport]; !ok {
		fmt.Println(sport, " not currently offered at mozzart")
		return nil
	}

	fmt.Println("...scraping mozzart - ", sport)
	var odds [][8]string

	switch sport {
	case "Tenis":
		odds = odds_parsers.TennisOddsParser(getIDByNameMap[sport], allSubgamesResponse)
	case "Košarka":
		odds = odds_parsers.BasketballOddsParser(getIDByNameMap[sport], allSubgamesResponse)
	case "Fudbal":
		odds = odds_parsers.SoccerOddsParser(getIDByNameMap[sport], allSubgamesResponse)
	default:
		panic("Sport offered at maxbet, but I dont offer it, why am I trying to scrape it?")
	}

	for _, el := range odds {
		fmt.Println(el)
	}
	//
	//standardize_tip_name = get_standardization_func_4_tip_names(sport)
	//col_name = df.columns[ExportIDX.KICKOFF]
	//df[col_name] = df[col_name].map(standardize_kickoff_time_string)
	//col_name = df.columns[ExportIDX.TIP1_NAME]
	//df[col_name] = df[col_name].map(standardize_tip_name)
	//col_name = df.columns[ExportIDX.TIP2_NAME]
	//df[col_name] = df[col_name].map(standardize_tip_name)
	//
	//print_to_file(df.to_string(index=False), f"mozz_{str(sport.toStandardName())}.txt")
	//export_for_merge(df, f"mozz_{str(sport.toStandardName())}.txt")
	fmt.Printf("--- %s seconds ---", time.Since(startTime))
	return odds
}
