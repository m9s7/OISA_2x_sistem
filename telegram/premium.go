package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
	"strings"
	"time"
)

func providePremiumService() {
	bot, err := tgbotapi.NewBotAPI("5649589726:AAHu4l02-AA0EsmSn5k-hQqnDQ6jgBoSNqg")
	if err != nil {
		log.Panic(err)
	}

	//bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	// Optional: wait for updates and clear them if you don't want to handle a large backlog of old messages
	time.Sleep(time.Millisecond * 500)
	updates.Clear()

	log.Printf("Begin work")
	for update := range updates {
		if update.Message == nil || !update.Message.IsCommand() {
			continue
		}

		fmt.Println(update.Message.Command())
		if update.Message.Command() == "start" {
			fmt.Println("New user!")
			user := strings.TrimPrefix(update.Message.Text, "/start ")
			fmt.Println(user)
			fmt.Println(update.Message.Chat.ID)
			// Write it to file, or better yet just send it to me on my phone
		}
		// parse command to int continue if it's not int
		capital, err := strconv.Atoi(update.Message.Command())
		if err != nil {
			continue
		}
		t1Investment, t2Investment := arbCalc(capital)

		// Reply to user msg
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, t1Investment+" "+t2Investment)
		msg.ReplyToMessageID = update.Message.MessageID

		_, err = bot.Send(msg)
		if err != nil {
			return
		}
		//fmt.Println(msg.ChatID)
	}
}

func arbCalc(capital int) (string, string) {
	return strconv.Itoa(capital), strconv.Itoa(capital)
}
