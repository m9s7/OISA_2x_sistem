package main

import (
	"OISA_2x_sistem/scrape/soccerbet"
	"OISA_2x_sistem/utility"
	"fmt"
)

func main() {

	//sportsToScrape := []string{
	//	utility.Tennis,
	//	utility.Basketball,
	//	utility.Soccer,
	//}
	//
	//bookies := []string{
	//	"mozzart",
	//	"maxbet",
	//	"soccerbet",
	//	"merkurxtip",
	//}
	//
	//find_arbs.OldArbsBySport = map[string][]arbitrage.Arb{
	//	utility.Tennis:     nil,
	//	utility.Basketball: nil,
	//	utility.Soccer:     nil,
	//}
	//
	//go service.ProvidePremiumService()
	////fmt.Println(runtime.GOMAXPROCS(-1))
	//
	//for {
	//
	//	sportsAtBookie := scrape.GetSportsCurrentlyOfferedAtEachBookie(bookies)
	//	for _, sport := range sportsToScrape {
	//
	//		if !scrape.IsInAtLeast2Bookies(sport, sportsAtBookie) {
	//			continue
	//		}
	//
	//		scrapedData := scrape.ScrapeDataFromEachBookie(sportsAtBookie, sport)
	//
	//		mergedData := merge.Merge(sport, scrapedData)
	//		if merge.IsEmpty(mergedData) {
	//			continue
	//		}
	//		//merge.PrintMergedData(mergedData)
	//
	//		arbs := find_arbs.FindArb(mergedData, sport)
	//		broadcastNewArbs(arbs, sport)
	//	}
	//	//telegram.BroadcastToPremium(arbitrage.ToStringPremium(arbitrage.GetExampleArbitrage(), "EXAMPLE SPORT"))
	//}

	export := soccerbet.Scrape(utility.Soccer)
	for _, e := range export {
		fmt.Println(*e)
	}

}

//func broadcastNewArbs(arbs []arbitrage.Arb, sport string) {
//	if len(arbs) == 0 {
//		find_arbs.OldArbsBySport[sport] = nil
//		return
//	}
//	for _, arb := range arbs {
//		if arb.IsIn(find_arbs.OldArbsBySport[sport]) || arb.ROI < 0.1 {
//			continue
//		}
//
//		if arb.ROI <= 1.5 {
//			telegram.BroadcastToFree(arb.ToStringFree())
//		}
//		if arb.ROI >= 1.0 {
//			telegram.BroadcastToPremium(arb.ToStringPremium(), premium_services.ChatIDs)
//		}
//	}
//	find_arbs.OldArbsBySport[sport] = arbs
//}
