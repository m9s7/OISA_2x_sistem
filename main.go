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
	fmt.Println(soccerbet.GetSportsCurrentlyOffered())
	fmt.Println(mozzart.GetSportsCurrentlyOffered())
	fmt.Println(merkurxtip.GetSportsCurrentlyOffered())

	sportsToScrape := [...]string{
		"Košarka",
		"Tenis",
		"Fudbal",
	}

	for _, sport := range sportsToScrape {
		mozzartData := mozzart.Scrape(sport)
		maxbetData := maxbet.Scrape(sport)
		soccerbetData := soccerbet.Scrape(sport)
		merkurxtipData := merkurxtip.Scrape(sport)

		printData(mozzartData)
		printData(maxbetData)
		printData(soccerbetData)
		printData(merkurxtipData)
	}

}

func printData(data []*[8]string) {
	for _, row := range data {
		fmt.Println(*row)
	}
}
