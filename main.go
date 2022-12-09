package main

import (
	"OISA_2x_sistem/arbitrage"
	"OISA_2x_sistem/maxbet"
	"OISA_2x_sistem/merge"
	"OISA_2x_sistem/merkurxtip"
	"OISA_2x_sistem/mozzart"
	"OISA_2x_sistem/soccerbet"
	"OISA_2x_sistem/telegram"
	"OISA_2x_sistem/utility"
	"fmt"
)

func main() {

	fmt.Println("Available sports: ")
	fmt.Println("maxbet:", maxbet.GetSportsCurrentlyOffered())
	fmt.Println("soccerbet:", soccerbet.GetSportsCurrentlyOffered())
	fmt.Println("mozzart:", mozzart.GetSportsCurrentlyOffered())
	fmt.Println("merkurxtip:", merkurxtip.GetSportsCurrentlyOffered())

	sportsToScrape := [...]string{
		utility.Tennis,
		utility.Basketball,
		utility.Soccer,
	}

	bookies := []string{
		"mozzart",
		"maxbet",
		"soccerbet",
		"merkurxtip",
	}

	oldArbs := map[string][]arbitrage.Arb{
		utility.Tennis:     nil,
		utility.Basketball: nil,
		utility.Soccer:     nil,
	}
	for {
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

			arbs := arbitrage.FindArb(mergedData)
			if len(arbs) == 0 {
				telegram.BroadcastToDev(`Nema arbe :\\( \\- ` + sport)
				oldArbs[sport] = nil
				continue
			}
		Loop:
			for _, arb := range arbs {
				for _, oldArb := range oldArbs[sport] {
					if arb.Equals(oldArb) {
						continue Loop
					}
				}
				telegram.BroadcastToDev("FRISKE ARBE")
				telegram.BroadcastToDev(arbitrage.ArbToString(arb, sport))
			}
			oldArbs[sport] = arbs
			//telegram.BroadcastToDev(arbitrage.ArbToString(arbitrage.GetExampleArbitrage(), "EXAMPLE SPORT"))
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
