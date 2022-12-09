package standardization

import (
	"strings"
)

func standardizeTipNameBasketball(tip string) string {
	if len(tip) == 0 {
		return ""
	}

	switch tip {
	case "KI 1":
		return "KI_1_w/OT"
	case "KI 2":
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

	case "KI 1":
		return "KI_1"
	case "KI 2":
		return "KI_2"

	case "Iset 1":
		return "FST_SET_1"
	case "Iset 2":
		return "FST_SET_2"

	case "TB DA":
		return "TIE_BREAK_YES"
	case "TB NE":
		return "TIE_BREAK_NO"

	default:
		panic("Unexpected tip_name: " + tip)
	}
}

func standardizeTipNameSoccer(tip string) string {
	if len(tip) == 0 {
		return ""
	}

	tip = strings.Trim(tip, " ")
	tipLen := len(tip)

	if strings.HasSuffix(tip, "+") {
		switch {

		case strings.HasPrefix(tip, "UG ") && tipLen == 5:
			return "UG_" + string(tip[3]) + "+"
		case strings.HasPrefix(tip, "UG I") && tipLen == 6:
			return "UG_1P_" + string(tip[4]) + "+"
		case strings.HasPrefix(tip, "UG II") && tipLen == 7:
			return "UG_2P_" + string(tip[5]) + "+"

		case strings.HasPrefix(tip, "D ") && tipLen == 4:
			return "UG_TIM1_" + string(tip[2]) + "+"
		case strings.HasPrefix(tip, "D I") && tipLen == 5:
			return "UG_1P_TIM1_" + string(tip[3]) + "+"
		case strings.HasPrefix(tip, "D II") && tipLen == 6:
			return "UG_2P_TIM1_" + string(tip[4]) + "+"

		case strings.HasPrefix(tip, "G ") && tipLen == 4:
			return "UG_TIM2_" + string(tip[2]) + "+"
		case strings.HasPrefix(tip, "G I") && tipLen == 5:
			return "UG_1P_TIM2_" + string(tip[3]) + "+"
		case strings.HasPrefix(tip, "G II") && tipLen == 6:
			return "UG_2P_TIM2_" + string(tip[4]) + "+"

		default:
			panic("Unexpected tip_name: " + tip)
		}
	}

	switch {

	case strings.HasPrefix(tip, "UG 0-") && tipLen == 6:
		return "UG_0-" + string(tip[5])
	case strings.HasPrefix(tip, "UG I0-") && tipLen == 7:
		return "UG_1P_0-" + string(tip[6])
	case strings.HasPrefix(tip, "UG II0-") && tipLen == 8:
		return "UG_2P_0-" + string(tip[7])

	case strings.HasPrefix(tip, "D 0-") && tipLen == 5:
		return "UG_TIM1_0-" + string(tip[4])
	case strings.HasPrefix(tip, "D I0-") && tipLen == 6:
		return "UG_1P_TIM1_0-" + string(tip[5])
	case strings.HasPrefix(tip, "D II0-") && tipLen == 7:
		return "UG_2P_TIM1_0-" + string(tip[6])

	case strings.HasPrefix(tip, "G 0-") && tipLen == 5:
		return "UG_TIM2_0-" + string(tip[4])
	case strings.HasPrefix(tip, "G I0-") && tipLen == 6:
		return "UG_1P_TIM2_0-" + string(tip[5])
	case strings.HasPrefix(tip, "G II0-") && tipLen == 7:
		return "UG_2P_TIM2_0-" + string(tip[6])
	}

	switch tip {
	case "UG I0":
		return "UG_1P_T0"
	case "UG II0":
		return "UG_2P_T0"
	case "D 0":
		return "UG_TIM1_T0"
	case "D I0":
		return "UG_1P_TIM1_T0"
	case "D II0":
		return "UG_2P_TIM1_T0"
	case "G 0":
		return "UG_TIM2_T0"
	case "G I0":
		return "UG_1P_TIM2_T0"
	case "G II0":
		return "UG_2P_TIM2_T0"
	case "GG":
		fallthrough
	case "NG":
		return tip
	default:
		panic("Unexpected tip_name: " + tip)
	}

}
