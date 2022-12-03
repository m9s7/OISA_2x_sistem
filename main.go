package main

import (
	"OISA_2x_sistem/maxbet"
	"fmt"
)

func main() {
	fmt.Println(maxbet.GetSportsCurrentlyOffered())

	maxbet.Scrape("Fudbal")
	maxbet.Scrape("Košarka")
	maxbet.Scrape("Tenis")
	maxbet.Scrape("Stoni Tenis")
	maxbet.Scrape("eSport")
}
