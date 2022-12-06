package server_response_parsers

func ParseGetMatchesIDsResponse(response []map[string]interface{}) []int {
	var matchIDs []int
	for _, league := range response {
		for _, match := range league["matchList"].([]interface{}) {
			matchIDs = append(matchIDs, int(match.(map[string]interface{})["id"].(float64)))
		}
	}
	return matchIDs
}
