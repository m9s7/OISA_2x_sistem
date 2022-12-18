package server_response_parsers

import "OISA_2x_sistem/requests_to_server/mozzart"

func GetSportIDByNameMap(response []mozzart.Sport) map[string]int {

	sportIDByName := map[string]int{}

	for _, sport := range response {
		sportIDByName[sport.Name] = sport.Id
	}

	return sportIDByName
}
