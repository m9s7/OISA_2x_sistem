package odds_parsers

import (
	"OISA_2x_sistem/scrape/soccerbet/requests_to_server"
	"OISA_2x_sistem/scrape/soccerbet/server_response_parsers"
	"OISA_2x_sistem/utility"
	"fmt"
	"log"
	"strconv"
)

func TennisOddsParser(
	sidebarLeagues []interface{},
	betgameByIdMap map[int]map[string]interface{},
	betgameOutcomeByIdMap map[int]map[string]interface{},
	betgameGroupByIdMap map[int]map[string]interface{},
) []*[8]string {

	matchesScrapedCounter := 0
	var export []*[8]string

	for _, league := range sidebarLeagues {
		league := league.(map[string]interface{})
		leagueID := int(league["Id"].(float64))

		response := requests_to_server.GetLeagueMatchesInfo(leagueID)
		if response == nil {
			fmt.Println("Soccerbet: GetLeagueMatchesInfo(leagueID:" + strconv.Itoa(leagueID) + ") is None, skipping it..")
			continue
		}
		matchInfoList := server_response_parsers.ParseGetLeagueMatchesInfo(response)

		for _, match := range matchInfoList {
			e1 := &[4]string{match["kickoff"].(string), league["Name"].(string), match["home"].(string), match["away"].(string)}

			matchID := int(match["match_id"].(float64))
			matchOdds := requests_to_server.GetMatchOddsValues(matchID)
			if matchOdds == nil {
				fmt.Println("Soccerbet: GetMatchOddsValues(matchID:" + strconv.Itoa(matchID) + ") is None, skipping it..")
				continue
			}
			exportMatchHelper := map[string]*[4]string{}

			for _, odds := range matchOdds {
				if !odds["IsEnabled"].(bool) {
					continue
				}

				outcome := betgameOutcomeByIdMap[int(odds["BetGameOutcomeId"].(float64))]
				betgame := betgameByIdMap[int(outcome["BetGameId"].(float64))]
				betgameGroup := betgameGroupByIdMap[int(betgame["BetGameGroupId"].(float64))]

				betgameName := betgame["Name"].(string)
				betgameGroupName := betgameGroup["Name"].(string)
				outcomeName := outcome["Name"].(string)

				tipComboKey := betgameGroupName + " " + betgameName
				tipVal := odds["Odds"].(float64)

				var exportMatchHelperKeys []string
				for key := range exportMatchHelper {
					exportMatchHelperKeys = append(exportMatchHelperKeys, key)
				}

				// KI
				if betgameGroupName == "MEČ" && betgameName == "Konačni Ishod" {

					if !utility.IsElInSliceSTR(tipComboKey, exportMatchHelperKeys) {
						exportMatchHelper[tipComboKey] = &[4]string{}
					}

					if outcomeName == "1" {
						exportMatchHelper[tipComboKey][0] = outcome["CodeForPrinting"].(string)
						exportMatchHelper[tipComboKey][1] = fmt.Sprintf("%.2f", tipVal)
					} else if outcomeName == "2" {
						exportMatchHelper[tipComboKey][2] = outcome["CodeForPrinting"].(string)
						exportMatchHelper[tipComboKey][3] = fmt.Sprintf("%.2f", tipVal)
					} else {
						log.Fatalln(tipComboKey, outcomeName, outcome["Description"].(string), outcome["CodeForPrinting"].(string), tipVal)
					}
				}

				// KI PRVI SET
				if betgameGroupName == "SET" && betgameName == "I Set" {

					if !utility.IsElInSliceSTR(tipComboKey, exportMatchHelperKeys) {
						exportMatchHelper[tipComboKey] = &[4]string{}
					}

					if outcomeName == "1" {
						exportMatchHelper[tipComboKey][0] = outcome["CodeForPrinting"].(string)
						exportMatchHelper[tipComboKey][1] = fmt.Sprintf("%.2f", tipVal)
					} else if outcomeName == "2" {
						exportMatchHelper[tipComboKey][2] = outcome["CodeForPrinting"].(string)
						exportMatchHelper[tipComboKey][3] = fmt.Sprintf("%.2f", tipVal)
					} else {
						log.Fatalln(tipComboKey, outcomeName, outcome["Description"].(string), outcome["CodeForPrinting"].(string), tipVal)
					}
				}

				// TIE BREAK
				if betgameGroupName == "MEČ" && betgameName == "Tie Break" {

					if !utility.IsElInSliceSTR(tipComboKey, exportMatchHelperKeys) {
						exportMatchHelper[tipComboKey] = &[4]string{}
					}

					if outcomeName == "DA" {
						exportMatchHelper[tipComboKey][0] = outcome["CodeForPrinting"].(string)
						exportMatchHelper[tipComboKey][1] = fmt.Sprintf("%.2f", tipVal)
					} else if outcomeName == "NE" {
						exportMatchHelper[tipComboKey][2] = outcome["CodeForPrinting"].(string)
						exportMatchHelper[tipComboKey][3] = fmt.Sprintf("%.2f", tipVal)
					} else {
						log.Fatalln(tipComboKey, outcomeName, outcome["Description"].(string), outcome["CodeForPrinting"].(string), tipVal)
					}
				}
			}

			for _, e2 := range exportMatchHelper {
				export = append(export, utility.MergeE1E2(e1, e2))
			}
			matchesScrapedCounter++
		}
	}

	fmt.Println("Matches scraped: ", matchesScrapedCounter)
	fmt.Println("Tips scraped: ", len(export))

	return export
}