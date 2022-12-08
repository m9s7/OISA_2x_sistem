package arbitrage

import (
	"fmt"
	"strings"
)

type Arb struct {
	Kickoff int
	League  string
	Team1   string
	Team2   string

	Tip1             string
	Bookie1          string
	Tip1Value        float64
	StakePercentage1 float64

	Tip2             string
	Bookie2          string
	Tip2Value        float64
	StakePercentage2 float64

	ROI float64
}

func ArbToString(a Arb, sport string) string {
	return strings.Join([]string{
		"```",
		strings.ToUpper(sport) + ", " + strings.ToUpper(a.League),
		a.Team1 + " vs " + a.Team2,
		strings.Repeat("=", len(a.Team1)+len(a.Team2)+4),
		strings.ToUpper(a.Tip1),
		"kvota: " + fmt.Sprintf("%.2f", a.Tip1Value) + " @ " + strings.ToUpper(a.Bookie1),
		"ulog = ukupno * " + fmt.Sprintf("%.3f", a.StakePercentage1),
		strings.Repeat("=", len(a.Team1)+len(a.Team2)+4),
		strings.ToUpper(a.Tip2),
		"kvota: " + fmt.Sprintf("%.2f", a.Tip2Value) + " @ " + strings.ToUpper(a.Bookie2),
		"ulog = ukupno * " + fmt.Sprintf("%.3f", a.StakePercentage2),
		strings.Repeat("=", len(a.Team1)+len(a.Team2)+4),
		"ROI: " + fmt.Sprintf("%.2f", a.ROI) + "%",
		"```",
	}, "\n")
}

func GetExampleArbitrage() Arb {
	return Arb{
		Kickoff: 63984384684,
		League:  "NBA",
		Team1:   "Milwaukee",
		Team2:   "Sacramento",

		Tip1:             "ki_1_w/ot",
		Bookie1:          "soccerbet",
		Tip1Value:        1.47,
		StakePercentage1: 0.6826845654,

		Tip2:             "ki_2_w/ot",
		Bookie2:          "mozz",
		Tip2Value:        3.15,
		StakePercentage2: 0.3173154346,

		ROI: 0.23,
	}
}
