package odds_parsers

import (
	"OISA_2x_sistem/requests_to_server/mozzart"
	"OISA_2x_sistem/utility"
	"fmt"
	"strings"
)

func getIDsForSubgameNames(offers []mozzart.Offer, subgameNames []string) []int {

	var subgameIDs []int

	for _, offer := range offers {

		if offer.Name != "Kompletna ponuda" {
			continue
		}

		for _, header := range offer.RegularHeaders {

			subgameName := header.GameName[0].ShortName

			if utility.IsElInSliceSTR(subgameName, subgameNames) {
				for _, subgame := range header.SubgameName {
					subgameIDs = append(subgameIDs, subgame.Id)
				}
			}

		}
	}
	
	return subgameIDs
}

func initExportHelp(matchesResponse []mozzart.Match) map[int]*[4]string {

	export := map[int]*[4]string{}

	for _, match := range matchesResponse {

		if match.SpecialType != 0 || len(match.Participants) != 2 {
			continue
		}

		export[match.Id] = &[4]string{
			fmt.Sprintf("%.0f", match.StartTime),
			match.Competition_name_sr,
			strings.Trim(match.Participants[0].Name, " "),
			strings.Trim(match.Participants[1].Name, " "),
		}

	}

	return export
}
