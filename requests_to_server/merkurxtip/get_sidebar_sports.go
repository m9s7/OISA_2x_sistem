package merkurxtip

import (
	"OISA_2x_sistem/requests_to_server"
	"fmt"
	"net/http"
)

type Sidebar struct {
	Categories []Category
}

type Category struct {
	Id   string
	Name string // nznm dal ce mi trebati
}

func getSidebarSportsNoRetry() (*Sidebar, error) {

	url := "https://www.merkurxtip.rs/restapi/offer/sr/categories/s"

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

	var response Sidebar
	err = requests_to_server.GetJson(requests_to_server.Merkurxtip, req, &response)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &response, nil
}

func GetSidebarSports() (*Sidebar, error) {
	for i := requests_to_server.RetryStrategy.Start(); ; {

		response, err := getSidebarSportsNoRetry()
		if err == nil {
			return response, nil
		}

		if !i.Next(nil) {
			return nil, fmt.Errorf("error getting sidebar sports after %d tries: %v", i.Count(), err)
		}

	}
}
