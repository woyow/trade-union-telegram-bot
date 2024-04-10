package telegram

import (
	"context"
	"github.com/NicoNex/echotron/v3"
	"github.com/sirupsen/logrus"
	"time"
	"trade-union-service/internal/domains/telegram/domain/entity"
	"trade-union-service/internal/domains/telegram/errs"
)

type bot struct {
	log      *logrus.Logger
	handlers []Handler
	service  service
	chatID   int64
	state    StateFn
}

func newBot(service service, log *logrus.Logger) func(chatID int64) echotron.Bot {
	return func(chatID int64) echotron.Bot {
		bot := &bot{
			service: service,
			chatID:  chatID,
			log:     log,
		}

		bot.handlers = getBotHandlers(bot)

		ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)
		defer cancel()

		chatCurrentState, err := bot.service.GetChatCurrentState(ctx, entity.GetChatCurrentStateServiceDTO{
			ChatID: bot.chatID,
		})
		if err != nil {
			switch err {
			case errs.ErrChatCurrentStateNotExists:
				if err := bot.service.CreateChatCurrentState(ctx, entity.CreateChatCurrentStateServiceDTO{
					State:  stateDefault,
					ChatID: bot.chatID,
				}); err != nil {
					bot.log.WithField(chatIDLoggingKey, bot.chatID).
						Error("bot: newBot - bot.service.CreateChatCurrentState error: ", err.Error())
				}
			}
			bot.state = bot.handleMessage
		} else {
			m := bot.getChatStates()
			state, ok := m[chatCurrentState.State]
			if ok {
				bot.state = state
				bot.log.WithField(chatIDLoggingKey, bot.chatID).Info("bot: newBot - Set " + chatCurrentState.State + " handler")
			} else {
				bot.state = bot.handleMessage
				bot.log.WithField(chatIDLoggingKey, bot.chatID).Info("bot: newBot - Set default handler")
			}
		}

		return bot
	}
}

func (b *bot) Update(update *echotron.Update) {
	//b.log.Debug("Update - message:", update.Message.Text)
	b.state = b.state(update)
}

type chatStatesMap map[string]StateFn

const (
	// Default
	stateEmpty = ""
	stateDefault = "default"

	// Start command
	stateStartCommand = "/start"

	// New command
	stateNewCommand = "/new"
	stateNewFirstName = "/new_first_name"
	stateNewLastName = "/new_last_name"
	stateNewMiddleName = "/new_middle_name"
	stateNewConfirmationCallback = "/new_confirmation_callback"
)

func (b *bot) getChatStates() chatStatesMap {
	m := chatStatesMap{
		// Default
		stateEmpty: b.handleMessage,
		stateDefault: b.handleMessage,

		// Start command
		stateStartCommand: b.handleStartCommand,

		// New command
		stateNewCommand: b.handleNewCommand,
		stateNewFirstName: b.handleNewFirstName,
		stateNewLastName: b.handleNewLastName,
		stateNewMiddleName: b.handleNewMiddleName,
		stateNewConfirmationCallback: b.handleNewConfirmationCallback,
	}

	return m
}