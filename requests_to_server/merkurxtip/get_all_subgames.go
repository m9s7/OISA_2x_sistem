package merkurxtip

import (
	"OISA_2x_sistem/requests_to_server"
	"fmt"
	"net/http"
)

type AllSubgamesResponse struct {
	BetPickMap map[string]BetPick
}

type BetPick struct {
	TipTypeCode int
	TipTypeName string
}

func getAllSubgamesNoRetry() (*AllSubgamesResponse, error) {

	url := "https://www.merkurxtip.rs/restapi/offer/sr/ttg_lang?locale=sr"

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

	var response AllSubgamesResponse
	err = requests_to_server.GetJson(requests_to_server.Merkurxtip, req, &response)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &response, nil
}

func GetAllSubgames() (*AllSubgamesResponse, error) {
	for i := requests_to_server.RetryStrategy.Start(); ; {

		response, err := getAllSubgamesNoRetry()
		if err == nil {
			return response, nil
		}

		if !i.Next(nil) {
			return nil, fmt.Errorf("error getting all subgames after %d tries: %v", i.Count(), err)
		}

	}
}
