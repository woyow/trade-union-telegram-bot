package telegram

import (
	"context"
	"time"
	"trade-union-service/internal/domains/telegram/metrics"

	"trade-union-service/internal/domains/telegram/domain/entity"

	"github.com/NicoNex/echotron/v3"
)

func (b *bot) handleStartCommand(update *echotron.Update) StateFn {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := b.service.StartCommand(entity.StartCommandServiceDTO{
		HandleCommand: entity.HandleCommand{
			Lang:   update.Message.From.LanguageCode,
			ChatID: b.chatID,
		},
	}); err != nil {
		b.log.WithField(chatIDLoggingKey, b.chatID).
			Error("bot: handleStartCommand error: ", err.Error())

		metrics.IncrementErrorTotal()

		return b.setState(ctx, stateDefault, b.handleDefault)
	}

	return b.setState(ctx, stateDefault, b.handleDefault)
}
