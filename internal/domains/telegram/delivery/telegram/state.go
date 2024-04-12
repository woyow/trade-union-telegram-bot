package telegram

import (
	"context"
	"github.com/NicoNex/echotron/v3"
	"trade-union-service/internal/domains/telegram/domain/entity"
)

func (b *bot) setStateAndCall(ctx context.Context, state string, stateFn StateFn, update *echotron.Update) StateFn {
	if err := b.service.SetChatCurrentState(ctx, entity.SetChatCurrentStateServiceDTO{
		State:  state,
		ChatID: b.chatID,
	}); err != nil {
		b.log.WithField(chatIDLoggingKey, b.chatID).
			Error("bot: setState - b.service.SetChatCurrentState error: ", err.Error())
	}

	return stateFn(update)
}

func (b *bot) setState(ctx context.Context, state string, stateFn StateFn) StateFn {
	if err := b.service.SetChatCurrentState(ctx, entity.SetChatCurrentStateServiceDTO{
		State:  state,
		ChatID: b.chatID,
	}); err != nil {
		b.log.WithField(chatIDLoggingKey, b.chatID).
			Error("bot: setState - b.service.SetChatCurrentState error: ", err.Error())
	}

	return stateFn
}
