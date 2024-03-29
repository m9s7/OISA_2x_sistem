package odds_parsers

import (
	"OISA_2x_sistem/requests_to_server/merkurxtip"
	"OISA_2x_sistem/utility"
	"fmt"
	"strconv"
	"strings"
)

func TennisOddsParser(matchIDs []int) []*[8]string {

	matchesScrapedCounter := 0
	var export []*[8]string

	// We'll see how hard-coding goes I doubt they change allSubgamesJSON too often if ever
	// maybe I can make a call each time and just check if something changed
	// but the json is pretty big and I don't know how efficient that can be
	tipTypeCodePairs := map[string]map[string]string{
		"1": {
			"tip1Name":            "KI_1",
			"tip2Name":            "KI_2",
			"matchingTipTypeCode": "3",
		},
		"50510": {
			"tip1Name":            "S1_1",
			"tip2Name":            "S1_2",
			"matchingTipTypeCode": "50511",
		},
		"50512": {
			"tip1Name":            "S2_1",
			"tip2Name":            "S2_2",
			"matchingTipTypeCode": "50513",
		},
		"50528": {
			"tip1Name":            "TIE_BREAK_YES",
			"tip2Name":            "TIE_BREAK_NO",
			"matchingTipTypeCode": "50529",
		},
	}

	for _, matchID := range matchIDs {
		match, err := merkurxtip.GetMatchOdds(matchID)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if err != nil {
			fmt.Println("Merkurxtip: GetMatchOdds(matchID:" + strconv.Itoa(matchID) + ") is None, skipping it..")
			continue
		}

		e1 := &[4]string{
			fmt.Sprintf("%.0f", match.KickOffTime),
			match.LeagueName,
			match.Home,
			match.Away,
		}

		getTipValByTipTypeCode := match.Odds
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
				m["tip1Name"], fmt.Sprintf("%.2f", tip1Val),
				m["tip2Name"], fmt.Sprintf("%.2f", tip2Val),
			}))
		}
		matchesScrapedCounter++
	}

	fmt.Println("@MERKURXTIP" + strings.Repeat("-", 26-len("@MERKURXTIP")))
	fmt.Println("Matches scraped: ", matchesScrapedCounter)
	fmt.Println("Tips scraped: ", len(export))

	return export
}
