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
	"github.com/go-gota/gota/dataframe"
)

func main() {
	//printAllAvailableSports()

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

				printScrapedData(scrapedData)
			}

			mergedData := merge.Merge(sport, scrapedData)
			if mergedData == nil || len(mergedData) == 1 {
				continue
			}
			printMergedData(mergedData)

			arbs := arbitrage.FindArb(mergedData)
			broadcastNewArbs(arbs, oldArbs, sport)
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

func printScrapedData(scrapedData map[string][]*[8]string) {
	for bookie, data := range scrapedData {
		fmt.Println(bookie)
		var kk [][]string
		for _, rec := range data {
			kk = append(kk, rec[:])
		}
		fmt.Println(dataframe.LoadRecords(kk).String())
	}
}

func printMergedData(mergedData [][]string) {
	fmt.Println(dataframe.LoadRecords(mergedData).Drop([]int{2, 3, 4}).String())
}

func broadcastNewArbs(arbs []arbitrage.Arb, oldArbs map[string][]arbitrage.Arb, sport string) {
	if len(arbs) == 0 {
		//telegram.BroadcastToDev(`Nema arbe :\\( \\- ` + sport)
		oldArbs[sport] = nil
		return
	}
	for _, arb := range arbs {
		if isArbInOldArbs(arb, oldArbs[sport]) {
			continue
		}
		telegram.BroadcastToDev("FRISKE ARBE")
		telegram.BroadcastToDev(arbitrage.ArbToString(arb, sport))
	}
	oldArbs[sport] = arbs
}

func isArbInOldArbs(arb arbitrage.Arb, oldArbs []arbitrage.Arb) bool {
	for _, oldArb := range oldArbs {
		if arb.Equals(oldArb) {
			return true
		}
	}
	return false
}

func printAllAvailableSports() {
	fmt.Println("Available sports: ")
	fmt.Println("maxbet:", maxbet.GetSportsCurrentlyOffered())
	fmt.Println("soccerbet:", soccerbet.GetSportsCurrentlyOffered())
	fmt.Println("mozzart:", mozzart.GetSportsCurrentlyOffered())
	fmt.Println("merkurxtip:", merkurxtip.GetSportsCurrentlyOffered())
}
