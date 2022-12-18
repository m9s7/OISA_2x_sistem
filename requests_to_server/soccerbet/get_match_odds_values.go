package soccerbet

import (
	"OISA_2x_sistem/requests_to_server"
	"fmt"
	"net/http"
	"strconv"
)

type MatchOdds struct {
	Id               int
	Odds             float64
	IsEnabled        bool
	BetGameOutcomeId int
}

func getMatchOddsValuesNoRetry(matchID int) ([]MatchOdds, error) {

	url := "https://soccerbet.rs/api/Prematch/GetMatchBets?matchId=" + strconv.Itoa(matchID)

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

	var response []MatchOdds
	err = requests_to_server.GetJson(requests_to_server.Mozzart, req, &response)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return response, nil
}

func GetMatchOddsValues(matchID int) ([]MatchOdds, error) {
	for i := requests_to_server.RetryStrategy.Start(); ; {

		response, err := getMatchOddsValuesNoRetry(matchID)
		if err == nil {
			return response, nil
		}

		if !i.Next(nil) {
			return nil, fmt.Errorf("error getting match odds values after %d tries: %v", i.Count(), err)
		}

	}
}
