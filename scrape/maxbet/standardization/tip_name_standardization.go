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

	if strings.HasSuffix(tip, "+") {
		switch {

		case strings.HasPrefix(tip, "ug ") && tipLen == 5:
			return "UG_" + string(tip[3]) + "+"
		case strings.HasPrefix(tip, "ug 1P") && tipLen == 7:
			return "UG_1P_" + string(tip[5]) + "+"
		case strings.HasPrefix(tip, "ug 2P") && tipLen == 7:
			return "UG_2P_" + string(tip[5]) + "+"

		case strings.HasPrefix(tip, "D") && tipLen == 3:
			return "UG_TIM1_" + string(tip[1]) + "+"
		case strings.HasPrefix(tip, "1D") && tipLen == 4:
			return "UG_1P_TIM1_" + string(tip[2]) + "+"
		case strings.HasPrefix(tip, "2D") && tipLen == 4:
			return "UG_2P_TIM1_" + string(tip[2]) + "+"

		case strings.HasPrefix(tip, "G") && tipLen == 3:
			return "UG_TIM2_" + string(tip[1]) + "+"
		case strings.HasPrefix(tip, "1G") && tipLen == 4:
			return "UG_1P_TIM2_" + string(tip[2]) + "+"
		case strings.HasPrefix(tip, "2G") && tipLen == 4:
			return "UG_2P_TIM2_" + string(tip[2]) + "+"

		default:
			panic("Unexpected tip_name: " + tip)
		}
	}

	switch {

	case strings.HasPrefix(tip, "ug 0-") && tipLen == 6:
		return "UG_0-" + string(tip[5])
	case strings.HasPrefix(tip, "ug 1P0-") && tipLen == 8:
		return "UG_1P_0-" + string(tip[7])
	case strings.HasPrefix(tip, "ug 2P0-") && tipLen == 8:
		return "UG_2P_0-" + string(tip[7])

	case strings.HasPrefix(tip, "D0-") && tipLen == 4:
		return "UG_TIM1_0-" + string(tip[3])
	case strings.HasPrefix(tip, "1D0-") && tipLen == 5:
		return "UG_1P_TIM1_0-" + string(tip[4])
	case strings.HasPrefix(tip, "2D0-") && tipLen == 5:
		return "UG_2P_TIM1_0-" + string(tip[4])

	case strings.HasPrefix(tip, "G0-") && tipLen == 4:
		return "UG_TIM2_0-" + string(tip[3])
	case strings.HasPrefix(tip, "1G0-") && tipLen == 5:
		return "UG_1P_TIM2_0-" + string(tip[4])
	case strings.HasPrefix(tip, "2G0-") && tipLen == 5:
		return "UG_2P_TIM2_0-" + string(tip[4])
	}

	switch tip {
	case "ug 1PT0":
		return "UG_1P_T0"
	case "ug 2PT0":
		return "UG_2P_T0"
	case "D0":
		return "UG_TIM1_T0"
	case "1D0":
		return "UG_1P_TIM1_T0"
	case "2D0":
		return "UG_2P_TIM1_T0"
	case "G0":
		return "UG_TIM2_T0"
	case "1G0":
		return "UG_1P_TIM2_T0"
	case "2G0":
		return "UG_2P_TIM2_T0"
	case "GG":
		fallthrough
	case "NG":
		fallthrough
	case "GG1":
		fallthrough
	case "NG1":
		fallthrough
	case "GG2":
		fallthrough
	case "NG2":
		return tip
	default:
		panic("Unexpected tip_name: " + tip)
	}

}
