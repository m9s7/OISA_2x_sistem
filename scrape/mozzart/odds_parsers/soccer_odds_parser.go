package odds_parsers

import (
	"OISA_2x_sistem/requests_to_server/mozzart"
	"OISA_2x_sistem/utility"
	"fmt"
	"strconv"
	"strings"
)

func SoccerOddsParser(sportID int, allSubgamesResponse map[string][]mozzart.Offer) []*[8]string {

	matchesScrapedCounter := 0
	var export []*[8]string

	matchesResponse, _ := mozzart.GetMatchIDs(sportID)
	exportHelp := initExportHelp(matchesResponse.Matches)

	var matchIDs []int
	for k := range exportHelp {
		matchIDs = append(matchIDs, k)
	}

	subgameIDs := getSoccerSubgameIDs(allSubgamesResponse[strconv.Itoa(sportID)])
	odds := mozzart.GetOdds(matchIDs, subgameIDs)

	for _, matchOdds := range odds {

		if matchOdds.Kodds == nil {
			continue
		}

		matchID := matchOdds.Id

		// Collect all tips I'm interested in (and their values)
		tip1, tip2 := collectFocusedSoccerTips(matchOdds)

		// Match collected tips
		for t1Key, t1Val := range tip1 {
			t1Game := t1Key[0]
			t1Subgame := t1Key[1]

			if strings.HasPrefix(t1Subgame, "0-") && len(t1Subgame) == 3 {
				x, _ := strconv.Atoi(string(t1Subgame[2]))
				t2Subgame := strconv.Itoa(x+1) + "+"
				if t2Val, ok := tip2[[2]string{t1Game, t2Subgame}]; ok {
					export = append(export, utility.MergeE1E2(exportHelp[matchID], &[4]string{
						strings.Join(t1Key[:], " "), t1Val,
						strings.Join([]string{t1Game, t2Subgame}, " "), t2Val,
					}))
				}
				continue
			}

			if t2Val, ok := tip2[[2]string{t1Game, "1+"}]; t1Subgame == "0" && ok {
				export = append(export, utility.MergeE1E2(exportHelp[matchID], &[4]string{
					strings.Join(t1Key[:], " "), t1Val,
					strings.Join([]string{t1Game, "1+"}, " "), t2Val,
				}))
				continue
			}

			if t2Val, ok := tip2[[2]string{"tgg", "ng"}]; t1Subgame == "gg" && ok {
				export = append(export, utility.MergeE1E2(exportHelp[matchID], &[4]string{"gg", t1Val, "ng", t2Val}))
				continue
			}
			if t2Val, ok := tip2[[2]string{"tgg", "1ng"}]; t1Subgame == "1gg" && ok {
				export = append(export, utility.MergeE1E2(exportHelp[matchID], &[4]string{"1gg", t1Val, "1ng", t2Val}))
				continue
			}
			if t2Val, ok := tip2[[2]string{"tgg", "2ng"}]; t1Subgame == "2gg" && ok {
				export = append(export, utility.MergeE1E2(exportHelp[matchID], &[4]string{"2gg", t1Val, "2ng", t2Val}))
				continue
			}

		}

		matchesScrapedCounter++
	}

	fmt.Println("@MOZZART" + strings.Repeat("-", 26-len("@MOZZART")))
	fmt.Println("Matches scraped: ", matchesScrapedCounter)
	fmt.Println("Tips scraped: ", len(export))
	return export
}

func ugConditionSatisfied(subgame string) bool {
	return strings.HasPrefix(subgame, "0-") ||
		strings.HasPrefix(subgame, "1-") ||
		strings.HasSuffix(subgame, "+") ||
		subgame == "0"
}

func getSoccerSubgameIDs(offers []mozzart.Offer) []int {
	//type Offer struct {
	//    Name           string
	//    RegularHeaders []Header
	//}

	var focusedSubgames []int

	for _, offer := range offers {

		if offer.Name != "Kompletna ponuda" {
			continue
		}

		for _, header := range offer.RegularHeaders {

			game := header.GameName[0].Name

			wantedGames := []string{
				"Ukupno golova na meču", "Tim 1 daje gol", "Tim 2 daje gol",
				"Ukupno golova prvo poluvreme", "Tim 1 golovi prvo poluvreme", "Tim 2 golovi prvo poluvreme",
				"Ukupno golova drugo poluvreme", "Tim 1 golovi drugo poluvreme", "Tim 2 golovi drugo poluvreme",
			}
			if utility.IsElInSliceSTR(game, wantedGames) {
				for _, subgame := range header.SubgameName {
					if ugConditionSatisfied(subgame.Name) {
						focusedSubgames = append(focusedSubgames, subgame.Id)
					}
				}
			}

			wantedGames = []string{
				"Tačan broj golova na meču",
				"Tačan broj golova prvo poluvreme",
				"Tačan broj golova drugo poluvreme",
			}
			if utility.IsElInSliceSTR(game, wantedGames) {
				for _, subgame := range header.SubgameName {
					if subgame.Name == "0" {
						focusedSubgames = append(focusedSubgames, subgame.Id)
					}
				}
			}

			wantedSubgames := []string{"gg", "ng", "1gg", "1ng", "2gg", "2ng"}
			if game == "Oba tima daju gol" {
				for _, subgame := range header.SubgameName {
					if utility.IsElInSliceSTR(subgame.Name, wantedSubgames) {
						focusedSubgames = append(focusedSubgames, subgame.Id)
					}
				}
			}

		}
	}

	return utility.RemoveDuplicates(&focusedSubgames)
}

func collectFocusedSoccerTips(odds mozzart.Odds) (map[[2]string]string, map[[2]string]string) {

	tip1 := map[[2]string]string{}
	tip2 := map[[2]string]string{}

	for _, subgameOdds := range odds.Kodds {

		game := subgameOdds.SubGame.GameShortName
		subgame := subgameOdds.SubGame.SubGameName
		key := [2]string{game, subgame}
		val := subgameOdds.Value

		focusedSubgames := []string{
			"ug", "1ug", "2ug", "tm1", "tm2",
			"1tm2", "1tm1", "2tm2", "2tm1",
			"tgg",
		}
		if utility.IsElInSliceSTR(game, focusedSubgames) {
			if strings.HasPrefix(subgame, "0-") || subgame == "0" || strings.HasSuffix(subgame, "gg") {
				tip1[key] = val
			}
			if strings.HasSuffix(subgame, "+") || strings.HasPrefix(subgame, "1-") || strings.HasSuffix(subgame, "ng") {
				tip2[key] = val
			}
		}
	}

	return tip1, tip2
}
