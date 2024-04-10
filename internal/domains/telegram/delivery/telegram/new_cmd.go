package telegram

import (
	"context"
	"github.com/NicoNex/echotron/v3"
	"time"
	"trade-union-service/internal/domains/telegram/domain/entity"
	"trade-union-service/internal/domains/telegram/errs"
)


func (b *bot) handleNewCommand(update *echotron.Update) StateFn {
	b.languageSupportMessage(update.Message)

	ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)
	defer cancel()

	if err := b.service.NewCommand(ctx, entity.NewCommandServiceDTO{
		HandleCommand: entity.HandleCommand{
			Lang:   update.Message.From.LanguageCode,
			ChatID: b.chatID,
		},
	}); err != nil {
		b.log.WithField(chatIDLoggingKey, b.chatID).Error("bot: handleNewCommand - b.service.NewCommand error: ", err.Error())
		return b.setState(ctx, stateDefault, b.handleMessage)
	}

	return b.setState(ctx, stateNewFirstName, b.handleNewFirstName)
}


func (b *bot) handleNewFirstName(update *echotron.Update) StateFn {
	ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)
	defer cancel()

	if update.Message == nil {
		return b.setState(ctx, stateNewFirstName, b.handleNewFirstName)
	}

	b.languageSupportMessage(update.Message)

	if err := b.service.NewCommandFirstNameState(ctx, entity.NewCommandFirstNameStateServiceDTO{
		HandleMessage: entity.HandleMessage{
			Lang:   update.Message.From.LanguageCode,
			ChatID: b.chatID,
			Text: update.Message.Text,
		},
	}); err != nil {
		b.log.WithField(chatIDLoggingKey, b.chatID).Error("bot: handleNewFirstName - b.service.NewCommandFirstNameState error: ", err.Error())
		return b.setState(ctx, stateDefault, b.handleMessage)
	}

	return b.setState(ctx, stateNewLastName, b.handleNewLastName)
}


func (b *bot) handleNewLastName(update *echotron.Update) StateFn {
	ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)
	defer cancel()

	if update.Message == nil {
		return b.setState(ctx, stateNewLastName, b.handleNewLastName)
	}

	b.languageSupportMessage(update.Message)

	if err := b.service.NewCommandLastNameState(ctx, entity.NewCommandLastNameStateServiceDTO{
		HandleMessage: entity.HandleMessage{
			Lang:   update.Message.From.LanguageCode,
			ChatID: b.chatID,
			Text:   update.Message.Text,
		},
	}); err != nil {
		b.log.WithField(chatIDLoggingKey, b.chatID).Error("bot: handleNewLastName - b.service.NewCommandLastNameStateerror: ", err.Error())
		return b.setState(ctx, stateDefault, b.handleMessage)
	}

	return b.setState(ctx, stateNewMiddleName, b.handleNewMiddleName)
}


func (b *bot) handleNewMiddleName(update *echotron.Update) StateFn {
	ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)
	defer cancel()

	if update.Message == nil {
		return b.setState(ctx, stateNewMiddleName, b.handleNewMiddleName)
	}

	b.languageSupportMessage(update.Message)

	if err := b.service.NewCommandMiddleNameState(ctx, entity.NewCommandMiddleNameStateServiceDTO{
		HandleMessage: entity.HandleMessage{
			Lang:   update.Message.From.LanguageCode,
			ChatID: b.chatID,
			Text: update.Message.Text,
		},
	}); err != nil {
		b.log.WithField(chatIDLoggingKey, b.chatID).Error("bot: handleNewLastName - b.service.NewCommandMiddleNameState error: ", err.Error())
		return b.setState(ctx, stateDefault, b.handleMessage)
	}

	return b.setState(ctx, stateNewConfirmationCallback, b.handleNewConfirmationCallback)
}

func (b *bot) handleNewConfirmationCallback(update *echotron.Update) StateFn {
	ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)
	defer cancel()

	if update.Message != nil {
		return b.setStateAndCall(ctx, stateDefault, b.handleMessage, update)
	}

	if update.CallbackQuery == nil {
		return b.setState(ctx, stateNewConfirmationCallback, b.handleNewConfirmationCallback)
	}

	b.languageSupportCallback(update.CallbackQuery)

	if err := b.service.NewCommandConfirmationState(ctx, entity.NewCommandConfirmationStateServiceDTO{
		HandleCallback: entity.HandleCallback{
			Lang:   update.CallbackQuery.From.LanguageCode,
			Data:   update.CallbackQuery.Data,
			ChatID: b.chatID,
		},
	}); err != nil {
		b.log.WithField(chatIDLoggingKey, b.chatID).Error("bot: handleNewConfirmationCallback - b.service.NewCommandConfirmationState error: ", err.Error())
		switch err {
		case errs.ErrUnknownAnswer:
			return b.setState(ctx, stateNewConfirmationCallback, b.handleNewConfirmationCallback)
		}
		return b.setState(ctx, stateDefault, b.handleMessage)
	}

	return b.setState(ctx, stateDefault, b.handleMessage)
}