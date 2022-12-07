package main

import (
	"OISA_2x_sistem/maxbet"
	"OISA_2x_sistem/merge"
	"OISA_2x_sistem/merkurxtip"
	"OISA_2x_sistem/mozzart"
	"OISA_2x_sistem/soccerbet"
	"fmt"
)

func main() {
	fmt.Println("Available sports: ")
	fmt.Println("maxbet:", maxbet.GetSportsCurrentlyOffered())
	fmt.Println("soccerbet:", soccerbet.GetSportsCurrentlyOffered())
	fmt.Println("mozzart:", mozzart.GetSportsCurrentlyOffered())
	fmt.Println("merkurxtip:", merkurxtip.GetSportsCurrentlyOffered())

	sportsToScrape := [...]string{
		//"Košarka",
		//"Tenis",
		"Fudbal",
	}

	bookies := []string{"mozzart", "maxbet", "soccerbet", "merkurxtip"}

	for _, sport := range sportsToScrape {
		scrapedData := map[string][]*[8]string{}

		for _, bookie := range bookies {
			scraper := getScraper(bookie)
			scrapedData[bookie] = scraper(sport)
		}

		mergedData := merge.Merge(sport, scrapedData)
		for _, row := range mergedData {
			fmt.Println(row)
		}
	}

}

func getScraper(bookie string) func(sport string) []*[8]string {
	switch bookie {
	case "mozzart":
		return mozzart.Scrape
	case "maxbet":
		return maxbet.Scrape
	case "soccerbet":
		return soccerbet.Scrape
	case "merkurxtip":
		return merkurxtip.Scrape
	default:
		panic("Bookie not supported")
	}
}
