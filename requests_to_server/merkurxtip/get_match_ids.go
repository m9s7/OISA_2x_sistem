package merkurxtip

import (
	"OISA_2x_sistem/requests_to_server"
	"fmt"
	"net/http"
)

type GetMatchIDsResponse struct {
	EsMatches []MatchID
}

type MatchID struct {
	Id      int
	Blocked bool
}

func getMatchIDsNoRetry(sportID string, groupID string) (*GetMatchIDsResponse, error) {

	url := "https://www.merkurxtip.rs/restapi/offer/sr/sport/" + sportID + "/league-group/" + groupID + "/desk?locale=sr"

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

	var response GetMatchIDsResponse
	err = requests_to_server.GetJson(requests_to_server.Merkurxtip, req, &response)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &response, nil
}

func GetMatchIDs(sportID string, groupID string) (*GetMatchIDsResponse, error) {
	for i := requests_to_server.RetryStrategy.Start(); ; {

		response, err := getMatchIDsNoRetry(sportID, groupID)
		if err == nil {
			return response, nil
		}

		if !i.Next(nil) {
			return nil, fmt.Errorf("error getting match ids after %d tries: %v", i.Count(), err)
		}

	}
}
