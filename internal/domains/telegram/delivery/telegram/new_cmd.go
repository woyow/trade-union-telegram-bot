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

func (b *bot) handleNewCommand(update *echotron.Update) StateFn {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := b.service.NewCommand(ctx, entity.NewCommandServiceDTO{
		HandleCommand: entity.HandleCommand{
			Lang:   update.Message.From.LanguageCode,
			ChatID: b.chatID,
		},
	}); err != nil {
		b.log.WithFields(logrus.Fields{
			chatIDLoggingKey: b.chatID,
			domainLoggingKey: domainLoggingValue,
			layerLoggingKey:  layerLoggingValue,
		}).Error("handleNewCommand - b.service.NewCommand error: ", err.Error())

		metrics.IncrementErrorTotal()

		return b.setState(ctx, stateDefault, b.handleDefault)
	}

	return b.setState(ctx, stateNewFullName, b.handleNewFullName)
}

func (b *bot) handleNewFullName(update *echotron.Update) StateFn {
	metrics.IncrementMessageTotal()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if update.Message == nil {
		return b.setState(ctx, stateNewFullName, b.handleNewFullName)
	}

	if err := b.service.NewCommandFullNameState(ctx, entity.NewCommandFullNameStateServiceDTO{
		HandleMessage: entity.HandleMessage{
			Lang:   update.Message.From.LanguageCode,
			ChatID: b.chatID,
			Text:   update.Message.Text,
		},
	}); err != nil {
		b.log.WithFields(logrus.Fields{
			chatIDLoggingKey: b.chatID,
			domainLoggingKey: domainLoggingValue,
			layerLoggingKey:  layerLoggingValue,
		}).Error("handleNewFullName - b.service.NewCommandFullNameState error: ", err.Error())

		metrics.IncrementErrorTotal()

		return b.setState(ctx, stateDefault, b.handleDefault)
	}

	return b.setState(ctx, stateNewSubjectCallback, b.handleNewSubjectCallback)
}

func (b *bot) handleNewSubjectCallback(update *echotron.Update) StateFn {
	metrics.IncrementCallbackQueryTotal()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if update.Message != nil {
		return b.setStateAndCall(ctx, stateDefault, b.handleDefault, update)
	}

	if update.CallbackQuery == nil {
		return b.setState(ctx, stateNewSubjectCallback, b.handleNewSubjectCallback)
	}

	if err := b.service.NewCommandSubjectState(ctx, entity.NewCommandSubjectStateServiceDTO{
		HandleCallback: entity.HandleCallback{
			Lang:   update.CallbackQuery.From.LanguageCode,
			Data:   update.CallbackQuery.Data,
			ChatID: b.chatID,
		},
	}); err != nil {
		b.log.WithFields(logrus.Fields{
			chatIDLoggingKey: b.chatID,
			domainLoggingKey: domainLoggingValue,
			layerLoggingKey:  layerLoggingValue,
		}).Error("handleNewSubjectCallback - b.service.NewCommandSubjectState error: ", err.Error())

		switch err {
		case errs.ErrUnknownAnswer:
			metrics.IncrementBusinessErrorTotal()
			return b.setState(ctx, stateNewSubjectCallback, b.handleNewSubjectCallback)
		default:
			metrics.IncrementErrorTotal()
			return b.setState(ctx, stateDefault, b.handleDefault)
		}
	}

	return b.setState(ctx, stateNewConfirmationCallback, b.handleNewConfirmationCallback)
}

func (b *bot) handleNewConfirmationCallback(update *echotron.Update) StateFn {
	metrics.IncrementCallbackQueryTotal()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if update.Message != nil {
		return b.setStateAndCall(ctx, stateDefault, b.handleDefault, update)
	}

	if update.CallbackQuery == nil {
		return b.setState(ctx, stateNewConfirmationCallback, b.handleNewConfirmationCallback)
	}

	if err := b.service.NewCommandConfirmationState(ctx, entity.NewCommandConfirmationStateServiceDTO{
		HandleCallback: entity.HandleCallback{
			Lang:   update.CallbackQuery.From.LanguageCode,
			Data:   update.CallbackQuery.Data,
			ChatID: b.chatID,
		},
	}); err != nil {
		b.log.WithFields(logrus.Fields{
			chatIDLoggingKey: b.chatID,
			domainLoggingKey: domainLoggingValue,
			layerLoggingKey:  layerLoggingValue,
		}).Error("handleNewConfirmationCallback - b.service.NewCommandConfirmationState error: ", err.Error())

		switch err {
		case errs.ErrUnknownAnswer:
			metrics.IncrementBusinessErrorTotal()
			return b.setState(ctx, stateNewConfirmationCallback, b.handleNewConfirmationCallback)
		default:
			metrics.IncrementErrorTotal()
			return b.setState(ctx, stateDefault, b.handleDefault)
		}
	}

	return b.setState(ctx, stateDefault, b.handleDefault)
}
