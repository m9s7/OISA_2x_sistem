package main

import (
	"OISA_2x_sistem/merkurxtip"
	"fmt"
)

func main() {
	//fmt.Println(maxbet.GetSportsCurrentlyOffered())
	//
	//maxbet.Scrape("Fudbal")
	//maxbet.Scrape("Košarka")
	//maxbet.Scrape("Tenis")
	//maxbet.Scrape("Stoni Tenis")
	//maxbet.Scrape("eSport")
	//
	//fmt.Println(soccerbet.GetSportsCurrentlyOffered())
	//
	//soccerbet.Scrape("Fudbal")
	//soccerbet.Scrape("Košarka")
	//soccerbet.Scrape("Tenis")
	//
	//fmt.Println(mozzart.GetSportsCurrentlyOffered())
	//
	//mozzart.Scrape("Tenis")
	//mozzart.Scrape("Košarka")
	//mozzart.Scrape("Fudbal")

	fmt.Println(merkurxtip.GetSportsCurrentlyOffered())

	//response := requests_to_server.GetSidebarSports()
	//
	//for k, _ := range response {
	//	fmt.Println(k)
	//}
}
