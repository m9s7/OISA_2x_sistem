package arbitrage

import (
	"OISA_2x_sistem/utility"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func FindArb(records [][]string) []Arb {

	startTime := time.Now()
	fmt.Println("... finding arbitrage opportunities\n-------------------------")

	var arbs []Arb

	x := colIdxByNameMap(records[0])
	// kick_off league_merkurxtip league_maxbet league_soccerbet league_mozzart 1 2
	//tip1 tip1_merkurxtip tip1_maxbet tip1_soccerbet tip1_mozzart
	//tip2 tip2_merkurxtip tip2_maxbet tip2_soccerbet tip2_mozzart

	for idx := range records[1:] {
		record := records[idx+1]

		relativeTip1Idx, tip1Max := getMaxTipIdxAndVal(record[x["tip1"]+1 : x["tip2"]])
		relativeTip2Idx, tip2Max := getMaxTipIdxAndVal(record[x["tip2"]+1 : len(x)])
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
		kickoff, err := strconv.Atoi(record[x["kick_off"]])
		if err != nil {
			log.Println("Kickoff string conversion to int failed! Kickoff string:", record[x["kick_off"]])
		}

		// Export
		a := Arb{
			Kickoff: kickoff,
			League:  record[x["kick_off"]+1],
			Team1:   record[x["1"]],
			Team2:   record[x["2"]],

			Tip1:             record[x["tip1"]],
			Bookie1:          strings.Split(records[0][x["tip1"]+relativeTip1Idx], "_")[1],
			Tip1Value:        tip1Max,
			StakePercentage1: IAP1 / outlay,

			Tip2:             record[x["tip2"]],
			Bookie2:          strings.Split(records[0][x["tip2"]+relativeTip2Idx], "_")[1],
			Tip2Value:        tip2Max,
			StakePercentage2: IAP2 / outlay,

			ROI: utility.ToFixed(100*((1/outlay)-1), 2),
		}
		fmt.Printf(ArbToString(a, "IDK"))
		arbs = append(arbs, a)
	}

	fmt.Println("--- ", time.Now().Sub(startTime), " ---")
	return arbs
}

func colIdxByNameMap(columns []string) map[string]int {
	colIndxByName := map[string]int{}
	for i, columnName := range columns {
		colIndxByName[columnName] = i
	}
	return colIndxByName
}

func getMaxTipIdxAndVal(tipVals []string) (int, float64) {
	maxTip := 0.0
	idx := 0
	for i := range tipVals {
		value, err := strconv.ParseFloat(tipVals[i], 32)
		if err != nil {
			continue
		}
		value = utility.ToFixed(value, 2)
		if maxTip < value {
			maxTip = value
			idx = i
		}
	}
	return idx + 1, maxTip
}