package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

type Bot struct {
	Commands      map[string]func(*tgbotapi.Message, *tgbotapi.BotAPI, *tgbotapi.MessageConfig)
	tgBot         *tgbotapi.BotAPI
	updateChannel tgbotapi.UpdatesChannel
}

func NewBot(key string, offset, timeout int) (*Bot, error) {
	tgbot, err := tgbotapi.NewBotAPI(key)
	if err != nil {
		return nil, err
	}
	tgbot.Debug = false
	u := tgbotapi.NewUpdate(offset)
	u.Timeout = 60
	updates, err := tgbot.GetUpdatesChan(u)

	return &Bot{
		tgBot:         tgbot,
		updateChannel: updates,
		Commands:      make(map[string]func(*tgbotapi.Message, *tgbotapi.BotAPI, *tgbotapi.MessageConfig)),
	}, nil
}

func (b *Bot) Start() {
	for update := range b.updateChannel {
		if update.Message == nil {
			continue
		}
		if update.Message.IsCommand() {
			reply := tgbotapi.NewMessage(update.Message.Chat.ID, "")

			cmd, ok := b.Commands[update.Message.Command()]
			if ok {
				go cmd(update.Message, b.tgBot, &reply)
			}
		}
	}
}

func (b *Bot) RegisterCommand(invocation string, cmd func(*tgbotapi.Message, *tgbotapi.BotAPI, *tgbotapi.MessageConfig)) {
	b.Commands[invocation] = cmd
}
