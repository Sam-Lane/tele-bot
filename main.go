package main

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"github.com/sam-lane/tele-bot/pkg/bot"
	"github.com/sam-lane/tele-bot/pkg/stackoverflow"
	"github.com/sam-lane/tele-bot/pkg/twitch"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("failed to load .env")
	}

	key := os.Getenv("TELEGRAMKEY")

	bot, err := bot.NewBot(key, 0, 60)

	if err != nil {
		log.Panic(err)
	}

	bot.RegisterCommand("stackoverflow", stackoverflow.StackOverFlowQuery)
	bot.RegisterCommand("twitchinfo", twitch.TwitchInfo)

	bot.RegisterCommand("roll", func(msg *tgbotapi.Message, bot *tgbotapi.BotAPI, reply *tgbotapi.MessageConfig) {
		rand.Seed(time.Now().UnixNano())
		reply.Text = strconv.FormatInt(int64(rand.Intn(101)), 10)
		bot.Send(reply)
	})

	bot.Start()
}
