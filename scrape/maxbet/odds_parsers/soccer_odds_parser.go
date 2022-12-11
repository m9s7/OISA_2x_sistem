package odds_parsers

import (
	"OISA_2x_sistem/scrape/maxbet/requests_to_server"
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

func GetSoccerOdds(matchIDs []int) []*[8]string {
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
				for _, tipCombo := range tipCombos {
					if !utility.IsElInSliceSTR(tipCombo[0], tipsKeys) || !utility.IsElInSliceSTR(tipCombo[1], tipsKeys) {
						continue
					}
					if tipsVals[tipCombo[0]] == 0 && tipsVals[tipCombo[1]] == 0 {
						continue
					}
					export = append(export, utility.MergeE1E2(e1, &[4]string{
						tipCombo[0], fmt.Sprintf("%.2f", tipsVals[tipCombo[0]]),
						tipCombo[1], fmt.Sprintf("%.2f", tipsVals[tipCombo[1]]),
					}))
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
			if !utility.IsElInSliceSTR(subgame["name"].(string), goloviSubgamesKeys) {
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
				tip1Value := TT[utility.IndexOf(tip1, tips)].(map[string]interface{})["value"].(float64)
				tip2Value := TT[indexOfTip2InTips].(map[string]interface{})["value"].(float64)

				if tip1Value == 0 && tip2Value == 0 {
					continue
				}

				export = append(export, utility.MergeE1E2(e1, &[4]string{
					tip1, fmt.Sprintf("%.2f", tip1Value),
					tip2, fmt.Sprintf("%.2f", tip2Value),
				}))
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
			tip1Value := TT[tip1IndexInTips].(map[string]interface{})["value"].(float64)

			// get tip2 value
			tip2IndexInTips := utility.IndexOf(tip2, tips)
			if tip1Value == 0 || tip2IndexInTips == -1 {
				continue
			}
			tip2Value := TT[tip2IndexInTips].(map[string]interface{})["value"].(float64)

			if tip1Value == 0 && tip2Value == 0 {
				continue
			}

			export = append(export, utility.MergeE1E2(e1, &[4]string{
				tip1, fmt.Sprintf("%.2f", tip1Value),
				tip2, fmt.Sprintf("%.2f", tip2Value),
			}))
		}
		matchesScrapedCounter++
	}

	fmt.Println("@MAXBET" + strings.Repeat("-", 26-len("@MAXBET")))
	fmt.Println("Matches scraped: ", matchesScrapedCounter)
	fmt.Println("Tips scraped: ", len(export))

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
