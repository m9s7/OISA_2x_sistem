package find_arbs

import (
	"OISA_2x_sistem/arbitrage"
	"OISA_2x_sistem/utility"
	"fmt"
	"strings"
	"time"
)

func FindArb(records [][]string, sport string) []arbitrage.Arb {

	startTime := time.Now()
	fmt.Println("... finding arbitrage opportunities\n-------------------------")

	var arbs []arbitrage.Arb

	x := colIdxByNameMap(records[0])
	// kick_off league_merkurxtip league_maxbet league_soccerbet league_mozzart 1 2
	//tip1 tip1_merkurxtip tip1_maxbet tip1_soccerbet tip1_mozzart
	//tip2 tip2_merkurxtip tip2_maxbet tip2_soccerbet tip2_mozzart

	tipLabels := getTipValLabels(records[0], x) // ex. [merkurxtip, maxbet, soccerbet, mozzart]

	for idx := range records[1:] {
		record := records[idx+1]
		tip1ValsAsStr := record[x["tip1"]+1 : x["tip2"]]
		tip2ValsAsStr := record[x["tip2"]+1 : len(x)]

		tip1Vals, relativeTip1Idx := parseTipValsAndGetMaxValIdx(tip1ValsAsStr)
		tip2Vals, relativeTip2Idx := parseTipValsAndGetMaxValIdx(tip2ValsAsStr)

		tip1Max := tip1Vals[relativeTip1Idx]
		tip2Max := tip2Vals[relativeTip2Idx]

		if tip1Max == 0.0 || tip2Max == 0.0 {
			continue
		}

		// Calculate Individual Arbitrage Percentage
		IAP1 := 1 / tip1Max
		IAP2 := 1 / tip2Max
		// Calculate outlay
		outlay := IAP1 + IAP2

		if outlay > 1.0 {
			continue
		}

		// Prep data for export

		bookie1 := strings.Split(records[0][x["tip1"]+relativeTip1Idx+1], "_")[1]
		bookie2 := strings.Split(records[0][x["tip2"]+relativeTip2Idx+1], "_")[1]

		var playFirst string
		tip1Deviation := tipDeviation(relativeTip1Idx, tip1Vals)
		tip2Deviation := tipDeviation(relativeTip2Idx, tip2Vals)

		if tip1Deviation > tip2Deviation {
			playFirst = bookie1
		} else {
			playFirst = bookie2
		}

		// Parse league names
		leagueNameColStartIdx := x["kick_off"] + 1
		leagueNameColEndIdx := x["1"] - 1
		leagueNameAtBookie := getLeagueNameByBookie(
			records[0][leagueNameColStartIdx:leagueNameColEndIdx],
			record[leagueNameColStartIdx:leagueNameColEndIdx],
		)

		// Export
		a := arbitrage.Arb{
			Kickoff:     kickoffStringToInt(record[x["kick_off"]]),
			Sport:       sport,
			League:      record[x["kick_off"]+1],
			LeagueNames: leagueNameAtBookie,

			Team1:          record[x["1"]],
			Team2:          record[x["2"]],
			TipValueLabels: tipLabels,

			Tip1:       record[x["tip1"]],
			Tip1Vals:   tip1Vals,
			Tip1MaxIdx: relativeTip1Idx,

			Bookie1:          bookie1,
			StakePercentage1: IAP1 / outlay,

			Tip2:       record[x["tip2"]],
			Tip2Vals:   tip2Vals,
			Tip2MaxIdx: relativeTip2Idx,

			Bookie2:          bookie2,
			StakePercentage2: IAP2 / outlay,

			PlayFirst: playFirst,
			ROI:       utility.ToFixed(100*((1/outlay)-1), 2),
		}

		fmt.Printf(a.ToStringWithExtra())
		arbs = append(arbs, a)
	}

	fmt.Println("--- ", time.Now().Sub(startTime), " ---")
	return arbs
}
