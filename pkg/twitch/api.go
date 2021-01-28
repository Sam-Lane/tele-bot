package twitch

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/nicklaw5/helix"
)

func TwitchInfo(msg *tgbotapi.Message, bot *tgbotapi.BotAPI, reply *tgbotapi.MessageConfig) {
	channel, err := getChannelInfo(msg.CommandArguments())
	if err != nil {
		reply.Text = err.Error()
	} else {
		var online = "💤 Offline"
		if channel.IsLive {
			online = "🔴 Online"
		}
		reply.Text = fmt.Sprintf("%s\n\n%s\nhttps://twitch.tv/%s", channel.Title, online, channel.DisplayName)
		reply.DisableWebPagePreview = true
	}
	bot.Send(reply)
}

func getChannelInfo(channel string) (helix.Channel, error) {
	client, err := helix.NewClient(&helix.Options{
		ClientID:     "13mplr22v7515sx1d5y67e8q73w3t8",
		ClientSecret: "10a49cejt9jloxqqp4hnq8jrhxy7bb",
	})
	if err != nil {
		return helix.Channel{}, err
	}

	res, err := client.RequestAppAccessToken([]string{"user:read:email"})
	if err != nil {
		return helix.Channel{}, err
	}

	client.SetAppAccessToken(res.Data.AccessToken)
	resp, err := client.SearchChannels(&helix.SearchChannelsParams{
		Channel: channel,
		First:   1,
	})

	if err != nil {
		return helix.Channel{}, err
	}

	if len(resp.Data.Channels) == 0 {
		return helix.Channel{}, fmt.Errorf("no channels found")
	}
	return resp.Data.Channels[0], nil
}

func newTwitchWebHook(channel string) error {
	client, err := helix.NewClient(&helix.Options{
		//TODO: pull from env
		ClientID:     os.Getenv("TWITCHID"),
		ClientSecret: os.Getenv("TWITCHSECRET"),
	})
	if err != nil {
		log.Printf("failed to query twitch api: %s", err.Error())
		return fmt.Errorf("Failed to get channel %s", channel)
	}
	client.PostWebhookSubscription(&helix.WebhookSubscriptionPayload{
		Mode: "subscribe",
	})
	return nil
}
