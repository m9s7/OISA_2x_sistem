package server_response_parsers

func ParseGetSidebarSportsAndLeagues(response []map[string]interface{}) map[string]int {
	getNameByIDMap := map[string]int{}
	for _, mapObj := range response {
		name := mapObj["name"].(string)
		id := int(mapObj["id"].(float64))
		getNameByIDMap[name] = id
	}
	return getNameByIDMap
}
