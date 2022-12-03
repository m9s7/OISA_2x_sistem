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

func GetMatchIds(leagueIDs []int) []map[string]interface{} {

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

	req.Header.Add("cookie", "SESSION=eefc750c-d31a-4f64-a746-f66acb11dbd4; utkn=4a5f9352453828af2c0f99e44aed5f57; c_stamp=dca0b108-2243-458b-bf06-0aefcb701edb; _gcl_au=1.1.936261194.1663309076; _hjSessionUser_1350202=eyJpZCI6IjBiZDc1ZTIzLWQ5NWYtNWUwMC1hMzlmLWFiNDBlOTk0MzViNyIsImNyZWF0ZWQiOjE2NjMzMDkwNzYxNjcsImV4aXN0aW5nIjp0cnVlfQ==; _gcl_aw=GCL.1663591068.CjwKCAjwpqCZBhAbEiwAa7pXeVSXQsZpTd3Oj7oLS9eL82c_mATnK2z2_-Y93vrUwnCkF-nP9DoZjxoCVBgQAvD_BwE; _gac_UA-61991021-7=1.1663591068.CjwKCAjwpqCZBhAbEiwAa7pXeVSXQsZpTd3Oj7oLS9eL82c_mATnK2z2_-Y93vrUwnCkF-nP9DoZjxoCVBgQAvD_BwE; _ga_5RBFPRP28H=GS1.1.1663591067.1.1.1663591078.0.0.0; _ga=GA1.2.649193112.1663309076; _gid=GA1.2.535184041.1665315269; org.springframework.web.servlet.i18n.CookieLocaleResolver.LOCALE=sr; ASP.NET_SessionId=i5kdqzaesig1djtrxw5hgxvq; _clck=1u50hws|1|f5k|0; _hjSession_1350202=eyJpZCI6ImZjN2RiYzM2LWQyZjUtNDU2Zi04ZTg1LTI2YTIyN2JhNmUyMiIsImNyZWF0ZWQiOjE2NjUzMTUyNzIyMzQsImluU2FtcGxlIjpmYWxzZX0=; _gat_UA-61991021-1=1; __cf_bm=UxnJKaEFrerBWR7NadBxg6XGPd682te1xtbz3zWIcHc-1665316239-0-AVz0cZ6bMOtACfdq4BzxD54DVYD9E3tvZ6xxu93ivK4IlBuwhJjFC0cKm0jfvndPU9avvSuCV2EB+FgV8cpvwxO6CvtcShQCFRYxnnoi4HVtBYC5kJZSTnwlVeJ9ZaGOig==; _gat=1; _clsk=1wswi4j|1665316242102|1|0|m.clarity.ms/collect")
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
