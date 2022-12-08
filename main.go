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
		"Košarka",
		"Tenis",
		"Fudbal",
	}

	bookies := []string{
		"mozzart",
		"maxbet",
		"soccerbet",
		"merkurxtip",
	}

	for _, sport := range sportsToScrape {
		scrapedData := map[string][]*[8]string{}

		for _, bookie := range bookies {
			scraper := getScraper(bookie)
			scrapedData[bookie] = scraper(sport)

			//// Print scraped data
			//var kk [][]string
			//for _, rec := range scrapedData[bookie] {
			//	kk = append(kk, rec[:])
			//}
			//fmt.Println(dataframe.LoadRecords(kk).String())
		}

		mergedData := merge.Merge(sport, scrapedData)
		if mergedData == nil || len(mergedData) == 1 {
			continue
		}

		//// Print merged data
		//fmt.Println(dataframe.LoadRecords(mergedData).Drop([]int{2, 3, 4}).String())

		FindArb(mergedData)
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
