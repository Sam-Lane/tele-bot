package deathroll

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func ExecCommand(msg *tgbotapi.Message, bot *tgbotapi.BotAPI, reply *tgbotapi.MessageConfig) {
	arg := strings.SplitN(msg.CommandArguments(), " ", 1)[0]

	var randNum int64 = 0

	if arg == "" {
		randNum = getRandomNumberWithMaxNumber(100)
	} else {
		num, err := strconv.Atoi(arg)

		if err != nil {
			reply.Text = "❌ Given argument is not a number!"
			bot.Send(reply)
			return
		}

		randNum = getRandomNumberWithMaxNumber(num)
	}

	reply.Text = strconv.FormatInt(randNum, 10)
	bot.Send(reply)
}

func getRandomNumberWithMaxNumber(maxNum int) int64 {
	rand.Seed(time.Now().UnixNano())
	// generate from 0-Max +1, so that rolling 0 is not possible
	return int64(rand.Intn(maxNum) + 1)
}
