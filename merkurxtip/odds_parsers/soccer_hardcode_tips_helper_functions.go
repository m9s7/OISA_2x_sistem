package odds_parsers

import (
	"OISA_2x_sistem/merkurxtip/requests_to_server"
	"strings"
)

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
// GA_0_X GA_X_PLUS
// G_NEDAJE_GOST GA_1_PLUS

// Ukupno golova gost (away, tim2) 1. poluvreme
// GA_0_X_P1 G1AX_PLUS
// G1A0 G1A1_PLUS

// Ukupno golova gost (away, tim2) 2. poluvreme
// GA_0_X_P2 G2AX_PLUS
// G2A0 G2A1_PLUS

// GG i NG
// GGG GNG

type tipFormat struct {
	length int
	prefix string
	suffix string
	pair   *tipFormat
}

func getSoccerTipFormats() ([]tipFormat, []tipFormat) {

	// if it ends in _PLUS it goes in the 2nd column, exception GNG
	tip1Format := []tipFormat{
		// UG fuck ups: GH_PLUS (code 356), GA_PLUS (358), GP_PLUS (229)
		{length: 4, prefix: "G0_", suffix: "", pair: &tipFormat{length: 7, prefix: "G", suffix: "_PLUS"}},        // UG
		{length: 6, prefix: "G1P0_", suffix: "", pair: &tipFormat{length: 9, prefix: "G1P", suffix: "_PLUS"}},    // UG 1P
		{length: 6, prefix: "G2P0_", suffix: "", pair: &tipFormat{length: 9, prefix: "G2P", suffix: "_PLUS"}},    // UG 2P
		{length: 6, prefix: "GH_0_", suffix: "", pair: &tipFormat{length: 9, prefix: "GH_", suffix: "_PLUS"}},    // UG D
		{length: 9, prefix: "GH_0_", suffix: "_P1", pair: &tipFormat{length: 9, prefix: "G1H", suffix: "_PLUS"}}, // UG D 1P
		{length: 9, prefix: "GH_0_", suffix: "_P2", pair: &tipFormat{length: 9, prefix: "G2H", suffix: "_PLUS"}}, // UG D 2P
		{length: 6, prefix: "GA_0_", suffix: "", pair: &tipFormat{length: 9, prefix: "GA_", suffix: "_PLUS"}},    // UG G
		{length: 9, prefix: "GA_0_", suffix: "_P1", pair: &tipFormat{length: 9, prefix: "G1A", suffix: "_PLUS"}}, // UG G 1P
		{length: 9, prefix: "GA_0_", suffix: "_P2", pair: &tipFormat{length: 9, prefix: "G2A", suffix: "_PLUS"}}, // UG G 2P
		{length: 3, prefix: "G0", suffix: "", pair: &tipFormat{length: 7, prefix: "G1_PLUS", suffix: ""}},        // GO, G1_PLUS nope
		{length: 4, prefix: "G1P0", suffix: "", pair: &tipFormat{length: 9, prefix: "G1P1_PLUS", suffix: ""}},
		{length: 4, prefix: "G2P0", suffix: "", pair: &tipFormat{length: 9, prefix: "G2P1_PLUS", suffix: ""}},
		{length: 16, prefix: "G_NEDAJE_DOMACIN", suffix: "", pair: &tipFormat{length: 9, prefix: "GH_1_PLUS", suffix: ""}}, //GH_1_PLUS nope
		{length: 4, prefix: "G1H0", suffix: "", pair: &tipFormat{length: 9, prefix: "G1H1_PLUS", suffix: ""}},
		{length: 4, prefix: "G2H0", suffix: "", pair: &tipFormat{length: 9, prefix: "G2H1_PLUS", suffix: ""}},           //G2H1_PLUS nope
		{length: 13, prefix: "G_NEDAJE_GOST", suffix: "", pair: &tipFormat{length: 9, prefix: "GA_1_PLUS", suffix: ""}}, //GA_1_PLUS nope
		{length: 4, prefix: "G1A0", suffix: "", pair: &tipFormat{length: 9, prefix: "G1A1_PLUS", suffix: ""}},
		{length: 4, prefix: "G2A0", suffix: "", pair: &tipFormat{length: 9, prefix: "G2A1_PLUS", suffix: ""}}, //G2A1_PLUS nope
		{length: 3, prefix: "GGG", suffix: "", pair: &tipFormat{length: 3, prefix: "GNG", suffix: ""}},
	}

	var tip2Format []tipFormat
	for _, format := range tip1Format {
		tip2Format = append(tip2Format, *format.pair)
	}

	return tip1Format, tip2Format
}

