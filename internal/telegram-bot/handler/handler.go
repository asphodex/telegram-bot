package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (bot *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) error {
	for update := range updates {
		if update.Message.IsCommand() {
			bot.handleCommand(update)
		}
	}
	return nil
}

func (bot *Bot) handleCommand(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, defaultMessage)
	msg.ParseMode = "HTML"
	switch update.Message.Command() {
		case "start":
			err := bot.start(&msg, &update)
			if err != nil {
				bot.log.Error("start error: ", err)
			}
			break
		case "search":
			err := bot.search(&msg, &update)
			if err != nil {
				bot.log.Error("search error: ", err)
			}
			break
		case "statistics":
			err := bot.statistics(&msg, &update)
			if err != nil {
				bot.log.Error("statistics error: ", err)
			}
			break
	}

	_, err := bot.bot.Send(msg)
	if err != nil {
		bot.log.Error("error sending message: ", err)
	}
}