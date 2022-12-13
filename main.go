package main

import (
	"OISA_2x_sistem/arbitrage"
	"OISA_2x_sistem/merge"
	"OISA_2x_sistem/scrape"
	"OISA_2x_sistem/service"
	"OISA_2x_sistem/service/find_arbs"
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

	service.OldArbsBySport = map[string][]arbitrage.Arb{
		utility.Tennis:     nil,
		utility.Basketball: nil,
		utility.Soccer:     nil,
	}

	go service.ProvidePremiumService()
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

			arbs := find_arbs.FindArb(mergedData, sport)
			service.BroadcastNewArbs(arbs, sport)
		}
		//telegram.BroadcastToPremium(arbitrage.ToStringPremium(arbitrage.GetExampleArbitrage(), "EXAMPLE SPORT"))
	}

}