func isTipFormat(tip string, format tipFormat) bool {
	if len(tip) != format.length {
		return false
	}
	if !strings.HasPrefix(tip, format.prefix) {
		return false
	}
	if !strings.HasSuffix(tip, format.suffix) {
		return false
	}
	return true
}

// GetFocusedTipsTypeCodes
func _() (map[string]int, map[string]int) {

	tip1 := map[string]int{}
	tip2 := map[string]int{}

	tip1Format, tip2Format := getSoccerTipFormats()

	allSubgames := requests_to_server.GetAllSubgamesBlocking()
	betPickMap := allSubgames["betPickMap"].(map[string]interface{})

	for k, v := range betPickMap {
		if !strings.HasSuffix(k, "_S") {
			continue
		}
		v := v.(map[string]interface{})

		tipTypeName := v["tipTypeName"].(string)
		tipCode := int(v["tipTypeCode"].(float64))

		if !strings.HasPrefix(tipTypeName, "G") {
			continue
		}

		for _, format := range tip1Format {
			if isTipFormat(tipTypeName, format) {
				tip1[tipTypeName] = tipCode
				break
			}
		}
		for _, format := range tip2Format {
			if isTipFormat(tipTypeName, format) {
				tip2[tipTypeName] = tipCode
				break
			}
		}

	}

	return tip1, tip2
}

