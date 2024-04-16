package telegram

import (
	"context"
	"github.com/NicoNex/echotron/v3"
	"github.com/sirupsen/logrus"
	"time"
	"trade-union-service/internal/domains/telegram/domain/entity"
	"trade-union-service/internal/domains/telegram/metrics"
)

type StateFn func(*echotron.Update) StateFn

type Handler struct {
	msgText  string
	handleFn StateFn
}

const (
	startCommand = "/start"
	newCommand   = "/new"
)

func getBotHandlers(bot *bot) []Handler {
	return []Handler{
		{
			msgText:  startCommand,
			handleFn: bot.handleStartCommand,
		},
		{
			msgText:  newCommand,
			handleFn: bot.handleNewCommand,
		},
	}
}

func (b *bot) handleDefault(update *echotron.Update) StateFn {
	if update.Message != nil {
		return b.handleMessage(update)
	}

	if update.CallbackQuery != nil {
		return b.handleCallbackQuery(update)
	}

	if update.EditedMessage != nil {
		return b.handleEditedMessage(update)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	return b.setState(ctx, stateDefault, b.handleDefault)
}

func (b *bot) handleMessage(update *echotron.Update) StateFn {
	metrics.IncrementMessageTotal()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	b.log.WithFields(logrus.Fields{
		chatIDLoggingKey: b.chatID,
		domainLoggingKey: domainLoggingValue,
		layerLoggingKey:  layerLoggingValue,
	}).Debug("handleMessage - Message Text: ", update.Message.Text)

	b.log.WithFields(logrus.Fields{
		chatIDLoggingKey: b.chatID,
		domainLoggingKey: domainLoggingValue,
		layerLoggingKey:  layerLoggingValue,
	}).Debug("handleMessage - Message LanguageCode: ", update.Message.From.LanguageCode)

	for i := range b.handlers {
		if update.Message.Text == b.handlers[i].msgText {
			ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)

			return b.setStateAndCall(ctx, StateKey(b.handlers[i].msgText), b.handlers[i].handleFn, update)
		}
	}

	if err := b.service.UnknownCommand(entity.UnknownCommandServiceDTO{
		HandleCommand: entity.HandleCommand{
			Lang:   update.Message.From.LanguageCode,
			ChatID: b.chatID,
		},
	}); err != nil {
		b.log.WithFields(logrus.Fields{
			chatIDLoggingKey: b.chatID,
			domainLoggingKey: domainLoggingValue,
			layerLoggingKey:  layerLoggingValue,
		}).Error("handleMessage - b.service.UnknownCommand error: ", err.Error())

		metrics.IncrementErrorTotal()

		return b.setState(ctx, stateDefault, b.handleDefault)
	}

	return b.setState(ctx, stateDefault, b.handleDefault)
}

func (b *bot) handleCallbackQuery(update *echotron.Update) StateFn {
	metrics.IncrementCallbackQueryTotal()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	b.log.WithFields(logrus.Fields{
		chatIDLoggingKey: b.chatID,
		domainLoggingKey: domainLoggingValue,
		layerLoggingKey:  layerLoggingValue,
	}).Debug("handleCallbackQuery - CallbackQuery Data: ", update.CallbackQuery.Data)

	return b.setState(ctx, stateDefault, b.handleDefault)
}

func (b *bot) handleEditedMessage(update *echotron.Update) StateFn {
	metrics.IncrementEditedMessageTotal()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	b.log.WithFields(logrus.Fields{
		chatIDLoggingKey: b.chatID,
		domainLoggingKey: domainLoggingValue,
		layerLoggingKey:  layerLoggingValue,
	}).Debug("handleEditedMessage - EditedMessage Text: ", update.EditedMessage.Text)

	return b.setState(ctx, stateDefault, b.handleDefault)
}
