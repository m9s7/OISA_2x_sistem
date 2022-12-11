package find_arb

import (
	"OISA_2x_sistem/telegram"
	"OISA_2x_sistem/utility"
	"fmt"
	"strings"
)

func ArbToString(a Arb, sport string) string {
	separatorLen := utility.Min(len(a.Team1)+len(a.Team2)+4, 33)
	separator := strings.Repeat("=", separatorLen)
	return strings.Join([]string{
		"```",
		strings.ToUpper(sport) + ", " + strings.ToUpper(a.League),
		a.Team1 + " vs " + a.Team2,
		separator,
		strings.ToUpper(a.Tip1),
		"kvota: " + fmt.Sprintf("%.2f", a.Tip1Value) + " @ " + strings.ToUpper(a.Bookie1),
		"ulog = ukupno * " + fmt.Sprintf("%.3f", a.StakePercentage1),
		separator,
		strings.ToUpper(a.Tip2),
		"kvota: " + fmt.Sprintf("%.2f", a.Tip2Value) + " @ " + strings.ToUpper(a.Bookie2),
		"ulog = ukupno * " + fmt.Sprintf("%.3f", a.StakePercentage2),
		separator,
		"Play first @ " + a.PlayFirst,
		"ROI: " + fmt.Sprintf("%.2f", a.ROI) + "%",
		"```",
	}, "\n")
}

func FreeArbToString(a Arb, sport string) string {
	separatorLen := utility.Min(len(a.Team1)+len(a.Team2)+4, 33)
	separator := strings.Repeat("=", separatorLen)
	return strings.Join([]string{
		"```",
		strings.ToUpper(sport) + ", " + strings.ToUpper(a.League),
		a.Team1 + " vs " + a.Team2,
		separator,
		strings.ToUpper(a.Tip1),
		"kvota: " + fmt.Sprintf("%.2f", a.Tip1Value) + " @ " + strings.ToUpper(a.Bookie1),
		"ulog = ukupno * " + fmt.Sprintf("%.3f", a.StakePercentage1),
		separator,
		strings.ToUpper(a.Tip2),
		"kvota: " + fmt.Sprintf("%.2f", a.Tip2Value) + " @ " + strings.ToUpper(a.Bookie2),
		"ulog = ukupno * " + fmt.Sprintf("%.3f", a.StakePercentage2),
		separator,
		"ROI: " + fmt.Sprintf("%.2f", a.ROI) + "%",
		"```",
	}, "\n")
}

func PremiumArbToString(a Arb, sport string) string {
	separatorLen := utility.Min(len(a.Team1)+len(a.Team2)+4, 33)
	separator := strings.Repeat("=", separatorLen)
	return strings.Join([]string{
		"```",
		strings.ToUpper(sport) + ", " + strings.ToUpper(a.League),
		a.Team1 + " vs " + a.Team2,
		separator,
		strings.ToUpper(a.Tip1),
		"kvota: " + fmt.Sprintf("%.2f", a.Tip1Value) + " @ " + strings.ToUpper(a.Bookie1),
		"ulog = ukupno * " + fmt.Sprintf("%.3f", a.StakePercentage1),
		separator,
		strings.ToUpper(a.Tip2),
		"kvota: " + fmt.Sprintf("%.2f", a.Tip2Value) + " @ " + strings.ToUpper(a.Bookie2),
		"ulog = ukupno * " + fmt.Sprintf("%.3f", a.StakePercentage2),
		separator,
		"ROI: " + fmt.Sprintf("%.2f", a.ROI) + "%\n",

		"Tip1 deviation table:\n" + a.tip1DeviationTable,
		"Tip2 deviation table:\n" + a.tip2DeviationTable,
		"Play first @ " + a.PlayFirst,

		"```",
	}, "\n")
}

func getTipDeviationTable(vals []string, labels []string, index int) string {

	header := make([]string, len(labels))
	body := make([]string, len(vals))

	for i := range labels {
		header[i] = " " + labels[i] + " "
		body[i] = fmt.Sprintf("%-6s", " "+vals[i])
	}
	labelsString := strings.Join(header, "|")
	valuesString := strings.Join(body, "|")
	arrowString := strings.Repeat(" ", index*7) + " ^^^^"

	return strings.Join([]string{labelsString, valuesString, arrowString}, "\n")
}

func BroadcastNewArbs(arbs []Arb, oldArbs map[string][]Arb, sport string, chatIDs []string) {
	if len(arbs) == 0 {
		//telegram.BroadcastToDev(`Nema arbe :\\( \\- ` + sport)
		oldArbs[sport] = nil
		return
	}
	for _, arb := range arbs {
		if isArbInOldArbs(arb, oldArbs[sport]) || arb.ROI < 0.1 {
			continue
		}

		if arb.ROI <= 1.5 {
			telegram.BroadcastToFree(FreeArbToString(arb, sport))
		}
		if arb.ROI >= 1.0 {
			telegram.BroadcastToPremium(PremiumArbToString(arb, sport))
		}
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

func bookieToLabel(bookie string) string {
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
