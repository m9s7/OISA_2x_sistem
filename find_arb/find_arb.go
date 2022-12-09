package find_arb

import (
	"OISA_2x_sistem/telegram"
	"OISA_2x_sistem/utility"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"time"
)

func FindArb(records [][]string) []Arb {

	startTime := time.Now()
	fmt.Println("... finding find_arb opportunities\n-------------------------")

	var arbs []Arb

	x := colIdxByNameMap(records[0])
	// kick_off league_merkurxtip league_maxbet league_soccerbet league_mozzart 1 2
	//tip1 tip1_merkurxtip tip1_maxbet tip1_soccerbet tip1_mozzart
	//tip2 tip2_merkurxtip tip2_maxbet tip2_soccerbet tip2_mozzart

	for idx := range records[1:] {
		record := records[idx+1]
		tip1Vals := record[x["tip1"]+1 : x["tip2"]]
		tip2Vals := record[x["tip2"]+1 : len(x)]

		relativeTip1Idx, tip1Max := getMaxTipIdxAndVal(tip1Vals)
		relativeTip2Idx, tip2Max := getMaxTipIdxAndVal(tip2Vals)
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

		// Export
		a := Arb{
			Kickoff: kickoffStringToInt(record[x["kick_off"]]),
			League:  record[x["kick_off"]+1],
			Team1:   record[x["1"]],
			Team2:   record[x["2"]],

			Tip1:             record[x["tip1"]],
			Bookie1:          bookie1,
			Tip1Value:        tip1Max,
			StakePercentage1: IAP1 / outlay,

			Tip2:             record[x["tip2"]],
			Bookie2:          bookie2,
			Tip2Value:        tip2Max,
			StakePercentage2: IAP2 / outlay,

			PlayFirst: playFirst,
			ROI:       utility.ToFixed(100*((1/outlay)-1), 2),
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

func kickoffStringToInt(kickoffStr string) int {
	kickoff, err := strconv.Atoi(kickoffStr)
	if err != nil {
		log.Println("Kickoff string conversion to int failed! Kickoff string:", kickoffStr)
	}
	return kickoff
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
	return idx, maxTip
}

func tipDeviation(tipIdx int, tipVals []string) float64 {
	deviation := 0.0
	tipVal, err := strconv.ParseFloat(tipVals[tipIdx], 32)
	if err != nil {
		log.Fatalln("tipDeviation: tipVal string conversion to float failed! TipVal string:", tipVals[tipIdx])
	}

	for i := range tipVals {
		if i == tipIdx {
			continue
		}

		value, err := strconv.ParseFloat(tipVals[i], 32)
		if err != nil {
			continue
		}
		value = utility.ToFixed(value, 2)
		deviation += math.Abs(tipVal - value)
	}
	return deviation / float64(len(tipVals)-1)
}

//TODO: move to telegram package

func BroadcastNewArbs(arbs []Arb, oldArbs map[string][]Arb, sport string) {
	if len(arbs) == 0 {
		//telegram.BroadcastToDev(`Nema arbe :\\( \\- ` + sport)
		oldArbs[sport] = nil
		return
	}
	for _, arb := range arbs {
		if isArbInOldArbs(arb, oldArbs[sport]) {
			continue
		}
		telegram.BroadcastToDev("FRISKE ARBE")
		telegram.BroadcastToDev(ArbToString(arb, sport))
	}
	oldArbs[sport] = arbs
}

func isArbInOldArbs(arb Arb, oldArbs []Arb) bool {
	for _, oldArb := range oldArbs {
		if arb.Equals(oldArb) {
			return true
		}
	}
	return false
}
