package odds_parsers

import (
	"OISA_2x_sistem/requests_to_server/mozzart"
	"OISA_2x_sistem/utility"
	"fmt"
	"strconv"
	"strings"
)

func BasketballOddsParser(sportID int, allSubgamesResponse map[string][]mozzart.Offer) []*[8]string {

	matchesScrapedCounter := 0
	var export []*[8]string

	matchesResponse, err := mozzart.GetMatchIDs(sportID)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	exportHelp := initExportHelp(matchesResponse.Matches)

	var matchIDs []int
	for k := range exportHelp {
		matchIDs = append(matchIDs, k)
	}

	focusedSubgames := []string{"pobm"}
	subgameIDs := getIDsForSubgameNames(allSubgamesResponse[strconv.Itoa(sportID)], focusedSubgames)
	odds := mozzart.GetOdds(matchIDs, subgameIDs)

	for _, matchOdds := range odds {

		if matchOdds.Kodds == nil {
			continue
		}

		e1 := exportHelp[matchOdds.Id]
		exportMatchHelper := map[string]*[4]string{}

		for _, subgameOdds := range matchOdds.Kodds {

			game := subgameOdds.SubGame.GameShortName
			subgame := subgameOdds.SubGame.SubGameName
			val := subgameOdds.Value

			if game == "pobm" {
				var exportMatchHelperKeys []string
				for k := range exportMatchHelper {
					exportMatchHelperKeys = append(exportMatchHelperKeys, k)
				}
				if !utility.IsElInSliceSTR(game, exportMatchHelperKeys) {
					exportMatchHelper[game] = &[4]string{}
				}

				if subgame == "1" {
					exportMatchHelper[game][0] = game + " " + subgame
					exportMatchHelper[game][1] = val
				}
				if subgame == "2" {
					exportMatchHelper[game][2] = game + " " + subgame
					exportMatchHelper[game][3] = val
				}
			}

		}

		for _, e2 := range exportMatchHelper {
			e := utility.MergeE1E2(e1, e2)
			export = append(export, e)
		}
		matchesScrapedCounter++
	}

	fmt.Println("@MOZZART" + strings.Repeat("-", 26-len("@MOZZART")))
	fmt.Println("Matches scraped: ", matchesScrapedCounter)
	fmt.Println("Tips scraped: ", len(export))

	return export
}
