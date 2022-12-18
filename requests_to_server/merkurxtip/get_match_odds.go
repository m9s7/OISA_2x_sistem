package merkurxtip

import (
	"OISA_2x_sistem/requests_to_server"
	"fmt"
	"net/http"
	"strconv"
)

type Match struct {
	Id          int
	Home        string
	Away        string
	KickOffTime float64
	Blocked     bool
	Odds        map[string]float64
	LeagueName  string
	LeagueShort string
}

func getMatchOddsNoRetry(matchID int) (*Match, error) {

	url := "https://www.merkurxtip.rs/restapi/offer/sr/match/" + strconv.Itoa(matchID) + "?locale=sr"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	req.Header.Add("authority", "www.merkurxtip.rs")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "en-US,en;q=0.9,bs;q=0.8")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")

	var response Match
	err = requests_to_server.GetJson(requests_to_server.Merkurxtip, req, &response)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &response, nil
}

func GetMatchOdds(matchID int) (*Match, error) {
	for i := requests_to_server.RetryStrategy.Start(); ; {

		response, err := getMatchOddsNoRetry(matchID)
		if err == nil {
			return response, nil
		}

		if !i.Next(nil) {
			return nil, fmt.Errorf("error getting match odds after %d tries: %v", i.Count(), err)
		}

	}
}
