package main

import (
	"OISA_2x_sistem/merkurxtip"
	"fmt"
)

func main() {
	//fmt.Println(maxbet.GetSportsCurrentlyOffered())
	//maxbet.Scrape("Fudbal")
	//maxbet.Scrape("Košarka")
	//maxbet.Scrape("Tenis")

	//fmt.Println(soccerbet.GetSportsCurrentlyOffered())
	//soccerbet.Scrape("Fudbal")
	//soccerbet.Scrape("Košarka")
	//soccerbet.Scrape("Tenis")

	//fmt.Println(mozzart.GetSportsCurrentlyOffered())
	//mozzart.Scrape("Tenis")
	//mozzart.Scrape("Košarka")
	//mozzart.Scrape("Fudbal")

	fmt.Println(merkurxtip.GetSportsCurrentlyOffered())
	//
	//merkurxtip.Scrape("Tenis")
	//merkurxtip.Scrape("Košarka")
	merkurxtip.Scrape("Fudbal")

}

// Ukupno golova
// G0_X GX_PLUS
// ??nope   G1_PLUS

// Ukupno golova 1. poluvreme
// G1P0_X G1PX_PLUS
// G1P0 G1P1_PLUS

// Ukupno golova 2. poluvreme
// G2P0_X G2PX_PLUS
// G2P0 G2P1_PLUS

// Ukupno golova domacin (home, tim1)
// GH_0_X GH_X_PLUS
// G_NEDAJE_DOMACIN GH_1_PLUS

// Ukupno golova domacin (home, tim1) 1. poluvreme
// GH_0_X_P1 G1HX_PLUS
// G1H0 G1H1_PLUS

// Ukupno golova domacin (home, tim1) 2. poluvreme
// GH_0_X_P2 G2HX_PLUS
// G2H0 G2H1_PLUS

// Ukupno golova gost (away, tim2)
// Ukupno golova gost (away, tim2) 1. poluvreme
// Ukupno golova gost (away, tim2) 2. poluvreme
// isto samo zamenis H sa A
//G_NEDAJE_GOST

// GG NG
// GGG GNG
// GG NG 1/2 Not offered

//if _, err := strconv.Atoi(string(tipTypeName[1])); err == nil &&
//	strings.HasPrefix(tipTypeName, "G") &&
//	strings.HasSuffix(tipTypeName, "_PLUS") &&
//	len(tipTypeName) == 7 {
//	fmt.Println(tipTypeName)
//}

//	G0_1 21 - G2_PLUS 242
//	G0_2 22 - G3_PLUS 24
//	G0_3 219 - G4_PLUS 25
//	G0_4 453 - G5_PLUS 27
//	G0_5 266 - G6_PLUS 223
//	G1A0 338 - G1A1_PLUS 308
//	G1H0 337 - G1H1_PLUS 307
//	G1P0 267 - G1P1_PLUS 207
//	G1P0_1 211 - G1P2_PLUS 208
//	G1P0_2 472 - G1P3_PLUS 209
//	G2P0 269 - G2P1_PLUS 213
//	G2P0_1 217 - G2P2_PLUS 214
//	G2P0_2 474 - G2P3_PLUS 215
//	GA_0_1 249 - GA_2_PLUS 250
//	GH_0_3 553 - GH_4_PLUS 555
//	GA_0_3 554 - GA_4_PLUS 556
//	GGG 272 - GNG 273
//	GH_0_1 247 - GH_2_PLUS 248
//
//	GH_0_1_P1 817 -
//	GH_0_1_P2 825 -
//	GH_0_2 551 -
//	GA_0_1_P1 821 -
//	GA_0_1_P2 829 -
//	GA_0_2 552 -
//	G_NEDAJE_DOMACIN 239 -
//	G_NEDAJE_GOST 241 -
//	G2H0 339 -
//	G2A0 340 -
//
//			   - G1P4_PLUS 210
//				- G2A2_PLUS 298
//				- G2A3_PLUS 352
//				- G2A4_PLUS 50241
//				- G2H2_PLUS 297
//				- G2H3_PLUS 351
//				- G2H4_PLUS 50231
//				- G7_PLUS 28
//	 			- G2P4_PLUS 216
//				- G1H2_PLUS 274
//				- G1H3_PLUS 349
//				- G1H4_PLUS 50229
//				- G1A2_PLUS 275
//				- G1A3_PLUS 350
//				- G1A4_PLUS 50239
