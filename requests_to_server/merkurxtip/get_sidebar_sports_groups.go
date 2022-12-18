package merkurxtip

import (
	"OISA_2x_sistem/requests_to_server"
	"fmt"
	"net/http"
)

type SidebarSportGroupsResponse struct {
	Categories []Group
}

type Group struct {
	Id    string
	Type  string
	Count int
}

func getSidebarSportGroupsNoRetry(sportID string) (*SidebarSportGroupsResponse, error) {

	url := "https://www.merkurxtip.rs/restapi/offer/sr/categories/sport/" + sportID + "/g?locale=sr"

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

	var response SidebarSportGroupsResponse
	err = requests_to_server.GetJson(requests_to_server.Merkurxtip, req, &response)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &response, nil
}

func GetSidebarSportGroups(sportID string) (*SidebarSportGroupsResponse, error) {
	for i := requests_to_server.RetryStrategy.Start(); ; {

		response, err := getSidebarSportGroupsNoRetry(sportID)
		if err == nil {
			return response, nil
		}

		if !i.Next(nil) {
			return nil, fmt.Errorf("error getting sidebar sport groups after %d tries: %v", i.Count(), err)
		}

	}
}
