package telegram

import (
	"context"
	"github.com/NicoNex/echotron/v3"
	"trade-union-service/internal/domains/telegram/domain/entity"
)

type chatStatesMap map[string]StateFn

const (
	// Default.
	stateEmpty   = ""
	stateDefault = "default"

	// Start command.
	stateStartCommand = "/start"

	// New command.
	stateNewCommand              = "/new"
	stateNewFirstName            = "/new_first_name"
	stateNewLastName             = "/new_last_name"
	stateNewMiddleName           = "/new_middle_name"
	stateNewSubjectCallback      = "/new_subject_callback"
	stateNewConfirmationCallback = "/new_confirmation_callback"
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
		stateNewFirstName:            b.handleNewFirstName,
		stateNewLastName:             b.handleNewLastName,
		stateNewMiddleName:           b.handleNewMiddleName,
		stateNewSubjectCallback:      b.handleNewSubjectCallback,
		stateNewConfirmationCallback: b.handleNewConfirmationCallback,
	}

	return chatStates
}

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
