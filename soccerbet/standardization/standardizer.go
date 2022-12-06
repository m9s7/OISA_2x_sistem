package standardization

import (
	"OISA_2x_sistem/utility"
)

func StandardizeData(data []*[8]string, sport string) {
	standardizeTipName := getStandardizationFunc4TipNames(sport)
	for _, row := range data {
		row[utility.Kickoff] = standardizeKickoffTime(row[utility.Kickoff])
		row[utility.Tip1Name] = standardizeTipName(row[utility.Tip1Name])
		row[utility.Tip2Name] = standardizeTipName(row[utility.Tip2Name])
	}
}

func standardizeKickoffTime(kickoff string) string {
	return kickoff[:len(kickoff)-3]
}

func getStandardizationFunc4TipNames(sport string) func(tip string) string {
	switch sport {
	case "Košarka":
		return standardizeTipNameBasketball
	case "Tenis":
		return standardizeTipNameTennis
	case "Fudbal":
		return standardizeTipNameSoccer
	default:
		panic("No tip name standardization function for sport: " + sport)
	}
}
