package odds_parsers

import (
	"OISA_2x_sistem/merkurxtip/requests_to_server"
	"OISA_2x_sistem/utility"
	"fmt"
	"strings"
)

func SoccerOddsParser(matchIDs []int) [][8]string {

	matchesScrapedCounter := 0
	var export [][8]string

	tipTypeCodePairs, leftovers := getTipTypeCodes()

	for _, matchID := range matchIDs {
		match := requests_to_server.GetMatchOdds(matchID)

		e1 := [4]string{
			fmt.Sprintf("%.0f", match["kickOffTime"].(float64)),
			match["leagueName"].(string),
			match["home"].(string),
			match["away"].(string),
		}

		getTipValByTipTypeCode := match["odds"].(map[string]interface{})
		for tip1Code, m := range tipTypeCodePairs {

			tip1Val, ok := getTipValByTipTypeCode[tip1Code]
			if !ok {
				continue
			}
			tip2Val, ok := getTipValByTipTypeCode[m["matchingTipTypeCode"]]
			if !ok {
				continue
			}

			e2 := [4]string{
				m["tip1Name"], fmt.Sprintf("%f", tip1Val.(float64)),
				m["tip2Name"], fmt.Sprintf("%f", tip2Val.(float64)),
			}
			export = append(export, utility.MergeE1E2(e1, e2))
		}

		for tip2Name, tip2Code := range leftovers {
			tip2Val, ok := getTipValByTipTypeCode[tip2Code]
			if !ok {
				continue
			}
			export = append(export, utility.MergeE1E2(e1, [4]string{
				"None", "0.0",
				tip2Name, fmt.Sprintf("%f", tip2Val.(float64)),
			}))
		}

		matchesScrapedCounter++
	}

	fmt.Println("Matches scraped: ", matchesScrapedCounter)
	fmt.Println("Tips scraped: ", len(export))

	return export
}

type tipFormat struct {
	length int
	prefix string
	suffix string
	pair   *tipFormat
}

func getSoccerTipFormats() ([]tipFormat, []tipFormat) {

	// ako se zavrsava na _PLUS ide u drugu kolonu
	tip1Format := []tipFormat{
		// UG fucks up: GH_PLUS (code 356), GA_PLUS (358), GP_PLUS (229)
		{length: 4, prefix: "G0_", suffix: "", pair: &tipFormat{length: 7, prefix: "", suffix: "_PLUS"}},         // UG
		{length: 6, prefix: "G1P0_", suffix: "", pair: &tipFormat{length: 9, prefix: "G1P", suffix: "_PLUS"}},    // UG 1P
		{length: 6, prefix: "G2P0_", suffix: "", pair: &tipFormat{length: 9, prefix: "G2P", suffix: "_PLUS"}},    // UG 2P
		{length: 6, prefix: "GH_0_", suffix: "", pair: &tipFormat{length: 9, prefix: "GH_", suffix: "_PLUS"}},    // UG D
		{length: 9, prefix: "GH_0_", suffix: "_P1", pair: &tipFormat{length: 9, prefix: "G1H", suffix: "_PLUS"}}, // UG D 1P
		{length: 9, prefix: "GH_0_", suffix: "_P2", pair: &tipFormat{length: 9, prefix: "G2H", suffix: "_PLUS"}}, // UG D 2P
		{length: 6, prefix: "GA_0_", suffix: "", pair: &tipFormat{length: 9, prefix: "GA_", suffix: "_PLUS"}},    // UG G
		{length: 9, prefix: "GA_0_", suffix: "_P1", pair: &tipFormat{length: 9, prefix: "G1A", suffix: "_PLUS"}}, // UG G 1P
		{length: 9, prefix: "GA_0_", suffix: "_P2", pair: &tipFormat{length: 9, prefix: "G2A", suffix: "_PLUS"}}, // UG G 2P
		{length: 4, prefix: "G1P0", suffix: "", pair: &tipFormat{length: 9, prefix: "G1P1_PLUS", suffix: ""}},
		{length: 4, prefix: "G2P0", suffix: "", pair: &tipFormat{length: 9, prefix: "G2P1_PLUS", suffix: ""}},
		{length: 16, prefix: "G_NEDAJE_DOMACIN", suffix: "", pair: &tipFormat{length: 9, prefix: "GH_1_PLUS", suffix: ""}}, //GH_1_PLUS nope
		{length: 4, prefix: "G1H0", suffix: "", pair: &tipFormat{length: 9, prefix: "G1H1_PLUS", suffix: ""}},
		{length: 4, prefix: "G2H0", suffix: "", pair: &tipFormat{length: 9, prefix: "G2H1_PLUS", suffix: ""}},           //G2H1_PLUS nope
		{length: 13, prefix: "G_NEDAJE_GOST", suffix: "", pair: &tipFormat{length: 9, prefix: "GA_1_PLUS", suffix: ""}}, //GA_1_PLUS nope
		{length: 4, prefix: "G1A0", suffix: "", pair: &tipFormat{length: 9, prefix: "G1A1_PLUS", suffix: ""}},
		{length: 4, prefix: "G2A0", suffix: "", pair: &tipFormat{length: 9, prefix: "G2A1_PLUS", suffix: ""}}, //G2A1_PLUS nema
		{length: 3, prefix: "GGG", suffix: "", pair: &tipFormat{length: 3, prefix: "GNG", suffix: ""}},
	}

	var tip2Format []tipFormat
	for _, format := range tip1Format {
		tip2Format = append(tip2Format, *format.pair)
	}

	return tip1Format, tip2Format
}

func GetFocusedTipsTypeCodes() (map[string]int, map[string]int) {

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

func getTipTypeCodes() (map[string]map[string]string, map[string]string) {

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

	leftovers := map[string]string{
		"G1P4_PLUS": "210",
		"G2A2_PLUS": "298",
		"G2A3_PLUS": "352",
		"G2A4_PLUS": "50241",
		"G2H2_PLUS": "297",
		"G2H3_PLUS": "351",
		"G2H4_PLUS": "50231",
		"G7_PLUS":   "28",
		"G2P4_PLUS": "216",
		"G1H2_PLUS": "274",
		"G1H3_PLUS": "349",
		"G1H4_PLUS": "50229",
		"G1A2_PLUS": "275",
		"G1A3_PLUS": "350",
		"G1A4_PLUS": "50239",
	}

	return pairs, leftovers
}

// I have the code -> value mapping (are codes unique??)
// - determine is it tip1 or tip2 and which tip1 tip2 pairing (I used game as key and x as val in other books)
