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

type MessageHandler struct {
	text     string
	handleFn StateFn
}

type CallbackHandler struct {
	data     string
	handleFn StateFn
}

const (
	startCommand = "/start"
	newCommand   = "/new"
)

func getBotMessageHandlers(bot *bot) []MessageHandler {
	return []MessageHandler{
		{
			text:     startCommand,
			handleFn: bot.handleStartCommand,
		},
		{
			text:     newCommand,
			handleFn: bot.handleNewCommand,
		},
	}
}

func getBotCallbackHandlers(bot *bot) []CallbackHandler {
	return []CallbackHandler{}
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
	}).Debug("handleMessage - Message text: ", update.Message.Text)

	b.log.WithFields(logrus.Fields{
		chatIDLoggingKey: b.chatID,
		domainLoggingKey: domainLoggingValue,
		layerLoggingKey:  layerLoggingValue,
	}).Debug("handleMessage - Message LanguageCode: ", update.Message.From.LanguageCode)

	for i := range b.messageHandlers {
		if update.Message.Text == b.messageHandlers[i].text {
			ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)

			return b.setStateAndCall(ctx, StateKey(b.messageHandlers[i].text), b.messageHandlers[i].handleFn, update)
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

	for i := range b.callbackHandlers {
		if update.CallbackQuery.Data == b.callbackHandlers[i].data {
			ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)

			return b.setStateAndCall(ctx, StateKey(b.callbackHandlers[i].data), b.callbackHandlers[i].handleFn, update)
		}
	}

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
	}).Debug("handleEditedMessage - EditedMessage text: ", update.EditedMessage.Text)

	return b.setState(ctx, stateDefault, b.handleDefault)
}
