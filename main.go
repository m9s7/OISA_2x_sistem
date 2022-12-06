package main

import (
	"OISA_2x_sistem/maxbet"
	"OISA_2x_sistem/merkurxtip"
	"OISA_2x_sistem/mozzart"
	"OISA_2x_sistem/soccerbet"
	"fmt"
)

func main() {
	fmt.Println(maxbet.GetSportsCurrentlyOffered())
	maxbet.Scrape("Fudbal")
	maxbet.Scrape("Košarka")
	maxbet.Scrape("Tenis")

	fmt.Println(soccerbet.GetSportsCurrentlyOffered())
	soccerbet.Scrape("Fudbal")
	soccerbet.Scrape("Košarka")
	soccerbet.Scrape("Tenis")

	fmt.Println(mozzart.GetSportsCurrentlyOffered())
	mozzart.Scrape("Tenis")
	mozzart.Scrape("Košarka")
	mozzart.Scrape("Fudbal")

	fmt.Println(merkurxtip.GetSportsCurrentlyOffered())
	merkurxtip.Scrape("Tenis")
	merkurxtip.Scrape("Košarka")
	merkurxtip.Scrape("Fudbal")

}
