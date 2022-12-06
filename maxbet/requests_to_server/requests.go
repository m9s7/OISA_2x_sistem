package requests_to_server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func GetSidebar() []map[string]interface{} {

	url := "https://www.maxbet.rs/ibet/offer/sportsAndLeagues/-1.json?v=4.48.18&locale=sr"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Add("cookie", "org.springframework.web.servlet.i18n.CookieLocaleResolver.LOCALE=sr")
	req.Header.Add("authority", "www.maxbet.rs")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "en-US,en;q=0.9,bs;q=0.8")
	req.Header.Add("referer", "https://www.maxbet.rs/ibet-web-client/")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(res.Body)

	var resultJSON []map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&resultJSON)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return resultJSON
}

func GetMatchIDs(leagueIDs []int) []map[string]interface{} {

	url := "https://www.maxbet.rs/ibet/offer/leagues//-1/0.json"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return nil
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

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(res.Body)

	var resultJSON []map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&resultJSON)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return resultJSON

}

func GetMatchData(matchId int) map[string]interface{} {

	url := "https://www.maxbet.rs/ibet/offer/special/undefined/" + strconv.Itoa(matchId) + ".json?v=4.50.1&locale=sr"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	req.Header.Add("cookie", "org.springframework.web.servlet.i18n.CookieLocaleResolver.LOCALE=sr")
	req.Header.Add("authority", "www.maxbet.rs")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "en-US,en;q=0.9,bs;q=0.8")
	req.Header.Add("referer", "https://www.maxbet.rs/ibet-web-client/")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(res.Body)

	var resultJSON map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&resultJSON)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return resultJSON
}
