package soccerbet

import (
	"OISA_2x_sistem/requests_to_server"
	"fmt"
	"net/http"
)

type MasterData struct {
	CompetitionsData    CompetitionsMasterData
	BetGameOutcomesData BetGameOutcomesMasterData
}

type CompetitionsMasterData struct {
	Sports       []SportMasterData
	Competitions []CompetitionMasterData
}

type SportMasterData struct {
	Id   int
	Name string
}

type CompetitionMasterData struct {
	Id      int
	Name    string
	SportId int
}

type BetGameOutcomesMasterData struct {
	BetGameGroups   []BetgameGroup
	BetGames        []Betgame
	BetGameOutcomes []BetgameOutcome
}

type Betgame struct {
	Id             int
	Name           string
	BetGameGroupId int
}

type BetgameOutcome struct {
	Id              int
	Name            string
	Description     string
	CodeForPrinting string
	BetGameId       int
}

type BetgameGroup struct {
	Id   int
	Name string
}

func getMasterDataNoRetry() (*MasterData, error) {
	url := "https://soccerbet.rs/api/MasterData/GetMasterData"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("Accept-Language", "en-US,en;q=0.9,bs;q=0.8")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Referer", "https://soccerbet.rs/")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")

	var response MasterData
	err = requests_to_server.GetJson(requests_to_server.Mozzart, req, &response)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &response, nil
}

func GetMasterData() (*MasterData, error) {
	for i := requests_to_server.RetryStrategy.Start(); ; {

		response, err := getMasterDataNoRetry()
		if err == nil {
			return response, nil
		}

		if !i.Next(nil) {
			return nil, fmt.Errorf("error getting master data after %d tries: %v", i.Count(), err)
		}

	}
}
