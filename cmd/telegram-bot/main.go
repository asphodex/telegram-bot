package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
	"os"
	"telegram-bot/internal/pkg/logger"
	"telegram-bot/internal/telegram-bot/config"
	"telegram-bot/internal/telegram-bot/handler"
	"telegram-bot/internal/telegram-bot/repository"
	"telegram-bot/internal/telegram-bot/repository/postgres"
	"telegram-bot/internal/telegram-bot/service"
)

var (
	botTimeout = 60
)

func main() {
	cfg := config.MustLoad()
	log := sl.SetupLogger(cfg.Env)

	storage, err := postgres.NewPostgresDB(&cfg.PostgreSQL, os.Getenv("PG_PASS"))
	if err != nil {
		log.Error("failed to connect to db", sl.Err(err), slog.String("PG_PASS", os.Getenv("PG_PASS")))
		os.Exit(1)
	}

	log.Info("init layers", slog.String("env", cfg.Env))

	repo := repository.NewRepository(storage)
	services := service.NewService(repo)

	log.Info("starting the telegram bot", slog.String("env", cfg.Env))

	bot, err := setupBot(false, os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		log.Error("unable setup the bot: %v", sl.Err(err))
		os.Exit(1)
	}

	if err = handler.NewBot(bot, log, services).Start(botTimeout); err != nil {
		log.Error("unable start the bot: %v", sl.Err(err))
		os.Exit(1)
	}
}

func setupBot(debug bool, token string) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	bot.Debug = debug

	return bot, nil
}

