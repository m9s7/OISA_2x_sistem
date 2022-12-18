package odds_parsers

import (
	"OISA_2x_sistem/requests_to_server/merkurxtip"
	"OISA_2x_sistem/utility"
	"fmt"
	"strconv"
	"strings"
)

func SoccerOddsParser(matchIDs []int) []*[8]string {

	matchesScrapedCounter := 0
	var export []*[8]string

	tipTypeCodePairs, leftovers := getHardcodedTipTypeCodes()

	for _, matchID := range matchIDs {

		match, err := merkurxtip.GetMatchOdds(matchID)
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

		for _, m := range leftovers {
			tip2Val, ok := getTipValByTipTypeCode[m["matchingTipTypeCode"]]
			if !ok {
				continue
			}
			export = append(export, utility.MergeE1E2(e1, &[4]string{
				m["tip1Name"], "0.0",
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
