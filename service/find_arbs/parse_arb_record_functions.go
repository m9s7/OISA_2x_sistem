package find_arbs

import (
	"OISA_2x_sistem/utility"
	"log"
	"math"
	"strconv"
	"strings"
)

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

func getLeagueNameByBookie(colName []string, leagueNames []string) map[string]string {

	leagueNameAtBookie := map[string]string{}
	for i := range colName {
		bookie := strings.Split(colName[i], "_")[1]
		leagueNameAtBookie[bookie] = leagueNames[i]
	}
	return leagueNameAtBookie

}

func getTipValLabels(mergedRecordsLabelCol []string, colIdxByNameMap map[string]int) []string {
	var tipLabels []string
	for _, val := range mergedRecordsLabelCol[colIdxByNameMap["tip1"]+1 : colIdxByNameMap["tip2"]] {
		bookie := strings.TrimPrefix(val, "tip1_")
		label := bookieToLaBL(bookie)

		tipLabels = append(tipLabels, label)
	}
	return tipLabels
}

func bookieToLaBL(bookie string) string {
	switch bookie {
	case "merkurxtip":
		return "MRxT"
	case "maxbet":
		return "MAXB"
	case "soccerbet":
		return "SOCC"
	case "mozzart":
		return "MOZZ"
	default:
		panic("Unknown bookie: " + bookie)
	}
}

func tipDeviation(tipIdx int, tipVals []float64) float64 {
	deviation := 0.0
	tipVal := tipVals[tipIdx]

	for i := range tipVals {
		if i == tipIdx {
			continue
		}

		deviation += math.Abs(tipVal - tipVals[i])
	}
	return deviation / float64(len(tipVals)-1)
}

func parseTipValsAndGetMaxValIdx(tipVals []string) ([]float64, int) {

	tipValsParsed := make([]float64, len(tipVals))

	maxTip := 0.0
	idx := 0
	for i := range tipVals {
		value, err := strconv.ParseFloat(tipVals[i], 32)
		if err != nil {
			tipValsParsed[i] = 0.0
			continue
		}
		value = utility.ToFixed(value, 2)
		tipValsParsed[i] = value

		if maxTip < value {
			maxTip = value
			idx = i
		}
	}

	return tipValsParsed, idx
}
