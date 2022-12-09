package standardization

import (
	"OISA_2x_sistem/utility"
	"fmt"
	"strconv"
	"time"
)

func StandardizeData(records []*[8]string, sport string) {
	standardizeTipName := getStandardizationFunc4TipNames(sport)
	for _, record := range records {
		for i := range record {
			record[i] = utility.TrimWhiteSpace(record[i])
		}
		record[utility.Kickoff] = standardizeKickoffTime(record[utility.Kickoff])
		record[utility.Tip1Name] = standardizeTipName(record[utility.Tip1Name])
		record[utility.Tip2Name] = standardizeTipName(record[utility.Tip2Name])
	}
}

func standardizeKickoffTime(kickoff string) string {
	layout := "2006-01-02T15:04:05"
	t, err := time.Parse(layout, kickoff)
	if err != nil {
		fmt.Println(err)
	}
	return strconv.Itoa(int(t.Unix()))
}

func getStandardizationFunc4TipNames(sport string) func(tip string) string {
	switch sport {
	case utility.Basketball:
		return standardizeTipNameBasketball
	case utility.Tennis:
		return standardizeTipNameTennis
	case utility.Soccer:
		return standardizeTipNameSoccer
	default:
		panic("No tip name standardization function for sport: " + sport)
	}
}
