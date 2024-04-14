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

		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
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

			bot.state = bot.handleDefault
		} else {
			m := bot.getChatStates()

			state, ok := m[chatCurrentState.State]
			if ok {
				bot.state = state
				bot.log.WithField(chatIDLoggingKey, bot.chatID).
					Info("bot: newBot - Set " + chatCurrentState.State + " handler")
			} else {
				bot.state = bot.handleDefault
				bot.log.WithField(chatIDLoggingKey, bot.chatID).
					Info("bot: newBot - Set default handler")
			}
		}

		return bot
	}
}

func (b *bot) Update(update *echotron.Update) {
	b.state = b.state(update)
}
