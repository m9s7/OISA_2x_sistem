package service

import (
	"OISA_2x_sistem/telegram"
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
	bot, err := tgbotapi.NewBotAPI(telegram.GetBotToken())
	if err != nil {
		log.Panic(err)
	}

	ChatIDs = loadPremiumUsers("service\\users.txt")

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
				telegram.BroadcastToDev("Error while writing user to file: "+err.Error(), "HTML")
			}
			//:TODO lock this when writing to it, but it happens so rarely that it's not worth it
			ChatIDs = append(ChatIDs, newUser)
			continue
		}

		if update.Message.ReplyToMessage == nil {
			continue
		}

		// ulog
		if strings.EqualFold(update.Message.Command(), "ulog") {
			err := ProvideArbCalculatorService(update.Message)
			if err != nil {
				log.Println(err.Error())
				errorMsg := "Error while providing arb calculator service: " + err.Error() + "msg sent: " + update.Message.Text
				telegram.BroadcastToDev(errorMsg, "HTML")
				continue
			}
		}

		// extra
		if strings.EqualFold(update.Message.Command(), "extra") {
			ProvideArbExtraDetailsService(update.Message)
		}

	}
}

func registerNewUser(msgText string, chatID int64) (string, error) {

	user := strings.TrimPrefix(msgText, "/register ")
	chatIDStr := strconv.Itoa(int(chatID))

	// TODO: move to premium, else move to premium_service package
	telegram.BroadcastToDev(
		telegram.EscapeTelegramSpecChars("User: "+user+" has started the bot!\nChat ID: "+chatIDStr), "HTML")

	err := utility.AppendToFile("service\\users.txt", user+": "+chatIDStr+"\n")
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

// TODO: keep this here, other functions to premium_service package

func BroadcastToPremium(msg string) {

	for _, user := range ChatIDs {
		response := telegram.BroadcastToTelegram(msg, user, "MarkdownV2")
		telegram.CheckIfSent(response, msg, "service")
	}

	premiumChannel := "-1001701172026"
	response := telegram.BroadcastToTelegram(msg, premiumChannel, "MarkdownV2")
	telegram.CheckIfSent(response, msg, "service")
}
