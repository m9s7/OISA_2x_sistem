package arbitrage

type Arb struct {
	Kickoff     int
	Sport       string
	League      string
	LeagueNames map[string]string

	Team1          string
	Team2          string
	TipValueLabels []string

	Tip1       string
	Tip1Vals   []float64
	Tip1MaxIdx int

	Bookie1          string
	StakePercentage1 float64

	Tip2       string
	Tip2Vals   []float64
	Tip2MaxIdx int

	Bookie2          string
	StakePercentage2 float64

	PlayFirst string
	ROI       float64
}

func (a Arb) Equals(b Arb) bool {
	epsilon := 0.001
	return a.Sport == b.Sport &&
		a.Team1 == b.Team1 &&
		a.Team2 == b.Team2 &&
		// Compare tip1
		a.Tip1 == b.Tip1 &&
		a.Bookie1 == b.Bookie1 &&
		a.Tip1Vals[a.Tip1MaxIdx]-b.Tip1Vals[b.Tip1MaxIdx] < epsilon &&
		// Compare tip2
		a.Tip2 == b.Tip2 &&
		a.Bookie2 == b.Bookie2 &&
		a.Tip2Vals[a.Tip2MaxIdx]-b.Tip2Vals[b.Tip2MaxIdx] < epsilon
}

func (a Arb) IsIn(oldArbs []Arb) bool {
	for _, oldArb := range oldArbs {
		if a.Equals(oldArb) {
			return true
		}
	}
	return false
}

//func GetExampleArbitrage() Arb {
//	return Arb{
//		Kickoff: 63984384684,
//		League:  "NBA",
//		Team1:   "Milwaukee",
//		Team2:   "Sacramento",
//
//		Tip1:             "ki_1_w/ot",
//		Bookie1:          "soccerbet",
//		Tip1Value:        1.47,
//		StakePercentage1: 0.6826845654,
//
//		Tip2:             "ki_2_w/ot",
//		Bookie2:          "mozz",
//		Tip2Value:        3.15,
//		StakePercentage2: 0.3173154346,
//
//		ROI: 0.23,
//	}
//}
