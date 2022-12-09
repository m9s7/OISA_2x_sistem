package odds_parsers

import (
	"OISA_2x_sistem/scrape/mozzart/requests_to_server"
	"OISA_2x_sistem/utility"
	"fmt"
	"strconv"
)

func BasketballOddsParser(sportID int, allSubgamesResponse map[string]interface{}) []*[8]string {

	matchesScrapedCounter := 0
	var export []*[8]string

	matchesResponse := requests_to_server.GetMatchIDsBlocking(sportID)
	exportHelp := initExportHelp(matchesResponse["matches"].([]interface{}))

	var exportHelpKeys []int
	for k := range exportHelp {
		exportHelpKeys = append(exportHelpKeys, k)
	}

	focusedSubgames := []string{"pobm"}
	subgameIDs := getIDsForSubgameNames(allSubgamesResponse[strconv.Itoa(sportID)], focusedSubgames)
	odds := requests_to_server.GetOddsBlocking(exportHelpKeys, subgameIDs)

	for _, o := range odds {
		if _, ok := o["kodds"]; !ok {
			continue
		}

		matchID := int(o["id"].(float64))
		e1 := exportHelp[matchID]
		exportMatchHelper := map[string]*[4]string{}

		for _, sg := range o["kodds"].(map[string]interface{}) {
			sg, ok := sg.(map[string]interface{})
			if !ok {
				continue
			}
			_, ok = sg["subGame"]
			if !ok {
				continue
			}

			game := sg["subGame"].(map[string]interface{})["gameShortName"].(string)
			subgame := sg["subGame"].(map[string]interface{})["subGameName"].(string)
			val := sg["value"].(string)

			if game == "pobm" {
				var exportMatchHelperKeys []string
				for k := range exportMatchHelper {
					exportMatchHelperKeys = append(exportMatchHelperKeys, k)
				}
				if !utility.IsElInSliceSTR(game, exportMatchHelperKeys) {
					exportMatchHelper[game] = &[4]string{}
				}

				if subgame == "1" {
					exportMatchHelper[game][0] = game + " " + subgame
					exportMatchHelper[game][1] = val
				} else if subgame == "2" {
					exportMatchHelper[game][2] = game + " " + subgame
					exportMatchHelper[game][3] = val
				} else {
					fmt.Println("Mozzart: Two-outcome game with third outcome" + game + subgame + "found, value=" + val)
					continue
				}
			}

		}

		for _, e2 := range exportMatchHelper {
			e := utility.MergeE1E2(e1, e2)
			export = append(export, e)
		}
		matchesScrapedCounter++
	}

	fmt.Println("Matches scraped: ", matchesScrapedCounter)
	fmt.Println("Tips scraped: ", len(export))

	return export
}
