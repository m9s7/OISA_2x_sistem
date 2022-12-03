package maxbet

import (
	"OISA_2x_sistem/maxbet/odds_parsers"
	"OISA_2x_sistem/maxbet/requests_to_server"
	"OISA_2x_sistem/utility"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func GetSportsCurrentlyOffered() []string {
	response := requests_to_server.GetSidebar()
	for response == nil {
		fmt.Println("Stuck on getting maxbet sidebar info")
		response = requests_to_server.GetSidebar()
	}
	var sportsInSidebar []string
	for _, sport := range response {
		sportsInSidebar = append(sportsInSidebar, sport["name"].(string))
	}
	return sportsInSidebar
}

// Scrape arg: "Fudbal", "Košarka"
func Scrape(sport string) {
	startTime := time.Now()

	response := requests_to_server.GetSidebar()
	for response == nil {
		fmt.Println("Stuck on getting maxbet sidebar info")
		response = requests_to_server.GetSidebar()
	}
	sidebar := parseSidebar(response)

	_, ok := sidebar[sport]
	if !ok {
		fmt.Println(sport, " not currently offered at maxbet")
		return
	}
	fmt.Println("...scraping maxb - ", sport)

	responseJson := requests_to_server.GetMatchIds(sidebar[sport])
	for responseJson == nil {
		responseJson = requests_to_server.GetMatchIds(sidebar[sport])
	}
	matchIDs := parseGetMatchesIDsResponse(responseJson)

	var scrapedData [][8]string
	if sport == "Fudbal" {
		scrapedData = odds_parsers.GetSoccerOdds(matchIDs)
	}
	if sport == "Košarka" {
		scrapedData = odds_parsers.Get2outcomeOdds(matchIDs, []string{"Konačan ishod sa produžecima"})
	}
	if sport == "Tenis" {
		scrapedData = odds_parsers.Get2outcomeOdds(matchIDs, []string{"Konačan ishod", "Prvi set", "Drugi set", "Tie Break", "Tie Break prvi set"})
	}
	if sport == "eSport" {
		scrapedData = odds_parsers.Get2outcomeOdds(matchIDs, []string{"Konačan ishod"})
	}
	if sport == "Stoni Tenis" {
		scrapedData = odds_parsers.Get2outcomeOdds(matchIDs, []string{"Konačan ishod", "Prvi set"})
	}

	// i onda jos ovo...
	// Standardize tip names
	// print_to_file(df.to_string(index=False), f"maxb_{str(sport_standard_name)}.txt")
	// export_for_merge(df, f"maxb_{str(sport_standard_name)}.txt")

	writeCSVtoFile(scrapedData, "maxb_"+sport)
	log.Printf("--- %s seconds ---", time.Since(startTime))
}

// Returns a dict { key = sport_name, val = [LeagueBetId] }
func parseSidebar(sidebarSportsJSON []map[string]interface{}) map[string][]int {
	sports := make(map[string][]int)
	for _, sport := range sidebarSportsJSON {

		var leagueBetIds []int

		for _, leagueDict := range sport["leagues"].([]interface{}) {
			leagueDict := leagueDict.(map[string]interface{})
			if strings.HasPrefix(leagueDict["name"].(string), "Max Bonus Tip") {
				continue
			}
			leagueBetIds = append(leagueBetIds, int(leagueDict["betLeagueId"].(float64)))
		}

		if len(leagueBetIds) != 0 {
			sports[sport["name"].(string)] = leagueBetIds
		}
	}
	return sports
}

func parseGetMatchesIDsResponse(response []map[string]interface{}) []int {
	var matchIDs []int
	for _, league := range response {
		for _, match := range league["matchList"].([]interface{}) {
			matchIDs = append(matchIDs, int(match.(map[string]interface{})["id"].(float64)))
		}
	}
	return matchIDs
}

func writeCSVtoFile(export [][8]string, sportName string) {
	file, err := os.Create("C:\\Users\\Matija\\GolandProjects\\OISA_2x_sistem\\IO\\dfs_for_import\\" + sportName + ".csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(file)
	w := csv.NewWriter(file)
	defer w.Flush()

	if err := w.Write(utility.ScrapeColsNames[:]); err != nil {
		log.Fatalln("error writing column names record to file", err)
	}

	//err = w.WriteAll(export)
	for _, e := range export {
		if err := w.Write(e[:]); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
	//if err != nil {
	//	log.Fatalln("error writing records to file", err)
	//}
}
