package mozzart

import (
	"OISA_2x_sistem/requests_to_server"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Odds struct {
	Id    int
	Kodds map[string]SubgameOdds
}

type SubgameOdds struct {
	Value   string
	SubGame SubgameDetails
}

type SubgameDetails struct {
	GameShortName string
	SubGameName   string
}

func getOddsForLimitedNumOfMatchesNoRetry(matchIDs []int, subgameIDs []int) ([]Odds, error) {
	url := "https://www.mozzartbet.com/getBettingOdds"

	var matchIDsAsStrings []string
	for _, matchID := range matchIDs {
		matchIDsAsStrings = append(matchIDsAsStrings, strconv.Itoa(matchID))
	}

	var subgameIDsAsStrings []string
	for _, subgameID := range subgameIDs {
		subgameIDsAsStrings = append(subgameIDsAsStrings, strconv.Itoa(subgameID))
	}

	payload := strings.NewReader(
		"{\"matchIds\":[" + strings.Join(matchIDsAsStrings, ", ") + "]," +
			"\"subgames\":[" + strings.Join(subgameIDsAsStrings, ", ") + "]}")

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	req.Header.Add("cookie", "i18next=sr")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "en-US,en;q=0.9,bs;q=0.8")
	req.Header.Add("content-type", "application/json;charset=UTF-8")
	req.Header.Add("origin", "https://www.mozzartbet.com")
	req.Header.Add("referer", "https://www.mozzartbet.com/sr")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("x-requested-with", "XMLHttpRequest")

	var response []Odds
	err = requests_to_server.GetJson(requests_to_server.Mozzart, req, &response)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return response, nil
}

func getOddsForLimitedNumOfMatches(matchIDs []int, subgameIDs []int) ([]Odds, error) {
	for i := requests_to_server.RetryStrategy.Start(); ; {

		response, err := getOddsForLimitedNumOfMatchesNoRetry(matchIDs, subgameIDs)
		if err == nil {
			return response, nil
		}

		if !i.Next(nil) {
			return nil, fmt.Errorf("error getting odds after %d tries: %v", i.Count(), err)
		}

	}
}

func GetOdds(matchIDs []int, subgameIDs []int) []Odds {
	limit := 49

	var result []Odds

	for len(matchIDs) > 0 {
		if len(matchIDs) <= limit {
			limit = len(matchIDs)
		}

		matchIDsBatch := matchIDs[:limit]
		matchIDs = matchIDs[limit:]

		currentBatchResult, err := getOddsForLimitedNumOfMatches(matchIDsBatch, subgameIDs)
		if err != nil {
			fmt.Println(err)
			continue
		}

		result = append(result, currentBatchResult...)
	}

	return result
}
