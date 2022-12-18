package maxbet

import (
	"OISA_2x_sistem/requests_to_server"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type LeagueMatches struct {
	MatchList []Match
}

type Match struct {
	Id int
}

func getMatchIDsNoRetry(leagueIDs []int) ([]LeagueMatches, error) {

	url := "https://www.maxbet.rs/ibet/offer/leagues//-1/0.json"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	q := req.URL.Query()
	q.Add("v", "4.50.1")
	q.Add("locale", "sr")
	var leagueIDsAsStrings []string
	for _, id := range leagueIDs {
		leagueIDsAsStrings = append(leagueIDsAsStrings, strconv.Itoa(id))
	}
	q.Add("token", strings.Join(leagueIDsAsStrings, "#"))
	q.Add("ttgIds", "")
	req.URL.RawQuery = q.Encode()

	req.Header.Add("cookie", "org.springframework.web.servlet.i18n.CookieLocaleResolver.LOCALE=sr")
	req.Header.Add("authority", "www.maxbet.rs")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "en-US,en;q=0.9,bs;q=0.8")
	req.Header.Add("referer", "https://www.maxbet.rs/ibet-web-client/")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")

	var response []LeagueMatches
	err = requests_to_server.GetJson(requests_to_server.Maxbet, req, &response)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return response, nil
}

func GetMatchIDs(leagueIDs []int) ([]LeagueMatches, error) {
	for i := requests_to_server.RetryStrategy.Start(); ; {

		response, err := getMatchIDsNoRetry(leagueIDs)
		if err == nil {
			return response, nil
		}

		if !i.Next(nil) {
			return nil, fmt.Errorf("error getting match ids after %d tries: %v", i.Count(), err)
		}

	}
}
