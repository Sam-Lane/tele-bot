package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sam-lane/tele-bot/pkg/bot"
	"github.com/sam-lane/tele-bot/pkg/deathroll"
	"github.com/sam-lane/tele-bot/pkg/eightball"
	"github.com/sam-lane/tele-bot/pkg/stackoverflow"
	"github.com/sam-lane/tele-bot/pkg/twitch"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("failed to load .env")
	}

	key := os.Getenv("TELEGRAMKEY")

	bot, err := bot.NewBot(key, 0, 60)

	if err != nil {
		log.Panic(err)
	}

	bot.RegisterCommand("stackoverflow", stackoverflow.StackOverFlowQuery)
	bot.RegisterCommand("twitchinfo", twitch.TwitchInfo)
	bot.RegisterCommand("roll", deathroll.ExecCommand)
	bot.RegisterCommand("8ball", eightball.Ball)

	bot.Start()
}
