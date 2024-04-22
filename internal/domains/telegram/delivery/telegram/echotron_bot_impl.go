package telegram

import (
	"context"
	"github.com/NicoNex/echotron/v3"
	"github.com/sirupsen/logrus"
	"time"
	"trade-union-service/internal/domains/telegram/domain/entity"
	"trade-union-service/internal/domains/telegram/errs"
	"trade-union-service/internal/domains/telegram/metrics"
)

type bot struct {
	log              *logrus.Logger
	messageHandlers  []MessageHandler
	callbackHandlers []CallbackHandler
	service          service
	chatID           int64
	state            StateFn
}

func newBot(destructCh chan destructChatID, service service, log *logrus.Logger) func(chatID int64) echotron.Bot {
	return func(chatID int64) echotron.Bot {
		bot := &bot{
			log:              log,
			messageHandlers:  nil,
			callbackHandlers: nil,
			service:          service,
			chatID:           chatID,
			state:            nil,
		}

		bot.messageHandlers = getBotMessageHandlers(bot)
		bot.callbackHandlers = getBotCallbackHandlers(bot)

		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		chatCurrentState, err := bot.service.GetChatCurrentState(ctx, entity.GetChatCurrentStateServiceDTO{
			ChatID: bot.chatID,
		})
		if err != nil {
			switch err {
			case errs.ErrChatCurrentStateNotExists:
				if err := bot.service.CreateChatCurrentState(ctx, entity.CreateChatCurrentStateServiceDTO{
					State:  string(stateDefault),
					ChatID: bot.chatID,
				}); err != nil {
					switch err {
					case errs.ErrChatCurrentStateAlreadyExists:
						bot.log.WithFields(logrus.Fields{
							chatIDLoggingKey: bot.chatID,
							domainLoggingKey: domainLoggingValue,
							layerLoggingKey:  layerLoggingValue,
						}).Error("newBot - bot.service.CreateChatCurrentState error: ", err.Error())
					default:
						metrics.IncrementErrorTotal()

						bot.log.WithFields(logrus.Fields{
							chatIDLoggingKey: bot.chatID,
							domainLoggingKey: domainLoggingValue,
							layerLoggingKey:  layerLoggingValue,
						}).Error("newBot - bot.service.CreateChatCurrentState error: ", err.Error())
					}
				}
			default:
				metrics.IncrementErrorTotal()

				bot.log.WithFields(logrus.Fields{
					chatIDLoggingKey: bot.chatID,
					domainLoggingKey: domainLoggingValue,
					layerLoggingKey:  layerLoggingValue,
				}).Error("newBot - bot.service.CreateChatCurrentState error: ", err.Error())
			}

			bot.state = bot.handleDefault
		} else {
			m := bot.getChatStates()

			state, ok := m[StateKey(chatCurrentState.State)]
			if ok {
				bot.state = state

				bot.log.WithFields(logrus.Fields{
					chatIDLoggingKey: bot.chatID,
					domainLoggingKey: domainLoggingValue,
					layerLoggingKey:  layerLoggingValue,
				}).Info("newBot - Set " + chatCurrentState.State + " handler")
			} else {
				bot.state = bot.handleDefault

				bot.log.WithFields(logrus.Fields{
					chatIDLoggingKey: bot.chatID,
					domainLoggingKey: domainLoggingValue,
					layerLoggingKey:  layerLoggingValue,
				}).Info("newBot - Set default handler")
			}
		}

		go func() {
			<-time.After(10 * time.Minute)

			bot.log.WithFields(logrus.Fields{
				chatIDLoggingKey: bot.chatID,
				domainLoggingKey: domainLoggingValue,
				layerLoggingKey:  layerLoggingValue,
			}).Info("newBot - send destruct bot")

			destructCh <- destructChatID(bot.chatID)
		}()

		return bot
	}
}

func (b *bot) Update(update *echotron.Update) {
	b.state = b.state(update)
}
