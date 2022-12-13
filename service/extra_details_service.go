package service

import (
	"OISA_2x_sistem/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"strings"
)

// TODO: same as with the calculator, works just with text, no imports, then move them to a new package, premium services

func ProvideArbExtraDetailsService(updateMessage *tgbotapi.Message) {

	repliedToArbString := updateMessage.ReplyToMessage.Text
	reply := generateArbExtrasReply(repliedToArbString)

	response := telegram.ReplyToMsg(reply, updateMessage.MessageID, updateMessage.Chat.ID)
	telegram.CheckIfSent(response, "Error while sending arb extra details", strconv.FormatInt(updateMessage.Chat.ID, 10))
}

func generateArbExtrasReply(repliedToArbString string) string {

	if strings.Contains(repliedToArbString, "deviation table") {
		return repliedToArbString
	}

	for _, oldArbs := range OldArbsBySport {
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
