package telegram

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func broadcastToTelegram(msg string, chatID string, markdown string) map[string]interface{} {
	url := "https://api.telegram.org/bot" + GetBotToken() + "/sendMessage"

	payload := strings.NewReader(
		"{\"text\":\"" + msg + "\"," +
			"\"parse_mode\":\"" + markdown + "\"," +
			"\"disable_web_page_preview\":false," +
			"\"disable_notification\":false," +
			"\"reply_to_message_id\":null," +
			"\"chat_id\":\"" + chatID + "\"}")

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(res.Body)

	var resultJSON map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&resultJSON)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return resultJSON
}

func checkIfSent(response map[string]interface{}, msg string, groupName string) {
	if response == nil {
		log.Println("Message to", groupName, "group WAS NOT sent!")
		log.Println("Message: ", msg)
		return
	}
	if answer, ok := response["ok"]; !ok || answer.(bool) != true {
		log.Println("Message to", groupName, "group WAS NOT sent!")
		log.Println("Message: ", msg)
		log.Println("Error code:", response["error_code"].(float64), "Description:", response["description"].(string))
	}
}

func BroadcastToFree(msg string) {
	chatID := "-1001875397817"
	response := broadcastToTelegram(msg, chatID, "MarkdownV2")
	checkIfSent(response, msg, "free")
}

func BroadcastToPremium(msg string) {

	for _, user := range ChatIDs {
		response := broadcastToTelegram(msg, user, "MarkdownV2")
		checkIfSent(response, msg, "premium")
	}

	premiumChannel := "-1001701172026"
	response := broadcastToTelegram(msg, premiumChannel, "MarkdownV2")
	checkIfSent(response, msg, "premium")
}

func BroadcastToDev(msg string, markdown string) {
	chatID := "1678076367"
	response := broadcastToTelegram(msg, chatID, markdown)
	checkIfSent(response, msg, "dev")
}
