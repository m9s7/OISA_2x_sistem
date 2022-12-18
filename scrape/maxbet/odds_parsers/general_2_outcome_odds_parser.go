package odds_parsers

import (
	"OISA_2x_sistem/requests_to_server/maxbet"
	"OISA_2x_sistem/utility"
	"fmt"
	"strings"
)

func Get2outcomeOdds(matchIDs []int, subgameNames []string) []*[8]string {
	matchesScrapedCounter := 0
	var export []*[8]string

	for _, matchID := range matchIDs {
		match, err := maxbet.GetMatchData(matchID)
		if err != nil {
			fmt.Println(err)
			continue
		}

		e1 := &[4]string{
			fmt.Sprintf("%.0f", match.KickOffTime),
			match.LeagueName,
			match.Home,
			match.Away,
		}

		for _, subgame := range match.OdBetPickGroups {

			if !utility.IsElInSliceSTR(subgame.Name, subgameNames) || len(subgame.TipTypes) != 2 {
				continue
			}

			TT1 := subgame.TipTypes[0]
			TT2 := subgame.TipTypes[1]

			if TT1.Value == 0 && TT2.Value == 0 {
				continue
			}
			export = append(export, utility.MergeE1E2(e1, &[4]string{
				TT1.TipType,
				fmt.Sprintf("%.2f", TT1.Value),
				TT2.TipType,
				fmt.Sprintf("%.2f", TT2.Value),
			}))
		}
		matchesScrapedCounter++
	}
	fmt.Println("@MAXBET" + strings.Repeat("-", 26-len("@MAXBET")))
	fmt.Println("Matches scraped: ", matchesScrapedCounter)
	fmt.Println("Tips scraped: ", len(export))

	return export
}
