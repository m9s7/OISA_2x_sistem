package soccerbet

import (
	"OISA_2x_sistem/requests_to_server"
	"fmt"
	"net/http"
	"strconv"
)

type LeagueMatchInfo struct {
	Id                 int
	HomeCompetitorName string
	AwayCompetitorName string
	StartDate          string
}

func getLeagueMatchesInfoNoRetry(leagueID int) ([]LeagueMatchInfo, error) {

	url := "https://soccerbet.rs/api/Prematch/GetCompetitionMatches?competitionId=" + strconv.Itoa(leagueID)

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

	var response []LeagueMatchInfo
	err = requests_to_server.GetJson(requests_to_server.Soccerbet, req, &response)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return response, nil
}

func GetLeagueMatchesInfo(leagueID int) ([]LeagueMatchInfo, error) {
	for i := requests_to_server.RetryStrategy.Start(); ; {

		response, err := getLeagueMatchesInfoNoRetry(leagueID)
		if err == nil {
			return response, nil
		}

		if !i.Next(nil) {
			return nil, fmt.Errorf("error getting league matches info after %d tries: %v", i.Count(), err)
		}

	}
}
