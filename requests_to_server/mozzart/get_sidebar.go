package mozzart

import (
	"OISA_2x_sistem/requests_to_server"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Sport struct {
	Id   int
	Name string
}

func GetSidebarNoRetry() ([]Sport, error) {

	url := "https://www.mozzartbet.com/getRegularGroups"

	currentDate := fmt.Sprint(time.Now().Format("01-02-2006"))
	payload := strings.NewReader(
		"{\"date\":\"" + currentDate + "\"," +
			"\"sportIds\":[]," +
			"\"competitionIds\":[]," +
			"\"sort\":\"bycompetition\"," +
			"\"specials\":null," +
			"\"subgames\":[]," +
			"\"size\":1000," +
			"\"mostPlayed\":false," +
			"\"type\":\"betting\"," +
			"\"numberOfGames\":0," +
			"\"activeCompleteOffer\":false," +
			"\"lang\":\"sr\"," +
			"\"offset\":0}")

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	req.Header.Add("cookie", "i18next=sr")
	req.Header.Add("authority", "www.mozzartbet.com")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "en-US,en;q=0.9,bs;q=0.8")
	req.Header.Add("content-type", "application/json;charset=UTF-8")
	req.Header.Add("origin", "https://www.mozzartbet.com")
	req.Header.Add("referer", "https://www.mozzartbet.com/sr")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("x-requested-with", "XMLHttpRequest")

	var response []Sport
	err = requests_to_server.GetJson(requests_to_server.Mozzart, req, &response)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return response, nil
}

func GetSidebar() ([]Sport, error) {
	for i := requests_to_server.RetryStrategy.Start(); ; {

		response, err := GetSidebarNoRetry()
		if err == nil {
			return response, nil
		}

		if !i.Next(nil) {
			return nil, fmt.Errorf("error getting sidebar after %d tries: %v", i.Count(), err)
		}

	}
}
