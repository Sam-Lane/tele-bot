package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"github.com/sam-lane/tele-bot/pkg/chat"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("failed to load .env")
	}

	key := os.Getenv("TELEGRAMKEY")

	bot, err := tgbotapi.NewBotAPI(key)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s]-[%d] %s", update.Message.From.UserName, update.Message.Chat.ID, update.Message.Text)

		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

			switch update.Message.Command() {
			case "stackoverflow":
				chat.StackOverFlowQuery(update.Message.CommandArguments(), bot, msg)
			case "twitchinfo":
				chat.TwitchInfo(update.Message.CommandArguments(), bot, msg)
			}
			//bot.Send(msg)
			//sticker := tgbotapi.NewStickerShare(msg.ChatID, "CAACAgIAAxkBAAMdX3sVqJuYfKiWnANg2_P3RVw5bIQAAtkCAALoPPca_LbYWdCxUOcbBA")

		}

	}
}
