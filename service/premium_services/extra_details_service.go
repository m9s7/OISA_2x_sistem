package premium_services

import (
	"OISA_2x_sistem/service/find_arbs"
	"strings"
)

func GenerateArbExtrasReply(repliedToArbString string) string {

	if strings.Contains(repliedToArbString, "deviation table") {
		return repliedToArbString
	}

	for _, oldArbs := range find_arbs.OldArbsBySport {
		for _, oldArb := range oldArbs {

			oldArbString := strings.Trim(oldArb.ToStringPremium(), "`\n")

			if repliedToArbString != oldArbString {
				continue
			}
			return oldArb.ToStringWithExtra()
		}
	}
	return "Arbitraža je istekla, kvote su se promenile"
}
