package odds_parsers

import (
	"OISA_2x_sistem/requests_to_server/soccerbet"
	"OISA_2x_sistem/utility"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func TennisOddsParser(
	sidebarLeagues []soccerbet.CompetitionMasterData,
	betgameByIdMap map[int]soccerbet.Betgame,
	betgameOutcomeByIdMap map[int]soccerbet.BetgameOutcome,
	betgameGroupByIdMap map[int]soccerbet.BetgameGroup,
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

			matchOdds, err := soccerbet.GetMatchOddsValues(match.Id)
			if err != nil {
				fmt.Println("Soccerbet: GetMatchOddsValues(matchID:" + strconv.Itoa(match.Id) + ") is None, skipping it..")
				continue
			}
			exportMatchHelper := map[string]*[4]string{}

			for _, odds := range matchOdds {

				if !odds.IsEnabled {
					continue
				}

				outcome := betgameOutcomeByIdMap[odds.BetGameOutcomeId]
				betgame := betgameByIdMap[outcome.BetGameId]
				betgameGroup := betgameGroupByIdMap[betgame.BetGameGroupId]

				tipVal := odds.Odds

				var exportMatchHelperKeys []string
				for key := range exportMatchHelper {
					exportMatchHelperKeys = append(exportMatchHelperKeys, key)
				}

				tipComboKey := betgameGroup.Name + " " + betgame.Name

				// KI
				if betgameGroup.Name == "MEČ" && betgame.Name == "Konačni Ishod" {

					if !utility.IsElInSliceSTR(tipComboKey, exportMatchHelperKeys) {
						exportMatchHelper[tipComboKey] = &[4]string{}
					}

					if outcome.Name == "1" {
						exportMatchHelper[tipComboKey][0] = outcome.CodeForPrinting
						exportMatchHelper[tipComboKey][1] = fmt.Sprintf("%.2f", tipVal)
					} else if outcome.Name == "2" {
						exportMatchHelper[tipComboKey][2] = outcome.CodeForPrinting
						exportMatchHelper[tipComboKey][3] = fmt.Sprintf("%.2f", tipVal)
					} else {
						log.Fatalln(tipComboKey, outcome.Name, outcome.Description, outcome.CodeForPrinting, tipVal)
					}
				}

				// KI PRVI SET
				if betgameGroup.Name == "SET" && betgame.Name == "I Set" {

					if !utility.IsElInSliceSTR(tipComboKey, exportMatchHelperKeys) {
						exportMatchHelper[tipComboKey] = &[4]string{}
					}

					if outcome.Name == "1" {
						exportMatchHelper[tipComboKey][0] = outcome.CodeForPrinting
						exportMatchHelper[tipComboKey][1] = fmt.Sprintf("%.2f", tipVal)
					} else if outcome.Name == "2" {
						exportMatchHelper[tipComboKey][2] = outcome.CodeForPrinting
						exportMatchHelper[tipComboKey][3] = fmt.Sprintf("%.2f", tipVal)
					} else {
						log.Fatalln(tipComboKey, outcome.Name, outcome.Description, outcome.CodeForPrinting, tipVal)
					}
				}

				// TIE BREAK
				if betgameGroup.Name == "MEČ" && betgame.Name == "Tie Break" {

					if !utility.IsElInSliceSTR(tipComboKey, exportMatchHelperKeys) {
						exportMatchHelper[tipComboKey] = &[4]string{}
					}

					if outcome.Name == "DA" {
						exportMatchHelper[tipComboKey][0] = outcome.CodeForPrinting
						exportMatchHelper[tipComboKey][1] = fmt.Sprintf("%.2f", tipVal)
					} else if outcome.Name == "NE" {
						exportMatchHelper[tipComboKey][2] = outcome.CodeForPrinting
						exportMatchHelper[tipComboKey][3] = fmt.Sprintf("%.2f", tipVal)
					} else {
						log.Fatalln(tipComboKey, outcome.Name, outcome.Description, outcome.CodeForPrinting, tipVal)
					}
				}
			}

			for _, e2 := range exportMatchHelper {
				export = append(export, utility.MergeE1E2(e1, e2))
			}
			matchesScrapedCounter++
		}
	}

	fmt.Println("@SOCCERBET" + strings.Repeat("-", 26-len("@SOCCERBET")))
	fmt.Println("Matches scraped: ", matchesScrapedCounter)
	fmt.Println("Tips scraped: ", len(export))

	return export
}
