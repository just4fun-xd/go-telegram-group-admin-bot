package main

import (
	"log"
	"os"

	"go-telegram-group-admin-bot/internal/bot"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN не найден в .env")
	}

	botAPI, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	botAPI.Debug = true
	log.Printf("Авторизован под %s", botAPI.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := botAPI.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			bot.HandleMessage(botAPI, update)
		}

		if update.CallbackQuery != nil {
			bot.HandleCallback(botAPI, update)
		}
	}
}
