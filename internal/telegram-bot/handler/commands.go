package handler

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
	"strconv"
	"strings"
	sl "telegram-bot/internal/pkg/logger"
	"telegram-bot/internal/pkg/raketa"
)

func (bot *Bot) startCommand(msg *tgbotapi.MessageConfig, update *tgbotapi.Update) error {
	if err := bot.services.Statistic.AddUser(update.Message.Chat.ID, update.SentFrom().UserName); err != nil {
		return err
	}

	msg.Text = startMessage
	return nil
}

func (bot *Bot) searchCommand(msg *tgbotapi.MessageConfig, update *tgbotapi.Update) error {
	track := strings.TrimSpace(update.Message.CommandArguments())
	if track == "" || strings.Contains(track, " ") {
		msg.Text = emptyTrackMessage
		return nil
	}

	trace, err := raketa.GetDeliveryTrace(track)
	if err != nil {
		bot.log.Error("error occurred while getting delivery step from raketa: ", sl.Err(err))
		return fmt.Errorf("error occurred while getting delivery step from raketa")
	}

	var message string

	if trace == nil {
		msg.Text = noOrdersMessage
		return nil
	}

	for _, step := range trace.Step {
		message = message + fmt.Sprintf("[<code>%s</code>] %s\n", step.Date.Format("02.01 15:04"), step.Status)
	}

	msg.Text = fmt.Sprintf("%s\n%s", track, message)
	return nil
}

func (bot *Bot) statisticsCommand(msg *tgbotapi.MessageConfig, update *tgbotapi.Update) error {
	if strconv.FormatInt(update.SentFrom().ID, 10) != os.Getenv("ADMIN") {
		msg.Text = forbiddenMessage
		return nil
	}

	uniqUsers, err := bot.services.GetCountOfUsers()
	if err != nil {
		msg.Text = errorMessage
		return err
	}

	msg.Text = fmt.Sprintf("%s\nУникальные пользователи: %d", statisticsMessage, uniqUsers)
	return nil
}

