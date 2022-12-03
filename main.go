package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	//fmt.Println(maxbet.GetSportsCurrentlyOffered())

	//maxbet.Scrape("Fudbal")
	//maxbet.Scrape("Košarka")
	//maxbet.Scrape("Tenis")
	//maxbet.Scrape("Stoni Tenis")
	//maxbet.Scrape("eSport")

	//fmt.Println(soccerbet.GetSportsCurrentlyOffered())

	//soccerbet.Scrape("Fudbal")
	//soccerbet.Scrape("Košarka")
	//soccerbet.Scrape("Tenis")

	//url := "https://www.mozzartbet.com/getRegularGroups"

	currentDate := fmt.Sprint(time.Now().Format("01-02-2006"))
	payload := strings.NewReader("{\"date\":\"" + currentDate + "\",\"sportIds\":[],\"competitionIds\":[],\"sort\":\"bycompetition\",\"specials\":null,\"subgames\":[],\"size\":1000,\"mostPlayed\":false,\"type\":\"betting\",\"numberOfGames\":0,\"activeCompleteOffer\":false,\"lang\":\"sr\",\"offset\":0}")

	//req, _ := http.NewRequest("POST", url, payload)

	fmt.Println(payload)
}
