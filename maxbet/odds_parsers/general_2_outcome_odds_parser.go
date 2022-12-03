package odds_parsers

import (
	"OISA_2x_sistem/maxbet/requests_to_server"
	"OISA_2x_sistem/utility"
	"fmt"
	"strconv"
)

func Get2outcomeOdds(matchIDs []int, subgameNames []string) [][8]string {
	matchesScrapedCounter := 0
	export := make([][8]string, 0)
	for _, matchID := range matchIDs {
		match := requests_to_server.GetMatchData(matchID)
		if match == nil {
			continue
		}
		matchesScrapedCounter++
		var e [8]string
		e[utility.Kickoff] = strconv.Itoa(int(match["kickOffTime"].(float64)))
		e[utility.League] = match["leagueName"].(string)
		e[utility.Team1] = match["home"].(string)
		e[utility.Team2] = match["away"].(string)

		for _, subgame := range match["odBetPickGroups"].([]interface{}) {
			subgame := subgame.(map[string]interface{})
			if subgameName, ok := subgame["name"]; !ok || !utility.IsInSlice(subgameName.(string), subgameNames) {
				continue
			}
			if len(subgame["tipTypes"].([]interface{})) != 2 {
				continue
			}
			TT := subgame["tipTypes"].([]interface{})
			TT1 := TT[0].(map[string]interface{})
			TT2 := TT[1].(map[string]interface{})

			if TT1["value"] == 0 && TT2["value"] == 0 {
				continue
			}
			e[utility.Tip1Name] = TT1["tipType"].(string)
			e[utility.Tip1Value] = fmt.Sprintf("%f", TT1["value"].(float64))
			e[utility.Tip2Name] = TT2["tipType"].(string)
			e[utility.Tip2Value] = fmt.Sprintf("%f", TT2["value"].(float64))

			//fmt.Printf("%v\n", e[:])
			export = append(export, e)
		}
	}
	fmt.Println("Matches scraped: ", matchesScrapedCounter)
	return export
	// - no need to write to file if everything is happening in go, but I need to standardize it without turning it into DF
	// or try how fast standardization with DF and just going through [][]string is
	// this all needs to be done on soccer because it has a LOT of data and I can really see speed improvements

	//columns := []string{"kick_off", "league", "1", "2", "tip1_name", "tip1_val", "tip2_name", "tip2_val"}
	//df = pd.DataFrame(export, columns=columns)
	//err = dataframe.LoadRecords(export, dataframe.Names(columns...)).WriteCSV(f)
	//if err != nil {
	//	fmt.Println("ERROR WRITING CSV TO FILE")
	//	return
	//}
}
