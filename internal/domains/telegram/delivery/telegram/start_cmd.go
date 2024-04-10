package telegram

import (
	"context"
	"time"

	"trade-union-service/internal/domains/telegram/domain/entity"

	"github.com/NicoNex/echotron/v3"
)

func (b *bot) handleStartCommand(update *echotron.Update) StateFn {
	b.languageSupportMessage(update.Message)

	if err := b.service.StartCommand(entity.StartCommandServiceDTO{
		HandleCommand: entity.HandleCommand{
			Lang:   update.Message.From.LanguageCode,
			ChatID: b.chatID,
		},
	}); err != nil {
		b.log.WithField(chatIDLoggingKey, b.chatID).Error("bot: handleStartCommand error: ", err.Error())
		return b.handleMessage
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)
	defer cancel()

	return b.setState(ctx, stateDefault, b.handleMessage)
}
