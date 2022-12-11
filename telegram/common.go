package telegram

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

func GetBotToken() string {
	err := godotenv.Load("telegram\\.env")
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
