package service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	tr "goWin/services/bot_service/internal/translate"
	"log"
)

func commands(command string) string {
	m := map[string]string{
		"help":       "Привет данный бот создан для помощи изучения английского языка",
		"my_profile": "тут в будущем будет ваша статистика",
	}
	return m[command]
}

func Bot() {

	bot, err := tgbotapi.NewBotAPI("8317425992:AAFAB8YsyD7rENOuFhWfXmwV3y5KNU1dxFw")

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		if update.Message.IsCommand() {
			msg.Text = commands(update.Message.Command())
			_, _ = bot.Send(msg)
		} else {
			re_text := tr.Translate(update.Message.Text)

			msg = tgbotapi.NewMessage(update.Message.Chat.ID, re_text)

			_, _ = bot.Send(msg)
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	}
}
