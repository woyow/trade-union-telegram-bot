package telegram

import (
	"context"
	"github.com/NicoNex/echotron/v3"
	"trade-union-service/internal/domains/telegram/domain/entity"
	"trade-union-service/internal/domains/telegram/metrics"
)

type StateKey string
type chatStatesMap map[StateKey]StateFn

const (
	// Default.
	stateEmpty   = StateKey("")
	stateDefault = StateKey("default")

	// Start command.
	stateStartCommand = StateKey("/start")

	// New command.
	stateNewCommand              = StateKey("/new")
	stateNewFullName             = StateKey("/new_full_name")
	stateNewSubjectCallback      = StateKey("/new_subject_callback")
	stateNewConfirmationCallback = StateKey("/new_confirmation_callback")
)

func (b *bot) getChatStates() chatStatesMap {
	chatStates := chatStatesMap{
		// Default
		stateEmpty:   b.handleDefault,
		stateDefault: b.handleDefault,

		// Start command
		stateStartCommand: b.handleStartCommand,

		// New command
		stateNewCommand:              b.handleNewCommand,
		stateNewFullName:             b.handleNewFullName,
		stateNewSubjectCallback:      b.handleNewSubjectCallback,
		stateNewConfirmationCallback: b.handleNewConfirmationCallback,
	}

	return chatStates
}

func (b *bot) setStateAndCall(ctx context.Context, state StateKey, stateFn StateFn, update *echotron.Update) StateFn {
	metrics.IncrementSetStateAndCallTotal()

	if err := b.service.SetChatCurrentState(ctx, entity.SetChatCurrentStateServiceDTO{
		State:  string(state),
		ChatID: b.chatID,
	}); err != nil {
		b.log.WithField(chatIDLoggingKey, b.chatID).
			Error("bot: setState - b.service.SetChatCurrentState error: ", err.Error())

		metrics.IncrementSetStateAndCallErrorTotal()
	}

	return stateFn(update)
}

func (b *bot) setState(ctx context.Context, state StateKey, stateFn StateFn) StateFn {
	metrics.IncrementSetStateTotal()

	if err := b.service.SetChatCurrentState(ctx, entity.SetChatCurrentStateServiceDTO{
		State:  string(state),
		ChatID: b.chatID,
	}); err != nil {
		b.log.WithField(chatIDLoggingKey, b.chatID).
			Error("bot: setState - b.service.SetChatCurrentState error: ", err.Error())

		metrics.IncrementSetStateErrorTotal()
	}

	return stateFn
}
