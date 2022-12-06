package standardization

import "strings"

func standardizeTipNameBasketball(tip string) string {
	//:TODO this is ok, just change in merge function that only one tip needs to match not both
	if len(tip) == 0 {
		return ""
	}

	switch tip {
	case "pobm 1":
		return "KI_1_w/OT"
	case "pobm 2":
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
	case "ki 1":
		return "KI_1"
	case "ki 2":
		return "KI_2"
	case "1s 1":
		return "FST_SET_1"
	case "1s 2":
		return "FST_SET_2"
	case "2s 1":
		return "SND_SET_1"
	case "2s 2":
		return "SND_SET_2"
	case "ug1s da 13":
		return "TIE_BREAK_FST_SET_YES"
	case "ug1s ne 13":
		return "TIE_BREAK_FST_SET_NO"
	case "ug2s da 13":
		return "TIE_BREAK_SND_SET_YES"
	case "ug2s ne 13":
		return "TIE_BREAK_SND_SET_NO"
	case "tb da":
		return "TIE_BREAK_YES"
	case "tb ne":
		return "TIE_BREAK_NO"
	default:
		panic("Unexpected tip name: " + tip)
	}
}

func standardizeTipNameSoccer(tip string) string {
	if len(tip) == 0 {
		return ""
	}
	tipLen := len(tip)

	if strings.HasSuffix(tip, "+") {
		switch {
		case strings.HasPrefix(tip, "ug ") && tipLen == 5:
			return "UG_" + string(tip[3]) + "+"
		case strings.HasPrefix(tip, "1ug ") && tipLen == 6:
			return "UG_1P_" + string(tip[4]) + "+"
		case strings.HasPrefix(tip, "2ug ") && tipLen == 6:
			return "UG_2P_" + string(tip[4]) + "+"
		case strings.HasPrefix(tip, "tm1 ") && tipLen == 6:
			return "UG_TIM1_" + string(tip[4]) + "+"
		case strings.HasPrefix(tip, "1tm1 ") && tipLen == 7:
			return "UG_1P_TIM1_" + string(tip[5]) + "+"
		case strings.HasPrefix(tip, "2tm1 ") && tipLen == 7:
			return "UG_2P_TIM1_" + string(tip[5]) + "+"
		case strings.HasPrefix(tip, "tm2 ") && tipLen == 6:
			return "UG_TIM2_" + string(tip[4]) + "+"
		case strings.HasPrefix(tip, "1tm2 ") && tipLen == 7:
			return "UG_1P_TIM2_" + string(tip[5]) + "+"
		case strings.HasPrefix(tip, "2tm2 ") && tipLen == 7:
			return "UG_2P_TIM2_" + string(tip[5]) + "+"
		default:
			panic("Unexpected tip_name: " + tip)
		}
	}

	switch {
	case strings.HasPrefix(tip, "ug 0-") && tipLen == 6:
		return "UG_0-" + string(tip[5])
	case strings.HasPrefix(tip, "1ug 0-") && tipLen == 7:
		return "UG_1P_0-" + string(tip[6])
	case strings.HasPrefix(tip, "2ug 0-") && tipLen == 7:
		return "UG_2P_0-" + string(tip[6])
	case strings.HasPrefix(tip, "tm1 0-") && tipLen == 7:
		return "UG_TIM1_0-" + string(tip[6])
	case strings.HasPrefix(tip, "1tm1 0-") && tipLen == 8:
		return "UG_1P_TIM1_0-" + string(tip[7])
	case strings.HasPrefix(tip, "2tm1 0-") && tipLen == 8:
		return "UG_2P_TIM1_0-" + string(tip[7])
	case strings.HasPrefix(tip, "tm2 0-") && tipLen == 7:
		return "UG_TIM2_0-" + string(tip[6])
	case strings.HasPrefix(tip, "1tm2 0-") && tipLen == 8:
		return "UG_1P_TIM2_0-" + string(tip[7])
	case strings.HasPrefix(tip, "2tm2 0-") && tipLen == 8:
		return "UG_2P_TIM2_0-" + string(tip[7])
	}

	switch tip {
	case "1ug 0":
		return "UG_1P_T0"
	case "2ug 0":
		return "UG_2P_T0"
	case "tm1 0":
		return "UG_TIM1_T0"
	case "1tm1 0":
		return "UG_1P_TIM1_T0"
	case "2tm1 0":
		return "UG_2P_TIM1_T0"
	case "tm2 0":
		return "UG_TIM2_T0"
	case "1tm2 0":
		return "UG_1P_TIM2_T0"
	case "2tm2 0":
		return "UG_2P_TIM2_T0"
	case "gg":
		fallthrough
	case "ng":
		fallthrough
	case "1gg":
		fallthrough
	case "1ng":
		fallthrough
	case "2gg":
		fallthrough
	case "2ng":
		return strings.ToUpper(tip)
	default:
		panic("Unexpected tip name: " + tip)
	}
}
