package maxbet

import (
	"OISA_2x_sistem/requests_to_server"
	"fmt"
	"net/http"
)

type SidebarSport struct {
	Name    string
	Leagues []League
}

type League struct {
	BetLeagueId int
	Name        string
	Active      bool
	Blocked     bool
}

func GetSidebarNoRetry() ([]SidebarSport, error) {

	url := "https://www.maxbet.rs/ibet/offer/sportsAndLeagues/-1.json?v=4.50.1&locale=sr"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("cookie", "org.springframework.web.servlet.i18n.CookieLocaleResolver.LOCALE=sr")
	req.Header.Add("authority", "www.maxbet.rs")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "en-US,en;q=0.9,bs;q=0.8")
	req.Header.Add("referer", "https://www.maxbet.rs/ibet-web-client/")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")

	var response []SidebarSport
	err = requests_to_server.GetJson(requests_to_server.Maxbet, req, &response)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return response, nil
}

func GetSidebar() ([]SidebarSport, error) {
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