func getHardcodedTipTypeCodes() (map[string]map[string]string, []map[string]string) {

	pairs := map[string]map[string]string{
		"21": {
			"tip1Name":            "G0_1",
			"tip2Name":            "G2_PLUS",
			"matchingTipTypeCode": "242",
		},
		"22": {
			"tip1Name":            "G0_2",
			"tip2Name":            "G3_PLUS",
			"matchingTipTypeCode": "24",
		},
		"219": {
			"tip1Name":            "G0_3",
			"tip2Name":            "G4_PLUS",
			"matchingTipTypeCode": "25",
		},
		"453": {
			"tip1Name":            "G0_4",
			"tip2Name":            "G5_PLUS",
			"matchingTipTypeCode": "27",
		},
		"266": {
			"tip1Name":            "G0_5",
			"tip2Name":            "G6_PLUS",
			"matchingTipTypeCode": "223",
		},
		"338": {
			"tip1Name":            "G1A0",
			"tip2Name":            "G1A1_PLUS",
			"matchingTipTypeCode": "308",
		},
		"337": {
			"tip1Name":            "G1H0",
			"tip2Name":            "G1H1_PLUS",
			"matchingTipTypeCode": "307",
		},
		"267": {
			"tip1Name":            "G1P0",
			"tip2Name":            "G1P1_PLUS",
			"matchingTipTypeCode": "207",
		},
		"211": {
			"tip1Name":            "G1P0_1",
			"tip2Name":            "G1P2_PLUS",
			"matchingTipTypeCode": "208",
		},
		"472": {
			"tip1Name":            "G1P0_2",
			"tip2Name":            "G1P3_PLUS",
			"matchingTipTypeCode": "209",
		},
		"269": {
			"tip1Name":            "G2P0",
			"tip2Name":            "G2P1_PLUS",
			"matchingTipTypeCode": "213",
		},
		"217": {
			"tip1Name":            "G2P0_1",
			"tip2Name":            "G2P2_PLUS",
			"matchingTipTypeCode": "214",
		},
		"474": {
			"tip1Name":            "G2P0_2",
			"tip2Name":            "G2P3_PLUS",
			"matchingTipTypeCode": "215",
		},
		"249": {
			"tip1Name":            "GA_0_1",
			"tip2Name":            "GA_2_PLUS",
			"matchingTipTypeCode": "250",
		},
		"553": {
			"tip1Name":            "GH_0_3",
			"tip2Name":            "GH_4_PLUS",
			"matchingTipTypeCode": "555",
		},
		"554": {
			"tip1Name":            "GA_0_3",
			"tip2Name":            "GA_4_PLUS",
			"matchingTipTypeCode": "556",
		},
		"272": {
			"tip1Name":            "GGG",
			"tip2Name":            "GNG",
			"matchingTipTypeCode": "273",
		},
		"247": {
			"tip1Name":            "GH_0_1",
			"tip2Name":            "GH_2_PLUS",
			"matchingTipTypeCode": "248",
		},
		"817": {
			"tip1Name":            "GH_0_1_P1",
			"tip2Name":            "",
			"matchingTipTypeCode": "",
		},
		"825": {
			"tip1Name":            "GH_0_1_P2",
			"tip2Name":            "",
			"matchingTipTypeCode": "",
		},
		"551": {
			"tip1Name":            "GH_0_2",
			"tip2Name":            "",
			"matchingTipTypeCode": "",
		},
		"821": {
			"tip1Name":            "GA_0_1_P1",
			"tip2Name":            "",
			"matchingTipTypeCode": "",
		},
		"829": {
			"tip1Name":            "GA_0_1_P2",
			"tip2Name":            "",
			"matchingTipTypeCode": "",
		},
		"552": {
			"tip1Name":            "GA_0_2",
			"tip2Name":            "",
			"matchingTipTypeCode": "",
		},
		"239": {
			"tip1Name":            "G_NEDAJE_DOMACIN",
			"tip2Name":            "",
			"matchingTipTypeCode": "",
		},
		"241": {
			"tip1Name":            "G_NEDAJE_GOST",
			"tip2Name":            "",
			"matchingTipTypeCode": "",
		},
		"339": {
			"tip1Name":            "G2H0",
			"tip2Name":            "",
			"matchingTipTypeCode": "",
		},
		"340": {
			"tip1Name":            "G2A0",
			"tip2Name":            "",
			"matchingTipTypeCode": "",
		},
	}

	leftovers := []map[string]string{
		{
			"tip1Name":            "G1P0_3",
			"tip2Name":            "G1P4_PLUS",
			"matchingTipTypeCode": "210",
		},
		{
			"tip1Name":            "GA_0_1_P2",
			"tip2Name":            "G2A2_PLUS",
			"matchingTipTypeCode": "298",
		},
		{
			"tip1Name":            "GA_0_2_P2",
			"tip2Name":            "G2A3_PLUS",
			"matchingTipTypeCode": "352",
		},
		{
			"tip1Name":            "GA_0_3_P2",
			"tip2Name":            "G2A4_PLUS",
			"matchingTipTypeCode": "50241",
		},
		{
			"tip1Name":            "GH_0_1_P2",
			"tip2Name":            "G2H2_PLUS",
			"matchingTipTypeCode": "297",
		},
		{
			"tip1Name":            "GH_0_2_P2",
			"tip2Name":            "G2H3_PLUS",
			"matchingTipTypeCode": "351",
		},
		{
			"tip1Name":            "GH_0_3_P2",
			"tip2Name":            "G2H4_PLUS",
			"matchingTipTypeCode": "50231",
		},
		{
			"tip1Name":            "G0_6",
			"tip2Name":            "G7_PLUS",
			"matchingTipTypeCode": "28",
		},
		{
			"tip1Name":            "G2P0_3",
			"tip2Name":            "G2P4_PLUS",
			"matchingTipTypeCode": "216",
		},
		{
			"tip1Name":            "GH_0_1_P1",
			"tip2Name":            "G1H2_PLUS",
			"matchingTipTypeCode": "274",
		},
		{
			"tip1Name":            "GH_0_2_P1",
			"tip2Name":            "G1H3_PLUS",
			"matchingTipTypeCode": "349",
		},
		{
			"tip1Name":            "GH_0_3_P1",
			"tip2Name":            "G1H4_PLUS",
			"matchingTipTypeCode": "50229",
		},
		{
			"tip1Name":            "GA_0_1_P1",
			"tip2Name":            "G1A2_PLUS",
			"matchingTipTypeCode": "275",
		},
		{
			"tip1Name":            "GA_0_2_P1",
			"tip2Name":            "G1A3_PLUS",
			"matchingTipTypeCode": "350",
		},
		{
			"tip1Name":            "GA_0_3_P1",
			"tip2Name":            "G1A4_PLUS",
			"matchingTipTypeCode": "50239",
		},
	}

	return pairs, leftovers
}

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
