package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
	"telegram-bot/internal/telegram-bot/service"
)

type Bot struct {
	bot *tgbotapi.BotAPI
	log *slog.Logger
	services *service.Service
}

func NewBot(bot *tgbotapi.BotAPI, log *slog.Logger, services *service.Service) *Bot {
	return &Bot{bot: bot, log: log, services: services}
}


func (bot *Bot) Start(botTimeout int) error {
	bot.log.Info("authorized", slog.String("account", bot.bot.Self.UserName))

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = botTimeout

	updates := bot.bot.GetUpdatesChan(updateConfig)

	if err := bot.handleUpdates(updates); err != nil {
		return err
	}

	return nil
}