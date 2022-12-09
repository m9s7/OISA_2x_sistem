package odds_parsers

import (
	"OISA_2x_sistem/utility"
	"fmt"
	"strings"
)

func getIDsForSubgameNames(offers interface{}, subgameNames []string) []int {
	var subgameIDs []int

	for _, offer := range offers.([]interface{}) {
		offer := offer.(map[string]interface{})
		if offer["name"].(string) != "Kompletna ponuda" {
			continue
		}
		for _, header := range offer["regularHeaders"].([]interface{}) {
			header := header.(map[string]interface{})
			game := header["gameName"].([]interface{})[0].(map[string]interface{})
			if utility.IsElInSliceSTR(game["shortName"].(string), subgameNames) {
				for _, subgame := range header["subGameName"].([]interface{}) {
					subgame := subgame.(map[string]interface{})
					subgameIDs = append(subgameIDs, int(subgame["id"].(float64)))
				}
			}
		}
	}
	return subgameIDs
}

func initExportHelp(matchesResponse []interface{}) map[int]*[4]string {

	export := map[int]*[4]string{}
	for _, match := range matchesResponse {
		match := match.(map[string]interface{})
		participants := match["participants"].([]interface{})

		if int(match["specialType"].(float64)) != 0 || len(participants) != 2 {
			continue
		}

		matchID := int(match["id"].(float64))
		export[matchID] = &[4]string{
			fmt.Sprintf("%.0f", match["startTime"].(float64)),
			match["competition_name_sr"].(string),
			strings.Trim(participants[0].(map[string]interface{})["name"].(string), " "),
			strings.Trim(participants[1].(map[string]interface{})["name"].(string), " "),
		}
	}
	return export
}
