package service

import (
	"OISA_2x_sistem/service/premium_services"
	"OISA_2x_sistem/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
	"strings"
	"time"
)

func ProvidePremiumService() {
	bot, err := tgbotapi.NewBotAPI(telegram.GetBotToken())
	if err != nil {
		log.Panic(err)
	}

	premium_services.ChatIDs = premium_services.LoadPremiumUsers()

	//bot.Debug = true

	log.Printf("Authorized on account %s\n", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	time.Sleep(time.Millisecond * 500)
	updates.Clear()

	log.Printf("Begin work")
	for update := range updates {
		if update.Message == nil || !update.Message.IsCommand() {
			continue
		}
		command := update.Message.Command()
		commandText := update.Message.Text
		chatID := update.Message.Chat.ID

		// register
		if strings.EqualFold(command, "register") {
			newUser, err := premium_services.RegisterNewUser(commandText, chatID)
			if err != nil {
				telegram.BroadcastToDev(err.Error(), "HTML")
			} else {
				telegram.BroadcastToDev("New user "+newUser+" registered", "HTML")
			}
			//:TODO lock this when writing to it, but it happens so rarely that it's not worth it
			premium_services.ChatIDs = append(premium_services.ChatIDs, newUser)
			continue
		}

		if update.Message.ReplyToMessage == nil {
			continue
		}

		arbString := update.Message.ReplyToMessage.Text

		// ulog
		if strings.EqualFold(command, "ulog") {
			reply, err := premium_services.ProvideArbCalculatorService(arbString, commandText)
			if err != nil {
				telegram.BroadcastToDev("ARB CALC service error: "+err.Error()+"\nmsg sent: "+commandText, "HTML")
				continue
			}

			response := telegram.ReplyToMsg(reply, update.Message.MessageID, chatID)
			telegram.CheckIfSent(response, "Calculating ulog good, sending back failed", strconv.FormatInt(chatID, 10))
		}

		// extra
		if strings.EqualFold(command, "extra") {
			reply := premium_services.GenerateArbExtrasReply(arbString)
			response := telegram.ReplyToMsg(reply, update.Message.MessageID, chatID)
			telegram.CheckIfSent(response, "Error while sending arb extra details", strconv.FormatInt(chatID, 10))
		}

	}
}
