package scrape

import (
	"OISA_2x_sistem/scrape/maxbet"
	"OISA_2x_sistem/scrape/merkurxtip"
	"OISA_2x_sistem/scrape/mozzart"
	"OISA_2x_sistem/scrape/soccerbet"
	"fmt"
	"github.com/go-gota/gota/dataframe"
)

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

func GetSportsCurrentlyOffered(bookie string) []string {
	switch bookie {
	case "mozzart":
		return mozzart.GetSportsCurrentlyOffered()
	case "maxbet":
		return maxbet.GetSportsCurrentlyOffered()
	case "soccerbet":
		return soccerbet.GetSportsCurrentlyOffered()
	case "merkurxtip":
		return merkurxtip.GetSportsCurrentlyOffered()
	default:
		panic("Bookie not supported")
	}
}

func GetSportsCurrentlyOfferedAtEachBookie(bookies []string) map[string][]string {
	sports := map[string][]string{}
	for _, bookie := range bookies {
		sports[bookie] = GetSportsCurrentlyOffered(bookie)
	}
	return sports
}

func ScrapeDataFromEachBookie(sportsAtBookie map[string][]string, sport string) map[string][]*[8]string {

	scrapedData := map[string][]*[8]string{}

	for bookie, sports := range sportsAtBookie {
		for _, s := range sports {
			if s != sport {
				continue
			}
			scrapedData[bookie] = getScraper(bookie)(sport)
		}
	}

	return scrapedData
}

func IsInAtLeast2Bookies(sport string, sportsAtBookie map[string][]string) bool {
	count := 0
	for _, sports := range sportsAtBookie {
		for _, s := range sports {
			if s == sport {
				count++
				break
			}
		}
	}
	return count >= 2
}

func PrintScrapedData(scrapedData []*[8]string) {
	var kk [][]string
	for _, record := range scrapedData {
		kk = append(kk, record[:])
		fmt.Println(dataframe.LoadRecords(kk).String())
	}
}
