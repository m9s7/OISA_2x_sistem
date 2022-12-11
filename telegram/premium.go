package telegram

import (
	"OISA_2x_sistem/utility"
	"bufio"
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var ChatIDs []string

func ProvidePremiumService() {
	bot, err := tgbotapi.NewBotAPI(GetBotToken())
	if err != nil {
		log.Panic(err)
	}

	ChatIDs = loadPremiumUsers("telegram\\users.txt")

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

		// register
		if strings.EqualFold(update.Message.Command(), "register") {
			newUser, err := registerNewUser(update.Message.Text, update.Message.Chat.ID)
			if err != nil {
				log.Println(err.Error())
				BroadcastToDev("Error while writing user to file: "+err.Error(), "HTML")
			}
			ChatIDs = append(ChatIDs, newUser)
			continue
		}

		// ulog
		if update.Message.ReplyToMessage != nil && strings.EqualFold(update.Message.Command(), "ulog") {
			msg, err := provideArbCalculatorService(update.Message)
			if err != nil {
				log.Println(err.Error())
				BroadcastToDev("Error while providing arb calculator service: "+err.Error(), "HTML")
				continue
			}
			_, err = bot.Send(msg)
			if err != nil {
				return
			}
		}

	}
}

func registerNewUser(msgText string, chatID int64) (string, error) {

	user := strings.TrimPrefix(msgText, "/register ")
	chatIDStr := strconv.Itoa(int(chatID))

	BroadcastToDev(
		EscapeTelegramSpecChars("User: "+user+" has started the bot!\nChat ID: "+chatIDStr), "HTML")

	err := utility.AppendToFile("telegram\\users.txt", user+": "+chatIDStr+"\n")
	if err != nil {
		return " ", errors.New("error appending to file")
	}
	return chatIDStr, nil
}

func loadPremiumUsers(premiumUsersFilePath string) []string {

	file, err := os.Open(premiumUsersFilePath)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(file)

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err.Error())
	}

	var premiumUsersCharIDs []string
	for _, line := range lines {
		kvPair := strings.Split(line, ": ")
		premiumUsersCharIDs = append(premiumUsersCharIDs, kvPair[1])
	}

	return utility.RemoveDuplicateStrings(premiumUsersCharIDs)
}
