package server_response_parsers

func GetSportNameByIDMap(masterData map[string]interface{}) map[int]string {
	result := map[int]string{}

	competitionsData := masterData["CompetitionsData"].(map[string]interface{})
	sports := competitionsData["Sports"].([]interface{})
	for _, sport := range sports {
		sport := sport.(map[string]interface{})
		result[int(sport["Id"].(float64))] = sport["Name"].(string)
	}
	return result
}

func GetBetgameByIdMap(masterData map[string]interface{}) map[int]map[string]interface{} {
	result := map[int]map[string]interface{}{}

	betGameOutcomesData := masterData["BetGameOutcomesData"].(map[string]interface{})
	betGames := betGameOutcomesData["BetGames"].([]interface{})
	for _, betGame := range betGames {
		betGame := betGame.(map[string]interface{})
		betGameID := int(betGame["Id"].(float64))
		result[betGameID] = betGame
	}
	return result
}

func GetBetgameOutcomeByIdMap(masterData map[string]interface{}) map[int]map[string]interface{} {
	result := map[int]map[string]interface{}{}

	betGameOutcomesData := masterData["BetGameOutcomesData"].(map[string]interface{})
	betGameOutcomes := betGameOutcomesData["BetGameOutcomes"].([]interface{})
	for _, betGameOutcome := range betGameOutcomes {
		betGameOutcome := betGameOutcome.(map[string]interface{})
		betGameOutcomeID := int(betGameOutcome["Id"].(float64))
		result[betGameOutcomeID] = betGameOutcome
	}
	return result
}

func GetBetgameGroupByIdMap(masterData map[string]interface{}) map[int]map[string]interface{} {
	result := map[int]map[string]interface{}{}
	betGameOutcomesData := masterData["BetGameOutcomesData"].(map[string]interface{})
	betGameGroups := betGameOutcomesData["BetGameGroups"].([]interface{})
	for _, betGameGroup := range betGameGroups {
		betGameGroup := betGameGroup.(map[string]interface{})
		betGameGroupID := int(betGameGroup["Id"].(float64))
		result[betGameGroupID] = betGameGroup
	}
	return result
}
