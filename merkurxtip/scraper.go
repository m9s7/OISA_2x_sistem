package merkurxtip

import (
	"OISA_2x_sistem/merkurxtip/requests_to_server"
	"OISA_2x_sistem/merkurxtip/server_response_parsers"
)

func GetSportsCurrentlyOffered() []string {
	response := requests_to_server.GetSidebarSports()
	sidebarSportsIDsByName := server_response_parsers.ParseGetSidebarSports(response)

	var sports []string
	for sport, _ := range sidebarSportsIDsByName {
		sports = append(sports, sport)
	}
	return sports
}
