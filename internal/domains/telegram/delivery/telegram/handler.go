package telegram

import (
	"context"
	"github.com/NicoNex/echotron/v3"
	"time"
	"trade-union-service/internal/domains/telegram/domain/entity"
)

type StateFn func(*echotron.Update) StateFn

type Handler struct {
	msgText  string
	handleFn StateFn
}

func getBotHandlers(bot *bot) []Handler {
	return []Handler{
		{
			msgText:  "/start",
			handleFn: bot.handleStartCommand,
		},
		{
			msgText:  "/new",
			handleFn: bot.handleNewCommand,
		},
	}
}

const (
	langRu = "ru"
	langEn = "en"
	defaultLang = langRu
)

func (b *bot) languageSupportMessage(msg *echotron.Message) {
	switch msg.From.LanguageCode {
	case langEn, langRu:
	default:
		msg.From.LanguageCode = defaultLang
	}
}

func (b *bot) languageSupportCallback(callback *echotron.CallbackQuery) {
	switch callback.From.LanguageCode {
	case langEn, langRu:
	default:
		callback.From.LanguageCode = defaultLang
	}
}

func (b *bot) handleMessage(update *echotron.Update) StateFn {
	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	defer cancel()

	if update.Message != nil {
		b.log.WithField(chatIDLoggingKey, b.chatID).Debug("handleMessage - Message Text: ", update.Message.Text)
		b.log.WithField(chatIDLoggingKey, b.chatID).Debug("handleMessage - Message LanguageCode: ", update.Message.From.LanguageCode)

		b.languageSupportMessage(update.Message)

		for i := range b.handlers {
			switch update.Message.Text {
			case b.handlers[i].msgText:
				ctx, _ := context.WithTimeout(context.Background(), 1 * time.Second)
				return b.setStateAndCall(ctx, b.handlers[i].msgText, b.handlers[i].handleFn, update)
			}
		}

		if err := b.service.UnknownCommand(entity.UnknownCommandServiceDTO{
			HandleCommand: entity.HandleCommand{
				Lang:   update.Message.From.LanguageCode,
				ChatID: b.chatID,
			},
		}); err != nil {
			b.log.WithField(chatIDLoggingKey, b.chatID).Error("bot: handleMessage - b.service.UnknownCommand error: ", err.Error())
			return b.setState(ctx, stateDefault, b.handleMessage)
		}
	}

	return b.setState(ctx, stateDefault, b.handleMessage)
}