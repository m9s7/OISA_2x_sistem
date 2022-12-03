package main

import (
	"OISA_2x_sistem/maxbet"
	"OISA_2x_sistem/soccerbet"
	"fmt"
)

func main() {
	fmt.Println(maxbet.GetSportsCurrentlyOffered())

	//maxbet.Scrape("Fudbal")
	//maxbet.Scrape("Košarka")
	//maxbet.Scrape("Tenis")
	//maxbet.Scrape("Stoni Tenis")
	//maxbet.Scrape("eSport")

	fmt.Println(soccerbet.GetSportsCurrentlyOffered())

	//soccerbet.Scrape("Fudbal")
	//soccerbet.Scrape("Košarka")
	soccerbet.Scrape("Tenis")

}
