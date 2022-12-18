package server_response_parsers

import "OISA_2x_sistem/requests_to_server/soccerbet"

func GetSportNameByID(masterData *soccerbet.MasterData) map[int]string {
	result := map[int]string{}

	for _, sport := range masterData.CompetitionsData.Sports {
		result[sport.Id] = sport.Name
	}
	return result
}

func GetBetgameById(masterData *soccerbet.MasterData) map[int]*soccerbet.Betgame {
	result := map[int]*soccerbet.Betgame{}

	for _, betGame := range masterData.BetGameOutcomesData.BetGames {
		result[betGame.Id] = &betGame
	}

	return result
}

func GetBetgameOutcomeById(masterData *soccerbet.MasterData) map[int]*soccerbet.BetgameOutcome {
	result := map[int]*soccerbet.BetgameOutcome{}

	for _, betGameOutcome := range masterData.BetGameOutcomesData.BetGameOutcomes {
		result[betGameOutcome.Id] = &betGameOutcome
	}

	return result
}

func GetBetgameGroupById(masterData *soccerbet.MasterData) map[int]*soccerbet.BetgameGroup {
	result := map[int]*soccerbet.BetgameGroup{}

	for _, betGameGroup := range masterData.BetGameOutcomesData.BetGameGroups {
		result[betGameGroup.Id] = &betGameGroup
	}
	return result
}
