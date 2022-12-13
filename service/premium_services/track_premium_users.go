package premium_services

import (
	"OISA_2x_sistem/utility"
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var ChatIDs []string
var usersFilePath = "service\\premium_services\\users.txt"

func RegisterNewUser(msgText string, chatID int64) (string, error) {

	user := strings.TrimLeft(msgText, "/RrEeGgIiSsTtEe ")
	chatIDStr := strconv.Itoa(int(chatID))

	err := utility.AppendToFile(usersFilePath, user+": "+chatIDStr+"\n")
	if err != nil {
		return "", errors.New("error appending to file")
	}
	return chatIDStr, nil
}

func LoadPremiumUsers() []string {

	file, err := os.Open(usersFilePath)
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
