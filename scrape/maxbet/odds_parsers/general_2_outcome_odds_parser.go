package odds_parsers

import (
	"OISA_2x_sistem/scrape/maxbet/requests_to_server"
	"OISA_2x_sistem/utility"
	"fmt"
	"strings"
)

func Get2outcomeOdds(matchIDs []int, subgameNames []string) []*[8]string {
	matchesScrapedCounter := 0
	var export []*[8]string

	for _, matchID := range matchIDs {
		match := requests_to_server.GetMatchData(matchID)
		if match == nil {
			continue
		}

		e1 := &[4]string{
			fmt.Sprintf("%.0f", match["kickOffTime"].(float64)),
			match["leagueName"].(string),
			match["home"].(string),
			match["away"].(string),
		}

		for _, subgame := range match["odBetPickGroups"].([]interface{}) {
			subgame := subgame.(map[string]interface{})
			if subgameName, ok := subgame["name"]; !ok || !utility.IsElInSliceSTR(subgameName.(string), subgameNames) {
				continue
			}
			if len(subgame["tipTypes"].([]interface{})) != 2 {
				continue
			}
			TT := subgame["tipTypes"].([]interface{})
			TT1 := TT[0].(map[string]interface{})
			TT2 := TT[1].(map[string]interface{})

			if TT1["value"].(float64) == 0 && TT2["value"].(float64) == 0 {
				continue
			}
			export = append(export, utility.MergeE1E2(e1, &[4]string{
				TT1["tipType"].(string),
				fmt.Sprintf("%.2f", TT1["value"].(float64)),
				TT2["tipType"].(string),
				fmt.Sprintf("%.2f", TT2["value"].(float64)),
			}))
		}
		matchesScrapedCounter++
	}
	fmt.Println("@MAXBET" + strings.Repeat("-", 26-len("@MAXBET")))
	fmt.Println("Matches scraped: ", matchesScrapedCounter)
	fmt.Println("Tips scraped: ", len(export))

	return export
}
