package odds_parsers

import (
	"OISA_2x_sistem/mozzart/requests_to_server"
	"OISA_2x_sistem/utility"
	"fmt"
	"strconv"
	"strings"
)

func SoccerOddsParser(sportID int, allSubgamesResponse map[string]interface{}) []*[8]string {

	matchesScrapedCounter := 0
	var export []*[8]string

	matchesResponse := requests_to_server.GetMatchIDsBlocking(sportID)
	exportHelp := initExportHelp(matchesResponse["matches"].([]interface{}))

	var exportHelpKeys []int
	for k := range exportHelp {
		exportHelpKeys = append(exportHelpKeys, k)
	}

	subgameIDs := getSoccerSubgameIDs(allSubgamesResponse[strconv.Itoa(sportID)])
	//fmt.Println(subgameIDs)
	odds := requests_to_server.GetOddsBlocking(exportHelpKeys, subgameIDs)

	for _, o := range odds {

		if _, ok := o["kodds"]; !ok {
			continue
		}
		matchID := int(o["id"].(float64))

		// Collect all tips I'm interested in
		tip1, tip2 := collectFocusedSoccerTips(o)

		// Match collected tips
		for t1Key, t1Val := range tip1 {
			t1Game := t1Key[0]
			t1Subgame := t1Key[1]

			if strings.HasPrefix(t1Subgame, "0-") && len(t1Subgame) == 3 {
				x, _ := strconv.Atoi(string(t1Subgame[2]))
				t2Subgame := strconv.Itoa(x+1) + "+"
				if t2Val, ok := tip2[[2]string{t1Game, t2Subgame}]; ok {
					export = append(export, utility.MergeE1E2(exportHelp[matchID], &[4]string{
						strings.Join(t1Key[:], " "), t1Val,
						strings.Join([]string{t1Game, t2Subgame}, " "), t2Val,
					}))
				}
				continue
			}

			if t2Val, ok := tip2[[2]string{t1Game, "1+"}]; t1Subgame == "0" && ok {
				export = append(export, utility.MergeE1E2(exportHelp[matchID], &[4]string{
					strings.Join(t1Key[:], " "), t1Val,
					strings.Join([]string{t1Game, "1+"}, " "), t2Val,
				}))
				continue
			}

			if t2Val, ok := tip2[[2]string{"tgg", "ng"}]; t1Subgame == "gg" && ok {
				export = append(export, utility.MergeE1E2(exportHelp[matchID], &[4]string{"gg", t1Val, "ng", t2Val}))
				continue
			}
			if t2Val, ok := tip2[[2]string{"tgg", "1ng"}]; t1Subgame == "1gg" && ok {
				export = append(export, utility.MergeE1E2(exportHelp[matchID], &[4]string{"1gg", t1Val, "1ng", t2Val}))
				continue
			}
			if t2Val, ok := tip2[[2]string{"tgg", "2ng"}]; t1Subgame == "2gg" && ok {
				export = append(export, utility.MergeE1E2(exportHelp[matchID], &[4]string{"2gg", t1Val, "2ng", t2Val}))
				continue
			}

		}

		matchesScrapedCounter++
	}

	fmt.Println("Matches scraped: ", matchesScrapedCounter)
	fmt.Println("Tips scraped: ", len(export))
	return export
}

func ugConditionSatisfied(subgame string) bool {
	return strings.HasPrefix(subgame, "0-") ||
		strings.HasPrefix(subgame, "1-") ||
		strings.HasSuffix(subgame, "+") ||
		subgame == "0"
}

func getSoccerSubgameIDs(offers interface{}) []int {
	var focusedSubgames []int

	for _, offer := range offers.([]interface{}) {
		offer := offer.(map[string]interface{})

		if offerName, ok := offer["name"]; !ok || offerName != "Kompletna ponuda" {
			continue
		}

		for _, header := range offer["regularHeaders"].([]interface{}) {
			header := header.(map[string]interface{})

			game := header["gameName"].([]interface{})[0].(map[string]interface{})["name"].(string)

			wantedGames := []string{
				"Ukupno golova na meču", "Tim 1 daje gol", "Tim 2 daje gol",
				"Ukupno golova prvo poluvreme", "Tim 1 golovi prvo poluvreme", "Tim 2 golovi prvo poluvreme",
				"Ukupno golova drugo poluvreme", "Tim 1 golovi drugo poluvreme", "Tim 2 golovi drugo poluvreme",
			}
			if utility.IsElInSliceSTR(game, wantedGames) {
				for _, subgame := range header["subGameName"].([]interface{}) {
					subgame := subgame.(map[string]interface{})
					if ugConditionSatisfied(subgame["name"].(string)) {
						focusedSubgames = append(focusedSubgames, int(subgame["id"].(float64)))
					}
				}
			}

			wantedGames = []string{
				"Tačan broj golova na meču",
				"Tačan broj golova prvo poluvreme",
				"Tačan broj golova drugo poluvreme",
			}
			if utility.IsElInSliceSTR(game, wantedGames) {
				for _, subgame := range header["subGameName"].([]interface{}) {
					subgame := subgame.(map[string]interface{})
					if subgame["name"].(string) == "0" {
						focusedSubgames = append(focusedSubgames, int(subgame["id"].(float64)))
					}
				}
			}

			wantedSubgames := []string{"gg", "ng", "1gg", "1ng", "2gg", "2ng"}
			if game == "Oba tima daju gol" {
				for _, subgame := range header["subGameName"].([]interface{}) {
					subgame := subgame.(map[string]interface{})
					if utility.IsElInSliceSTR(subgame["name"].(string), wantedSubgames) {
						focusedSubgames = append(focusedSubgames, int(subgame["id"].(float64)))
					}
				}
			}

		}
	}

	return utility.RemoveDuplicates(&focusedSubgames)
}

func collectFocusedSoccerTips(odds map[string]interface{}) (map[[2]string]string, map[[2]string]string) {

	tip1 := map[[2]string]string{}
	tip2 := map[[2]string]string{}

	for _, sg := range odds["kodds"].(map[string]interface{}) {
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
		key := [2]string{game, subgame}
		val := sg["value"].(string)

		focusedSubgames := []string{
			"ug", "1ug", "2ug", "tm1", "tm2",
			"1tm2", "1tm1", "2tm2", "2tm1",
			"tgg",
		}
		if utility.IsElInSliceSTR(game, focusedSubgames) {
			if strings.HasPrefix(subgame, "0-") || subgame == "0" || strings.HasSuffix(subgame, "gg") {
				tip1[key] = val
			}
			if strings.HasSuffix(subgame, "+") || strings.HasPrefix(subgame, "1-") || strings.HasSuffix(subgame, "ng") {
				tip2[key] = val
			}
		}
	}

	return tip1, tip2
}
