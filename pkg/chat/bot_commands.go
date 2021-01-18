package chat

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sam-lane/tele-bot/pkg/chat/stackoverflow"
	"github.com/sam-lane/tele-bot/pkg/twitch"
)

// TwitchInfo bot sends to chat info on a twitch channel
func TwitchInfo(channelName string, bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig) {
	channel, err := twitch.GetChannelInfo(channelName)
	if err != nil {
		msg.Text = err.Error()
	} else {
		var online = "💤"
		if channel.IsLive {
			online = "🔴"
		}
		msg.Text = fmt.Sprintf("%s\n\n%s\nhttps://twitch.tv/%s", channel.Title, online, channelName)
	}
	bot.Send(msg)
}

func StackOverFlowQuery(query string, bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig) {
	res, err := stackoverflow.Query(query)
	if err != nil {
		msg.Text = err.Error()
		bot.Send(msg)
		return
	}
	msg.Text = res.Items[0].Link
	bot.Send(msg)
}
