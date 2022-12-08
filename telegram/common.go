package telegram

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func getBotToken() string {
	err := godotenv.Load("telegram\\.env")
	if err != nil {
		log.Fatalln("Telegram: Error loading .env")
	}
	return os.Getenv("BOT_TOKEN")
}
