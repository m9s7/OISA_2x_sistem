package mozzart

import (
	"OISA_2x_sistem/requests_to_server"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Matches struct {
	Matches []Match
	Count   int
}

type Match struct {
	Id                  int
	Participants        []Participant
	SpecialType         int
	StartTime           float64
	Competition_name_sr string
}

type Participant struct {
	Name string
}

func getMatchIDsNoRetry(sportID int) (*Matches, error) {

	url := "https://www.mozzartbet.com/betOffer2"

	currentDate := fmt.Sprint(time.Now().Format("01-02-2006"))
	payload := strings.NewReader(
		"{\"date\":\"" + currentDate +
			"\",\"sportIds\":[" + strconv.Itoa(sportID) + "]," +
			"\"competitionIds\":[]," +
			"\"sort\":\"bycompetition\"," +
			"\"specials\":false," +
			"\"subgames\":[]," +
			"\"size\":1000," +
			"\"mostPlayed\":false," +
			"\"type\":\"betting\"," +
			"\"numberOfGames\":1000," +
			"\"activeCompleteOffer\":false," +
			"\"lang\":\"sr\"," +
			"\"offset\":0}")

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	req.Header.Add("cookie", "i18next=sr")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "en-US,en;q=0.9,bs;q=0.8")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("content-type", "application/json;charset=UTF-8")
	req.Header.Add("origin", "https://www.mozzartbet.com")
	req.Header.Add("referer", "https://www.mozzartbet.com/sr")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("x-requested-with", "XMLHttpRequest")

	var response Matches
	err = requests_to_server.GetJson(requests_to_server.Mozzart, req, &response)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &response, nil
}

func GetMatchIDs(sportID int) (*Matches, error) {
	for i := requests_to_server.RetryStrategy.Start(); ; {

		response, err := getMatchIDsNoRetry(sportID)
		if err == nil {
			return response, nil
		}

		if !i.Next(nil) {
			return nil, fmt.Errorf("error getting matchIDs after %d tries: %v", i.Count(), err)
		}

	}
}
