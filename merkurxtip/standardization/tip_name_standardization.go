package standardization

import (
	"OISA_2x_sistem/utility"
	"strings"
)

func standardizeTipNameBasketball(tip string) string {
	if len(tip) == 0 {
		return ""
	}

	switch tip {
	case "FT_OT_1":
		return "KI_1_w/OT"
	case "FT_OT_2":
		return "KI_2_w/OT"
	default:
		panic("Unexpected tip name: " + tip)
	}
}

func standardizeTipNameTennis(tip string) string {
	if len(tip) == 0 {
		return ""
	}

	switch tip {
	case "S1_1":
		return "FST_SET_1"
	case "S1_2":
		return "FST_SET_2"
	case "S2_1":
		return "SND_SET_1"
	case "S2_2":
		return "SND_SET_2"
	case "TIE_BREAK_S1_YES":
		return "TIE_BREAK_FST_SET_YES"
	case "TIE_BREAK_S1_NO":
		return "TIE_BREAK_FST_SET_NO"
	case "TIE_BREAK_S2_YES":
		return "TIE_BREAK_SND_SET_YES"
	case "TIE_BREAK_S2_NO":
		return "TIE_BREAK_SND_SET_NO"
	}

	if utility.IsElInSliceSTR(tip, []string{"KI_1", "KI_2", "TIE_BREAK_YES", "TIE_BREAK_NO"}) {
		return tip
	}

	panic("Unexpected tip_name: " + tip)
}

func standardizeTipNameSoccer(tip string) string {
	if len(tip) == 0 {
		return ""
	}

	tip = strings.Trim(tip, " ")
	tipLen := len(tip)

	if strings.HasSuffix(tip, "_PLUS") {
		switch {

		case strings.HasPrefix(tip, "G") && tipLen == 7:
			return "UG_" + string(tip[1]) + "+"
		case strings.HasPrefix(tip, "G1P") && tipLen == 9:
			return "UG_1P_" + string(tip[3]) + "+"
		case strings.HasPrefix(tip, "G2P") && tipLen == 9:
			return "UG_2P_" + string(tip[3]) + "+"

		case strings.HasPrefix(tip, "GH_") && tipLen == 9:
			return "UG_TIM1_" + string(tip[3]) + "+"
		case strings.HasPrefix(tip, "G1H") && tipLen == 9:
			return "UG_1P_TIM1_" + string(tip[3]) + "+"
		case strings.HasPrefix(tip, "G2H") && tipLen == 9:
			return "UG_2P_TIM1_" + string(tip[3]) + "+"

		case strings.HasPrefix(tip, "GA_") && tipLen == 9:
			return "UG_TIM2_" + string(tip[3]) + "+"
		case strings.HasPrefix(tip, "G1A") && tipLen == 9:
			return "UG_1P_TIM2_" + string(tip[3]) + "+"
		case strings.HasPrefix(tip, "G2A") && tipLen == 9:
			return "UG_2P_TIM2_" + string(tip[3]) + "+"

		default:
			panic("Unexpected tip_name: " + tip)
		}
	}

	switch {

	case strings.HasPrefix(tip, "G0_") && tipLen == 4:
		return "UG_0-" + string(tip[3])
	case strings.HasPrefix(tip, "G1P0_") && tipLen == 6:
		return "UG_1P_0-" + string(tip[5])
	case strings.HasPrefix(tip, "G2P0_") && tipLen == 6:
		return "UG_2P_0-" + string(tip[5])

	case strings.HasPrefix(tip, "GH_0_") && tipLen == 6:
		return "UG_TIM1_0-" + string(tip[5])
	case strings.HasPrefix(tip, "GH_0_") && strings.HasSuffix(tip, "_P1") && tipLen == 9:
		return "UG_1P_TIM1_0-" + string(tip[5])
	case strings.HasPrefix(tip, "GH_0_") && strings.HasSuffix(tip, "_P2") && tipLen == 9:
		return "UG_2P_TIM1_0-" + string(tip[5])

	case strings.HasPrefix(tip, "GA_0_") && tipLen == 6:
		return "UG_TIM2_0-" + string(tip[5])
	case strings.HasPrefix(tip, "GA_0_") && strings.HasSuffix(tip, "_P1") && tipLen == 9:
		return "UG_1P_TIM2_0-" + string(tip[5])
	case strings.HasPrefix(tip, "GA_0_") && strings.HasSuffix(tip, "_P2") && tipLen == 9:
		return "UG_2P_TIM2_0-" + string(tip[5])
	}

	switch tip {

	case "G1P0":
		return "UG_1P_T0"
	case "G2P0":
		return "UG_2P_T0"

	case "G_NEDAJE_DOMACIN":
		return "UG_TIM1_T0"
	case "G1H0":
		return "UG_1P_TIM1_T0"
	case "G2H0":
		return "UG_2P_TIM1_T0"

	case "G_NEDAJE_GOST":
		return "UG_TIM2_T0"
	case "G1A0":
		return "UG_1P_TIM2_T0"
	case "G2A0":
		return "UG_2P_TIM2_T0"

	case "GGG":
		return "GG"
	case "GNG":
		return "NG"
		
	default:
		panic("Unexpected tip_name: " + tip)
	}
}
