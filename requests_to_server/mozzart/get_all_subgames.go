package mozzart

import (
	"OISA_2x_sistem/requests_to_server"
	"fmt"
	"net/http"
)

type Offer struct {
	Name           string
	RegularHeaders []Header
}

type Header struct {
	GameName    [1]Game
	SubgameName []Subgame
}

type Game struct {
	Name      string
	ShortName string
}

type Subgame struct {
	Id   int
	Name string
}

func getAllSubgamesNoRetry() (map[string][]Offer, error) {

	url := "https://www.mozzartbet.com/getAllGames"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	req.Header.Add("cookie", "i18next=sr")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "en-US,en;q=0.9,bs;q=0.8")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("referer", "https://www.mozzartbet.com/sr")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")

	var response map[string][]Offer
	err = requests_to_server.GetJson(requests_to_server.Mozzart, req, &response)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return response, nil
}

func GetAllSubgames() (map[string][]Offer, error) {
	for i := requests_to_server.RetryStrategy.Start(); ; {

		response, err := getAllSubgamesNoRetry()
		if err == nil {
			return response, nil
		}

		if !i.Next(nil) {
			return nil, fmt.Errorf("error getting allSubgames after %d tries: %v", i.Count(), err)
		}

	}
}
