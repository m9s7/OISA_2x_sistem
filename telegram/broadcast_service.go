package telegram

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func BroadcastToTelegram(msg string, chatID string, markdown string) map[string]interface{} {
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

func ReplyToMsg(reply string, msgID int, chatID int64) map[string]interface{} {
	url := "https://api.telegram.org/bot" + GetBotToken() + "/sendMessage"

	payload := strings.NewReader("{\"text\":\"" + reply +
		"\",\"parse_mode\":\"Markdown\"," +
		"\"disable_web_page_preview\":false," +
		"\"disable_notification\":false," +
		"\"reply_to_message_id\":" + strconv.Itoa(msgID) + "," +
		"\"chat_id\":\"" + strconv.FormatInt(chatID, 10) + "\"}")

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

func BroadcastToFree(msg string) {
	chatID := "-1001875397817"
	response := BroadcastToTelegram(msg, chatID, "MarkdownV2")
	CheckIfSent(response, msg, "free")
}

func BroadcastToDev(msg string, markdown string) {
	chatID := "1678076367"
	response := BroadcastToTelegram(msg, chatID, markdown)
	CheckIfSent(response, msg, "dev")
}
