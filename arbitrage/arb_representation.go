package arbitrage

import (
	"OISA_2x_sistem/utility"
	"fmt"
	"strings"
)

func (a Arb) ToStringFree() string {
	separatorLen := utility.Min(len(a.Team1)+len(a.Team2)+4, 33)
	separator := strings.Repeat("=", separatorLen)
	return strings.Join([]string{
		"```",
		strings.ToUpper(a.Sport) + ", " + strings.ToUpper(a.League),
		a.Team1 + " vs " + a.Team2,
		separator,
		strings.ToUpper(a.Tip1),
		"kvota: " + fmt.Sprintf("%.2f", a.Tip1Vals[a.Tip1MaxIdx]) + " @ " + strings.ToUpper(a.Bookie1),
		"ulog = ukupno * " + fmt.Sprintf("%.3f", a.StakePercentage1),
		separator,
		strings.ToUpper(a.Tip2),
		"kvota: " + fmt.Sprintf("%.2f", a.Tip2Vals[a.Tip2MaxIdx]) + " @ " + strings.ToUpper(a.Bookie2),
		"ulog = ukupno * " + fmt.Sprintf("%.3f", a.StakePercentage2),
		separator,
		"ROI: " + fmt.Sprintf("%.2f", a.ROI) + "%",
		"```",
	}, "\n")
}

func (a Arb) ToStringPremium() string {
	separatorLen := utility.Min(len(a.Team1)+len(a.Team2)+4, 33)
	separator := strings.Repeat("=", separatorLen)
	return strings.Join([]string{
		"```",
		strings.ToUpper(a.Sport) + ", " + strings.ToUpper(a.League),
		a.Team1 + " vs " + a.Team2,
		separator,
		strings.ToUpper(a.Tip1),
		"kvota: " + fmt.Sprintf("%.2f", a.Tip1Vals[a.Tip1MaxIdx]) + " @ " + strings.ToUpper(a.Bookie1),
		"ulog = ukupno * " + fmt.Sprintf("%.3f", a.StakePercentage1),
		separator,
		strings.ToUpper(a.Tip2),
		"kvota: " + fmt.Sprintf("%.2f", a.Tip2Vals[a.Tip2MaxIdx]) + " @ " + strings.ToUpper(a.Bookie2),
		"ulog = ukupno * " + fmt.Sprintf("%.3f", a.StakePercentage2),
		separator,
		"Play first @ " + a.PlayFirst,
		"ROI: " + fmt.Sprintf("%.2f", a.ROI) + "%",
		"```",
	}, "\n")
}

func (a Arb) ToStringWithExtra() string {
	separatorLen := utility.Min(len(a.Team1)+len(a.Team2)+4, 33)
	separator := strings.Repeat("=", separatorLen)
	return strings.Join([]string{
		"```",
		strings.ToUpper(a.Sport) + ", " + strings.ToUpper(a.League),
		a.Team1 + " vs " + a.Team2,
		separator,
		strings.ToUpper(a.Tip1),
		"kvota: " + fmt.Sprintf("%.2f", a.Tip1Vals[a.Tip1MaxIdx]) + " @ " + strings.ToUpper(a.Bookie1),
		"ulog = ukupno * " + fmt.Sprintf("%.3f", a.StakePercentage1),
		separator,
		strings.ToUpper(a.Tip2),
		"kvota: " + fmt.Sprintf("%.2f", a.Tip2Vals[a.Tip2MaxIdx]) + " @ " + strings.ToUpper(a.Bookie2),
		"ulog = ukupno * " + fmt.Sprintf("%.3f", a.StakePercentage2),
		separator,
		"ROI: " + fmt.Sprintf("%.2f", a.ROI) + "%\n",

		"Tip1 deviation table:\n" + a.getTip1DeviationTable(),
		"Tip2 deviation table:\n" + a.getTip2DeviationTable(),
		"Play first @ " + a.PlayFirst,

		"```",
	}, "\n")
}
