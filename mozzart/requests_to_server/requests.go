package requests_to_server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func GetSidebarSportsAndLeagues() []map[string]interface{} {

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
		return nil
	}

	req.Header.Add("cookie", "i18next=sr; SERVERID=MB-N7")
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

func getAllSubgames() map[string]interface{} {

	url := "https://www.mozzartbet.com/getAllGames"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	req.Header.Add("cookie", "i18next=sr; SERVERID=MB-N7")
	req.Header.Add("authority", "www.mozzartbet.com")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "en-US,en;q=0.9,bs;q=0.8")
	req.Header.Add("referer", "https://www.mozzartbet.com/sr")
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

func getMatchIDs(sportID int) map[string]interface{} {

	url := "https://www.mozzartbet.com/betOffer2"

	currentDate := fmt.Sprint(time.Now().Format("01-02-2006"))
	payload := strings.NewReader(
		"{\"date\":\"" + currentDate +
			"\",\"sportIds\":[" + strconv.Itoa(sportID) + "]," +
			"\"competitionIds\":[]," +
			"\"sort\":\"bycompetition\"," +
			"\"specials\":false," +
			"\"subgames\":[]," +
			"\"size\":1000," +
			"\"mostPlayed\":false," +
			"\"type\":\"betting\"," +
			"\"numberOfGames\":1000," +
			"\"activeCompleteOffer\":false," +
			"\"lang\":\"sr\"," +
			"\"offset\":0}")

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	req.Header.Add("cookie", "i18next=sr; SERVERID=MB-N7")
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

func getOddsForLimitedNumOfMatches(matchIDs []int, subgameIDs []int) []map[string]interface{} {
	url := "https://www.mozzartbet.com/getBettingOdds"

	var matchIDsAsStrings []string
	for _, matchID := range matchIDs {
		matchIDsAsStrings = append(matchIDsAsStrings, strconv.Itoa(matchID))
	}

	var subgameIDsAsStrings []string
	for _, subgameID := range subgameIDs {
		subgameIDsAsStrings = append(subgameIDsAsStrings, strconv.Itoa(subgameID))
	}

	payload := strings.NewReader(
		"{\"matchIds\":[" + strings.Join(matchIDsAsStrings, ", ") + "]," +
			"\"subgames\":[" + strings.Join(subgameIDsAsStrings, ", ") + "]}")

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	req.Header.Add("cookie", "i18next=sr; SERVERID=MB-N7")
	req.Header.Add("authority", "www.mozzartbet.com")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "en-US,en;q=0.9,bs;q=0.8")
	req.Header.Add("content-type", "application/json;charset=UTF-8")
	req.Header.Add("origin", "https://www.mozzartbet.com")
	req.Header.Add("referer", "https://www.mozzartbet.com/sr")
	req.Header.Add("x-requested-with", "XMLHttpRequest")

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

func getOdds(matchIDs []int, subgameIDs []int) []map[string]interface{} {
	limit := 49

	var result []map[string]interface{}

	for len(matchIDs) > 0 {
		matchIDsBatch := matchIDs[:limit]
		matchIDs = matchIDs[limit:]

		currentBatchResult := getOddsForLimitedNumOfMatches(matchIDsBatch, subgameIDs)
		for currentBatchResult == nil {
			fmt.Println("Mozzart requests: Stuck on getOddsForLimitedNumOfMatches(matchIDs, subgameIDs)")
			fmt.Println("matchIDs batch: ", matchIDsBatch)
			fmt.Println("matchIDs batch length is", len(matchIDsBatch))
			fmt.Println("subgameIDs: ", subgameIDs)
			fmt.Println("subgameIDs length is", len(subgameIDs))

			currentBatchResult = getOddsForLimitedNumOfMatches(matchIDsBatch, subgameIDs)
		}

		result = append(result, currentBatchResult...)
	}

	return result
}
