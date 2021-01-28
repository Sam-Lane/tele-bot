package eightball

import (
	"math/rand"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func Ball(msg *tgbotapi.Message, bot *tgbotapi.BotAPI, reply *tgbotapi.MessageConfig) {
	if len(msg.CommandArguments()) == 0 {
		reply.Text = "That's not a question for the magic 8 ball 🎱"
		bot.Send(reply)
		return
	}
	rand.Seed(time.Now().UnixNano())
	resp := []string{
		"🟢 It is certain.",
		"🟢 It is decidedly so.",
		"🟢 Without a doubt.",
		"🟢 Yes – definitely.",
		"🟢 You may rely on it.",
		"🟢 As I see it, yes.",
		"🟢 Most likely.",
		"🟢 Outlook good.",
		"🟢 Yes.",
		"🟢 Signs point to yes",
		"🟡 Reply hazy, try again.",
		"🟡 Ask again later.",
		"🟡 Better not tell you now.",
		"🟡 Cannot predict now.",
		"🟡 Concentrate and ask again.",
		"🔴 Don't count on it.",
		"🔴 My reply is no.",
		"🔴 My sources say no.",
		"🔴 Outlook not so good.",
		"🔴 Very doubtful.",
	}
	n := rand.Int() % len(resp)
	reply.Text = resp[n]
	bot.Send(reply)
}
