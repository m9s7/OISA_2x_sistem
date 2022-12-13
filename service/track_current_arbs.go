package service

import (
	"OISA_2x_sistem/arbitrage"
	"OISA_2x_sistem/telegram"
)

var OldArbsBySport map[string][]arbitrage.Arb

func BroadcastNewArbs(arbs []arbitrage.Arb, sport string) {
	if len(arbs) == 0 {
		OldArbsBySport[sport] = nil
		return
	}
	for _, arb := range arbs {
		if arb.IsIn(OldArbsBySport[sport]) || arb.ROI < 0.1 {
			continue
		}

		if arb.ROI <= 1.5 {
			telegram.BroadcastToFree(arb.ToStringFree())
		}
		if arb.ROI >= 1.0 {
			BroadcastToPremium(arb.ToStringPremium())
		}
	}
	OldArbsBySport[sport] = arbs
}
