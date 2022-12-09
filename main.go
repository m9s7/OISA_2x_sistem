package main

import (
	"OISA_2x_sistem/find_arb"
	"OISA_2x_sistem/merge"
	"OISA_2x_sistem/scrape"
	"OISA_2x_sistem/utility"
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

	for {

		sportsAtBookie := scrape.GetSportsCurrentlyOfferedAtEachBookie(bookies)

		for _, sport := range sportsToScrape {

			if !scrape.IsInAtLeast2Bookies(sport, sportsAtBookie) {
				continue
			}

			scrapedSportAtBookie := scrape.ScrapeDataFromEachBookie(sportsAtBookie, sport)

			mergedData := merge.Merge(sport, scrapedSportAtBookie)
			if merge.IsEmpty(mergedData) {
				continue
			}
			//merge.PrintMergedData(mergedData)

			arbs := find_arb.FindArb(mergedData)
			find_arb.BroadcastNewArbs(arbs, oldArbs, sport)
			//telegram.BroadcastToDev(find_arb.ArbToString(find_arb.GetExampleArbitrage(), "EXAMPLE SPORT"))
		}
	}

}
