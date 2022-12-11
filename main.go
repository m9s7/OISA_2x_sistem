package main

import (
	"OISA_2x_sistem/find_arb"
	"OISA_2x_sistem/merge"
	"OISA_2x_sistem/scrape"
	"OISA_2x_sistem/telegram"
	"OISA_2x_sistem/utility"
	"log"
)

func main() {

	sportsToScrape := []string{
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

	oldArbs := map[string][]find_arb.Arb{
		utility.Tennis:     nil,
		utility.Basketball: nil,
		utility.Soccer:     nil,
	}

	go telegram.ProvidePremiumService()
	//fmt.Println(runtime.GOMAXPROCS(-1))

	log.Println("Starting scraping...")
	for {

		sportsAtBookie := scrape.GetSportsCurrentlyOfferedAtEachBookie(bookies)
		for _, sport := range sportsToScrape {

			if !scrape.IsInAtLeast2Bookies(sport, sportsAtBookie) {
				continue
			}

			scrapedData := scrape.ScrapeDataFromEachBookie(sportsAtBookie, sport)

			mergedData := merge.Merge(sport, scrapedData)
			if merge.IsEmpty(mergedData) {
				continue
			}
			//merge.PrintMergedData(mergedData)

			arbs := find_arb.FindArb(mergedData)
			find_arb.BroadcastNewArbs(arbs, oldArbs, sport, telegram.ChatIDs)
		}
		//telegram.BroadcastToPremium(find_arb.PremiumArbToString(find_arb.GetExampleArbitrage(), "EXAMPLE SPORT"))
	}

}
