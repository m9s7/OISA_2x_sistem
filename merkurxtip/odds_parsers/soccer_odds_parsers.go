package odds_parsers

import (
	"OISA_2x_sistem/merkurxtip/requests_to_server"
	"OISA_2x_sistem/utility"
	"fmt"
)

func SoccerOddsParser(matchIDs []int) []*[8]string {

	matchesScrapedCounter := 0
	var export []*[8]string

	tipTypeCodePairs, leftovers := getHardcodedTipTypeCodes()

	for _, matchID := range matchIDs {
		match := requests_to_server.GetMatchOdds(matchID)

		e1 := &[4]string{
			fmt.Sprintf("%.0f", match["kickOffTime"].(float64)),
			match["leagueName"].(string),
			match["home"].(string),
			match["away"].(string),
		}

		getTipValByTipTypeCode := match["odds"].(map[string]interface{})
		for tip1Code, m := range tipTypeCodePairs {

			tip1Val, ok := getTipValByTipTypeCode[tip1Code]
			if !ok {
				continue
			}
			tip2Val, ok := getTipValByTipTypeCode[m["matchingTipTypeCode"]]
			if !ok {
				continue
			}

			export = append(export, utility.MergeE1E2(e1, &[4]string{
				m["tip1Name"], fmt.Sprintf("%.2f", tip1Val.(float64)),
				m["tip2Name"], fmt.Sprintf("%.2f", tip2Val.(float64)),
			}))
		}

		for _, m := range leftovers {
			tip2Val, ok := getTipValByTipTypeCode[m["matchingTipTypeCode"]]
			if !ok {
				continue
			}
			export = append(export, utility.MergeE1E2(e1, &[4]string{
				m["tip1Name"], "0.0",
				m["tip2Name"], fmt.Sprintf("%.2f", tip2Val.(float64)),
			}))
		}

		matchesScrapedCounter++
	}

	fmt.Println("Matches scraped: ", matchesScrapedCounter)
	fmt.Println("Tips scraped: ", len(export))

	return export
}
