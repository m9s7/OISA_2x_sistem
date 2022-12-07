package odds_parsers

import (
	"OISA_2x_sistem/soccerbet/requests_to_server"
	"OISA_2x_sistem/soccerbet/server_response_parsers"
	"OISA_2x_sistem/utility"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var goloviSubgames = []string{"Ukupno Golova", "I Pol. Uk. Golova", "II Pol. Uk. Golova",
	"Domaćin Ukupno Golova", "I Pol. Domaćin Uk. Golova",
	"II Pol. Domaćin Uk. Golova",
	"Gost Ukupno Golova", "I Pol. Gost Uk. Golova", "II Pol. Gost Uk. Golova"}

func SoccerOddsParser(
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
			tip1, tip2, err := parseMatchTips(matchID, betgameByIdMap, betgameOutcomeByIdMap, betgameGroupByIdMap)
			if err != nil {
				continue
			}

			// Match tips
			for t1Key, t1Val := range tip1 {
				t1Game := t1Key[0]
				t1Subgame := t1Key[1]
				t1Code := t1Key[2]

				// Process T0 and 1+ combo
				if t1Subgame == "0" {
					for t2Key, t2Val := range tip2 {
						t2Game := t2Key[0]
						t2Subgame := t2Key[1]
						t2Code := t2Key[2]

						if t2Game == t1Game && t2Subgame == "1+" {
							e2 := &[4]string{
								t1Code,
								fmt.Sprintf("%.2f", t1Val),
								t2Code,
								fmt.Sprintf("%.2f", t2Val),
							}
							export = append(export, utility.MergeE1E2(e1, e2))
						}
					}
				}
				// Process 0-x and x+ combos
				if strings.HasPrefix(t1Subgame, "0-") && len(t1Subgame) == 3 {
					x, _ := strconv.Atoi(string(t1Subgame[2]))
					for t2Key, t2Val := range tip2 {
						t2Game := t2Key[0]
						t2Subgame := t2Key[1]
						t2Code := t2Key[2]

						if t2Game == t1Game && t2Subgame == strconv.Itoa(x+1)+"+" {
							e2 := &[4]string{
								t1Code,
								fmt.Sprintf("%.2f", t1Val),
								t2Code,
								fmt.Sprintf("%.2f", t2Val),
							}
							export = append(export, utility.MergeE1E2(e1, e2))
						}
					}
				}
				// Process GG NG combo
				if t1Subgame == "GG" {
					for t2Key, t2Val := range tip2 {
						t2Game := t2Key[0]
						t2Subgame := t2Key[1]
						t2Code := t2Key[2]

						if t2Game == t1Game && t2Subgame == "NG" {
							e2 := &[4]string{
								t1Code,
								fmt.Sprintf("%.2f", t1Val),
								t2Code,
								fmt.Sprintf("%.2f", t2Val),
							}
							export = append(export, utility.MergeE1E2(e1, e2))
						}
					}
				}
			}

			matchesScrapedCounter++
		}
	}

	fmt.Println("Matches scraped: ", matchesScrapedCounter)
	fmt.Println("Tips scraped: ", len(export))

	return export
}

func parseMatchTips(
	matchID int,
	betgameByIdMap map[int]map[string]interface{},
	betgameOutcomeByIdMap map[int]map[string]interface{},
	betgameGroupByIdMap map[int]map[string]interface{}) (map[[3]string]float64, map[[3]string]float64, error) {

	matchOdds := requests_to_server.GetMatchOddsValues(matchID)
	if matchOdds == nil {
		fmt.Println("Soccerbet: GetMatchOddsValues(matchID:" + strconv.Itoa(matchID) + ") is None, skipping it..")
		return nil, nil, errors.New("skipping match")
	}

	tip1 := map[[3]string]float64{}
	tip2 := map[[3]string]float64{}

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

		tipKey := [3]string{
			betgameName,
			outcomeName,
			outcome["CodeForPrinting"].(string),
		}
		tipVal := odds["Odds"].(float64)

		// GG/NG - GG1/2 is never offered
		if betgameGroupName == "OBA TIMA DAJU GOL" && betgameName == "Oba Tima Daju Gol" {
			if outcomeName == "GG" {
				tip1[tipKey] = tipVal
			}
			if outcomeName == "NG" {
				tip2[tipKey] = tipVal
			}
			continue
		}
		// UKUPNO GOLOVA
		if betgameGroupName != "UKUPNO GOLOVA" &&
			betgameGroupName != "DOMAĆIN UK. GOLOVA" &&
			betgameGroupName != "GOST UK. GOLOVA" {
			continue
		}
		if utility.IsElInSliceSTR(betgameName, goloviSubgames) {
			if strings.HasPrefix(outcomeName, "0-") || outcomeName == "0" {
				tip1[tipKey] = tipVal
			} else if strings.HasSuffix(outcomeName, "+") {
				tip2[tipKey] = tipVal
			}
		}
	}
	return tip1, tip2, nil
}
