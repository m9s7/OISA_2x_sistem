package telegram

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

func GetBotToken() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Telegram: Error loading .env")
	}
	return os.Getenv("BOT_TOKEN")
}

func EscapeTelegramSpecChars(str string) string {
	specialCharacters := []string{"_", "*", "[", "]", "<", ">", "\"", "(", ")", "~", "`", "#", "+", "-", "=", "|", "{", "}", ".", "!"}

	for _, char := range specialCharacters {
		str = strings.Replace(str, char, "\\"+char, -1)
	}

	return str
}

func CheckIfSent(response map[string]interface{}, msg string, groupName string) {
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
