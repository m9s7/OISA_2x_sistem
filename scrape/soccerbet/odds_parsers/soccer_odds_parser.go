package odds_parsers

import (
	"OISA_2x_sistem/requests_to_server/soccerbet"
	"OISA_2x_sistem/utility"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var goloviSubgames = []string{"Ukupno Golova", "I Pol. Uk. Golova", "II Pol. Uk. Golova",
	"Domaćin Ukupno Golova", "I Pol. Domaćin Uk. Golova",
	"II Pol. Domaćin Uk. Golova",
	"Gost Ukupno Golova", "I Pol. Gost Uk. Golova", "II Pol. Gost Uk. Golova"}

func SoccerOddsParser(
	sidebarLeagues []soccerbet.CompetitionMasterData,
	betgameByIdMap map[int]*soccerbet.Betgame,
	betgameOutcomeByIdMap map[int]*soccerbet.BetgameOutcome,
	betgameGroupByIdMap map[int]*soccerbet.BetgameGroup,
) []*[8]string {

	matchesScrapedCounter := 0
	var export []*[8]string

	for _, league := range sidebarLeagues {

		matchesInfo, err := soccerbet.GetLeagueMatchesInfo(league.Id)
		if err != nil {
			fmt.Println("Soccerbet: GetLeagueMatchesInfo(leagueID:" + strconv.Itoa(league.Id) + ") is None, skipping it..")
			continue
		}

		for _, match := range matchesInfo {
			e1 := &[4]string{match.StartDate, league.Name, match.HomeCompetitorName, match.AwayCompetitorName}

			tip1, tip2, err := parseMatchTips(match.Id, betgameByIdMap, betgameOutcomeByIdMap, betgameGroupByIdMap)
			if err != nil {
				continue
			}

			// Match tips
			for t1Key, t1Val := range tip1 {
				t1Game := t1Key[0]
				t1Subgame := t1Key[1]
				t1Code := t1Key[2]

				// Process T0 and 1+ combo
				if t1Subgame == "0" {
					for t2Key, t2Val := range tip2 {
						t2Game := t2Key[0]
						t2Subgame := t2Key[1]
						t2Code := t2Key[2]

						if t2Game == t1Game && t2Subgame == "1+" {
							e2 := &[4]string{
								t1Code,
								fmt.Sprintf("%.2f", t1Val),
								t2Code,
								fmt.Sprintf("%.2f", t2Val),
							}
							export = append(export, utility.MergeE1E2(e1, e2))
						}
					}
				}
				// Process 0-x and x+ combos
				if strings.HasPrefix(t1Subgame, "0-") && len(t1Subgame) == 3 {
					x, _ := strconv.Atoi(string(t1Subgame[2]))
					for t2Key, t2Val := range tip2 {
						t2Game := t2Key[0]
						t2Subgame := t2Key[1]
						t2Code := t2Key[2]

						if t2Game == t1Game && t2Subgame == strconv.Itoa(x+1)+"+" {
							e2 := &[4]string{
								t1Code,
								fmt.Sprintf("%.2f", t1Val),
								t2Code,
								fmt.Sprintf("%.2f", t2Val),
							}
							export = append(export, utility.MergeE1E2(e1, e2))
						}
					}
				}
				// Process GG NG combo
				if t1Subgame == "GG" {
					for t2Key, t2Val := range tip2 {
						t2Game := t2Key[0]
						t2Subgame := t2Key[1]
						t2Code := t2Key[2]

						if t2Game == t1Game && t2Subgame == "NG" {
							e2 := &[4]string{
								t1Code,
								fmt.Sprintf("%.2f", t1Val),
								t2Code,
								fmt.Sprintf("%.2f", t2Val),
							}
							export = append(export, utility.MergeE1E2(e1, e2))
						}
					}
				}
			}

			matchesScrapedCounter++
		}
	}

	fmt.Println("@SOCCERBET" + strings.Repeat("-", 26-len("@SOCCERBET")))
	fmt.Println("Matches scraped: ", matchesScrapedCounter)
	fmt.Println("Tips scraped: ", len(export))

	return export
}

func parseMatchTips(
	matchID int,
	betgameByIdMap map[int]*soccerbet.Betgame,
	betgameOutcomeByIdMap map[int]*soccerbet.BetgameOutcome,
	betgameGroupByIdMap map[int]*soccerbet.BetgameGroup) (map[[3]string]float64, map[[3]string]float64, error) {

	matchOdds, err := soccerbet.GetMatchOddsValues(matchID)
	if err != nil {
		fmt.Println("Soccerbet: GetMatchOddsValues(matchID:" + strconv.Itoa(matchID) + ") is None, skipping it..")
		return nil, nil, errors.New("skipping match")
	}

	tip1 := map[[3]string]float64{}
	tip2 := map[[3]string]float64{}

	for _, odds := range matchOdds {
		if !odds.IsEnabled {
			continue
		}

		outcome := betgameOutcomeByIdMap[odds.BetGameOutcomeId]
		betgame := betgameByIdMap[outcome.BetGameId]
		betgameGroup := betgameGroupByIdMap[betgame.BetGameGroupId]

		tipKey := [3]string{
			betgame.Name,
			outcome.Name,
			outcome.CodeForPrinting,
		}
		tipVal := odds.Odds

		// GG/NG - GG1/2 is never offered
		if betgameGroup.Name == "OBA TIMA DAJU GOL" && betgame.Name == "Oba Tima Daju Gol" {
			if outcome.Name == "GG" {
				tip1[tipKey] = tipVal
			}
			if outcome.Name == "NG" {
				tip2[tipKey] = tipVal
			}
			continue
		}
		// UKUPNO GOLOVA
		if betgameGroup.Name != "UKUPNO GOLOVA" &&
			betgameGroup.Name != "DOMAĆIN UK. GOLOVA" &&
			betgameGroup.Name != "GOST UK. GOLOVA" {
			continue
		}
		if utility.IsElInSliceSTR(betgame.Name, goloviSubgames) {
			if strings.HasPrefix(outcome.Name, "0-") || outcome.Name == "0" {
				tip1[tipKey] = tipVal
			} else if strings.HasSuffix(outcome.Name, "+") {
				tip2[tipKey] = tipVal
			}
		}
	}
	return tip1, tip2, nil
}
