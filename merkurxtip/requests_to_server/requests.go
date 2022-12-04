package requests_to_server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func GetSidebarSports() map[string]interface{} {

	url := "https://www.merkurxtip.rs/restapi/offer/sr/categories/s"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	req.Header.Add("authority", "www.merkurxtip.rs")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "en-US,en;q=0.9,bs;q=0.8")
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

func GetAllSubgames() map[string]interface{} {

	url := "https://www.merkurxtip.rs/restapi/offer/sr/ttg_lang?locale=sr"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	req.Header.Add("authority", "www.merkurxtip.rs")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "en-US,en;q=0.9,bs;q=0.8")
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

func GetSidebarSportGroups(sportID string) map[string]interface{} {

	url := "https://www.merkurxtip.rs/restapi/offer/sr/categories/sport/" + sportID + "/g?locale=sr"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	req.Header.Add("authority", "www.merkurxtip.rs")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "en-US,en;q=0.9,bs;q=0.8")
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

func GetSidebarSportGroupLeagues(sportID string, groupID string) map[string]interface{} {

	url := "https://www.merkurxtip.rs/restapi/offer/sr/categories/sport/" + sportID + "/group/" + groupID + "/l?locale=sr"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	req.Header.Add("authority", "www.merkurxtip.rs")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "en-US,en;q=0.9,bs;q=0.8")
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

func GetMatchIDs(sportID string, leagueID string) map[string]interface{} {

	url := "https://www.merkurxtip.rs/restapi/offer/sr/sport/" + sportID + "/league/" + leagueID + "/desk?locale=sr"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	req.Header.Add("authority", "www.merkurxtip.rs")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "en-US,en;q=0.9,bs;q=0.8")
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

func GetMatchOdds(matchID int) map[string]interface{} {

	url := "https://www.merkurxtip.rs/restapi/offer/sr/match/" + strconv.Itoa(matchID) + "?locale=sr"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	req.Header.Add("authority", "www.merkurxtip.rs")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "en-US,en;q=0.9,bs;q=0.8")
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
