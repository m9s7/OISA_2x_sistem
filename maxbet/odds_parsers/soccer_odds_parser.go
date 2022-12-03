package odds_parsers

import (
	"OISA_2x_sistem/maxbet/requests_to_server"
	"OISA_2x_sistem/utility"
	"fmt"
	"log"
	"strconv"
	"strings"
)

var goloviSubgames = map[string]map[string]interface{}{
	"Ukupno golova 90'": {
		"prefix":       "ug ",
		"tip1_length":  6,
		"tip1_special": "ug T0",
	},
	"Ukupno golova prvo poluvreme": {
		"prefix":       "ug 1P",
		"tip1_length":  8,
		"tip1_special": "ug 1PT0",
	},
	"Ukupno golova drugo poluvreme": {
		"prefix":       "ug 2P",
		"tip1_length":  8,
		"tip1_special": "ug 2PT0",
	},
	"Domaćin golovi": {
		"prefix":       "D",
		"tip1_length":  4,
		"tip1_special": "D0",
	},
	"Domaćin golovi prvo poluvreme": {
		"prefix":       "1D",
		"tip1_length":  5,
		"tip1_special": "1D0",
	},
	"Domaćin golovi drugo poluvreme": {
		"prefix":       "2D",
		"tip1_length":  5,
		"tip1_special": "2D0",
	},
	"Gost golovi": {
		"prefix":       "G",
		"tip1_length":  4,
		"tip1_special": "G0",
	},
	"Gost golovi prvo poluvreme": {
		"prefix":       "1G",
		"tip1_length":  5,
		"tip1_special": "1G0",
	},
	"Gost golovi drugo poluvreme": {
		"prefix":       "2G",
		"tip1_length":  5,
		"tip1_special": "2G0",
	},
}

func GetSoccerOdds(matchIDs []int) [][8]string {
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

			// Process NG GG
			if subgame["name"].(string) == "GG - NG " {
				var tipsKeys []string
				tipsVals := map[string]float64{}
				for _, tip := range subgame["tipTypes"].([]interface{}) {
					tip := tip.(map[string]interface{})
					tipsVals[tip["name"].(string)] = tip["value"].(float64)
					tipsKeys = append(tipsKeys, tip["name"].(string))
				}
				tipCombos := [3][2]string{{"GG", "NG"}, {"GG1", "NG1"}, {"GG2", "NG2"}}
				for _, combo := range tipCombos {
					if !utility.IsInSlice(combo[0], tipsKeys) || !utility.IsInSlice(combo[1], tipsKeys) {
						continue
					}
					e[utility.Tip1Name] = combo[0]
					e[utility.Tip1Value] = fmt.Sprintf("%f", tipsVals[combo[0]])
					e[utility.Tip2Name] = combo[1]
					e[utility.Tip2Value] = fmt.Sprintf("%f", tipsVals[combo[1]])
					export = append(export, e)
				}
				continue
			}

			// Check if subgame is in ug subgames we are interested in
			goloviSubgamesKeys := make([]string, len(goloviSubgames))
			i := 0
			for k := range goloviSubgames {
				goloviSubgamesKeys[i] = k
				i++
			}
			if !utility.IsInSlice(subgame["name"].(string), goloviSubgamesKeys) {
				continue
			}

			// Preprocess for UG subgames
			tips := getAllSubgameTips(subgame)
			subgameName := subgame["name"].(string)

			// Process 0-x and x+ combinations
			for _, tip1 := range tips {
				if !isOXtip(tip1, subgameName) {
					continue
				}
				tip2 := makeXPlusTipFromOXTip(tip1, subgameName)

				indexOfTip2InTips := utility.IndexOf(tip2, tips)
				if indexOfTip2InTips == -1 {
					continue
				}

				TT := subgame["tipTypes"].([]interface{})
				tip1Value := TT[utility.IndexOf(tip1, tips)].(map[string]interface{})["value"]
				tip2Value := TT[indexOfTip2InTips].(map[string]interface{})["value"]

				e[utility.Tip1Name] = tip1
				e[utility.Tip1Value] = fmt.Sprintf("%f", tip1Value.(float64))
				e[utility.Tip2Name] = tip2
				e[utility.Tip2Value] = fmt.Sprintf("%f", tip2Value.(float64))
				export = append(export, e)
			}

			// Process T0 and 1+ combo
			s := goloviSubgames[subgame["name"].(string)]
			tip1 := s["tip1_special"].(string)
			tip2 := s["prefix"].(string) + "1+"

			tip1IndexInTips := utility.IndexOf(tip1, tips)
			if tip1IndexInTips == -1 {
				continue
			}
			TT := subgame["tipTypes"].([]interface{})
			tip1Value := TT[tip1IndexInTips].(map[string]interface{})["value"]

			// get tip2 value
			tip2IndexInTips := utility.IndexOf(tip2, tips)
			if tip1Value == 0 || tip2IndexInTips == -1 {
				continue
			}
			tip2Value := TT[tip2IndexInTips].(map[string]interface{})["value"]
			e[utility.Tip1Name] = tip1
			e[utility.Tip1Value] = fmt.Sprintf("%f", tip1Value.(float64))
			e[utility.Tip2Name] = tip2
			e[utility.Tip2Value] = fmt.Sprintf("%f", tip2Value.(float64))

			export = append(export, e)
		}
	}
	fmt.Println("Matches scraped: ", matchesScrapedCounter)
	return export
}

func getAllSubgameTips(subgame map[string]interface{}) []string {
	tips := make([]string, len(subgame["tipTypes"].([]interface{})))
	for i, tip := range subgame["tipTypes"].([]interface{}) {
		tip := tip.(map[string]interface{})
		tips[i] = tip["name"].(string)
	}
	return tips
}

func isOXtip(tip string, subgameName string) bool {
	if len(tip) != goloviSubgames[subgameName]["tip1_length"] {
		return false
	}
	if !strings.HasPrefix(tip, goloviSubgames[subgameName]["prefix"].(string)+"0-") {
		return false
	}
	return true
}

func parseXFromOXTip(tip string, subgameName string) (int, error) {
	return strconv.Atoi(string(tip[len(goloviSubgames[subgameName]["prefix"].(string))+2]))
}

func makeXPlusTipFromOXTip(tip string, subgameName string) string {
	x, err := parseXFromOXTip(tip, subgameName)
	if err != nil {
		log.Fatalln(err)
	}
	return goloviSubgames[subgameName]["prefix"].(string) + strconv.Itoa(x+1) + "+"
}
